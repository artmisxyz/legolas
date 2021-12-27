package database

import (
	"context"
	"fmt"
	"github.com/artmisxyz/legolas/configs"
	"github.com/artmisxyz/legolas/domain"
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/ent/event"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolswap"
	"github.com/artmisxyz/uniswap-go/factory"
	"github.com/artmisxyz/uniswap-go/nftpositionmanager"
	"github.com/artmisxyz/uniswap-go/pool"
	"github.com/ethereum/go-ethereum/core/types"
	_ "github.com/lib/pq"
	"time"
)

type Storage struct {
	db *ent.Client
}

const dataSourceSchema = "host=%s port=%d user=%s dbname=%s password=%s sslmode=%s"

func NewPostgresClient(config configs.Config) (*ent.Client, error) {
	dn := fmt.Sprintf(dataSourceSchema,
		config.Database.Host, config.Database.Port, config.Database.User,
		config.Database.DatabaseName, config.Database.Password, config.Database.SSLMode)

	c, err := ent.Open(config.Database.Driver, dn)
	return c, err
}

func NewPostgresStorage(db *ent.Client) *Storage {
	return &Storage{
		db: db,
	}
}

func (m *Storage) CreateEvent(t time.Time, name, signature string, log types.Log) (int, error) {
	id, err := m.db.Event.Create().
		SetTime(t).
		SetAddress(log.Address.String()).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(name).
		SetSignature(signature).
		OnConflictColumns(event.FieldTxHash, event.FieldTxIndex).
		ID(context.Background())

	if err != nil {
		return -1, err
	}
	return id, nil
}

func (m *Storage) CreateIncreaseLiquidity(eventId int, il *nftpositionmanager.NftpositionmanagerIncreaseLiquidity) error {
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

func (m *Storage) CreateDecreaseLiquidity(eventId int, dl *nftpositionmanager.NftpositionmanagerDecreaseLiquidity) error {
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

func (m *Storage) CreateTransfer(eventId int, t *nftpositionmanager.NftpositionmanagerTransfer) error {
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

func (m *Storage) CreateCollect(eventId int, c *nftpositionmanager.NftpositionmanagerCollect) error {
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

func (m *Storage) CreatePoolCreated(eventId int, pc *factory.FactoryPoolCreated) error {
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

func (m *Storage) CreatePoolInitialize(eventId int, pi *pool.PoolInitialize) error {
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

func (m *Storage) CreatePoolSwap(eventId int, ps *pool.PoolSwap) error {
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

func (m *Storage) CreatePoolBurn(eventId int, pb *pool.PoolBurn) error {
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

func (m *Storage) CreatePoolFlash(eventId int, pf *pool.PoolFlash) error {
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

func (m *Storage) CreatePoolMint(eventId int, pm *pool.PoolMint) error {
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

func (m *Storage) GetEventsByBlockNumber(from, to uint64, cursor, limit int) ([]*domain.Event, error) {
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

func (m *Storage) GetEvents(cursor, limit int) ([]*domain.Event, error) {
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

func (m *Storage) GetSwaps(cursor, limit int) ([]*domain.Swap, error) {
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
