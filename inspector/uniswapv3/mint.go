package uniswapv3

import (
	"github.com/artmisxyz/legolas/database"
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/inspector"
	"github.com/artmisxyz/uniswap-go/pool"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	PoolMintEventName      = "UniswapV3_Pool_Mint"
	PoolMintEventSignature = "0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde"
)

type mintEventHandler struct {
	binding *pool.Pool
	state   *database.Storage
}

func (m *mintEventHandler) Name() string {
	return PoolMintEventName
}

func (m *mintEventHandler) Signature() string {
	return PoolMintEventSignature
}

func (m *mintEventHandler) ParseAndSavePayload(eventID int, log types.Log) error {
	event, err := m.binding.ParseMint(log)
	if err != nil {
		return err
	}
	return m.state.CreatePoolMint(eventID, event)
}

func NewMintEventHandler(address common.Address, backend bind.ContractBackend, db *ent.Client) inspector.EventHandler {
	binding, err := pool.NewPool(address, backend)
	if err != nil {
		panic(err)
	}
	return &mintEventHandler{
		binding: binding,
		state:   database.NewPostgresStorage(db),
	}
}
