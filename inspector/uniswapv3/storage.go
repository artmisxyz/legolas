package uniswapv3

import (
	"context"
	"github.com/artmisxyz/blockinspector/ent"
	"github.com/artmisxyz/blockinspector/ent/schema"
	"github.com/artmisxyz/blockinspector/pkg/hashing"
	"github.com/artmisxyz/uniswap-go/factory"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
)

type Storage interface {
	CreateIncreaseLiquidity(event *nftpositionmanager.NftpositionmanagerIncreaseLiquidity) error
	CreateDecreaseLiquidity(event *nftpositionmanager.NftpositionmanagerDecreaseLiquidity) error
	CreateTransfer(event *nftpositionmanager.NftpositionmanagerTransfer) error
	CreateCollect(event *nftpositionmanager.NftpositionmanagerCollect) error
	CreatePoolCreated(event *factory.FactoryPoolCreated) error
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
		SetLiquidity(&schema.BigInt{*event.Liquidity}).
		SetTokenID(&schema.BigInt{*event.TokenId}).
		SetAmount0(&schema.BigInt{*event.Amount0}).
		SetAmount1(&schema.BigInt{*event.Amount1}).
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
		SetLiquidity(&schema.BigInt{*event.Liquidity}).
		SetTokenID(&schema.BigInt{*event.TokenId}).
		SetAmount0(&schema.BigInt{*event.Amount0}).
		SetAmount1(&schema.BigInt{*event.Amount1}).
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
		SetTokenID(&schema.BigInt{*event.TokenId}).
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
		SetTokenID(&schema.BigInt{*event.TokenId}).
		SetAmount0(&schema.BigInt{*event.Amount0}).
		SetAmount1(&schema.BigInt{*event.Amount1}).
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
		SetFee(&schema.BigInt{*event.Fee}).
		SetTickSpacing(&schema.BigInt{*event.TickSpacing}).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}