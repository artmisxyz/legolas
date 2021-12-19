package uniswapv3

import (
	"context"
	"github.com/artmisxyz/blockinspector/ent"
	"github.com/artmisxyz/blockinspector/ent/schema"
	"github.com/artmisxyz/blockinspector/pkg/hashing"
	"github.com/artmisxyz/uniswap-go/factory"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
	"github.com/artmisxyz/uniswap-go/pool"
)

type Storage interface {
	CreateIncreaseLiquidity(event *nftpositionmanager.NftpositionmanagerIncreaseLiquidity) error
	CreateDecreaseLiquidity(event *nftpositionmanager.NftpositionmanagerDecreaseLiquidity) error
	CreateTransfer(event *nftpositionmanager.NftpositionmanagerTransfer) error
	CreateCollect(event *nftpositionmanager.NftpositionmanagerCollect) error
	CreatePoolCreated(event *factory.FactoryPoolCreated) error
	CreatePoolInitialize(event *pool.PoolInitialize) error
	CreatePoolSwap(event *pool.PoolSwap) error
	CreatePoolBurn(event *pool.PoolBurn) error
	CreatePoolFlash(event *pool.PoolFlash) error
	CreatePoolMint(event *pool.PoolMint) error
}

type Postgres struct {
	db *ent.Client
}

func NewPostgres(db *ent.Client) Storage {
	return &Postgres{
		db: db,
	}
}

