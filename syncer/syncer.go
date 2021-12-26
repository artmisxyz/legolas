package syncer

import (
	"context"
	"fmt"
	"github.com/artmisxyz/legolas/blockposition"
	"github.com/artmisxyz/legolas/configs"
	"github.com/artmisxyz/legolas/connections"
	"github.com/artmisxyz/legolas/database"
	"github.com/artmisxyz/legolas/inspector"
	"github.com/artmisxyz/legolas/inspector/uniswapv3"
	"github.com/artmisxyz/legolas/logger"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"math/big"
)

type Syncer struct {
	newHeaders chan *types.Header
	ws         *ethclient.Client
	rpc        *ethclient.Client

	logger *zap.Logger

	blockPositionHolder blockposition.BlockPositionHolder

	currentBlockPosition uint64
	lag                  uint64

	inspectors map[string]inspector.Inspector
}

func (s *Syncer) Init(conf configs.Config) {
	l := logger.New(conf)
	s.logger = l.Named("syncer")

	s.ws = connections.RPC(conf.Node.Websocket)
	s.rpc = connections.Websocket(conf.Node.RPC)

	newHeaders, err := connections.CurrentBlockChan(s.logger, s.ws)
	if err != nil {
		panic(err)
	}
	s.newHeaders = newHeaders

	db, err := database.NewPostgresClient(conf)
	if err != nil {
		panic(err)
	}

	s.inspectors = make(map[string]inspector.Inspector)
	univ3Inspector := uniswapv3.NewUniswapV3(s.logger, s.ws, db)
	s.registerInspectors(univ3Inspector)

	s.blockPositionHolder = blockposition.
		NewRedisPositionHolder(
			univ3Inspector.Name(),
			conf.General.StartBlockNumber,
			database.NewRedis(conf))
	if !s.blockPositionHolder.Exists() {
		fmt.Println("block position does not exist. creating a new one.")
		err = s.blockPositionHolder.Create()
		if err != nil {
			panic(err)
		}
	}
	s.currentBlockPosition, err = s.blockPositionHolder.Read()
	if err != nil {
		panic(err)
	}
	fmt.Printf("starting from block: %d\n", s.currentBlockPosition)
}

func (s *Syncer) Sync() {
	currentBlockPosition, err := s.blockPositionHolder.Read()
	if err != nil {
		s.logger.Error("error reading block position")
		return
	}
	s.currentBlockPosition = currentBlockPosition
	newHeader := <-s.newHeaders

	for s.currentBlockPosition < newHeader.Number.Uint64()-s.lag {
		block, err := s.rpc.BlockByNumber(context.Background(), big.NewInt(int64(s.currentBlockPosition)))
		if err != nil {
			s.logger.Error("error getting block", zap.Error(err))
			break
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
		err = s.blockPositionHolder.Incr()
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
