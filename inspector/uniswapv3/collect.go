package uniswapv3

import (
	"fmt"
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/inspector"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type collectEventHandler struct {
	binding *nftpositionmanager.Nftpositionmanager
	state   *Postgres
}

const (
	CollectEventName      = "UniswapV3_Collect"
	CollectEventSignature = "0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01"
)

func NewCollectEventHandler(address common.Address, backend bind.ContractBackend, db *ent.Client) inspector.EventHandler {
	binding, err := nftpositionmanager.NewNftpositionmanager(address, backend)
	if err != nil {
		panic(err)
	}
	return &collectEventHandler{
		binding: binding,
		state:   NewPostgres(db),
	}
}

func (c *collectEventHandler) Save(log types.Log) error {
	event, err := c.binding.ParseCollect(log)
	if err != nil {
		return fmt.Errorf("error parsing collect. %w", err)
	}
	return c.state.CreateCollect(event)
}

func (c *collectEventHandler) Signature() string {
	return "0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01"
}
