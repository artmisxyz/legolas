package uniswapv3

import (
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/inspector"
	"github.com/artmisxyz/uniswap-go/pool"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	PoolInitializeEventName      = "UniswapV3_Pool_Initialize"
	PoolInitializeEventSignature = "0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95"
)

type initializeEventHandler struct {
	binding *pool.Pool
	state   Storage
}

func NewInitializeEventHandler(address common.Address, backend bind.ContractBackend, db *ent.Client) inspector.EventHandler {
	binding, err := pool.NewPool(address, backend)
	if err != nil {
		panic(err)
	}
	return &initializeEventHandler{
		binding: binding,
		state:   NewPostgres(db),
	}
}

func (i *initializeEventHandler) Signature() string {
	return PoolInitializeEventSignature
}

func (i *initializeEventHandler) Save(log types.Log) error {
	event, err := i.binding.ParseInitialize(log)
	if err != nil {
		return err
	}
	return i.state.CreatePoolInitialize(event)
}
