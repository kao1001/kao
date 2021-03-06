// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pmn-kao/app/ent/degree"
	"github.com/pmn-kao/app/ent/department"
	"github.com/pmn-kao/app/ent/doctor"
	"github.com/pmn-kao/app/ent/nametitle"
	"github.com/pmn-kao/app/ent/predicate"
)

// DoctorUpdate is the builder for updating Doctor entities.
type DoctorUpdate struct {
	config
	hooks      []Hook
	mutation   *DoctorMutation
	predicates []predicate.Doctor
}

// Where adds a new predicate for the builder.
func (du *DoctorUpdate) Where(ps ...predicate.Doctor) *DoctorUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetEmail sets the Email field.
func (du *DoctorUpdate) SetEmail(s string) *DoctorUpdate {
	du.mutation.SetEmail(s)
	return du
}

// SetPassword sets the Password field.
func (du *DoctorUpdate) SetPassword(s string) *DoctorUpdate {
	du.mutation.SetPassword(s)
	return du
}

// SetName sets the Name field.
func (du *DoctorUpdate) SetName(s string) *DoctorUpdate {
	du.mutation.SetName(s)
	return du
}

// SetTel sets the Tel field.
func (du *DoctorUpdate) SetTel(s string) *DoctorUpdate {
	du.mutation.SetTel(s)
	return du
}

