package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// ETCRecord holds the schema definition for the ETCRecord entity.
type ETCRecord struct {
	ent.Schema
}

func (ETCRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "etc_record"},
	}
}

// Fields of the ETCRecord.
func (ETCRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint",
		}).Unique(),
		field.String("username").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("entry").SchemaType(map[string]string{
			dialect.MySQL: "varchar(128)",
		}),
		field.Int32("entry_year").SchemaType(map[string]string{
			dialect.MySQL: "int(10)",
		}),
		field.Int32("entry_month").SchemaType(map[string]string{
			dialect.MySQL: "int(10)",
		}),
		field.Int32("entry_day").SchemaType(map[string]string{
			dialect.MySQL: "int(10)",
		}),
		field.String("exit").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("exit_date").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("exit_time").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.Int32("total_price").SchemaType(map[string]string{
			dialect.MySQL: "int(10)",
		}),
		field.Int32("discount_price").SchemaType(map[string]string{
			dialect.MySQL: "int(10)",
		}),
		field.Int32("paid_price").SchemaType(map[string]string{
			dialect.MySQL: "int(10)",
		}),
		field.Int8("car_type").SchemaType(map[string]string{
			dialect.MySQL: "int(2)",
		}),
		field.String("car_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)",
		}),
		field.String("card_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("status").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("comment").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}).Optional().Nillable(),
		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now()),
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now()),
	}
}

// Edges of the ETCRecord.
//func (ETCRecord) Edges() []ent.Edge {
//	return nil
//}
