// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/Magic-Circle/ent/direction"
	"github.com/0B1t322/Magic-Circle/ent/predicate"
	"github.com/0B1t322/Magic-Circle/ent/variant"
)

// DirectionUpdate is the builder for updating Direction entities.
type DirectionUpdate struct {
	config
	hooks    []Hook
	mutation *DirectionMutation
}

// Where appends a list predicates to the DirectionUpdate builder.
func (du *DirectionUpdate) Where(ps ...predicate.Direction) *DirectionUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetName sets the "name" field.
func (du *DirectionUpdate) SetName(s string) *DirectionUpdate {
	du.mutation.SetName(s)
	return du
}

// AddVariantIDs adds the "Variants" edge to the Variant entity by IDs.
func (du *DirectionUpdate) AddVariantIDs(ids ...int) *DirectionUpdate {
	du.mutation.AddVariantIDs(ids...)
	return du
}

// AddVariants adds the "Variants" edges to the Variant entity.
func (du *DirectionUpdate) AddVariants(v ...*Variant) *DirectionUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return du.AddVariantIDs(ids...)
}

// Mutation returns the DirectionMutation object of the builder.
func (du *DirectionUpdate) Mutation() *DirectionMutation {
	return du.mutation
}

// ClearVariants clears all "Variants" edges to the Variant entity.
func (du *DirectionUpdate) ClearVariants() *DirectionUpdate {
	du.mutation.ClearVariants()
	return du
}

// RemoveVariantIDs removes the "Variants" edge to Variant entities by IDs.
func (du *DirectionUpdate) RemoveVariantIDs(ids ...int) *DirectionUpdate {
	du.mutation.RemoveVariantIDs(ids...)
	return du
}

// RemoveVariants removes "Variants" edges to Variant entities.
func (du *DirectionUpdate) RemoveVariants(v ...*Variant) *DirectionUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return du.RemoveVariantIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DirectionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DirectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DirectionUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DirectionUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DirectionUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DirectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   direction.Table,
			Columns: direction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: direction.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: direction.FieldName,
		})
	}
	if du.mutation.VariantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   direction.VariantsTable,
			Columns: []string{direction.VariantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedVariantsIDs(); len(nodes) > 0 && !du.mutation.VariantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   direction.VariantsTable,
			Columns: []string{direction.VariantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.VariantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   direction.VariantsTable,
			Columns: []string{direction.VariantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{direction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DirectionUpdateOne is the builder for updating a single Direction entity.
type DirectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DirectionMutation
}

// SetName sets the "name" field.
func (duo *DirectionUpdateOne) SetName(s string) *DirectionUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// AddVariantIDs adds the "Variants" edge to the Variant entity by IDs.
func (duo *DirectionUpdateOne) AddVariantIDs(ids ...int) *DirectionUpdateOne {
	duo.mutation.AddVariantIDs(ids...)
	return duo
}

// AddVariants adds the "Variants" edges to the Variant entity.
func (duo *DirectionUpdateOne) AddVariants(v ...*Variant) *DirectionUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return duo.AddVariantIDs(ids...)
}

// Mutation returns the DirectionMutation object of the builder.
func (duo *DirectionUpdateOne) Mutation() *DirectionMutation {
	return duo.mutation
}

// ClearVariants clears all "Variants" edges to the Variant entity.
func (duo *DirectionUpdateOne) ClearVariants() *DirectionUpdateOne {
	duo.mutation.ClearVariants()
	return duo
}

// RemoveVariantIDs removes the "Variants" edge to Variant entities by IDs.
func (duo *DirectionUpdateOne) RemoveVariantIDs(ids ...int) *DirectionUpdateOne {
	duo.mutation.RemoveVariantIDs(ids...)
	return duo
}

// RemoveVariants removes "Variants" edges to Variant entities.
func (duo *DirectionUpdateOne) RemoveVariants(v ...*Variant) *DirectionUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return duo.RemoveVariantIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DirectionUpdateOne) Select(field string, fields ...string) *DirectionUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Direction entity.
func (duo *DirectionUpdateOne) Save(ctx context.Context) (*Direction, error) {
	var (
		err  error
		node *Direction
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DirectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DirectionUpdateOne) SaveX(ctx context.Context) *Direction {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DirectionUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DirectionUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DirectionUpdateOne) sqlSave(ctx context.Context) (_node *Direction, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   direction.Table,
			Columns: direction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: direction.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Direction.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, direction.FieldID)
		for _, f := range fields {
			if !direction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != direction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: direction.FieldName,
		})
	}
	if duo.mutation.VariantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   direction.VariantsTable,
			Columns: []string{direction.VariantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedVariantsIDs(); len(nodes) > 0 && !duo.mutation.VariantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   direction.VariantsTable,
			Columns: []string{direction.VariantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.VariantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   direction.VariantsTable,
			Columns: []string{direction.VariantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Direction{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{direction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
