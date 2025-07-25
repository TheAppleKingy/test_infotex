// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"test_infotex/ent/transaction"
	"test_infotex/ent/wallet"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WalletCreate is the builder for creating a Wallet entity.
type WalletCreate struct {
	config
	mutation *WalletMutation
	hooks    []Hook
}

// SetAddress sets the "address" field.
func (wc *WalletCreate) SetAddress(s string) *WalletCreate {
	wc.mutation.SetAddress(s)
	return wc
}

// SetBalance sets the "balance" field.
func (wc *WalletCreate) SetBalance(i int) *WalletCreate {
	wc.mutation.SetBalance(i)
	return wc
}

// AddSentTransactionIDs adds the "sent_transactions" edge to the Transaction entity by IDs.
func (wc *WalletCreate) AddSentTransactionIDs(ids ...int) *WalletCreate {
	wc.mutation.AddSentTransactionIDs(ids...)
	return wc
}

// AddSentTransactions adds the "sent_transactions" edges to the Transaction entity.
func (wc *WalletCreate) AddSentTransactions(t ...*Transaction) *WalletCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddSentTransactionIDs(ids...)
}

// AddRecievedTransactionIDs adds the "recieved_transactions" edge to the Transaction entity by IDs.
func (wc *WalletCreate) AddRecievedTransactionIDs(ids ...int) *WalletCreate {
	wc.mutation.AddRecievedTransactionIDs(ids...)
	return wc
}

// AddRecievedTransactions adds the "recieved_transactions" edges to the Transaction entity.
func (wc *WalletCreate) AddRecievedTransactions(t ...*Transaction) *WalletCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wc.AddRecievedTransactionIDs(ids...)
}

// Mutation returns the WalletMutation object of the builder.
func (wc *WalletCreate) Mutation() *WalletMutation {
	return wc.mutation
}

// Save creates the Wallet in the database.
func (wc *WalletCreate) Save(ctx context.Context) (*Wallet, error) {
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WalletCreate) SaveX(ctx context.Context) *Wallet {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WalletCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WalletCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WalletCreate) check() error {
	if _, ok := wc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Wallet.address"`)}
	}
	if v, ok := wc.mutation.Address(); ok {
		if err := wallet.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Wallet.address": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`ent: missing required field "Wallet.balance"`)}
	}
	if v, ok := wc.mutation.Balance(); ok {
		if err := wallet.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf(`ent: validator failed for field "Wallet.balance": %w`, err)}
		}
	}
	return nil
}

func (wc *WalletCreate) sqlSave(ctx context.Context) (*Wallet, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WalletCreate) createSpec() (*Wallet, *sqlgraph.CreateSpec) {
	var (
		_node = &Wallet{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(wallet.Table, sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeInt))
	)
	if value, ok := wc.mutation.Address(); ok {
		_spec.SetField(wallet.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := wc.mutation.Balance(); ok {
		_spec.SetField(wallet.FieldBalance, field.TypeInt, value)
		_node.Balance = value
	}
	if nodes := wc.mutation.SentTransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.SentTransactionsTable,
			Columns: []string{wallet.SentTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.RecievedTransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.RecievedTransactionsTable,
			Columns: []string{wallet.RecievedTransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WalletCreateBulk is the builder for creating many Wallet entities in bulk.
type WalletCreateBulk struct {
	config
	err      error
	builders []*WalletCreate
}

// Save creates the Wallet entities in the database.
func (wcb *WalletCreateBulk) Save(ctx context.Context) ([]*Wallet, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Wallet, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WalletMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WalletCreateBulk) SaveX(ctx context.Context) []*Wallet {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WalletCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WalletCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
