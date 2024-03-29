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
	PoolBurnEventName      = "UniswapV3_Pool_Burn"
	PoolBurnEventSignature = "0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c"
)

type burnEventHandler struct {
	binding *pool.Pool
	state   *database.Storage
}

func (b *burnEventHandler) Name() string {
	return PoolBurnEventName
}

func (b *burnEventHandler) Signature() string {
	return PoolBurnEventSignature
}

func (b *burnEventHandler) ParseAndSavePayload(eventID int, log types.Log) error {
	burn, err := b.binding.ParseBurn(log)
	if err != nil {
		return err
	}
	return b.state.CreatePoolBurn(eventID, burn)
}

func NewBurnEventHandler(address common.Address, backend bind.ContractBackend, db *ent.Client) inspector.EventHandler {
	binding, err := pool.NewPool(address, backend)
	if err != nil {
		panic(err)
	}
	return &burnEventHandler{
		state:   database.NewPostgresStorage(db),
		binding: binding,
	}
}
