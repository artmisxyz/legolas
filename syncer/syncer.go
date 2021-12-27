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
	id   int
	Head chan *types.Header
	ws   *ethclient.Client
	rpc  *ethclient.Client

	logger *zap.Logger

	blockPositionHolder blockposition.BlockPositionHolder

	lag int

	inspectors map[string]inspector.Inspector
}

func (s *Syncer) Init(id int, startBlock, finishBlock uint64, conf configs.Config) {
	s.id = id
	s.lag = conf.General.BlockLag

	l := logger.New(conf)
	s.logger = l.Named(fmt.Sprintf("syncer:%d", s.id))

	s.ws = connections.RPC(conf.Node.Websocket)
	s.rpc = connections.Websocket(conf.Node.RPC)

	head, err := connections.CurrentBlockChan(s.logger, s.ws)
	if err != nil {
		panic(err)
	}
	s.Head = head

	db, err := database.NewPostgresClient(conf)
	if err != nil {
		panic(err)
	}

	s.inspectors = make(map[string]inspector.Inspector)
	univ3Inspector := uniswapv3.NewUniswapV3(s.logger, s.ws, db)
	s.registerInspectors(univ3Inspector)

	s.blockPositionHolder = blockposition.
		NewRedisPositionHolder(
			univ3Inspector.Name()+fmt.Sprintf(":%d", id),
			startBlock,
			finishBlock,
			database.NewRedis(conf))

	if !s.blockPositionHolder.Exists() {
		fmt.Printf("block position for syncer %d does not exist. creating a new one.", id)
		err = s.blockPositionHolder.Create()
		if err != nil {
			panic(err)
		}
	}
}

func (s *Syncer) Sync() {
	startBlock, err := s.blockPositionHolder.Read(blockposition.Start)
	if err != nil {
		s.logger.Error("error reading block position")
		return
	}
	finishBlock, err := s.blockPositionHolder.Read(blockposition.Finish)
	if err != nil {
		s.logger.Error("error reading block position")
		return
	}
	fmt.Printf("syncer %d from %d to %d\n", s.id, startBlock, finishBlock)
	lag := uint64(0)
	if finishBlock == 0 {
		head := <-s.Head
		finishBlock = head.Number.Uint64()
		lag = uint64(s.lag)
	}

	for i := startBlock; i < finishBlock-lag; i++ {
		block, err := s.rpc.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			s.logger.Error("error getting block", zap.Uint64("block_number", i), zap.Error(err))
			break
		}
		s.logger.
			Info("processing block",
				zap.Uint64("current_block", i),
				zap.Float64("progress", float64(i)/float64(finishBlock)))

		err = s.Inspect(block)
		if err != nil {
			s.logger.Error("error inspecting block", zap.Uint64("block_number", i), zap.Error(err))
			break
		}
		err = s.blockPositionHolder.Update(i)
		if err != nil {
			s.logger.Error("error updating block position", zap.Error(err))
			break
		}
	}
}

func (s *Syncer) Start() {
	go func() {
		for {
			s.Sync()
		}
	}()
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
