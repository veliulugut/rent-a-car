// Code generated by ent, DO NOT EDIT.

package ent

import (
	"rent-a-car/ent/car"
	"rent-a-car/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescCreatedAt is the schema descriptor for created_at field.
	carDescCreatedAt := carFields[5].Descriptor()
	// car.DefaultCreatedAt holds the default value on creation for the created_at field.
	car.DefaultCreatedAt = carDescCreatedAt.Default.(time.Time)
	// carDescUpdatedAt is the schema descriptor for updated_at field.
	carDescUpdatedAt := carFields[6].Descriptor()
	// car.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	car.DefaultUpdatedAt = carDescUpdatedAt.Default.(time.Time)
}
