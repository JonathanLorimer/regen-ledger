// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package marketplacev1beta1

import (
	context "context"
	ormdb "github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	ormlist "github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ormtable "github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	ormerrors "github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type SellOrderStore interface {
	Insert(ctx context.Context, sellOrder *SellOrder) error
	Update(ctx context.Context, sellOrder *SellOrder) error
	Save(ctx context.Context, sellOrder *SellOrder) error
	Delete(ctx context.Context, sellOrder *SellOrder) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	Get(ctx context.Context, id uint64) (*SellOrder, error)
	List(ctx context.Context, prefixKey SellOrderIndexKey, opts ...ormlist.Option) (SellOrderIterator, error)
	ListRange(ctx context.Context, from, to SellOrderIndexKey, opts ...ormlist.Option) (SellOrderIterator, error)

	doNotImplement()
}

type SellOrderIterator struct {
	ormtable.Iterator
}

func (i SellOrderIterator) Value() (*SellOrder, error) {
	var sellOrder SellOrder
	err := i.UnmarshalMessage(&sellOrder)
	return &sellOrder, err
}

type SellOrderIndexKey interface {
	id() uint32
	values() []interface{}
	sellOrderIndexKey()
}

// primary key starting index..
type SellOrderPrimaryKey = SellOrderIdIndexKey

type SellOrderIdIndexKey struct {
	vs []interface{}
}

func (x SellOrderIdIndexKey) id() uint32            { return 0 }
func (x SellOrderIdIndexKey) values() []interface{} { return x.vs }
func (x SellOrderIdIndexKey) sellOrderIndexKey()    {}

func (this SellOrderIdIndexKey) WithId(id uint64) SellOrderIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type SellOrderBatchDenomIndexKey struct {
	vs []interface{}
}

func (x SellOrderBatchDenomIndexKey) id() uint32            { return 1 }
func (x SellOrderBatchDenomIndexKey) values() []interface{} { return x.vs }
func (x SellOrderBatchDenomIndexKey) sellOrderIndexKey()    {}

func (this SellOrderBatchDenomIndexKey) WithBatchDenom(batch_denom string) SellOrderBatchDenomIndexKey {
	this.vs = []interface{}{batch_denom}
	return this
}

type SellOrderSellerIndexKey struct {
	vs []interface{}
}

func (x SellOrderSellerIndexKey) id() uint32            { return 2 }
func (x SellOrderSellerIndexKey) values() []interface{} { return x.vs }
func (x SellOrderSellerIndexKey) sellOrderIndexKey()    {}

func (this SellOrderSellerIndexKey) WithSeller(seller []byte) SellOrderSellerIndexKey {
	this.vs = []interface{}{seller}
	return this
}

type SellOrderExpirationIndexKey struct {
	vs []interface{}
}

func (x SellOrderExpirationIndexKey) id() uint32            { return 3 }
func (x SellOrderExpirationIndexKey) values() []interface{} { return x.vs }
func (x SellOrderExpirationIndexKey) sellOrderIndexKey()    {}

func (this SellOrderExpirationIndexKey) WithExpiration(expiration *timestamppb.Timestamp) SellOrderExpirationIndexKey {
	this.vs = []interface{}{expiration}
	return this
}

type sellOrderStore struct {
	table ormtable.Table
}

func (this sellOrderStore) Insert(ctx context.Context, sellOrder *SellOrder) error {
	return this.table.Insert(ctx, sellOrder)
}

func (this sellOrderStore) Update(ctx context.Context, sellOrder *SellOrder) error {
	return this.table.Update(ctx, sellOrder)
}

func (this sellOrderStore) Save(ctx context.Context, sellOrder *SellOrder) error {
	return this.table.Save(ctx, sellOrder)
}

func (this sellOrderStore) Delete(ctx context.Context, sellOrder *SellOrder) error {
	return this.table.Delete(ctx, sellOrder)
}

func (this sellOrderStore) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this sellOrderStore) Get(ctx context.Context, id uint64) (*SellOrder, error) {
	var sellOrder SellOrder
	found, err := this.table.PrimaryKey().Get(ctx, &sellOrder, id)
	if !found {
		return nil, err
	}
	return &sellOrder, err
}

