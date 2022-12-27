// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"myetc.lantron.ltd/m/ent/etcsyncrecord"
)

// ETCSyncRecord is the model entity for the ETCSyncRecord schema.
type ETCSyncRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// SyncData holds the value of the "sync_data" field.
	SyncData string `json:"sync_data,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ETCSyncRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case etcsyncrecord.FieldID:
			values[i] = new(sql.NullInt64)
		case etcsyncrecord.FieldUsername, etcsyncrecord.FieldSyncData:
			values[i] = new(sql.NullString)
		case etcsyncrecord.FieldCreatedAt, etcsyncrecord.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ETCSyncRecord", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ETCSyncRecord fields.
func (esr *ETCSyncRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case etcsyncrecord.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			esr.ID = int64(value.Int64)
		case etcsyncrecord.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				esr.Username = value.String
			}
		case etcsyncrecord.FieldSyncData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sync_data", values[i])
			} else if value.Valid {
				esr.SyncData = value.String
			}
		case etcsyncrecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				esr.CreatedAt = value.Time
			}
		case etcsyncrecord.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				esr.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ETCSyncRecord.
// Note that you need to call ETCSyncRecord.Unwrap() before calling this method if this ETCSyncRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (esr *ETCSyncRecord) Update() *ETCSyncRecordUpdateOne {
	return (&ETCSyncRecordClient{config: esr.config}).UpdateOne(esr)
}

// Unwrap unwraps the ETCSyncRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (esr *ETCSyncRecord) Unwrap() *ETCSyncRecord {
	_tx, ok := esr.config.driver.(*txDriver)
	if !ok {
		panic("ent: ETCSyncRecord is not a transactional entity")
	}
	esr.config.driver = _tx.drv
	return esr
}

// String implements the fmt.Stringer.
func (esr *ETCSyncRecord) String() string {
	var builder strings.Builder
	builder.WriteString("ETCSyncRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", esr.ID))
	builder.WriteString("username=")
	builder.WriteString(esr.Username)
	builder.WriteString(", ")
	builder.WriteString("sync_data=")
	builder.WriteString(esr.SyncData)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(esr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(esr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ETCSyncRecords is a parsable slice of ETCSyncRecord.
type ETCSyncRecords []*ETCSyncRecord

func (esr ETCSyncRecords) config(cfg config) {
	for _i := range esr {
		esr[_i].config = cfg
	}
}