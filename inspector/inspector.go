package inspector

import "github.com/ethereum/go-ethereum/core/types"

type Inspector interface {
	InspectTransaction(tx *types.Transaction) error
	InspectBlock(block *types.Block) error
	Filter(block *types.Block) []*types.Transaction
	Name() string
}
