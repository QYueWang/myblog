package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable().Unique(),
		field.String("title").MinLen(5).MaxLen(60),
		field.String("content"),
		field.Time("create_at").Immutable().Default(time.Now()),
		field.Time("update_at").Default(time.Now()),
		field.Time("delete_at").Default(time.Now()),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
		edge.To("tags", Tag.Type),
	}
}