func (this sellOrderStore) List(ctx context.Context, prefixKey SellOrderIndexKey, opts ...ormlist.Option) (SellOrderIterator, error) {
	opts = append(opts, ormlist.Prefix(prefixKey.values()...))
	it, err := this.table.GetIndexByID(prefixKey.id()).Iterator(ctx, opts...)
	return SellOrderIterator{it}, err
}

func (this sellOrderStore) ListRange(ctx context.Context, from, to SellOrderIndexKey, opts ...ormlist.Option) (SellOrderIterator, error) {
	opts = append(opts, ormlist.Start(from.values()...), ormlist.End(to.values()...))
	it, err := this.table.GetIndexByID(from.id()).Iterator(ctx, opts...)
	return SellOrderIterator{it}, err
}

func (this sellOrderStore) doNotImplement() {}

var _ SellOrderStore = sellOrderStore{}

func NewSellOrderStore(db ormdb.ModuleDB) (SellOrderStore, error) {
	table := db.GetTable(&SellOrder{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&SellOrder{}).ProtoReflect().Descriptor().FullName()))
	}
	return sellOrderStore{table}, nil
}

type BuyOrderStore interface {
	Insert(ctx context.Context, buyOrder *BuyOrder) error
	Update(ctx context.Context, buyOrder *BuyOrder) error
	Save(ctx context.Context, buyOrder *BuyOrder) error
	Delete(ctx context.Context, buyOrder *BuyOrder) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	Get(ctx context.Context, id uint64) (*BuyOrder, error)
	List(ctx context.Context, prefixKey BuyOrderIndexKey, opts ...ormlist.Option) (BuyOrderIterator, error)
	ListRange(ctx context.Context, from, to BuyOrderIndexKey, opts ...ormlist.Option) (BuyOrderIterator, error)

	doNotImplement()
}

type BuyOrderIterator struct {
	ormtable.Iterator
}

func (i BuyOrderIterator) Value() (*BuyOrder, error) {
	var buyOrder BuyOrder
	err := i.UnmarshalMessage(&buyOrder)
	return &buyOrder, err
}

type BuyOrderIndexKey interface {
	id() uint32
	values() []interface{}
	buyOrderIndexKey()
}

// primary key starting index..
type BuyOrderPrimaryKey = BuyOrderIdIndexKey

type BuyOrderIdIndexKey struct {
	vs []interface{}
}

func (x BuyOrderIdIndexKey) id() uint32            { return 0 }
func (x BuyOrderIdIndexKey) values() []interface{} { return x.vs }
func (x BuyOrderIdIndexKey) buyOrderIndexKey()     {}

func (this BuyOrderIdIndexKey) WithId(id uint64) BuyOrderIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type BuyOrderBuyerIndexKey struct {
	vs []interface{}
}

func (x BuyOrderBuyerIndexKey) id() uint32            { return 1 }
func (x BuyOrderBuyerIndexKey) values() []interface{} { return x.vs }
func (x BuyOrderBuyerIndexKey) buyOrderIndexKey()     {}

func (this BuyOrderBuyerIndexKey) WithBuyer(buyer []byte) BuyOrderBuyerIndexKey {
	this.vs = []interface{}{buyer}
	return this
}

type BuyOrderExpirationIndexKey struct {
	vs []interface{}
}

func (x BuyOrderExpirationIndexKey) id() uint32            { return 2 }
func (x BuyOrderExpirationIndexKey) values() []interface{} { return x.vs }
func (x BuyOrderExpirationIndexKey) buyOrderIndexKey()     {}

func (this BuyOrderExpirationIndexKey) WithExpiration(expiration *timestamppb.Timestamp) BuyOrderExpirationIndexKey {
	this.vs = []interface{}{expiration}
	return this
}

type buyOrderStore struct {
	table ormtable.Table
}

func (this buyOrderStore) Insert(ctx context.Context, buyOrder *BuyOrder) error {
	return this.table.Insert(ctx, buyOrder)
}

func (this buyOrderStore) Update(ctx context.Context, buyOrder *BuyOrder) error {
	return this.table.Update(ctx, buyOrder)
}

func (this buyOrderStore) Save(ctx context.Context, buyOrder *BuyOrder) error {
	return this.table.Save(ctx, buyOrder)
}

func (this buyOrderStore) Delete(ctx context.Context, buyOrder *BuyOrder) error {
	return this.table.Delete(ctx, buyOrder)
}

