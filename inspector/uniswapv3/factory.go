package uniswapv3

import (
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/ethereum/go-ethereum/core/types"
)

type FactoryState struct {
	poolCount                    int
	totalVolumeETH               int
	totalVolumeUSD               int
	untrackedVolumeUSD           int
	totalFeesUSD                 int
	totalFeesETH                 int
	totalValueLockedETH          int
	totalValueLockedUSD          int
	totalValueLockedUSDUntracked int
	totalValueLockedETHUntracked int
	txCount                      int
}

type poolCreatedEventHandler struct {
	factoryState *FactoryState
}

func NewPoolCreatedEventHandler() inspector.EventHandler {
	return &poolCreatedEventHandler{
		factoryState: &FactoryState{},
	}
}

func (p *poolCreatedEventHandler) Handle(event types.Log) error {
	panic("implement me")
}

func (p *poolCreatedEventHandler) Signature() string {
	// crypto.Keccak256Hash([]byte("PoolCreated(unit256,int64)")).String()
	return "0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118"
}
