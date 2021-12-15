package uniswapv3

import (
	"context"
	"fmt"
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
	UniswapV3Factory           = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	SwapRouter                 = "0xE592427A0AEce92De3Edee1F18E0157C05861564"
	NonfungiblePositionManager = "0xC36442b4a4522E871399CD717aBDD847Ab11FE88"
	V3Migrator                 = "0xA5644E29708357803b5A882D272c41cC0dF92B34"
)

func NewUniswapV3(logger *zap.Logger, ws *ethclient.Client) inspector.Inspector {
	v := &uniswapV3{
		logger:        logger.Named(Name),
		ws:            ws,
		eventHandlers: make(map[string]inspector.EventHandler),
	}

	//v.registerEventHandlers(UniswapV3Factory,
	//	NewPoolCreatedEventHandler())

	v.registerAddress(common.HexToAddress(NonfungiblePositionManager))

	v.registerEventHandlers(
		NewIncreaseLiquidityEventHandler(common.HexToAddress(NonfungiblePositionManager), ws),
		NewDecreaseLiquidityEventHandler(common.HexToAddress(NonfungiblePositionManager), ws),
		NewCollectEventHandler(common.HexToAddress(NonfungiblePositionManager), ws),
		NewTransferEventHandler(common.HexToAddress(NonfungiblePositionManager), ws))

	//v.registerEventHandlers(SwapRouter, "SwapRouter")
	//
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
			return fmt.Errorf("event handler not registered")
		}
		err := eventHandler.Handle(log)
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
