package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable().Unique(),
		field.String("name"),
		field.String("content"),
		field.Time("create_at").Immutable().Default(time.Now()),
		field.Time("update_at").Default(time.Now()),
		field.Time("delete_at").Default(time.Now()),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).Ref("comments").Unique(),
		edge.From("user", User.Type).Ref("comments").Unique(),
	}
}