func (m *Postgres) CreateIncreaseLiquidity(event *nftpositionmanager.NftpositionmanagerIncreaseLiquidity) error {
	log := event.Raw
	hash := hashing.LogHash(log)
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
		SetHash(hash).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(IncreaseLiquidityEventName).
		SetSignature(IncreaseLiquidityEventSignature).
		Save(context.Background())
	if err != nil {
		return err
	}

	_, err = m.db.UniswapV3IncreaseLiqudity.Create().
		SetLiquidity(&schema.BigInt{Int: *event.Liquidity}).
		SetTokenID(&schema.BigInt{Int: *event.TokenId}).
		SetAmount0(&schema.BigInt{Int: *event.Amount0}).
		SetAmount1(&schema.BigInt{Int: *event.Amount1}).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreateDecreaseLiquidity(event *nftpositionmanager.NftpositionmanagerDecreaseLiquidity) error {
	log := event.Raw
	hash := hashing.LogHash(log)
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
		SetHash(hash).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(DecreaseLiquidityEventName).
		SetSignature(DecreaseLiquidityEventSignature).
		Save(context.Background())
	if err != nil {
		return err
	}

	_, err = m.db.UniswapV3DecreaseLiqudity.Create().
		SetLiquidity(&schema.BigInt{Int: *event.Liquidity}).
		SetTokenID(&schema.BigInt{Int: *event.TokenId}).
		SetAmount0(&schema.BigInt{Int: *event.Amount0}).
		SetAmount1(&schema.BigInt{Int: *event.Amount1}).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreateTransfer(event *nftpositionmanager.NftpositionmanagerTransfer) error {
	log := event.Raw
	hash := hashing.LogHash(log)
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
		SetHash(hash).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(TransferEventName).
		SetSignature(TransferEventSignature).
		Save(context.Background())
	if err != nil {
		return err
	}

	_, err = m.db.UniswapV3Transfer.Create().
		SetTokenID(&schema.BigInt{Int: *event.TokenId}).
		SetFrom(event.From.String()).
		SetTo(event.To.String()).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreateCollect(event *nftpositionmanager.NftpositionmanagerCollect) error {
	log := event.Raw
	hash := hashing.LogHash(log)
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
		SetHash(hash).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(CollectEventName).
		SetSignature(CollectEventSignature).
		Save(context.Background())
	if err != nil {
		return err
	}

	_, err = m.db.UniswapV3Collect.Create().
		SetTokenID(&schema.BigInt{Int: *event.TokenId}).
		SetAmount0(&schema.BigInt{Int: *event.Amount0}).
		SetAmount1(&schema.BigInt{Int: *event.Amount1}).
		SetRecipient(event.Recipient.String()).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolCreated(event *factory.FactoryPoolCreated) error {
	log := event.Raw
	hash := hashing.LogHash(log)
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
		SetHash(hash).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(PoolCreatedEventName).
		SetSignature(PoolCreatedEventSignature).
		Save(context.Background())
	if err != nil {
		return err
	}

	_, err = m.db.UniswapV3PoolCreated.Create().
		SetPool(event.Pool.String()).
		SetToken0(event.Token0.String()).
		SetToken1(event.Token1.String()).
		SetFee(&schema.BigInt{Int: *event.Fee}).
		SetTickSpacing(&schema.BigInt{Int: *event.TickSpacing}).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolInitialize(event *pool.PoolInitialize) error {
	log := event.Raw
	hash := hashing.LogHash(log)
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
		SetHash(hash).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(PoolInitializeEventName).
		SetSignature(PoolInitializeEventSignature).
		Save(context.Background())
	if err != nil {
		return err
	}

	_, err = m.db.UniswapV3PoolInitialize.Create().
		SetTick(&schema.BigInt{Int: *event.Tick}).
		SetSqrtPriceX96(&schema.BigInt{Int: *event.SqrtPriceX96}).
		SetEvent(instance).
		Save(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolSwap(event *pool.PoolSwap) error {
	log := event.Raw
	hash := hashing.LogHash(log)
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
		SetHash(hash).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(PoolSwapEventName).
		SetSignature(PoolSwapEventSignature).
		Save(context.Background())
	if err != nil {
		return err
	}

	_, err = m.db.UniswapV3PoolSwap.Create().
		SetRecipient(event.Recipient.String()).
		SetSender(event.Sender.String()).
		SetSqrtPriceX96(&schema.BigInt{Int: *event.SqrtPriceX96}).
		SetLiquidity(&schema.BigInt{Int: *event.Liquidity}).
		SetAmount0(&schema.BigInt{Int: *event.Amount0}).
		SetAmount1(&schema.BigInt{Int: *event.Amount1}).
		SetTick(&schema.BigInt{Int: *event.Tick}).
		SetEvent(instance).Save(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolBurn(event *pool.PoolBurn) error {
	log := event.Raw
	hash := hashing.LogHash(log)
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
		SetHash(hash).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(PoolBurnEventName).
		SetSignature(PoolBurnEventSignature).
		Save(context.Background())
	if err != nil {
		return err
	}

	_, err = m.db.UniswapV3PoolBurn.Create().
		SetOwner(event.Owner.String()).
		SetTickLower(&schema.BigInt{Int: *event.TickLower}).
		SetTickUpper(&schema.BigInt{Int: *event.TickUpper}).
		SetAmount(&schema.BigInt{Int: *event.Amount}).
		SetAmount0(&schema.BigInt{Int: *event.Amount0}).
		SetAmount1(&schema.BigInt{Int: *event.Amount1}).
		SetEvent(instance).
		Save(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolFlash(event *pool.PoolFlash) error {
	log := event.Raw
	hash := hashing.LogHash(log)
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
		SetHash(hash).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(PoolFlashEventName).
		SetSignature(PoolFlashEventSignature).
		Save(context.Background())
	if err != nil {
		return err
	}

	_, err = m.db.UniswapV3PoolFlash.Create().
		SetRecipient(event.Recipient.String()).
		SetSender(event.Sender.String()).
		SetAmount0(&schema.BigInt{Int: *event.Amount0}).
		SetAmount1(&schema.BigInt{Int: *event.Amount1}).
		SetPaid0(&schema.BigInt{Int: *event.Paid0}).
		SetPaid1(&schema.BigInt{Int: *event.Paid1}).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolMint(event *pool.PoolMint) error {
	log := event.Raw
	hash := hashing.LogHash(log)
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
		SetHash(hash).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(PoolMintEventName).
		SetSignature(PoolMintEventSignature).
		Save(context.Background())
	if err != nil {
		return err
	}

	_, err = m.db.UniswapV3PoolMint.Create().
		SetOwner(event.Owner.String()).
		SetTickLower(&schema.BigInt{Int: *event.TickLower}).
		SetTickUpper(&schema.BigInt{Int: *event.TickUpper}).
		SetAmount(&schema.BigInt{Int: *event.Amount}).
		SetAmount0(&schema.BigInt{Int: *event.Amount0}).
		SetAmount1(&schema.BigInt{Int: *event.Amount1}).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}