func (this buyOrderStore) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this buyOrderStore) Get(ctx context.Context, id uint64) (*BuyOrder, error) {
	var buyOrder BuyOrder
	found, err := this.table.PrimaryKey().Get(ctx, &buyOrder, id)
	if !found {
		return nil, err
	}
	return &buyOrder, err
}

func (this buyOrderStore) List(ctx context.Context, prefixKey BuyOrderIndexKey, opts ...ormlist.Option) (BuyOrderIterator, error) {
	opts = append(opts, ormlist.Prefix(prefixKey.values()...))
	it, err := this.table.GetIndexByID(prefixKey.id()).Iterator(ctx, opts...)
	return BuyOrderIterator{it}, err
}

func (this buyOrderStore) ListRange(ctx context.Context, from, to BuyOrderIndexKey, opts ...ormlist.Option) (BuyOrderIterator, error) {
	opts = append(opts, ormlist.Start(from.values()...), ormlist.End(to.values()...))
	it, err := this.table.GetIndexByID(from.id()).Iterator(ctx, opts...)
	return BuyOrderIterator{it}, err
}

func (this buyOrderStore) doNotImplement() {}

var _ BuyOrderStore = buyOrderStore{}

func NewBuyOrderStore(db ormdb.ModuleDB) (BuyOrderStore, error) {
	table := db.GetTable(&BuyOrder{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&BuyOrder{}).ProtoReflect().Descriptor().FullName()))
	}
	return buyOrderStore{table}, nil
}

type AllowedDenomStore interface {
	Insert(ctx context.Context, allowedDenom *AllowedDenom) error
	Update(ctx context.Context, allowedDenom *AllowedDenom) error
	Save(ctx context.Context, allowedDenom *AllowedDenom) error
	Delete(ctx context.Context, allowedDenom *AllowedDenom) error
	Has(ctx context.Context, bank_denom string) (found bool, err error)
	Get(ctx context.Context, bank_denom string) (*AllowedDenom, error)
	HasByDisplayDenom(ctx context.Context, display_denom string) (found bool, err error)
	GetByDisplayDenom(ctx context.Context, display_denom string) (*AllowedDenom, error)
	List(ctx context.Context, prefixKey AllowedDenomIndexKey, opts ...ormlist.Option) (AllowedDenomIterator, error)
	ListRange(ctx context.Context, from, to AllowedDenomIndexKey, opts ...ormlist.Option) (AllowedDenomIterator, error)

	doNotImplement()
}

type AllowedDenomIterator struct {
	ormtable.Iterator
}

func (i AllowedDenomIterator) Value() (*AllowedDenom, error) {
	var allowedDenom AllowedDenom
	err := i.UnmarshalMessage(&allowedDenom)
	return &allowedDenom, err
}

type AllowedDenomIndexKey interface {
	id() uint32
	values() []interface{}
	allowedDenomIndexKey()
}

// primary key starting index..
type AllowedDenomPrimaryKey = AllowedDenomBankDenomIndexKey

type AllowedDenomBankDenomIndexKey struct {
	vs []interface{}
}

func (x AllowedDenomBankDenomIndexKey) id() uint32            { return 0 }
func (x AllowedDenomBankDenomIndexKey) values() []interface{} { return x.vs }
func (x AllowedDenomBankDenomIndexKey) allowedDenomIndexKey() {}

func (this AllowedDenomBankDenomIndexKey) WithBankDenom(bank_denom string) AllowedDenomBankDenomIndexKey {
	this.vs = []interface{}{bank_denom}
	return this
}

type AllowedDenomDisplayDenomIndexKey struct {
	vs []interface{}
}

func (x AllowedDenomDisplayDenomIndexKey) id() uint32            { return 1 }
func (x AllowedDenomDisplayDenomIndexKey) values() []interface{} { return x.vs }
func (x AllowedDenomDisplayDenomIndexKey) allowedDenomIndexKey() {}

func (this AllowedDenomDisplayDenomIndexKey) WithDisplayDenom(display_denom string) AllowedDenomDisplayDenomIndexKey {
	this.vs = []interface{}{display_denom}
	return this
}

type allowedDenomStore struct {
	table ormtable.Table
}

