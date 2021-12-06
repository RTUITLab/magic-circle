// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/Magic-Circle/ent/direction"
)

// Direction is the model entity for the Direction schema.
type Direction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DirectionQuery when eager-loading is set.
	Edges DirectionEdges `json:"edges"`
}

// DirectionEdges holds the relations/edges for other nodes in the graph.
type DirectionEdges struct {
	// Variants holds the value of the Variants edge.
	Variants []*Variant `json:"Variants,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VariantsOrErr returns the Variants value or an error if the edge
// was not loaded in eager-loading.
func (e DirectionEdges) VariantsOrErr() ([]*Variant, error) {
	if e.loadedTypes[0] {
		return e.Variants, nil
	}
	return nil, &NotLoadedError{edge: "Variants"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Direction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case direction.FieldID:
			values[i] = new(sql.NullInt64)
		case direction.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Direction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Direction fields.
func (d *Direction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case direction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case direction.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		}
	}
	return nil
}

// QueryVariants queries the "Variants" edge of the Direction entity.
func (d *Direction) QueryVariants() *VariantQuery {
	return (&DirectionClient{config: d.config}).QueryVariants(d)
}

// Update returns a builder for updating this Direction.
// Note that you need to call Direction.Unwrap() before calling this method if this Direction
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Direction) Update() *DirectionUpdateOne {
	return (&DirectionClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Direction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Direction) Unwrap() *Direction {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Direction is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Direction) String() string {
	var builder strings.Builder
	builder.WriteString("Direction(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", name=")
	builder.WriteString(d.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Directions is a parsable slice of Direction.
type Directions []*Direction

func (d Directions) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
