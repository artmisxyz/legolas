package uniswapv3

import (
	"context"
	"github.com/artmisxyz/blockinspector/ent"
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type uniswapV3 struct {
	logger        *zap.Logger
	ws            *ethclient.Client
	eventHandlers map[string]inspector.EventHandler
	addresses     []common.Address
}

const Name = "uniswapV3:inspector"

const (
	Factory    = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	SwapRouter = "0xE592427A0AEce92De3Edee1F18E0157C05861564"
	NonfungiblePositionManager = "0xC36442b4a4522E871399CD717aBDD847Ab11FE88"
	V3Migrator                 = "0xA5644E29708357803b5A882D272c41cC0dF92B34"
)

func NewUniswapV3(logger *zap.Logger, ws *ethclient.Client, db *ent.Client) inspector.Inspector {
	v := &uniswapV3{
		logger:        logger.Named(Name),
		ws:            ws,
		eventHandlers: make(map[string]inspector.EventHandler),
	}

	v.registerAddress(common.HexToAddress(Factory))
	v.registerAddress(common.HexToAddress(NonfungiblePositionManager))

	v.registerEventHandlers(
		NewIncreaseLiquidityEventHandler(common.HexToAddress(NonfungiblePositionManager), ws, db),
		NewDecreaseLiquidityEventHandler(common.HexToAddress(NonfungiblePositionManager), ws, db),
		NewCollectEventHandler(common.HexToAddress(NonfungiblePositionManager), ws, db),
		NewTransferEventHandler(common.HexToAddress(NonfungiblePositionManager), ws, db),
		NewPoolCreatedEventHandler(common.HexToAddress(Factory), ws, db))

	//v.registerEventHandlers(SwapRouter, "SwapRouter")
	//v.registerEventHandlers(V3Migrator, "V3Migrator")
	return v
}

func (v *uniswapV3) Name() string {
	return Name
}

func (v *uniswapV3) InspectBlock(block *types.Block) error {
	logs, err := v.ws.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: block.Number(),
		ToBlock:   block.Number(),
		Addresses: v.addresses,
	})
	if err != nil {
		return err
	}

	for _, log := range logs {
		eventHandler, ok := v.eventHandlers[log.Topics[0].String()]
		if !ok {
			v.logger.
				Warn("event handler not registered",
					zap.Uint64("block_number", log.BlockNumber),
					zap.String("tx_hash", log.TxHash.String()),
					zap.Uint("event_index", log.Index))
			continue
		}
		err := eventHandler.Save(log)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *uniswapV3) registerEventHandlers(eventHandlers ...inspector.EventHandler) {
	for _, eventHandler := range eventHandlers {
		v.eventHandlers[eventHandler.Signature()] = eventHandler
	}
}

func (v *uniswapV3) registerAddress(addr common.Address) {
	v.addresses = append(v.addresses, addr)
}
