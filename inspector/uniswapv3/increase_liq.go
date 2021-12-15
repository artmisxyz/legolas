package uniswapv3

import (
	"fmt"
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type increaseLiquidityEventHandler struct {
	binding *nftpositionmanager.Nftpositionmanager
	state   State
}

func NewIncreaseLiquidityEventHandler(address common.Address, backend bind.ContractBackend) inspector.EventHandler {
	binding, err := nftpositionmanager.NewNftpositionmanager(address, backend)
	if err != nil {
		panic(err)
	}
	return &increaseLiquidityEventHandler{
		binding: binding,
		state:   NewMemoryState(),
	}
}

func (i *increaseLiquidityEventHandler) Signature() string {
	return "0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f"
}

func (i *increaseLiquidityEventHandler) Handle(log types.Log) error {
	event, err := i.binding.ParseIncreaseLiquidity(log)
	if err != nil {
		return fmt.Errorf("error parsing increase liquidity. %w", err)
	}
	return i.state.IncreaseLiquidity(event, log)
}
