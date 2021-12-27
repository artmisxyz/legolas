package main

import (
	"context"
	"fmt"
	"github.com/artmisxyz/legolas/configs"
	"github.com/artmisxyz/legolas/database"
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/ent/syncer"
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
	name string
	Head chan *types.Header
	ws   *ethclient.Client
	rpc  *ethclient.Client

	logger *zap.Logger

	db *ent.Client

	lag int

	inspectors map[string]inspector.Inspector

	current uint64
	finish  uint64
}

func (s *Syncer) Init(id int, startBlock, finishBlock uint64, conf configs.Config) {
	s.lag = conf.General.BlockLag

	l := logger.New(conf)
	s.logger = l.Named(fmt.Sprintf("syncer:%d", id))

	s.ws = Connect(conf.Node.Websocket)
	s.rpc = Connect(conf.Node.RPC)

	head, err := CurrentBlockChan(s.logger, s.ws)
	if err != nil {
		panic(err)
	}
	s.Head = head

	db, err := database.NewPostgresClient(conf)
	if err != nil {
		panic(err)
	}
	s.db = db

	s.inspectors = make(map[string]inspector.Inspector)
	univ3Inspector := uniswapv3.NewUniswapV3(s.logger, s.ws, db)
	s.registerInspectors(univ3Inspector)

	s.name = univ3Inspector.Name() + fmt.Sprintf(":%d", id)
	state, err := db.Syncer.Query().Where(syncer.NameEQ(s.name)).Only(context.Background())
	if err != nil {
		if !ent.IsNotFound(err) {
			panic(err)
		}
		state = db.Syncer.
			Create().
			SetName(s.name).
			SetStart(startBlock).
			SetCurrent(startBlock).
			SetFinish(finishBlock).
			SaveX(context.Background())
	}

	s.current = state.Current
	s.finish = state.Finish
	fmt.Printf("%s from %d to %d\n", s.name, s.current, s.finish)
}

func (s *Syncer) Sync() {
	lag := uint64(0)
	if s.finish == 0 {
		head := <-s.Head
		s.finish = head.Number.Uint64()
		lag = uint64(s.lag)
	}

	for s.current < s.finish-lag {
		block, err := s.rpc.BlockByNumber(context.Background(), big.NewInt(int64(s.current)))
		if err != nil {
			s.logger.Error("error getting block", zap.Uint64("block_number", s.current), zap.Error(err))
			break
		}
		s.logger.
			Info("processing block",
				zap.Uint64("current_block", s.current),
				zap.Float64("progress", float64(s.finish-s.current)/float64(s.finish)))

		err = s.Inspect(block)
		if err != nil {
			s.logger.Error("error inspecting block", zap.Uint64("block_number", s.current), zap.Error(err))
			break
		}
		_, err = s.db.Syncer.Update().Where(syncer.NameEQ(s.name)).SetCurrent(s.current).Save(context.Background())
		if err != nil {
			s.logger.Error("error updating syncer current block position")
			break
		}
		s.current++
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
