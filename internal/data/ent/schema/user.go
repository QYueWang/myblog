package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("account").Unique().Immutable(),
		field.String("name"),
		field.String("email"),
		field.String("password"),
		field.Time("create_at").Immutable().Default(time.Now()),
		field.Time("update_at").Default(time.Now()),
		field.Time("delete_at").Default(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("articles", Article.Type),
		edge.To("comments", Comment.Type),
	}
}
