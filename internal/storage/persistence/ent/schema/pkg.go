package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"

	"github.com/artefactual-sdps/enduro/internal/storage/types"
)

// Pkg holds the schema definition for the Pkg entity.
type Pkg struct {
	ent.Schema
}

// Annotations of the Pkg.
func (Pkg) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "package"},
	}
}

// Fields of the Pkg.
func (Pkg) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(entsql.Annotation{
				Size: 2048,
			}),
		field.UUID("aip_id", uuid.UUID{}).
			Unique(),
		field.Int("location_id").
			Optional(),
		field.Enum("status").
			GoType(types.StatusUnspecified),
		field.UUID("object_key", uuid.UUID{}).
			Unique(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
	}
}

// Edges of the Pkg.
func (Pkg) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("location", Location.Type).
			Field("location_id").
			Unique(),
	}
}

// Indexes of the Pkg.
func (Pkg) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("aip_id"),
		index.Fields("object_key"),
	}
}
