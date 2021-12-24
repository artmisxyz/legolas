package uniswapv3

import (
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/inspector"
	"github.com/artmisxyz/uniswap-go/factory"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type poolCreatedEventHandler struct {
	binding *factory.Factory
	storage *Postgres
}

const (
	PoolCreatedEventName      = "UniswapV3_PoolCreated"
	PoolCreatedEventSignature = "0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118"
)

func NewPoolCreatedEventHandler(address common.Address, backend bind.ContractBackend, db *ent.Client) inspector.EventHandler {
	binding, err := factory.NewFactory(address, backend)
	if err != nil {
		panic(err)
	}
	return &poolCreatedEventHandler{
		storage: NewPostgres(db),
		binding: binding,
	}
}

func (p *poolCreatedEventHandler) Save(log types.Log) error {
	event, err := p.binding.ParsePoolCreated(log)
	if err != nil {
		return err
	}
	return p.storage.CreatePoolCreated(event)
}

func (p *poolCreatedEventHandler) Signature() string {
	return PoolCreatedEventSignature
}
