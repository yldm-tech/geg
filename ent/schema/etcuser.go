package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// ETCUser holds the schema definition for the ETCUser entity.
type ETCUser struct {
	ent.Schema
}

func (ETCUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "etc_user"},
	}
}

// Fields of the ETCUser.
func (ETCUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint",
		}).Unique(),
		field.Int64("etc_username").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("etc_password").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)",
		}),
		field.String("point_username").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("point_password").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)",
		}),
		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now()),
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now()),
	}
}

// Edges of the ETCUser.
//func (ETCUser) Edges() []ent.Edge {
//	return nil
//}
