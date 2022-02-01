// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package orderbookv1beta1

import (
	context "context"
	ormdb "github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	ormlist "github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ormtable "github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	ormerrors "github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	v1beta1 "github.com/regen-ledger/regen-network/api/regen/ecocredit/marketplace/v1beta1"
)

type BuyOrderSellOrderMatchStore interface {
	Insert(ctx context.Context, buyOrderSellOrderMatch *BuyOrderSellOrderMatch) error
	Update(ctx context.Context, buyOrderSellOrderMatch *BuyOrderSellOrderMatch) error
	Save(ctx context.Context, buyOrderSellOrderMatch *BuyOrderSellOrderMatch) error
	Delete(ctx context.Context, buyOrderSellOrderMatch *BuyOrderSellOrderMatch) error
	Has(ctx context.Context, buy_order_id uint64, sell_order_id uint64) (found bool, err error)
	Get(ctx context.Context, buy_order_id uint64, sell_order_id uint64) (*BuyOrderSellOrderMatch, error)
	List(ctx context.Context, prefixKey BuyOrderSellOrderMatchIndexKey, opts ...ormlist.Option) (BuyOrderSellOrderMatchIterator, error)
	ListRange(ctx context.Context, from, to BuyOrderSellOrderMatchIndexKey, opts ...ormlist.Option) (BuyOrderSellOrderMatchIterator, error)

	doNotImplement()
}

type BuyOrderSellOrderMatchIterator struct {
	ormtable.Iterator
}

func (i BuyOrderSellOrderMatchIterator) Value() (*BuyOrderSellOrderMatch, error) {
	var buyOrderSellOrderMatch BuyOrderSellOrderMatch
	err := i.UnmarshalMessage(&buyOrderSellOrderMatch)
	return &buyOrderSellOrderMatch, err
}

type BuyOrderSellOrderMatchIndexKey interface {
	id() uint32
	values() []interface{}
	buyOrderSellOrderMatchIndexKey()
}

// primary key starting index..
type BuyOrderSellOrderMatchBuyOrderIdSellOrderIdIndexKey struct {
	vs []interface{}
}

func (x BuyOrderSellOrderMatchBuyOrderIdSellOrderIdIndexKey) id() uint32                      { return 1 }
func (x BuyOrderSellOrderMatchBuyOrderIdSellOrderIdIndexKey) values() []interface{}           { return x.vs }
func (x BuyOrderSellOrderMatchBuyOrderIdSellOrderIdIndexKey) buyOrderSellOrderMatchIndexKey() {}

func (this BuyOrderSellOrderMatchBuyOrderIdSellOrderIdIndexKey) WithBuyOrderId(buy_order_id uint64) BuyOrderSellOrderMatchBuyOrderIdSellOrderIdIndexKey {
	this.vs = []interface{}{buy_order_id}
	return this
}

func (this BuyOrderSellOrderMatchBuyOrderIdSellOrderIdIndexKey) WithBuyOrderIdSellOrderId(buy_order_id uint64, sell_order_id uint64) BuyOrderSellOrderMatchBuyOrderIdSellOrderIdIndexKey {
	this.vs = []interface{}{buy_order_id, sell_order_id}
	return this
}

type BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey struct {
	vs []interface{}
}

func (x BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey) id() uint32 {
	return 1
}
func (x BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey) values() []interface{} {
	return x.vs
}
func (x BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey) buyOrderSellOrderMatchIndexKey() {
}

func (this BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey) WithBidDenomId(bid_denom_id uint32) BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey {
	this.vs = []interface{}{bid_denom_id}
	return this
}

func (this BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey) WithBidDenomIdBidPriceComplement(bid_denom_id uint32, bid_price_complement uint64) BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey {
	this.vs = []interface{}{bid_denom_id, bid_price_complement}
	return this
}

func (this BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey) WithBidDenomIdBidPriceComplementBuyOrderId(bid_denom_id uint32, bid_price_complement uint64, buy_order_id uint64) BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey {
	this.vs = []interface{}{bid_denom_id, bid_price_complement, buy_order_id}
	return this
}

func (this BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey) WithBidDenomIdBidPriceComplementBuyOrderIdAskPrice(bid_denom_id uint32, bid_price_complement uint64, buy_order_id uint64, ask_price uint64) BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey {
	this.vs = []interface{}{bid_denom_id, bid_price_complement, buy_order_id, ask_price}
	return this
}

