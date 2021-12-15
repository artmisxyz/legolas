package uniswapv3

import (
	"fmt"
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type collectEventHandler struct {
	binding *nftpositionmanager.Nftpositionmanager
	state   State
}

func NewCollectEventHandler(address common.Address, backend bind.ContractBackend) inspector.EventHandler {
	binding, err := nftpositionmanager.NewNftpositionmanager(address, backend)
	if err != nil {
		panic(err)
	}
	return &collectEventHandler{
		binding: binding,
		state:   NewMemoryState(),
	}
}

func (c *collectEventHandler) Handle(log types.Log) error {
	event, err := c.binding.ParseCollect(log)
	if err != nil {
		return fmt.Errorf("error parsing collect. %w", err)
	}
	return c.state.Collect(event, log)
}

func (c *collectEventHandler) Signature() string {
	return "0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01"
}
