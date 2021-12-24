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

type transferEventHandler struct {
	binding *nftpositionmanager.Nftpositionmanager
	state   *Postgres
}

const (
	TransferEventName      = "UniswapV3_Transfer"
	TransferEventSignature = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

func (t *transferEventHandler) Name() string {
	return TransferEventName
}

func (t *transferEventHandler) Signature() string {
	return TransferEventSignature
}

func (t *transferEventHandler) ParseAndSavePayload(eventID int, log types.Log) error {
	event, err := t.binding.ParseTransfer(log)
	if err != nil {
		return fmt.Errorf("error parsing transfer. %w", err)
	}
	return t.state.CreateTransfer(eventID, event)
}

func NewTransferEventHandler(address common.Address, backend bind.ContractBackend, db *ent.Client) inspector.EventHandler {
	binding, err := nftpositionmanager.NewNftpositionmanager(address, backend)
	if err != nil {
		panic(err)
	}
	return &transferEventHandler{
		binding: binding,
		state:   NewPostgres(db),
	}
}
