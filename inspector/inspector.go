package inspector

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Inspector interface {
	Name() string
	InspectBlock(block *types.Block) error
}

type EventHandler interface {
	Signature() string
	Handle(log types.Log) error
}
