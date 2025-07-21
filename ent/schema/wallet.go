package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Wallet struct {
	ent.Schema
}

func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").Immutable().NotEmpty().Unique(),
		field.Int("balance").NonNegative(),
	}
}

func (Wallet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sent_transactions", Transaction.Type),
		edge.To("recieved_transactions", Transaction.Type),
	}
}
