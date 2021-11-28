package inspector

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Inspector interface {
	Name() string
	InspectBlock(block *types.Block) error
}

type FilterFunc func(blockNumber *big.Int) ([]types.Log, error)

func GetFilterFunc(ws *ethclient.Client, addresses []common.Address) FilterFunc {
	return func(blockNumber *big.Int) ([]types.Log, error) {
		logs, err := ws.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: blockNumber,
			ToBlock:   blockNumber,
			Addresses: addresses,
		})
		if err != nil {
			return nil, err
		}
		return logs, nil
	}
}
