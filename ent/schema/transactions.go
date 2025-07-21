package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Transaction struct {
	ent.Schema
}

func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount").Immutable().NonNegative(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Int("from_wallet_id"),
		field.Int("to_wallet_id"),
	}
}

func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("from_wallet", Wallet.Type).Ref("sent_transactions").Field("from_wallet_id").Unique().Required(),
		edge.From("to_wallet", Wallet.Type).Ref("recieved_transactions").Field("to_wallet_id").Unique().Required(),
	}
}