// SetDepartmentID sets the department edge to Department by id.
func (du *DoctorUpdate) SetDepartmentID(id int) *DoctorUpdate {
	du.mutation.SetDepartmentID(id)
	return du
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (du *DoctorUpdate) SetNillableDepartmentID(id *int) *DoctorUpdate {
	if id != nil {
		du = du.SetDepartmentID(*id)
	}
	return du
}

// SetDepartment sets the department edge to Department.
func (du *DoctorUpdate) SetDepartment(d *Department) *DoctorUpdate {
	return du.SetDepartmentID(d.ID)
}

// SetDegreeID sets the degree edge to Degree by id.
func (du *DoctorUpdate) SetDegreeID(id int) *DoctorUpdate {
	du.mutation.SetDegreeID(id)
	return du
}

// SetNillableDegreeID sets the degree edge to Degree by id if the given value is not nil.
func (du *DoctorUpdate) SetNillableDegreeID(id *int) *DoctorUpdate {
	if id != nil {
		du = du.SetDegreeID(*id)
	}
	return du
}

// SetDegree sets the degree edge to Degree.
func (du *DoctorUpdate) SetDegree(d *Degree) *DoctorUpdate {
	return du.SetDegreeID(d.ID)
}

// SetNametitleID sets the nametitle edge to Nametitle by id.
func (du *DoctorUpdate) SetNametitleID(id int) *DoctorUpdate {
	du.mutation.SetNametitleID(id)
	return du
}

// SetNillableNametitleID sets the nametitle edge to Nametitle by id if the given value is not nil.
func (du *DoctorUpdate) SetNillableNametitleID(id *int) *DoctorUpdate {
	if id != nil {
		du = du.SetNametitleID(*id)
	}
	return du
}

// SetNametitle sets the nametitle edge to Nametitle.
func (du *DoctorUpdate) SetNametitle(n *Nametitle) *DoctorUpdate {
	return du.SetNametitleID(n.ID)
}

// Mutation returns the DoctorMutation object of the builder.
func (du *DoctorUpdate) Mutation() *DoctorMutation {
	return du.mutation
}

// ClearDepartment clears the department edge to Department.
func (du *DoctorUpdate) ClearDepartment() *DoctorUpdate {
	du.mutation.ClearDepartment()
	return du
}

// ClearDegree clears the degree edge to Degree.
func (du *DoctorUpdate) ClearDegree() *DoctorUpdate {
	du.mutation.ClearDegree()
	return du
}

// ClearNametitle clears the nametitle edge to Nametitle.
func (du *DoctorUpdate) ClearNametitle() *DoctorUpdate {
	du.mutation.ClearNametitle()
	return du
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DoctorUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.Email(); ok {
		if err := doctor.EmailValidator(v); err != nil {
			return 0, &ValidationError{Name: "Email", err: fmt.Errorf("ent: validator failed for field \"Email\": %w", err)}
		}
	}
	if v, ok := du.mutation.Password(); ok {
		if err := doctor.PasswordValidator(v); err != nil {
			return 0, &ValidationError{Name: "Password", err: fmt.Errorf("ent: validator failed for field \"Password\": %w", err)}
		}
	}
	if v, ok := du.mutation.Name(); ok {
		if err := doctor.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}
	if v, ok := du.mutation.Tel(); ok {
		if err := doctor.TelValidator(v); err != nil {
			return 0, &ValidationError{Name: "Tel", err: fmt.Errorf("ent: validator failed for field \"Tel\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DoctorUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DoctorUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DoctorUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DoctorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctor.Table,
			Columns: doctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldEmail,
		})
	}
	if value, ok := du.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldPassword,
		})
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldName,
		})
	}
	if value, ok := du.mutation.Tel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldTel,
		})
	}
	if du.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.DepartmentTable,
			Columns: []string{doctor.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.DepartmentTable,
			Columns: []string{doctor.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.DegreeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.DegreeTable,
			Columns: []string{doctor.DegreeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: degree.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DegreeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.DegreeTable,
			Columns: []string{doctor.DegreeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: degree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.NametitleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.NametitleTable,
			Columns: []string{doctor.NametitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nametitle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.NametitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.NametitleTable,
			Columns: []string{doctor.NametitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nametitle.FieldID,
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
			err = &NotFoundError{doctor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DoctorUpdateOne is the builder for updating a single Doctor entity.
type DoctorUpdateOne struct {
	config
	hooks    []Hook
	mutation *DoctorMutation
}

// SetEmail sets the Email field.
func (duo *DoctorUpdateOne) SetEmail(s string) *DoctorUpdateOne {
	duo.mutation.SetEmail(s)
	return duo
}

// SetPassword sets the Password field.
func (duo *DoctorUpdateOne) SetPassword(s string) *DoctorUpdateOne {
	duo.mutation.SetPassword(s)
	return duo
}

// SetName sets the Name field.
func (duo *DoctorUpdateOne) SetName(s string) *DoctorUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetTel sets the Tel field.
func (duo *DoctorUpdateOne) SetTel(s string) *DoctorUpdateOne {
	duo.mutation.SetTel(s)
	return duo
}

// SetDepartmentID sets the department edge to Department by id.
func (duo *DoctorUpdateOne) SetDepartmentID(id int) *DoctorUpdateOne {
	duo.mutation.SetDepartmentID(id)
	return duo
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableDepartmentID(id *int) *DoctorUpdateOne {
	if id != nil {
		duo = duo.SetDepartmentID(*id)
	}
	return duo
}

// SetDepartment sets the department edge to Department.
func (duo *DoctorUpdateOne) SetDepartment(d *Department) *DoctorUpdateOne {
	return duo.SetDepartmentID(d.ID)
}

// SetDegreeID sets the degree edge to Degree by id.
func (duo *DoctorUpdateOne) SetDegreeID(id int) *DoctorUpdateOne {
	duo.mutation.SetDegreeID(id)
	return duo
}

// SetNillableDegreeID sets the degree edge to Degree by id if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableDegreeID(id *int) *DoctorUpdateOne {
	if id != nil {
		duo = duo.SetDegreeID(*id)
	}
	return duo
}

// SetDegree sets the degree edge to Degree.
func (duo *DoctorUpdateOne) SetDegree(d *Degree) *DoctorUpdateOne {
	return duo.SetDegreeID(d.ID)
}

// SetNametitleID sets the nametitle edge to Nametitle by id.
func (duo *DoctorUpdateOne) SetNametitleID(id int) *DoctorUpdateOne {
	duo.mutation.SetNametitleID(id)
	return duo
}

// SetNillableNametitleID sets the nametitle edge to Nametitle by id if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableNametitleID(id *int) *DoctorUpdateOne {
	if id != nil {
		duo = duo.SetNametitleID(*id)
	}
	return duo
}

// SetNametitle sets the nametitle edge to Nametitle.
func (duo *DoctorUpdateOne) SetNametitle(n *Nametitle) *DoctorUpdateOne {
	return duo.SetNametitleID(n.ID)
}

// Mutation returns the DoctorMutation object of the builder.
func (duo *DoctorUpdateOne) Mutation() *DoctorMutation {
	return duo.mutation
}

// ClearDepartment clears the department edge to Department.
func (duo *DoctorUpdateOne) ClearDepartment() *DoctorUpdateOne {
	duo.mutation.ClearDepartment()
	return duo
}

// ClearDegree clears the degree edge to Degree.
func (duo *DoctorUpdateOne) ClearDegree() *DoctorUpdateOne {
	duo.mutation.ClearDegree()
	return duo
}

// ClearNametitle clears the nametitle edge to Nametitle.
func (duo *DoctorUpdateOne) ClearNametitle() *DoctorUpdateOne {
	duo.mutation.ClearNametitle()
	return duo
}

// Save executes the query and returns the updated entity.
func (duo *DoctorUpdateOne) Save(ctx context.Context) (*Doctor, error) {
	if v, ok := duo.mutation.Email(); ok {
		if err := doctor.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "Email", err: fmt.Errorf("ent: validator failed for field \"Email\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Password(); ok {
		if err := doctor.PasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "Password", err: fmt.Errorf("ent: validator failed for field \"Password\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Name(); ok {
		if err := doctor.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Tel(); ok {
		if err := doctor.TelValidator(v); err != nil {
			return nil, &ValidationError{Name: "Tel", err: fmt.Errorf("ent: validator failed for field \"Tel\": %w", err)}
		}
	}

	var (
		err  error
		node *Doctor
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DoctorUpdateOne) SaveX(ctx context.Context) *Doctor {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DoctorUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DoctorUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DoctorUpdateOne) sqlSave(ctx context.Context) (d *Doctor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctor.Table,
			Columns: doctor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Doctor.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldEmail,
		})
	}
	if value, ok := duo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldPassword,
		})
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldName,
		})
	}
	if value, ok := duo.mutation.Tel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldTel,
		})
	}
	if duo.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.DepartmentTable,
			Columns: []string{doctor.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.DepartmentTable,
			Columns: []string{doctor.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.DegreeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.DegreeTable,
			Columns: []string{doctor.DegreeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: degree.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DegreeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.DegreeTable,
			Columns: []string{doctor.DegreeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: degree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.NametitleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.NametitleTable,
			Columns: []string{doctor.NametitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nametitle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.NametitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.NametitleTable,
			Columns: []string{doctor.NametitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nametitle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Doctor{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
