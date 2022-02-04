// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package testpb

import (
	context "context"

	ormlist "github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ormtable "github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	ormerrors "github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
)

type BalanceStore interface {
	Insert(ctx context.Context, balance *Balance) error
	Update(ctx context.Context, balance *Balance) error
	Save(ctx context.Context, balance *Balance) error
	Delete(ctx context.Context, balance *Balance) error
	Has(ctx context.Context, address string, denom string) (found bool, err error)
	Get(ctx context.Context, address string, denom string) (*Balance, error)
	List(ctx context.Context, prefixKey BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error)
	ListRange(ctx context.Context, from, to BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error)
	DeleteBy(ctx context.Context, prefixKey BalanceIndexKey) error
	DeleteRange(ctx context.Context, from, to BalanceIndexKey) error

	doNotImplement()
}

type BalanceIterator struct {
	ormtable.Iterator
}

func (i BalanceIterator) Value() (*Balance, error) {
	var balance Balance
	err := i.UnmarshalMessage(&balance)
	return &balance, err
}

type BalanceIndexKey interface {
	id() uint32
	values() []interface{}
	balanceIndexKey()
}

// primary key starting index..
type BalancePrimaryKey = BalanceAddressDenomIndexKey

type BalanceAddressDenomIndexKey struct {
	vs []interface{}
}

func (x BalanceAddressDenomIndexKey) id() uint32            { return 0 }
func (x BalanceAddressDenomIndexKey) values() []interface{} { return x.vs }
func (x BalanceAddressDenomIndexKey) balanceIndexKey()      {}

func (this BalanceAddressDenomIndexKey) WithAddress(address string) BalanceAddressDenomIndexKey {
	this.vs = []interface{}{address}
	return this
}

func (this BalanceAddressDenomIndexKey) WithAddressDenom(address string, denom string) BalanceAddressDenomIndexKey {
	this.vs = []interface{}{address, denom}
	return this
}

type BalanceDenomIndexKey struct {
	vs []interface{}
}

func (x BalanceDenomIndexKey) id() uint32            { return 1 }
func (x BalanceDenomIndexKey) values() []interface{} { return x.vs }
func (x BalanceDenomIndexKey) balanceIndexKey()      {}

func (this BalanceDenomIndexKey) WithDenom(denom string) BalanceDenomIndexKey {
	this.vs = []interface{}{denom}
	return this
}

type balanceStore struct {
	table ormtable.Table
}

func (this balanceStore) Insert(ctx context.Context, balance *Balance) error {
	return this.table.Insert(ctx, balance)
}

func (this balanceStore) Update(ctx context.Context, balance *Balance) error {
	return this.table.Update(ctx, balance)
}

func (this balanceStore) Save(ctx context.Context, balance *Balance) error {
	return this.table.Save(ctx, balance)
}

func (this balanceStore) Delete(ctx context.Context, balance *Balance) error {
	return this.table.Delete(ctx, balance)
}

func (this balanceStore) Has(ctx context.Context, address string, denom string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, address, denom)
}

func (this balanceStore) Get(ctx context.Context, address string, denom string) (*Balance, error) {
	var balance Balance
	found, err := this.table.PrimaryKey().Get(ctx, &balance, address, denom)
	if !found {
		return nil, err
	}
	return &balance, err
}

func (this balanceStore) List(ctx context.Context, prefixKey BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return BalanceIterator{it}, err
}

func (this balanceStore) ListRange(ctx context.Context, from, to BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return BalanceIterator{it}, err
}

func (this balanceStore) DeleteBy(ctx context.Context, prefixKey BalanceIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this balanceStore) DeleteRange(ctx context.Context, from, to BalanceIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this balanceStore) doNotImplement() {}

var _ BalanceStore = balanceStore{}

func NewBalanceStore(db ormtable.Schema) (BalanceStore, error) {
	table := db.GetTable(&Balance{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Balance{}).ProtoReflect().Descriptor().FullName()))
	}
	return balanceStore{table}, nil
}

