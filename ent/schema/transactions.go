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
		field.Uint("amount").Immutable().Positive(),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("from_wallet", Wallet.Type).Ref("sent_transactions").Unique().Required(),
		edge.From("to_wallet", Wallet.Type).Ref("recieved_transactions").Unique().Required(),
	}
}
