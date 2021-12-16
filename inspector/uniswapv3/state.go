package uniswapv3

import (
	"context"
	"fmt"
	"github.com/artmisxyz/blockinspector/ent"
	"github.com/artmisxyz/blockinspector/ent/schema"
	"github.com/artmisxyz/blockinspector/pkg/hashing"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
)

type State interface {
	CreateIncreaseLiquidity(event *nftpositionmanager.NftpositionmanagerIncreaseLiquidity) error
	CreateDecreaseLiquidity(event *nftpositionmanager.NftpositionmanagerDecreaseLiquidity) error
	CreateTransfer(event *nftpositionmanager.NftpositionmanagerTransfer) error
	CreateCollect(event *nftpositionmanager.NftpositionmanagerCollect) error
}

type Postgres struct {
	db ent.Client
}

func NewMemoryState(db ent.Client) State {
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
		AddEvent(instance).
		SetLiquidity(&schema.BigInt{*event.Liquidity}).
		SetTokenID(&schema.BigInt{*event.TokenId}).
		SetAmount0(&schema.BigInt{*event.Amount0}).
		SetAmount1(&schema.BigInt{*event.Amount1}).
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
		AddEvent(instance).
		SetLiquidity(&schema.BigInt{*event.Liquidity}).
		SetTokenID(&schema.BigInt{*event.TokenId}).
		SetAmount0(&schema.BigInt{*event.Amount0}).
		SetAmount1(&schema.BigInt{*event.Amount1}).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreateTransfer(event *nftpositionmanager.NftpositionmanagerTransfer) error {
	fmt.Println("transfer liquidity state save")
	return nil
}

func (m *Postgres) CreateCollect(event *nftpositionmanager.NftpositionmanagerCollect) error {
	fmt.Println("collect liquidity state save")
	return nil
}
