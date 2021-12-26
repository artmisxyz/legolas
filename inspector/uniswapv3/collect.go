package uniswapv3

import (
	"fmt"
	"github.com/artmisxyz/legolas/database"
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/inspector"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type collectEventHandler struct {
	binding *nftpositionmanager.Nftpositionmanager
	state   *database.Storage
}

const (
	CollectEventName      = "UniswapV3_Collect"
	CollectEventSignature = "0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01"
)

func (c *collectEventHandler) Name() string {
	return CollectEventName
}

func (c *collectEventHandler) Signature() string {
	return CollectEventSignature
}

func (c *collectEventHandler) ParseAndSavePayload(eventID int, log types.Log) error {
	collect, err := c.binding.ParseCollect(log)
	if err != nil {
		return fmt.Errorf("error parsing collect. %w", err)
	}
	return c.state.CreateCollect(eventID, collect)
}

func NewCollectEventHandler(address common.Address, backend bind.ContractBackend, db *ent.Client) inspector.EventHandler {
	binding, err := nftpositionmanager.NewNftpositionmanager(address, backend)
	if err != nil {
		panic(err)
	}
	return &collectEventHandler{
		binding: binding,
		state:   database.NewPostgresStorage(db),
	}
}
