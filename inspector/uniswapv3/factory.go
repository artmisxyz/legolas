package uniswapv3

import (
	"fmt"
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	Fee         int
	TickSpacing int
	Pool        common.Address
}

type poolCreatedEventHandler struct {
	factoryState *FactoryState
}

func NewPoolCreatedEventHandler() inspector.EventHandler {
	return &poolCreatedEventHandler{}
}

func (p *poolCreatedEventHandler) Signature() string {
	return crypto.Keccak256Hash([]byte("createPool(address,address,uint24)")).String()
}

func (p *poolCreatedEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) {
	var v PoolCreatedEvent
	err := abi.UnpackIntoInterface(&v, "PoolCreated", data)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("token0", v.Token0)
	fmt.Println("token1", v.Token1)
	p.factoryState.poolCount++
}
