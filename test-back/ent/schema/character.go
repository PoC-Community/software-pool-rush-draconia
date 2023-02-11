package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),

		// attributs normaux
		field.Int("combat").Min(0),
		field.Int("connaissance").Min(0),
		field.Int("discretion").Min(0),
		field.Int("endurance").Min(0),
		field.Int("force").Min(0),
		field.Int("habilite").Min(0),
		field.Int("mouvement").Min(0),
		field.Int("perception").Min(0),
		field.Int("sociabilite").Min(0),
		field.Int("survie").Min(0),
		field.Int("tir").Min(0),
		field.Int("volonte").Min(0),
		field.Int("exp").Min(0),

		// armes
		field.Int("armes_hast").Min(0),
		field.Int("armes_moine").Min(0),
		field.Int("armes_doubles").Min(0),
		field.Int("armes_naturelles").Min(0),
		field.Int("batons").Min(0),
		field.Int("cimeterres").Min(0),
		field.Int("fleaux").Min(0),
		field.Int("fouets").Min(0),
		field.Int("haches").Min(0),
		field.Int("katanas").Min(0),
		field.Int("lames_legeres").Min(0),
		field.Int("lames_lourdes").Min(0),
		field.Int("lances").Min(0),
		field.Int("marteaux").Min(0),
		field.Int("mains_nues").Min(0),

		// connaissances
		field.Int("mysteres").Min(0),
		field.Int("exploration_souterraine").Min(0),
		field.Int("ingenierie").Min(0),
		field.Int("geographie").Min(0),
		field.Int("histoire").Min(0),
		field.Int("folklore").Min(0),
		field.Int("nature").Min(0),
		field.Int("noblesse").Min(0),
		field.Int("plans").Min(0),
		field.Int("religon").Min(0),
		field.Int("anatomie").Min(0),
		field.Int("magie_theorique").Min(0),
		field.Int("economie").Min(0),
		field.Int("linguistique").Min(0),
	}
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("characters").Unique(),
	}
}
