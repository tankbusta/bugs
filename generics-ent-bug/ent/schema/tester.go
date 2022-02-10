package schema

import (
	"genericserr/ent/privacy"
	"genericserr/lib/permissions"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tester holds the schema definition for the Tester entity.
type Tester struct {
	ent.Schema
}

// Fields of the Tester.
func (Tester) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
		field.String("username").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Tester.
func (Tester) Edges() []ent.Edge {
	return nil
}

//Policy defines the privacy policy of the Tester.
// TODO(cschmitt): Comment this entire Policy block out and everything works
// Otherwise you get go generate
// Unexpected package creation during export data loading
// exit status 1
func (Tester) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			privacy.DenyIfNoViewer(permissions.Read),
		},
	}
}
