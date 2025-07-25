// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"test_infotex/ent/wallet"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Wallet is the model entity for the Wallet schema.
type Wallet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance int `json:"balance,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WalletQuery when eager-loading is set.
	Edges        WalletEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WalletEdges holds the relations/edges for other nodes in the graph.
type WalletEdges struct {
	// SentTransactions holds the value of the sent_transactions edge.
	SentTransactions []*Transaction `json:"sent_transactions,omitempty"`
	// RecievedTransactions holds the value of the recieved_transactions edge.
	RecievedTransactions []*Transaction `json:"recieved_transactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SentTransactionsOrErr returns the SentTransactions value or an error if the edge
// was not loaded in eager-loading.
func (e WalletEdges) SentTransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[0] {
		return e.SentTransactions, nil
	}
	return nil, &NotLoadedError{edge: "sent_transactions"}
}

// RecievedTransactionsOrErr returns the RecievedTransactions value or an error if the edge
// was not loaded in eager-loading.
func (e WalletEdges) RecievedTransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[1] {
		return e.RecievedTransactions, nil
	}
	return nil, &NotLoadedError{edge: "recieved_transactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Wallet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wallet.FieldID, wallet.FieldBalance:
			values[i] = new(sql.NullInt64)
		case wallet.FieldAddress:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Wallet fields.
func (w *Wallet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wallet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case wallet.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				w.Address = value.String
			}
		case wallet.FieldBalance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				w.Balance = int(value.Int64)
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Wallet.
// This includes values selected through modifiers, order, etc.
func (w *Wallet) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// QuerySentTransactions queries the "sent_transactions" edge of the Wallet entity.
func (w *Wallet) QuerySentTransactions() *TransactionQuery {
	return NewWalletClient(w.config).QuerySentTransactions(w)
}

// QueryRecievedTransactions queries the "recieved_transactions" edge of the Wallet entity.
func (w *Wallet) QueryRecievedTransactions() *TransactionQuery {
	return NewWalletClient(w.config).QueryRecievedTransactions(w)
}

// Update returns a builder for updating this Wallet.
// Note that you need to call Wallet.Unwrap() before calling this method if this Wallet
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Wallet) Update() *WalletUpdateOne {
	return NewWalletClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Wallet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Wallet) Unwrap() *Wallet {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Wallet is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Wallet) String() string {
	var builder strings.Builder
	builder.WriteString("Wallet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("address=")
	builder.WriteString(w.Address)
	builder.WriteString(", ")
	builder.WriteString("balance=")
	builder.WriteString(fmt.Sprintf("%v", w.Balance))
	builder.WriteByte(')')
	return builder.String()
}

// Wallets is a parsable slice of Wallet.
type Wallets []*Wallet
