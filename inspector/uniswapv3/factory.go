package uniswapv3

import (
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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

type PoolCreatedEvent struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Pool        common.Address
}

type poolCreatedEventHandler struct {
	factoryState *FactoryState
}

func NewPoolCreatedEventHandler() inspector.EventHandler {
	return &poolCreatedEventHandler{
		factoryState: &FactoryState{},
	}
}

func (p *poolCreatedEventHandler) Signature() string {
	return "0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118"
}

func (p *poolCreatedEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
	var v PoolCreatedEvent
	err := abi.UnpackIntoInterface(&v, "PoolCreated", data)
	if err != nil {
		return err
	}
	v.Token0 = common.BytesToAddress(topics[1].Bytes())
	v.Token1 = common.BytesToAddress(topics[2].Bytes())
	v.Fee = topics[3].Big()
	p.factoryState.poolCount++
	return nil
}