func (this BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey) WithBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderId(bid_denom_id uint32, bid_price_complement uint64, buy_order_id uint64, ask_price uint64, sell_order_id uint64) BuyOrderSellOrderMatchBidDenomIdBidPriceComplementBuyOrderIdAskPriceSellOrderIdIndexKey {
	this.vs = []interface{}{bid_denom_id, bid_price_complement, buy_order_id, ask_price, sell_order_id}
	return this
}

type BuyOrderSellOrderMatchSellOrderIdIndexKey struct {
	vs []interface{}
}

func (x BuyOrderSellOrderMatchSellOrderIdIndexKey) id() uint32                      { return 2 }
func (x BuyOrderSellOrderMatchSellOrderIdIndexKey) values() []interface{}           { return x.vs }
func (x BuyOrderSellOrderMatchSellOrderIdIndexKey) buyOrderSellOrderMatchIndexKey() {}

func (this BuyOrderSellOrderMatchSellOrderIdIndexKey) WithSellOrderId(sell_order_id uint64) BuyOrderSellOrderMatchSellOrderIdIndexKey {
	this.vs = []interface{}{sell_order_id}
	return this
}

type buyOrderSellOrderMatchStore struct {
	table ormtable.Table
}

func (this buyOrderSellOrderMatchStore) Insert(ctx context.Context, buyOrderSellOrderMatch *BuyOrderSellOrderMatch) error {
	return this.table.Insert(ctx, buyOrderSellOrderMatch)
}

func (this buyOrderSellOrderMatchStore) Update(ctx context.Context, buyOrderSellOrderMatch *BuyOrderSellOrderMatch) error {
	return this.table.Update(ctx, buyOrderSellOrderMatch)
}

func (this buyOrderSellOrderMatchStore) Save(ctx context.Context, buyOrderSellOrderMatch *BuyOrderSellOrderMatch) error {
	return this.table.Save(ctx, buyOrderSellOrderMatch)
}

func (this buyOrderSellOrderMatchStore) Delete(ctx context.Context, buyOrderSellOrderMatch *BuyOrderSellOrderMatch) error {
	return this.table.Delete(ctx, buyOrderSellOrderMatch)
}

func (this buyOrderSellOrderMatchStore) Has(ctx context.Context, buy_order_id uint64, sell_order_id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, buy_order_id, sell_order_id)
}

func (this buyOrderSellOrderMatchStore) Get(ctx context.Context, buy_order_id uint64, sell_order_id uint64) (*BuyOrderSellOrderMatch, error) {
	var buyOrderSellOrderMatch BuyOrderSellOrderMatch
	found, err := this.table.PrimaryKey().Get(ctx, &buyOrderSellOrderMatch, buy_order_id, sell_order_id)
	if !found {
		return nil, err
	}
	return &buyOrderSellOrderMatch, err
}

func (this buyOrderSellOrderMatchStore) List(ctx context.Context, prefixKey BuyOrderSellOrderMatchIndexKey, opts ...ormlist.Option) (BuyOrderSellOrderMatchIterator, error) {
	opts = append(opts, ormlist.Prefix(prefixKey.values()))
	it, err := this.table.GetIndexByID(prefixKey.id()).Iterator(ctx, opts...)
	return BuyOrderSellOrderMatchIterator{it}, err
}

func (this buyOrderSellOrderMatchStore) ListRange(ctx context.Context, from, to BuyOrderSellOrderMatchIndexKey, opts ...ormlist.Option) (BuyOrderSellOrderMatchIterator, error) {
	opts = append(opts, ormlist.Start(from.values()), ormlist.End(to))
	it, err := this.table.GetIndexByID(from.id()).Iterator(ctx, opts...)
	return BuyOrderSellOrderMatchIterator{it}, err
}

func (this buyOrderSellOrderMatchStore) doNotImplement() {}

var _ BuyOrderSellOrderMatchStore = buyOrderSellOrderMatchStore{}

func NewBuyOrderSellOrderMatchStore(db ormdb.ModuleDB) (BuyOrderSellOrderMatchStore, error) {
	table := db.GetTable(&BuyOrderSellOrderMatch{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&BuyOrderSellOrderMatch{}).ProtoReflect().Descriptor().FullName()))
	}
	return buyOrderSellOrderMatchStore{table}, nil
}

