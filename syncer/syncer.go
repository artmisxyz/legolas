package syncer

import (
	"context"
	"fmt"
	"github.com/artmisxyz/blockinspector/connections"
	"github.com/artmisxyz/blockinspector/ent"
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/artmisxyz/blockinspector/inspector/uniswapv3"
	"github.com/artmisxyz/blockinspector/position"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"math/big"
	"os"
)

type Syncer struct {
	newHeaders chan *types.Header
	ws         *ethclient.Client
	rpc        *ethclient.Client

	logger *zap.Logger

	blockPositionHolder position.BlockPositionHolder

	currentBlockPosition uint64
	lag                  uint64

	inspectors map[string]inspector.Inspector
}

func (s *Syncer) Init(conf Config) {
	var lvl zapcore.Level
	err := lvl.Set(conf.Logger.Level)
	if err != nil {
		fmt.Printf("cannot parse log level %s: %s", conf.Logger.Level, err)
		lvl = zapcore.WarnLevel
	}
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	defaultCore := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stderr)), lvl)
	cores := []zapcore.Core{
		defaultCore,
	}

	core := zapcore.NewTee(cores...)
	s.logger = zap.New(core, zap.AddCaller())

	s.ws = connections.RPC(conf.Node.Websocket)
	s.rpc = connections.Websocket(conf.Node.RPC)

	s.newHeaders, err = connections.CurrentBlockChan(s.logger, s.ws)
	if err != nil {
		panic(err)
	}
	s.blockPositionHolder = position.NewFilePosition(s.logger, conf.General.PosFileLocation, conf.General.PosFileName, conf.General.StartBlockNumber)
	if ok := s.blockPositionHolder.Exists(); !ok {
		err := s.blockPositionHolder.Create()
		if err != nil {
			panic(err)
		}
	}
	s.currentBlockPosition, err = s.blockPositionHolder.Read()
	if err != nil {
		panic(err)
	}
	if s.currentBlockPosition == 0 {
		s.currentBlockPosition = conf.General.StartBlockNumber
	}
	s.inspectors = make(map[string]inspector.Inspector)
	db, err := ent.Open("sqlite3", "file:db?cache=shared&_fk=1")
	if err != nil {
		s.logger.Fatal("failed opening connection to sqlite.", zap.Error(err))
	}
	// Run the auto migration tool.
	if err := db.Schema.Create(context.Background()); err != nil {
		s.logger.Fatal("failed creating schema resources", zap.Error(err))
	}
	s.registerInspectors(uniswapv3.NewUniswapV3(s.logger, s.ws, db))
}

func (s *Syncer) Sync() {
	newHeader := <-s.newHeaders

	for s.currentBlockPosition < newHeader.Number.Uint64()-s.lag {
		block, err := s.rpc.BlockByNumber(context.Background(), big.NewInt(int64(s.currentBlockPosition)))
		if err != nil {
			s.logger.Error("error getting block", zap.Error(err))
		}
		s.logger.
			Info("processing block",
				zap.Uint64("current_block", s.currentBlockPosition),
				zap.Float64("progress", float64(s.currentBlockPosition)/float64(newHeader.Number.Uint64())))

		err = s.Inspect(block)
		if err != nil {
			s.logger.Error("error inspecting block", zap.Uint64("block_number", s.currentBlockPosition), zap.Error(err))
			break
		}
		err = s.blockPositionHolder.Update(s.currentBlockPosition)
		if err != nil {
			s.logger.Error("error updating block position file", zap.Error(err))
			break
		}
		s.currentBlockPosition++
	}
}

func (s *Syncer) Inspect(block *types.Block) error {
	for _, v := range s.inspectors {
		err := v.InspectBlock(block)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Syncer) registerInspectors(inspectors ...inspector.Inspector) {
	for _, i := range inspectors {
		s.inspectors[i.Name()] = i
	}
}

func (s *Syncer) getInspector(name string) inspector.Inspector {
	return s.inspectors[name]
}
