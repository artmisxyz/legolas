package uniswapv3

import (
	"fmt"
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type decreaseLiquidityEventHandler struct {
	binding *nftpositionmanager.Nftpositionmanager
	state   State
}

func (d *decreaseLiquidityEventHandler) Handle(log types.Log) error {
	event, err := d.binding.ParseDecreaseLiquidity(log)
	if err != nil {
		return fmt.Errorf("error parsing decrease liquidty. %w", err)
	}
	return d.state.DecreaseLiquidity(event, log)
}

func (d *decreaseLiquidityEventHandler) Signature() string {
	return "0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4"
}

func NewDecreaseLiquidityEventHandler(address common.Address, backend bind.ContractBackend) inspector.EventHandler {
	binding, err := nftpositionmanager.NewNftpositionmanager(address, backend)
	if err != nil {
		panic(err)
	}
	return &decreaseLiquidityEventHandler{
		binding: binding,
		state:   NewMemoryState(),
	}
}
