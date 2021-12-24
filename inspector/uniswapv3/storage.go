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
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

type Postgres struct {
	db *ent.Client
}

func NewPostgres(db *ent.Client) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (m *Postgres) CreateEvent(t time.Time, name, signature string, log types.Log) (*domain.Event, error) {
	e, err := m.db.Event.Create().
		SetTime(t).
		SetAddress(log.Address.String()).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(name).
		SetSignature(signature).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return entToDomainEvent(e), nil
}

func (m *Postgres) CreateIncreaseLiquidity(eventId int, il *nftpositionmanager.NftpositionmanagerIncreaseLiquidity) error {
	_, err := m.db.UniswapV3IncreaseLiqudity.Create().
		SetLiquidity(il.Liquidity.String()).
		SetTokenID(il.TokenId.String()).
		SetAmount0(il.Amount0.String()).
		SetAmount1(il.Amount1.String()).
		SetEventID(eventId).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreateDecreaseLiquidity(eventId int, dl *nftpositionmanager.NftpositionmanagerDecreaseLiquidity) error {
	_, err := m.db.UniswapV3DecreaseLiqudity.Create().
		SetLiquidity(dl.Liquidity.String()).
		SetTokenID(dl.TokenId.String()).
		SetAmount0(dl.Amount0.String()).
		SetAmount1(dl.Amount1.String()).
		SetEventID(eventId).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreateTransfer(eventId int, t *nftpositionmanager.NftpositionmanagerTransfer) error {
	_, err := m.db.UniswapV3Transfer.Create().
		SetTokenID(t.TokenId.String()).
		SetFrom(t.From.String()).
		SetTo(t.To.String()).
		SetEventID(eventId).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreateCollect(eventId int, c *nftpositionmanager.NftpositionmanagerCollect) error {
	_, err := m.db.UniswapV3Collect.Create().
		SetTokenID(c.TokenId.String()).
		SetAmount0(c.Amount0.String()).
		SetAmount1(c.Amount1.String()).
		SetRecipient(c.Recipient.String()).
		SetEventID(eventId).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolCreated(eventId int, pc *factory.FactoryPoolCreated) error {
	_, err := m.db.UniswapV3PoolCreated.Create().
		SetPool(pc.Pool.String()).
		SetToken0(pc.Token0.String()).
		SetToken1(pc.Token1.String()).
		SetFee(pc.Fee.String()).
		SetTickSpacing(pc.TickSpacing.String()).
		SetEventID(eventId).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolInitialize(eventId int, pi *pool.PoolInitialize) error {
	_, err := m.db.UniswapV3PoolInitialize.Create().
		SetTick(pi.Tick.String()).
		SetSqrtPriceX96(pi.SqrtPriceX96.String()).
		SetEventID(eventId).
		Save(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolSwap(eventId int, ps *pool.PoolSwap) error {
	_, err := m.db.UniswapV3PoolSwap.Create().
		SetRecipient(ps.Recipient.String()).
		SetSender(ps.Sender.String()).
		SetSqrtPriceX96(ps.SqrtPriceX96.String()).
		SetLiquidity(ps.Liquidity.String()).
		SetAmount0(ps.Amount0.String()).
		SetAmount1(ps.Amount1.String()).
		SetTick(ps.Tick.String()).
		SetEventID(eventId).
		Save(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolBurn(eventId int, pb *pool.PoolBurn) error {
	_, err := m.db.UniswapV3PoolBurn.Create().
		SetOwner(pb.Owner.String()).
		SetTickLower(pb.TickLower.String()).
		SetTickUpper(pb.TickUpper.String()).
		SetAmount(pb.Amount.String()).
		SetAmount0(pb.Amount0.String()).
		SetAmount1(pb.Amount1.String()).
		SetEventID(eventId).
		Save(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolFlash(eventId int, pf *pool.PoolFlash) error {
	_, err := m.db.UniswapV3PoolFlash.Create().
		SetRecipient(pf.Recipient.String()).
		SetSender(pf.Sender.String()).
		SetAmount0(pf.Amount0.String()).
		SetAmount1(pf.Amount1.String()).
		SetPaid0(pf.Paid0.String()).
		SetPaid1(pf.Paid1.String()).
		SetEventID(eventId).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgres) CreatePoolMint(eventId int, pm *pool.PoolMint) error {
	_, err := m.db.UniswapV3PoolMint.Create().
		SetOwner(pm.Owner.String()).
		SetTickLower(pm.TickLower.String()).
		SetTickUpper(pm.TickUpper.String()).
		SetAmount(pm.Amount.String()).
		SetAmount0(pm.Amount0.String()).
		SetAmount1(pm.Amount1.String()).
		SetEventID(eventId).
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
		WithEvent().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.Swap
	for _, s := range swaps {
		if err != nil {
			return nil, err
		}
		swap := entToDomainSwap(s)
		swap.Event = entToDomainEvent(s.Edges.Event)
		ret = append(ret, swap)
	}
	return ret, nil
}

func entToDomainEvent(event *ent.Event) *domain.Event {
	return &domain.Event{
		ID:          event.ID,
		Time:        event.Time,
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
