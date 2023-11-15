// Code generated by ent, DO NOT EDIT.

package car

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the car type in the database.
	Label = "car"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCarName holds the string denoting the car_name field in the database.
	FieldCarName = "car_name"
	// FieldMilesPerGallon holds the string denoting the miles_per_gallon field in the database.
	FieldMilesPerGallon = "miles_per_gallon"
	// FieldCylinders holds the string denoting the cylinders field in the database.
	FieldCylinders = "cylinders"
	// FieldPower holds the string denoting the power field in the database.
	FieldPower = "power"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the car in the database.
	Table = "cars"
)

// Columns holds all SQL columns for car fields.
var Columns = []string{
	FieldID,
	FieldCarName,
	FieldMilesPerGallon,
	FieldCylinders,
	FieldPower,
	FieldYear,
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

// OrderOption defines the ordering options for the Car queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCarName orders the results by the car_name field.
func ByCarName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCarName, opts...).ToFunc()
}

// ByMilesPerGallon orders the results by the miles_per_gallon field.
func ByMilesPerGallon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMilesPerGallon, opts...).ToFunc()
}

// ByCylinders orders the results by the cylinders field.
func ByCylinders(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCylinders, opts...).ToFunc()
}

// ByPower orders the results by the power field.
func ByPower(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPower, opts...).ToFunc()
}

// ByYear orders the results by the year field.
func ByYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYear, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
