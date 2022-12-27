package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// ETCSyncRecord holds the schema definition for the ETCSyncRecord entity.
type ETCSyncRecord struct {
	ent.Schema
}

func (ETCSyncRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "etc_sync_record"},
	}
}

// Fields of the ETCSyncRecord.
func (ETCSyncRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint",
		}).Unique(),
		field.String("username").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("sync_data").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now()),
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now()),
	}
}

// Edges of the ETCSyncRecord.
//func (ETCSyncRecord) Edges() []ent.Edge {
//	return nil
//}
