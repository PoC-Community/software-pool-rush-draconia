package schema

import (
	"errors"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("username").Unique().NotEmpty(),
		field.String("password").NotEmpty().MinLen(8).Sensitive().Validate(func(pwd string) error {
			if strings.ToLower(pwd) == pwd {
				return errors.New("Password must be 8 chars long and contain uppercase letters.")
			}
			return nil
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("characters", Character.Type),
	}
}
