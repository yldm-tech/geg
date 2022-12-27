// Code generated by ent, DO NOT EDIT.

package etcrecord

import (
	"time"
)

const (
	// Label holds the string label denoting the etcrecord type in the database.
	Label = "etc_record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEntry holds the string denoting the entry field in the database.
	FieldEntry = "entry"
	// FieldEntryYear holds the string denoting the entry_year field in the database.
	FieldEntryYear = "entry_year"
	// FieldEntryMonth holds the string denoting the entry_month field in the database.
	FieldEntryMonth = "entry_month"
	// FieldEntryDay holds the string denoting the entry_day field in the database.
	FieldEntryDay = "entry_day"
	// FieldExit holds the string denoting the exit field in the database.
	FieldExit = "exit"
	// FieldExitDate holds the string denoting the exit_date field in the database.
	FieldExitDate = "exit_date"
	// FieldExitTime holds the string denoting the exit_time field in the database.
	FieldExitTime = "exit_time"
	// FieldTotalPrice holds the string denoting the total_price field in the database.
	FieldTotalPrice = "total_price"
	// FieldDiscountPrice holds the string denoting the discount_price field in the database.
	FieldDiscountPrice = "discount_price"
	// FieldPaidPrice holds the string denoting the paid_price field in the database.
	FieldPaidPrice = "paid_price"
	// FieldCarType holds the string denoting the car_type field in the database.
	FieldCarType = "car_type"
	// FieldCarNumber holds the string denoting the car_number field in the database.
	FieldCarNumber = "car_number"
	// FieldCardNumber holds the string denoting the card_number field in the database.
	FieldCardNumber = "card_number"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the etcrecord in the database.
	Table = "etc_record"
)

// Columns holds all SQL columns for etcrecord fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldEntry,
	FieldEntryYear,
	FieldEntryMonth,
	FieldEntryDay,
	FieldExit,
	FieldExitDate,
	FieldExitTime,
	FieldTotalPrice,
	FieldDiscountPrice,
	FieldPaidPrice,
	FieldCarType,
	FieldCarNumber,
	FieldCardNumber,
	FieldStatus,
	FieldComment,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
)
