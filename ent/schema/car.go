package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("car_name"),
		field.String("miles_per_gallon"),
		field.String("cylinders"),
		field.String("power"),
		field.String("year"),
		field.Time("created_at").Default(time.Now().UTC()),
		field.Time("updated_at").Default(time.Now().UTC()),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return nil
}
