// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EtcRecordColumns holds the columns for the "etc_record" table.
	EtcRecordColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "username", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "entry", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(128)"}},
		{Name: "entry_year", Type: field.TypeInt32, SchemaType: map[string]string{"mysql": "int(10)"}},
		{Name: "entry_month", Type: field.TypeInt32, SchemaType: map[string]string{"mysql": "int(10)"}},
		{Name: "entry_day", Type: field.TypeInt32, SchemaType: map[string]string{"mysql": "int(10)"}},
		{Name: "exit", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "exit_date", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "exit_time", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "total_price", Type: field.TypeInt32, SchemaType: map[string]string{"mysql": "int(10)"}},
		{Name: "discount_price", Type: field.TypeInt32, SchemaType: map[string]string{"mysql": "int(10)"}},
		{Name: "paid_price", Type: field.TypeInt32, SchemaType: map[string]string{"mysql": "int(10)"}},
		{Name: "car_type", Type: field.TypeInt8, SchemaType: map[string]string{"mysql": "int(2)"}},
		{Name: "car_number", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(32)"}},
		{Name: "card_number", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "status", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "comment", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// EtcRecordTable holds the schema information for the "etc_record" table.
	EtcRecordTable = &schema.Table{
		Name:       "etc_record",
		Columns:    EtcRecordColumns,
		PrimaryKey: []*schema.Column{EtcRecordColumns[0]},
	}
	// EtcSyncRecordColumns holds the columns for the "etc_sync_record" table.
	EtcSyncRecordColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "username", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "sync_data", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// EtcSyncRecordTable holds the schema information for the "etc_sync_record" table.
	EtcSyncRecordTable = &schema.Table{
		Name:       "etc_sync_record",
		Columns:    EtcSyncRecordColumns,
		PrimaryKey: []*schema.Column{EtcSyncRecordColumns[0]},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "etc_username", Type: field.TypeInt64, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "etc_password", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(512)"}},
		{Name: "point_username", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "point_password", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(512)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EtcRecordTable,
		EtcSyncRecordTable,
		UserTable,
	}
)

func init() {
	EtcRecordTable.Annotation = &entsql.Annotation{
		Table: "etc_record",
	}
	EtcSyncRecordTable.Annotation = &entsql.Annotation{
		Table: "etc_sync_record",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}
