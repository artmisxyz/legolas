package database

import (
	"context"
	"fmt"
	"github.com/artmisxyz/legolas/configs"
	"github.com/artmisxyz/legolas/domain"
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/ent/event"
	"github.com/artmisxyz/legolas/ent/uniswapv3collect"
	"github.com/artmisxyz/legolas/ent/uniswapv3decreaseliqudity"
	"github.com/artmisxyz/legolas/ent/uniswapv3increaseliqudity"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolburn"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolcreated"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolflash"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolinitialize"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolmint"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolswap"
	"github.com/artmisxyz/legolas/ent/uniswapv3transfer"
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
	onConflict := m.db.Event.Create().
		SetTime(t).
		SetAddress(log.Address.String()).
		SetBlockHash(log.BlockHash.String()).
		SetIndex(log.Index).
		SetBlockNumber(log.BlockNumber).
		SetTxIndex(log.TxIndex).
		SetTxHash(log.TxHash.String()).
		SetName(name).
		SetSignature(signature).
		OnConflictColumns(event.FieldIndex, event.FieldTxHash)

	err := onConflict.UpdateNewValues().Exec(context.Background())
	if err != nil {
		return -1, err
	}
	id, err := onConflict.ID(context.Background())
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (m *Storage) CreateIncreaseLiquidity(eventId int, il *nftpositionmanager.NftpositionmanagerIncreaseLiquidity) error {
	err := m.db.UniswapV3IncreaseLiqudity.Create().
		SetLiquidity(il.Liquidity.String()).
		SetTokenID(il.TokenId.String()).
		SetAmount0(il.Amount0.String()).
		SetAmount1(il.Amount1.String()).
		SetEventID(eventId).
		OnConflictColumns(uniswapv3increaseliqudity.EventColumn).
		UpdateNewValues().
		Exec(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Storage) CreateDecreaseLiquidity(eventId int, dl *nftpositionmanager.NftpositionmanagerDecreaseLiquidity) error {
	err := m.db.UniswapV3DecreaseLiqudity.Create().
		SetLiquidity(dl.Liquidity.String()).
		SetTokenID(dl.TokenId.String()).
		SetAmount0(dl.Amount0.String()).
		SetAmount1(dl.Amount1.String()).
		SetEventID(eventId).
		OnConflictColumns(uniswapv3decreaseliqudity.EventColumn).
		UpdateNewValues().
		Exec(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Storage) CreateTransfer(eventId int, t *nftpositionmanager.NftpositionmanagerTransfer) error {
	err := m.db.UniswapV3Transfer.Create().
		SetTokenID(t.TokenId.String()).
		SetFrom(t.From.String()).
		SetTo(t.To.String()).
		SetEventID(eventId).
		OnConflictColumns(uniswapv3transfer.EventColumn).
		UpdateNewValues().
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Storage) CreateCollect(eventId int, c *nftpositionmanager.NftpositionmanagerCollect) error {
	err := m.db.UniswapV3Collect.Create().
		SetTokenID(c.TokenId.String()).
		SetAmount0(c.Amount0.String()).
		SetAmount1(c.Amount1.String()).
		SetRecipient(c.Recipient.String()).
		SetEventID(eventId).
		OnConflictColumns(uniswapv3collect.EventColumn).
		UpdateNewValues().
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Storage) CreatePoolCreated(eventId int, pc *factory.FactoryPoolCreated) error {
	err := m.db.UniswapV3PoolCreated.Create().
		SetPool(pc.Pool.String()).
		SetToken0(pc.Token0.String()).
		SetToken1(pc.Token1.String()).
		SetFee(pc.Fee.String()).
		SetTickSpacing(pc.TickSpacing.String()).
		SetEventID(eventId).
		OnConflictColumns(uniswapv3poolcreated.EventColumn).
		UpdateNewValues().
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Storage) CreatePoolInitialize(eventId int, pi *pool.PoolInitialize) error {
	err := m.db.UniswapV3PoolInitialize.Create().
		SetTick(pi.Tick.String()).
		SetSqrtPriceX96(pi.SqrtPriceX96.String()).
		SetEventID(eventId).
		OnConflictColumns(uniswapv3poolinitialize.EventColumn).
		UpdateNewValues().
		Exec(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Storage) CreatePoolSwap(eventId int, ps *pool.PoolSwap) error {
	err := m.db.UniswapV3PoolSwap.Create().
		SetRecipient(ps.Recipient.String()).
		SetSender(ps.Sender.String()).
		SetSqrtPriceX96(ps.SqrtPriceX96.String()).
		SetLiquidity(ps.Liquidity.String()).
		SetAmount0(ps.Amount0.String()).
		SetAmount1(ps.Amount1.String()).
		SetTick(ps.Tick.String()).
		SetEventID(eventId).
		OnConflictColumns(uniswapv3poolswap.EventColumn).
		UpdateNewValues().
		Exec(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Storage) CreatePoolBurn(eventId int, pb *pool.PoolBurn) error {
	err := m.db.UniswapV3PoolBurn.Create().
		SetOwner(pb.Owner.String()).
		SetTickLower(pb.TickLower.String()).
		SetTickUpper(pb.TickUpper.String()).
		SetAmount(pb.Amount.String()).
		SetAmount0(pb.Amount0.String()).
		SetAmount1(pb.Amount1.String()).
		SetEventID(eventId).
		OnConflictColumns(uniswapv3poolburn.EventColumn).
		UpdateNewValues().
		Exec(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Storage) CreatePoolFlash(eventId int, pf *pool.PoolFlash) error {
	err := m.db.UniswapV3PoolFlash.Create().
		SetRecipient(pf.Recipient.String()).
		SetSender(pf.Sender.String()).
		SetAmount0(pf.Amount0.String()).
		SetAmount1(pf.Amount1.String()).
		SetPaid0(pf.Paid0.String()).
		SetPaid1(pf.Paid1.String()).
		SetEventID(eventId).
		OnConflictColumns(uniswapv3poolflash.EventColumn).
		UpdateNewValues().
		Exec(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Storage) CreatePoolMint(eventId int, pm *pool.PoolMint) error {
	err := m.db.UniswapV3PoolMint.Create().
		SetSender(pm.Sender.String()).
		SetOwner(pm.Owner.String()).
		SetTickLower(pm.TickLower.String()).
		SetTickUpper(pm.TickUpper.String()).
		SetAmount(pm.Amount.String()).
		SetAmount0(pm.Amount0.String()).
		SetAmount1(pm.Amount1.String()).
		SetEventID(eventId).
		OnConflictColumns(uniswapv3poolmint.EventColumn).
		UpdateNewValues().
		Exec(context.Background())

	if err != nil {
		return err
	}
	return nil
}

func (m *Storage) GetEventsByBlockNumber(from, to uint64, cursor, limit int) ([]*domain.Event, error) {
	events, err := m.db.Event.
		Query().
		Where(event.BlockNumberGTE(from), event.BlockNumberLTE(to), event.IDGT(cursor)).
		Order(ent.Desc(event.FieldBlockNumber)).
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
		Where(event.IDLT(cursor)).
		Order(ent.Desc(event.FieldID)).
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
		Order(ent.Asc(uniswapv3poolswap.FieldID)).
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

func (m *Storage) GetBurns(cursor, limit int) ([]*domain.Burn, error) {
	burns, err := m.db.UniswapV3PoolBurn.
		Query().
		Where(uniswapv3poolburn.IDGT(cursor)).
		Order(ent.Asc(uniswapv3poolburn.FieldID)).
		Limit(limit).
		WithEvent().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.Burn
	for _, b := range burns {
		if err != nil {
			return nil, err
		}
		burn := entToDomainBurn(b)
		burn.Event = entToDomainEvent(b.Edges.Event)
		ret = append(ret, burn)
	}
	return ret, nil
}

func (m *Storage) GetCollects(cursor, limit int) ([]*domain.Collect, error) {
	collects, err := m.db.UniswapV3Collect.
		Query().
		Where(uniswapv3collect.IDGT(cursor)).
		Order(ent.Asc(uniswapv3collect.FieldID)).
		Limit(limit).
		WithEvent().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.Collect
	for _, c := range collects {
		if err != nil {
			return nil, err
		}
		collect := entToDomainCollect(c)
		collect.Event = entToDomainEvent(c.Edges.Event)
		ret = append(ret, collect)
	}
	return ret, nil
}

func (m *Storage) GetDecreaseLiquidity(cursor, limit int) ([]*domain.DecreaseLiquidity, error) {
	dls, err := m.db.UniswapV3DecreaseLiqudity.
		Query().
		Where(uniswapv3decreaseliqudity.IDGT(cursor)).
		Order(ent.Asc(uniswapv3decreaseliqudity.FieldID)).
		Limit(limit).
		WithEvent().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.DecreaseLiquidity
	for _, dl := range dls {
		if err != nil {
			return nil, err
		}
		decreaseLiquidity := entToDomainDecreaseLiquidity(dl)
		decreaseLiquidity.Event = entToDomainEvent(dl.Edges.Event)
		ret = append(ret, decreaseLiquidity)
	}
	return ret, nil
}

func (m *Storage) GetIncreaseLiquidity(cursor, limit int) ([]*domain.IncreaseLiquidity, error) {
	ils, err := m.db.UniswapV3IncreaseLiqudity.
		Query().
		Where(uniswapv3increaseliqudity.IDGT(cursor)).
		Order(ent.Asc(uniswapv3increaseliqudity.FieldID)).
		Limit(limit).
		WithEvent().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.IncreaseLiquidity
	for _, il := range ils {
		if err != nil {
			return nil, err
		}
		increaseLiquidity := entToDomainIncreaseLiquidity(il)
		increaseLiquidity.Event = entToDomainEvent(il.Edges.Event)
		ret = append(ret, increaseLiquidity)
	}
	return ret, nil
}

func (m *Storage) GetPoolCreated(cursor, limit int) ([]*domain.PoolCreated, error) {
	pcs, err := m.db.UniswapV3PoolCreated.
		Query().
		Where(uniswapv3poolcreated.IDGT(cursor)).
		Order(ent.Asc(uniswapv3poolcreated.FieldID)).
		Limit(limit).
		WithEvent().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.PoolCreated
	for _, pc := range pcs {
		if err != nil {
			return nil, err
		}
		poolCreated := entToDomainPoolCreated(pc)
		poolCreated.Event = entToDomainEvent(pc.Edges.Event)
		ret = append(ret, poolCreated)
	}
	return ret, nil
}

func (m *Storage) GetFlash(cursor, limit int) ([]*domain.Flash, error) {
	flashes, err := m.db.UniswapV3PoolFlash.
		Query().
		Where(uniswapv3poolflash.IDGT(cursor)).
		Order(ent.Asc(uniswapv3poolflash.FieldID)).
		Limit(limit).
		WithEvent().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.Flash
	for _, f := range flashes {
		if err != nil {
			return nil, err
		}
		flash := entToDomainFlash(f)
		flash.Event = entToDomainEvent(f.Edges.Event)
		ret = append(ret, flash)
	}
	return ret, nil
}

func (m *Storage) GetPoolInitialize(cursor, limit int) ([]*domain.PoolInitialize, error) {
	pis, err := m.db.UniswapV3PoolInitialize.
		Query().
		Where(uniswapv3poolinitialize.IDGT(cursor)).
		Order(ent.Asc(uniswapv3poolinitialize.FieldID)).
		Limit(limit).
		WithEvent().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.PoolInitialize
	for _, pi := range pis {
		if err != nil {
			return nil, err
		}
		poolInitialize := entToDomainInitialize(pi)
		poolInitialize.Event = entToDomainEvent(pi.Edges.Event)
		ret = append(ret, poolInitialize)
	}
	return ret, nil
}

func (m *Storage) GetMints(cursor, limit int) ([]*domain.Mint, error) {
	mints, err := m.db.UniswapV3PoolMint.
		Query().
		Where(uniswapv3poolmint.IDGT(cursor)).
		Order(ent.Asc(uniswapv3poolmint.FieldID)).
		Limit(limit).
		WithEvent().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.Mint
	for _, m := range mints {
		if err != nil {
			return nil, err
		}
		mint := entToDomainMint(m)
		mint.Event = entToDomainEvent(m.Edges.Event)
		ret = append(ret, mint)
	}
	return ret, nil
}

func (m *Storage) GetTransfers(cursor, limit int) ([]*domain.Transfer, error) {
	transfers, err := m.db.UniswapV3Transfer.
		Query().
		Where(uniswapv3transfer.IDGT(cursor)).
		Order(ent.Asc(uniswapv3transfer.FieldID)).
		Limit(limit).
		WithEvent().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	var ret []*domain.Transfer
	for _, t := range transfers {
		if err != nil {
			return nil, err
		}
		mint := entToDomainTransfer(t)
		mint.Event = entToDomainEvent(t.Edges.Event)
		ret = append(ret, mint)
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

func entToDomainBurn(burn *ent.UniswapV3PoolBurn) *domain.Burn {
	return &domain.Burn{
		Owner:     burn.Owner,
		TickLower: burn.TickLower,
		TickUpper: burn.TickUpper,
		Amount:    burn.Amount,
		Amount0:   burn.Amount0,
		Amount1:   burn.Amount1,
	}
}

func entToDomainCollect(collect *ent.UniswapV3Collect) *domain.Collect {
	return &domain.Collect{
		TokenId:   collect.TokenID,
		Recipient: collect.Recipient,
		Amount0:   collect.Amount0,
		Amount1:   collect.Amount1,
	}
}

func entToDomainDecreaseLiquidity(dl *ent.UniswapV3DecreaseLiqudity) *domain.DecreaseLiquidity {
	return &domain.DecreaseLiquidity{
		TokenId:   dl.TokenID,
		Liquidity: dl.Liquidity,
		Amount0:   dl.Amount0,
		Amount1:   dl.Amount1,
	}
}

func entToDomainPoolCreated(pc *ent.UniswapV3PoolCreated) *domain.PoolCreated {
	return &domain.PoolCreated{
		Token0:      pc.Token0,
		Token1:      pc.Token1,
		Fee:         pc.Fee,
		TickSpacing: pc.TickSpacing,
		Pool:        pc.Pool,
	}
}

func entToDomainFlash(f *ent.UniswapV3PoolFlash) *domain.Flash {
	return &domain.Flash{
		Sender:    f.Sender,
		Recipient: f.Recipient,
		Amount0:   f.Amount0,
		Amount1:   f.Amount1,
		Paid0:     f.Paid0,
		Paid1:     f.Paid1,
	}
}

func entToDomainIncreaseLiquidity(il *ent.UniswapV3IncreaseLiqudity) *domain.IncreaseLiquidity {
	return &domain.IncreaseLiquidity{
		TokenId:   il.TokenID,
		Liquidity: il.Liquidity,
		Amount0:   il.Amount0,
		Amount1:   il.Amount1,
	}
}

func entToDomainInitialize(pi *ent.UniswapV3PoolInitialize) *domain.PoolInitialize {
	return &domain.PoolInitialize{
		SqrtPriceX96: pi.SqrtPriceX96,
		Tick:         pi.Tick,
	}
}

func entToDomainMint(mint *ent.UniswapV3PoolMint) *domain.Mint {
	return &domain.Mint{
		Sender:    mint.Sender,
		Owner:     mint.Owner,
		TickLower: mint.TickLower,
		TickUpper: mint.TickUpper,
		Amount:    mint.Amount,
		Amount0:   mint.Amount0,
		Amount1:   mint.Amount1,
	}
}

func entToDomainTransfer(transfer *ent.UniswapV3Transfer) *domain.Transfer {
	return &domain.Transfer{
		From:    transfer.From,
		To:      transfer.To,
		TokenId: transfer.TokenID,
	}
}
