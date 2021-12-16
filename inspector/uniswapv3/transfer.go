package uniswapv3

import (
	"fmt"
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type transferEventHandler struct {
	binding *nftpositionmanager.Nftpositionmanager
	state   State
}

func NewTransferEventHandler(address common.Address, backend bind.ContractBackend) inspector.EventHandler {
	binding, err := nftpositionmanager.NewNftpositionmanager(address, backend)
	if err != nil {
		panic(err)
	}
	return &transferEventHandler{
		binding: binding,
		state:   NewMemoryState(),
	}
}

func (t *transferEventHandler) Save(log types.Log) error {
	event, err := t.binding.ParseTransfer(log)
	if err != nil {
		return fmt.Errorf("error parsing transfer. %w", err)
	}
	return t.state.CreateTransfer(event)
}

func (t *transferEventHandler) Signature() string {
	return "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
}