func (this allowedDenomStore) Insert(ctx context.Context, allowedDenom *AllowedDenom) error {
	return this.table.Insert(ctx, allowedDenom)
}

func (this allowedDenomStore) Update(ctx context.Context, allowedDenom *AllowedDenom) error {
	return this.table.Update(ctx, allowedDenom)
}

func (this allowedDenomStore) Save(ctx context.Context, allowedDenom *AllowedDenom) error {
	return this.table.Save(ctx, allowedDenom)
}

func (this allowedDenomStore) Delete(ctx context.Context, allowedDenom *AllowedDenom) error {
	return this.table.Delete(ctx, allowedDenom)
}

func (this allowedDenomStore) Has(ctx context.Context, bank_denom string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, bank_denom)
}

func (this allowedDenomStore) Get(ctx context.Context, bank_denom string) (*AllowedDenom, error) {
	var allowedDenom AllowedDenom
	found, err := this.table.PrimaryKey().Get(ctx, &allowedDenom, bank_denom)
	if !found {
		return nil, err
	}
	return &allowedDenom, err
}

func (this allowedDenomStore) HasByDisplayDenom(ctx context.Context, display_denom string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		display_denom,
	)
}

func (this allowedDenomStore) GetByDisplayDenom(ctx context.Context, display_denom string) (*AllowedDenom, error) {
	var allowedDenom AllowedDenom
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &allowedDenom,
		display_denom,
	)
	if !found {
		return nil, err
	}
	return &allowedDenom, nil
}

func (this allowedDenomStore) List(ctx context.Context, prefixKey AllowedDenomIndexKey, opts ...ormlist.Option) (AllowedDenomIterator, error) {
	opts = append(opts, ormlist.Prefix(prefixKey.values()...))
	it, err := this.table.GetIndexByID(prefixKey.id()).Iterator(ctx, opts...)
	return AllowedDenomIterator{it}, err
}

func (this allowedDenomStore) ListRange(ctx context.Context, from, to AllowedDenomIndexKey, opts ...ormlist.Option) (AllowedDenomIterator, error) {
	opts = append(opts, ormlist.Start(from.values()...), ormlist.End(to.values()...))
	it, err := this.table.GetIndexByID(from.id()).Iterator(ctx, opts...)
	return AllowedDenomIterator{it}, err
}

func (this allowedDenomStore) doNotImplement() {}

var _ AllowedDenomStore = allowedDenomStore{}

func NewAllowedDenomStore(db ormdb.ModuleDB) (AllowedDenomStore, error) {
	table := db.GetTable(&AllowedDenom{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&AllowedDenom{}).ProtoReflect().Descriptor().FullName()))
	}
	return allowedDenomStore{table}, nil
}

type MarketStore interface {
	Insert(ctx context.Context, market *Market) error
	Update(ctx context.Context, market *Market) error
	Save(ctx context.Context, market *Market) error
	Delete(ctx context.Context, market *Market) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	Get(ctx context.Context, id uint64) (*Market, error)
	HasByCreditTypeBankDenom(ctx context.Context, credit_type string, bank_denom string) (found bool, err error)
	GetByCreditTypeBankDenom(ctx context.Context, credit_type string, bank_denom string) (*Market, error)
	List(ctx context.Context, prefixKey MarketIndexKey, opts ...ormlist.Option) (MarketIterator, error)
	ListRange(ctx context.Context, from, to MarketIndexKey, opts ...ormlist.Option) (MarketIterator, error)

	doNotImplement()
}

type MarketIterator struct {
	ormtable.Iterator
}

func (i MarketIterator) Value() (*Market, error) {
	var market Market
	err := i.UnmarshalMessage(&market)
	return &market, err
}

type MarketIndexKey interface {
	id() uint32
	values() []interface{}
	marketIndexKey()
}

// primary key starting index..
type MarketPrimaryKey = MarketIdIndexKey

type MarketIdIndexKey struct {
	vs []interface{}
}

func (x MarketIdIndexKey) id() uint32            { return 0 }
func (x MarketIdIndexKey) values() []interface{} { return x.vs }
func (x MarketIdIndexKey) marketIndexKey()       {}

func (this MarketIdIndexKey) WithId(id uint64) MarketIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type MarketCreditTypeBankDenomIndexKey struct {
	vs []interface{}
}