type UInt64SelectorBuyOrderStore interface {
	Insert(ctx context.Context, uInt64SelectorBuyOrder *UInt64SelectorBuyOrder) error
	Update(ctx context.Context, uInt64SelectorBuyOrder *UInt64SelectorBuyOrder) error
	Save(ctx context.Context, uInt64SelectorBuyOrder *UInt64SelectorBuyOrder) error
	Delete(ctx context.Context, uInt64SelectorBuyOrder *UInt64SelectorBuyOrder) error
	Has(ctx context.Context, buy_order_id uint64, selector_type v1beta1.SelectorType, value uint64) (found bool, err error)
	Get(ctx context.Context, buy_order_id uint64, selector_type v1beta1.SelectorType, value uint64) (*UInt64SelectorBuyOrder, error)
	List(ctx context.Context, prefixKey UInt64SelectorBuyOrderIndexKey, opts ...ormlist.Option) (UInt64SelectorBuyOrderIterator, error)
	ListRange(ctx context.Context, from, to UInt64SelectorBuyOrderIndexKey, opts ...ormlist.Option) (UInt64SelectorBuyOrderIterator, error)

	doNotImplement()
}

type UInt64SelectorBuyOrderIterator struct {
	ormtable.Iterator
}

func (i UInt64SelectorBuyOrderIterator) Value() (*UInt64SelectorBuyOrder, error) {
	var uInt64SelectorBuyOrder UInt64SelectorBuyOrder
	err := i.UnmarshalMessage(&uInt64SelectorBuyOrder)
	return &uInt64SelectorBuyOrder, err
}

type UInt64SelectorBuyOrderIndexKey interface {
	id() uint32
	values() []interface{}
	uInt64SelectorBuyOrderIndexKey()
}

// primary key starting index..
type UInt64SelectorBuyOrderBuyOrderIdSelectorTypeValueIndexKey struct {
	vs []interface{}
}

func (x UInt64SelectorBuyOrderBuyOrderIdSelectorTypeValueIndexKey) id() uint32 { return 2 }
func (x UInt64SelectorBuyOrderBuyOrderIdSelectorTypeValueIndexKey) values() []interface{} {
	return x.vs
}
func (x UInt64SelectorBuyOrderBuyOrderIdSelectorTypeValueIndexKey) uInt64SelectorBuyOrderIndexKey() {}

func (this UInt64SelectorBuyOrderBuyOrderIdSelectorTypeValueIndexKey) WithBuyOrderId(buy_order_id uint64) UInt64SelectorBuyOrderBuyOrderIdSelectorTypeValueIndexKey {
	this.vs = []interface{}{buy_order_id}
	return this
}

func (this UInt64SelectorBuyOrderBuyOrderIdSelectorTypeValueIndexKey) WithBuyOrderIdSelectorType(buy_order_id uint64, selector_type v1beta1.SelectorType) UInt64SelectorBuyOrderBuyOrderIdSelectorTypeValueIndexKey {
	this.vs = []interface{}{buy_order_id, selector_type}
	return this
}

func (this UInt64SelectorBuyOrderBuyOrderIdSelectorTypeValueIndexKey) WithBuyOrderIdSelectorTypeValue(buy_order_id uint64, selector_type v1beta1.SelectorType, value uint64) UInt64SelectorBuyOrderBuyOrderIdSelectorTypeValueIndexKey {
	this.vs = []interface{}{buy_order_id, selector_type, value}
	return this
}

type UInt64SelectorBuyOrderSelectorTypeValueIndexKey struct {
	vs []interface{}
}

func (x UInt64SelectorBuyOrderSelectorTypeValueIndexKey) id() uint32                      { return 1 }
func (x UInt64SelectorBuyOrderSelectorTypeValueIndexKey) values() []interface{}           { return x.vs }
func (x UInt64SelectorBuyOrderSelectorTypeValueIndexKey) uInt64SelectorBuyOrderIndexKey() {}

func (this UInt64SelectorBuyOrderSelectorTypeValueIndexKey) WithSelectorType(selector_type v1beta1.SelectorType) UInt64SelectorBuyOrderSelectorTypeValueIndexKey {
	this.vs = []interface{}{selector_type}
	return this
}

func (this UInt64SelectorBuyOrderSelectorTypeValueIndexKey) WithSelectorTypeValue(selector_type v1beta1.SelectorType, value uint64) UInt64SelectorBuyOrderSelectorTypeValueIndexKey {
	this.vs = []interface{}{selector_type, value}
	return this
}

type uInt64SelectorBuyOrderStore struct {
	table ormtable.Table
}

