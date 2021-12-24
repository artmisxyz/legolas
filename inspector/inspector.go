package inspector

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Inspector interface {
	Name() string
	InspectBlock(block *types.Block) error
}

type EventHandler interface {
	Name() string
	Signature() string
	ParseAndSavePayload(eventID int, log types.Log) error
}
