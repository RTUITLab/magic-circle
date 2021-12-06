// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/Magic-Circle/ent/adjacenttable"
	"github.com/0B1t322/Magic-Circle/ent/direction"
	"github.com/0B1t322/Magic-Circle/ent/institute"
	"github.com/0B1t322/Magic-Circle/ent/profile"
	"github.com/0B1t322/Magic-Circle/ent/variant"
)

// VariantCreate is the builder for creating a Variant entity.
type VariantCreate struct {
	config
	mutation *VariantMutation
	hooks    []Hook
}

// SetDirectionID sets the "direction_id" field.
func (vc *VariantCreate) SetDirectionID(i int) *VariantCreate {
	vc.mutation.SetDirectionID(i)
	return vc
}

// SetProfileID sets the "profile_id" field.
func (vc *VariantCreate) SetProfileID(i int) *VariantCreate {
	vc.mutation.SetProfileID(i)
	return vc
}

// SetInsituteID sets the "insitute_id" field.
func (vc *VariantCreate) SetInsituteID(i int) *VariantCreate {
	vc.mutation.SetInsituteID(i)
	return vc
}

// SetInsitute sets the "Insitute" edge to the Institute entity.
func (vc *VariantCreate) SetInsitute(i *Institute) *VariantCreate {
	return vc.SetInsituteID(i.ID)
}

// SetDirection sets the "Direction" edge to the Direction entity.
func (vc *VariantCreate) SetDirection(d *Direction) *VariantCreate {
	return vc.SetDirectionID(d.ID)
}

// SetProfile sets the "Profile" edge to the Profile entity.
func (vc *VariantCreate) SetProfile(p *Profile) *VariantCreate {
	return vc.SetProfileID(p.ID)
}

// AddAdjacentTableIDs adds the "AdjacentTables" edge to the AdjacentTable entity by IDs.
func (vc *VariantCreate) AddAdjacentTableIDs(ids ...int) *VariantCreate {
	vc.mutation.AddAdjacentTableIDs(ids...)
	return vc
}

// AddAdjacentTables adds the "AdjacentTables" edges to the AdjacentTable entity.
func (vc *VariantCreate) AddAdjacentTables(a ...*AdjacentTable) *VariantCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vc.AddAdjacentTableIDs(ids...)
}

// Mutation returns the VariantMutation object of the builder.
func (vc *VariantCreate) Mutation() *VariantMutation {
	return vc.mutation
}

// Save creates the Variant in the database.
func (vc *VariantCreate) Save(ctx context.Context) (*Variant, error) {
	var (
		err  error
		node *Variant
	)
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VariantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			if node, err = vc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			if vc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VariantCreate) SaveX(ctx context.Context) *Variant {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VariantCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VariantCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VariantCreate) check() error {
	if _, ok := vc.mutation.DirectionID(); !ok {
		return &ValidationError{Name: "direction_id", err: errors.New(`ent: missing required field "direction_id"`)}
	}
	if _, ok := vc.mutation.ProfileID(); !ok {
		return &ValidationError{Name: "profile_id", err: errors.New(`ent: missing required field "profile_id"`)}
	}
	if _, ok := vc.mutation.InsituteID(); !ok {
		return &ValidationError{Name: "insitute_id", err: errors.New(`ent: missing required field "insitute_id"`)}
	}
	if _, ok := vc.mutation.InsituteID(); !ok {
		return &ValidationError{Name: "Insitute", err: errors.New("ent: missing required edge \"Insitute\"")}
	}
	if _, ok := vc.mutation.DirectionID(); !ok {
		return &ValidationError{Name: "Direction", err: errors.New("ent: missing required edge \"Direction\"")}
	}
	if _, ok := vc.mutation.ProfileID(); !ok {
		return &ValidationError{Name: "Profile", err: errors.New("ent: missing required edge \"Profile\"")}
	}
	return nil
}

func (vc *VariantCreate) sqlSave(ctx context.Context) (*Variant, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vc *VariantCreate) createSpec() (*Variant, *sqlgraph.CreateSpec) {
	var (
		_node = &Variant{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: variant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: variant.FieldID,
			},
		}
	)
	if nodes := vc.mutation.InsituteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variant.InsituteTable,
			Columns: []string{variant.InsituteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: institute.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.InsituteID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.DirectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variant.DirectionTable,
			Columns: []string{variant.DirectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: direction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DirectionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variant.ProfileTable,
			Columns: []string{variant.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProfileID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.AdjacentTablesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   variant.AdjacentTablesTable,
			Columns: []string{variant.AdjacentTablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adjacenttable.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VariantCreateBulk is the builder for creating many Variant entities in bulk.
type VariantCreateBulk struct {
	config
	builders []*VariantCreate
}

// Save creates the Variant entities in the database.
func (vcb *VariantCreateBulk) Save(ctx context.Context) ([]*Variant, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Variant, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VariantMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VariantCreateBulk) SaveX(ctx context.Context) []*Variant {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VariantCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VariantCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