func (this uInt64SelectorBuyOrderStore) Insert(ctx context.Context, uInt64SelectorBuyOrder *UInt64SelectorBuyOrder) error {
	return this.table.Insert(ctx, uInt64SelectorBuyOrder)
}

func (this uInt64SelectorBuyOrderStore) Update(ctx context.Context, uInt64SelectorBuyOrder *UInt64SelectorBuyOrder) error {
	return this.table.Update(ctx, uInt64SelectorBuyOrder)
}

func (this uInt64SelectorBuyOrderStore) Save(ctx context.Context, uInt64SelectorBuyOrder *UInt64SelectorBuyOrder) error {
	return this.table.Save(ctx, uInt64SelectorBuyOrder)
}

func (this uInt64SelectorBuyOrderStore) Delete(ctx context.Context, uInt64SelectorBuyOrder *UInt64SelectorBuyOrder) error {
	return this.table.Delete(ctx, uInt64SelectorBuyOrder)
}

func (this uInt64SelectorBuyOrderStore) Has(ctx context.Context, buy_order_id uint64, selector_type v1beta1.SelectorType, value uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, buy_order_id, selector_type, value)
}

func (this uInt64SelectorBuyOrderStore) Get(ctx context.Context, buy_order_id uint64, selector_type v1beta1.SelectorType, value uint64) (*UInt64SelectorBuyOrder, error) {
	var uInt64SelectorBuyOrder UInt64SelectorBuyOrder
	found, err := this.table.PrimaryKey().Get(ctx, &uInt64SelectorBuyOrder, buy_order_id, selector_type, value)
	if !found {
		return nil, err
	}
	return &uInt64SelectorBuyOrder, err
}

func (this uInt64SelectorBuyOrderStore) List(ctx context.Context, prefixKey UInt64SelectorBuyOrderIndexKey, opts ...ormlist.Option) (UInt64SelectorBuyOrderIterator, error) {
	opts = append(opts, ormlist.Prefix(prefixKey.values()))
	it, err := this.table.GetIndexByID(prefixKey.id()).Iterator(ctx, opts...)
	return UInt64SelectorBuyOrderIterator{it}, err
}

func (this uInt64SelectorBuyOrderStore) ListRange(ctx context.Context, from, to UInt64SelectorBuyOrderIndexKey, opts ...ormlist.Option) (UInt64SelectorBuyOrderIterator, error) {
	opts = append(opts, ormlist.Start(from.values()), ormlist.End(to))
	it, err := this.table.GetIndexByID(from.id()).Iterator(ctx, opts...)
	return UInt64SelectorBuyOrderIterator{it}, err
}

func (this uInt64SelectorBuyOrderStore) doNotImplement() {}

var _ UInt64SelectorBuyOrderStore = uInt64SelectorBuyOrderStore{}

func NewUInt64SelectorBuyOrderStore(db ormdb.ModuleDB) (UInt64SelectorBuyOrderStore, error) {
	table := db.GetTable(&UInt64SelectorBuyOrder{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&UInt64SelectorBuyOrder{}).ProtoReflect().Descriptor().FullName()))
	}
	return uInt64SelectorBuyOrderStore{table}, nil
}

type MemoryStore interface {
	BuyOrderSellOrderMatchStore() BuyOrderSellOrderMatchStore
	UInt64SelectorBuyOrderStore() UInt64SelectorBuyOrderStore

	doNotImplement()
}

type memoryStore struct {
	buyOrderSellOrderMatch BuyOrderSellOrderMatchStore
	uInt64SelectorBuyOrder UInt64SelectorBuyOrderStore
}

func (x memoryStore) BuyOrderSellOrderMatchStore() BuyOrderSellOrderMatchStore {
	return x.buyOrderSellOrderMatch
}

func (x memoryStore) UInt64SelectorBuyOrderStore() UInt64SelectorBuyOrderStore {
	return x.uInt64SelectorBuyOrder
}

func (memoryStore) doNotImplement() {}

var _ MemoryStore = memoryStore{}

func NewMemoryStore(db ormdb.ModuleDB) (MemoryStore, error) {
	buyOrderSellOrderMatchStore, err := NewBuyOrderSellOrderMatchStore(db)
	if err != nil {
		return nil, err
	}

	uInt64SelectorBuyOrderStore, err := NewUInt64SelectorBuyOrderStore(db)
	if err != nil {
		return nil, err
	}

	return memoryStore{
		buyOrderSellOrderMatchStore,
		uInt64SelectorBuyOrderStore,
	}, nil
}
