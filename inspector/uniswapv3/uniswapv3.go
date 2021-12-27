package uniswapv3

import (
	"context"
	"github.com/artmisxyz/legolas/database"
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/inspector"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"time"
)

type uniswapV3 struct {
	logger        *zap.Logger
	ws            *ethclient.Client
	eventHandlers map[string]inspector.EventHandler
	addresses     []common.Address
	storage       *database.Storage
}

const Name = "uniswapV3:inspector"

const (
	Factory                    = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	SwapRouter                 = "0xE592427A0AEce92De3Edee1F18E0157C05861564"
	NonfungiblePositionManager = "0xC36442b4a4522E871399CD717aBDD847Ab11FE88"
	V3Migrator                 = "0xA5644E29708357803b5A882D272c41cC0dF92B34"
	ETHUSDC                    = "0x8ad599c3A0ff1De082011EFDDc58f1908eb6e6D8"
)

func NewUniswapV3(logger *zap.Logger, ws *ethclient.Client, db *ent.Client) inspector.Inspector {
	v := &uniswapV3{
		logger:        logger.Named(Name),
		ws:            ws,
		eventHandlers: make(map[string]inspector.EventHandler),
		storage:       database.NewPostgresStorage(db),
	}

	v.registerAddress(common.HexToAddress(Factory))
	v.registerAddress(common.HexToAddress(NonfungiblePositionManager))
	v.registerAddress(common.HexToAddress(ETHUSDC))

	v.registerEventHandlers(
		NewIncreaseLiquidityEventHandler(common.HexToAddress(NonfungiblePositionManager), ws, db),
		NewDecreaseLiquidityEventHandler(common.HexToAddress(NonfungiblePositionManager), ws, db),
		NewCollectEventHandler(common.HexToAddress(NonfungiblePositionManager), ws, db),
		NewTransferEventHandler(common.HexToAddress(NonfungiblePositionManager), ws, db),
		NewPoolCreatedEventHandler(common.HexToAddress(Factory), ws, db),
		NewInitializeEventHandler(common.HexToAddress(ETHUSDC), ws, db),
		NewSwapEventHandler(common.HexToAddress(ETHUSDC), ws, db),
		NewMintEventHandler(common.HexToAddress(ETHUSDC), ws, db),
		NewBurnEventHandler(common.HexToAddress(ETHUSDC), ws, db),
		NewFlashEventHandler(common.HexToAddress(ETHUSDC), ws, db),
	)

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
	t := time.Unix(int64(block.Time()), 0)

	for _, log := range logs {
		eventHandler, ok := v.eventHandlers[log.Topics[0].String()]
		if !ok {
			v.logger.
				Info("event handler not registered",
					zap.Uint64("block_number", log.BlockNumber),
					zap.String("tx_hash", log.TxHash.String()),
					zap.Uint("event_index", log.Index))
			continue
		}
		id, err := v.storage.CreateEvent(t, eventHandler.Name(), eventHandler.Signature(), log)
		if err != nil {
			if ent.IsConstraintError(err) {
				v.logger.
					Warn("failed to create event",
						zap.Error(err),
						zap.Uint64("block_number", log.BlockNumber),
						zap.String("tx_hash", log.TxHash.String()),
						zap.Uint("event_index", log.Index))
				continue
			}
			return err
		}
		err = eventHandler.ParseAndSavePayload(id, log)
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
