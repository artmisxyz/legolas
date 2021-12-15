package uniswapv3

import (
	"fmt"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
	"github.com/ethereum/go-ethereum/core/types"
)

type State interface {
	IncreaseLiquidity(event *nftpositionmanager.NftpositionmanagerIncreaseLiquidity, log types.Log) error
	DecreaseLiquidity(event *nftpositionmanager.NftpositionmanagerDecreaseLiquidity, log types.Log) error
	Transfer(event *nftpositionmanager.NftpositionmanagerTransfer, log types.Log) error
	Collect(event *nftpositionmanager.NftpositionmanagerCollect, log types.Log) error
}

type MemoryState struct {
}

func NewMemoryState() State {
	return &MemoryState{}
}

func (m MemoryState) IncreaseLiquidity(event *nftpositionmanager.NftpositionmanagerIncreaseLiquidity, log types.Log) error {
	fmt.Println("increase liquidity state save")
	return nil
}

func (m MemoryState) DecreaseLiquidity(event *nftpositionmanager.NftpositionmanagerDecreaseLiquidity, log types.Log) error {
	fmt.Println("decrease liquidity state save")
	return nil
}

func (m MemoryState) Transfer(event *nftpositionmanager.NftpositionmanagerTransfer, log types.Log) error {
	fmt.Println("transfer liquidity state save")
	return nil
}

func (m MemoryState) Collect(event *nftpositionmanager.NftpositionmanagerCollect, log types.Log) error {
	fmt.Println("collect liquidity state save")

	return nil
}