func (x MarketCreditTypeBankDenomIndexKey) id() uint32            { return 1 }
func (x MarketCreditTypeBankDenomIndexKey) values() []interface{} { return x.vs }
func (x MarketCreditTypeBankDenomIndexKey) marketIndexKey()       {}

func (this MarketCreditTypeBankDenomIndexKey) WithCreditType(credit_type string) MarketCreditTypeBankDenomIndexKey {
	this.vs = []interface{}{credit_type}
	return this
}

func (this MarketCreditTypeBankDenomIndexKey) WithCreditTypeBankDenom(credit_type string, bank_denom string) MarketCreditTypeBankDenomIndexKey {
	this.vs = []interface{}{credit_type, bank_denom}
	return this
}

type marketStore struct {
	table ormtable.Table
}

func (this marketStore) Insert(ctx context.Context, market *Market) error {
	return this.table.Insert(ctx, market)
}

func (this marketStore) Update(ctx context.Context, market *Market) error {
	return this.table.Update(ctx, market)
}

func (this marketStore) Save(ctx context.Context, market *Market) error {
	return this.table.Save(ctx, market)
}

func (this marketStore) Delete(ctx context.Context, market *Market) error {
	return this.table.Delete(ctx, market)
}

func (this marketStore) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this marketStore) Get(ctx context.Context, id uint64) (*Market, error) {
	var market Market
	found, err := this.table.PrimaryKey().Get(ctx, &market, id)
	if !found {
		return nil, err
	}
	return &market, err
}

func (this marketStore) HasByCreditTypeBankDenom(ctx context.Context, credit_type string, bank_denom string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		credit_type,
		bank_denom,
	)
}

func (this marketStore) GetByCreditTypeBankDenom(ctx context.Context, credit_type string, bank_denom string) (*Market, error) {
	var market Market
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &market,
		credit_type,
		bank_denom,
	)
	if !found {
		return nil, err
	}
	return &market, nil
}

func (this marketStore) List(ctx context.Context, prefixKey MarketIndexKey, opts ...ormlist.Option) (MarketIterator, error) {
	opts = append(opts, ormlist.Prefix(prefixKey.values()...))
	it, err := this.table.GetIndexByID(prefixKey.id()).Iterator(ctx, opts...)
	return MarketIterator{it}, err
}

func (this marketStore) ListRange(ctx context.Context, from, to MarketIndexKey, opts ...ormlist.Option) (MarketIterator, error) {
	opts = append(opts, ormlist.Start(from.values()...), ormlist.End(to.values()...))
	it, err := this.table.GetIndexByID(from.id()).Iterator(ctx, opts...)
	return MarketIterator{it}, err
}

func (this marketStore) doNotImplement() {}

var _ MarketStore = marketStore{}

func NewMarketStore(db ormdb.ModuleDB) (MarketStore, error) {
	table := db.GetTable(&Market{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Market{}).ProtoReflect().Descriptor().FullName()))
	}
	return marketStore{table}, nil
}

type StateStore interface {
	SellOrderStore() SellOrderStore
	BuyOrderStore() BuyOrderStore
	AllowedDenomStore() AllowedDenomStore
	MarketStore() MarketStore

	doNotImplement()
}

type stateStore struct {
	sellOrder    SellOrderStore
	buyOrder     BuyOrderStore
	allowedDenom AllowedDenomStore
	market       MarketStore
}

func (x stateStore) SellOrderStore() SellOrderStore {
	return x.sellOrder
}

func (x stateStore) BuyOrderStore() BuyOrderStore {
	return x.buyOrder
}

func (x stateStore) AllowedDenomStore() AllowedDenomStore {
	return x.allowedDenom
}

func (x stateStore) MarketStore() MarketStore {
	return x.market
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormdb.ModuleDB) (StateStore, error) {
	sellOrderStore, err := NewSellOrderStore(db)
	if err != nil {
		return nil, err
	}

	buyOrderStore, err := NewBuyOrderStore(db)
	if err != nil {
		return nil, err
	}

	allowedDenomStore, err := NewAllowedDenomStore(db)
	if err != nil {
		return nil, err
	}

	marketStore, err := NewMarketStore(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		sellOrderStore,
		buyOrderStore,
		allowedDenomStore,
		marketStore,
	}, nil
}
