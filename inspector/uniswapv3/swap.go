package uniswapv3

import (
	"github.com/artmisxyz/blockinspector/ent"
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/artmisxyz/uniswap-go/pool"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	PoolSwapEventName      = "UniswapV3_Pool_Swap"
	PoolSwapEventSignature = "0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"
)

type swapEventHandler struct {
	binding *pool.Pool
	state   Storage
}

func NewSwapEventHandler(address common.Address, backend bind.ContractBackend, db *ent.Client) inspector.EventHandler {
	binding, err := pool.NewPool(address, backend)
	if err != nil {
		panic(err)
	}
	return &swapEventHandler{
		binding: binding,
		state:   NewPostgres(db),
	}
}

func (s *swapEventHandler) Signature() string {
	return PoolSwapEventSignature
}

func (s *swapEventHandler) Save(log types.Log) error {
	event, err := s.binding.ParseSwap(log)
	if err != nil {
		panic(err)
	}
	return s.state.CreatePoolSwap(event)
}
