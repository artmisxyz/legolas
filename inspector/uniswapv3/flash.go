package uniswapv3

//
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
	PoolFlashEventName      = "UniswapV3_Pool_Flash"
	PoolFlashEventSignature = "0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633"
)

type flashEventHandler struct {
	binding *pool.Pool
	state   *database.Storage
}

func (f *flashEventHandler) Name() string {
	return PoolFlashEventName
}

func (f *flashEventHandler) Signature() string {
	return PoolFlashEventSignature
}

func (f *flashEventHandler) ParseAndSavePayload(eventID int, log types.Log) error {
	pf, err := f.binding.ParseFlash(log)
	if err != nil {
		return err
	}
	return f.state.CreatePoolFlash(eventID, pf)
}

func NewFlashEventHandler(address common.Address, backend bind.ContractBackend, db *ent.Client) inspector.EventHandler {
	binding, err := pool.NewPool(address, backend)
	if err != nil {
		panic(err)
	}
	return &flashEventHandler{
		binding: binding,
		state:   database.NewPostgresStorage(db),
	}
}