type SupplyStore interface {
	Insert(ctx context.Context, supply *Supply) error
	Update(ctx context.Context, supply *Supply) error
	Save(ctx context.Context, supply *Supply) error
	Delete(ctx context.Context, supply *Supply) error
	Has(ctx context.Context, denom string) (found bool, err error)
	Get(ctx context.Context, denom string) (*Supply, error)
	List(ctx context.Context, prefixKey SupplyIndexKey, opts ...ormlist.Option) (SupplyIterator, error)
	ListRange(ctx context.Context, from, to SupplyIndexKey, opts ...ormlist.Option) (SupplyIterator, error)
	DeleteBy(ctx context.Context, prefixKey SupplyIndexKey) error
	DeleteRange(ctx context.Context, from, to SupplyIndexKey) error

	doNotImplement()
}

type SupplyIterator struct {
	ormtable.Iterator
}

func (i SupplyIterator) Value() (*Supply, error) {
	var supply Supply
	err := i.UnmarshalMessage(&supply)
	return &supply, err
}

type SupplyIndexKey interface {
	id() uint32
	values() []interface{}
	supplyIndexKey()
}

// primary key starting index..
type SupplyPrimaryKey = SupplyDenomIndexKey

type SupplyDenomIndexKey struct {
	vs []interface{}
}

func (x SupplyDenomIndexKey) id() uint32            { return 0 }
func (x SupplyDenomIndexKey) values() []interface{} { return x.vs }
func (x SupplyDenomIndexKey) supplyIndexKey()       {}

func (this SupplyDenomIndexKey) WithDenom(denom string) SupplyDenomIndexKey {
	this.vs = []interface{}{denom}
	return this
}

type supplyStore struct {
	table ormtable.Table
}

func (this supplyStore) Insert(ctx context.Context, supply *Supply) error {
	return this.table.Insert(ctx, supply)
}

func (this supplyStore) Update(ctx context.Context, supply *Supply) error {
	return this.table.Update(ctx, supply)
}

func (this supplyStore) Save(ctx context.Context, supply *Supply) error {
	return this.table.Save(ctx, supply)
}

func (this supplyStore) Delete(ctx context.Context, supply *Supply) error {
	return this.table.Delete(ctx, supply)
}

func (this supplyStore) Has(ctx context.Context, denom string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, denom)
}

func (this supplyStore) Get(ctx context.Context, denom string) (*Supply, error) {
	var supply Supply
	found, err := this.table.PrimaryKey().Get(ctx, &supply, denom)
	if !found {
		return nil, err
	}
	return &supply, err
}

func (this supplyStore) List(ctx context.Context, prefixKey SupplyIndexKey, opts ...ormlist.Option) (SupplyIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return SupplyIterator{it}, err
}

func (this supplyStore) ListRange(ctx context.Context, from, to SupplyIndexKey, opts ...ormlist.Option) (SupplyIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return SupplyIterator{it}, err
}

func (this supplyStore) DeleteBy(ctx context.Context, prefixKey SupplyIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this supplyStore) DeleteRange(ctx context.Context, from, to SupplyIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this supplyStore) doNotImplement() {}

var _ SupplyStore = supplyStore{}

func NewSupplyStore(db ormtable.Schema) (SupplyStore, error) {
	table := db.GetTable(&Supply{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Supply{}).ProtoReflect().Descriptor().FullName()))
	}
	return supplyStore{table}, nil
}

type BankStore interface {
	BalanceStore() BalanceStore
	SupplyStore() SupplyStore

	doNotImplement()
}

type bankStore struct {
	balance BalanceStore
	supply  SupplyStore
}

func (x bankStore) BalanceStore() BalanceStore {
	return x.balance
}

func (x bankStore) SupplyStore() SupplyStore {
	return x.supply
}

func (bankStore) doNotImplement() {}

var _ BankStore = bankStore{}

func NewBankStore(db ormtable.Schema) (BankStore, error) {
	balanceStore, err := NewBalanceStore(db)
	if err != nil {
		return nil, err
	}

	supplyStore, err := NewSupplyStore(db)
	if err != nil {
		return nil, err
	}

	return bankStore{
		balanceStore,
		supplyStore,
	}, nil
}
