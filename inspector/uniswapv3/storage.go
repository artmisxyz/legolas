package uniswapv3

import (
	"context"
	"github.com/artmisxyz/legolas/domain"
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/ent/event"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolswap"
	"github.com/artmisxyz/uniswap-go/factory"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
	"github.com/artmisxyz/uniswap-go/pool"
)

type Postgres struct {
	db *ent.Client
}

func NewPostgres(db *ent.Client) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (m *Postgres) CreateIncreaseLiquidity(event *nftpositionmanager.NftpositionmanagerIncreaseLiquidity) error {
	log := event.Raw
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
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
		SetLiquidity(event.Liquidity.String()).
		SetTokenID(event.TokenId.String()).
		SetAmount0(event.Amount0.String()).
		SetAmount1(event.Amount1.String()).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreateDecreaseLiquidity(event *nftpositionmanager.NftpositionmanagerDecreaseLiquidity) error {
	log := event.Raw
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
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
		SetLiquidity(event.Liquidity.String()).
		SetTokenID(event.TokenId.String()).
		SetAmount0(event.Amount0.String()).
		SetAmount1(event.Amount1.String()).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreateTransfer(event *nftpositionmanager.NftpositionmanagerTransfer) error {
	log := event.Raw
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
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
		SetTokenID(event.TokenId.String()).
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
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
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
		SetTokenID(event.TokenId.String()).
		SetAmount0(event.Amount0.String()).
		SetAmount1(event.Amount1.String()).
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
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
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
		SetFee(event.Fee.String()).
		SetTickSpacing(event.TickSpacing.String()).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolInitialize(event *pool.PoolInitialize) error {
	log := event.Raw
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
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
		SetTick(event.Tick.String()).
		SetSqrtPriceX96(event.SqrtPriceX96.String()).
		SetEvent(instance).
		Save(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolSwap(event *pool.PoolSwap) error {
	log := event.Raw
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
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
		SetSqrtPriceX96(event.SqrtPriceX96.String()).
		SetLiquidity(event.Liquidity.String()).
		SetAmount0(event.Amount0.String()).
		SetAmount1(event.Amount1.String()).
		SetTick(event.Tick.String()).
		SetEvent(instance).Save(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolBurn(event *pool.PoolBurn) error {
	log := event.Raw
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
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
		SetTickLower(event.TickLower.String()).
		SetTickUpper(event.TickUpper.String()).
		SetAmount(event.Amount.String()).
		SetAmount0(event.Amount0.String()).
		SetAmount1(event.Amount1.String()).
		SetEvent(instance).
		Save(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolFlash(event *pool.PoolFlash) error {
	log := event.Raw
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
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
		SetAmount0(event.Amount0.String()).
		SetAmount1(event.Amount1.String()).
		SetPaid0(event.Paid0.String()).
		SetPaid1(event.Paid1.String()).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolMint(event *pool.PoolMint) error {
	log := event.Raw
	instance, err := m.db.Event.Create().
		SetAddress(log.Address.String()).
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
		SetTickLower(event.TickLower.String()).
		SetTickUpper(event.TickUpper.String()).
		SetAmount(event.Amount.String()).
		SetAmount0(event.Amount0.String()).
		SetAmount1(event.Amount1.String()).
		SetEvent(instance).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) GetEventsByBlockNumber(from, to uint64, cursor, limit int) ([]*domain.Event, error) {
	events, err := m.db.Event.
		Query().
		Where(event.BlockNumberGTE(from), event.BlockNumberLTE(to), event.IDGT(cursor)).
		Limit(limit).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.Event
	for _, e := range events {
		ret = append(ret, entToDomainEvent(e))
	}
	return ret, nil
}

func (m *Postgres) GetEvents(cursor, limit int) ([]*domain.Event, error) {
	events, err := m.db.Event.
		Query().
		Where(event.IDGT(cursor)).
		Limit(limit).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.Event
	for _, e := range events {
		ret = append(ret, entToDomainEvent(e))
	}
	return ret, nil
}

func (m *Postgres) GetSwaps(cursor, limit int) ([]*domain.Swap, error) {
	swaps, err := m.db.UniswapV3PoolSwap.
		Query().
		Where(uniswapv3poolswap.IDGT(cursor)).
		Limit(limit).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.Swap
	for _, s := range swaps {
		e, err := s.QueryEvent().Only(context.Background())
		if err != nil {
			return nil, err
		}
		swap := entToDomainSwap(s)
		swap.Event = entToDomainEvent(e)
		ret = append(ret, swap)
	}
	return ret, nil
}

func entToDomainEvent(event *ent.Event) *domain.Event {
	return &domain.Event{
		Name:        event.Name,
		Signature:   event.Signature,
		Address:     event.Address,
		BlockNumber: event.BlockNumber,
		TxHash:      event.TxHash,
		TxIndex:     event.TxIndex,
		BlockHash:   event.BlockHash,
		Index:       event.Index,
	}
}

func entToDomainSwap(swap *ent.UniswapV3PoolSwap) *domain.Swap {
	return &domain.Swap{
		Sender:       swap.Sender,
		Recipient:    swap.Recipient,
		Amount0:      swap.Amount0,
		Amount1:      swap.Amount1,
		SqrtPriceX96: swap.SqrtPriceX96,
		Liquidity:    swap.Liquidity,
		Tick:         swap.Tick,
	}
}
