// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pmn-kao/app/ent/degree"
	"github.com/pmn-kao/app/ent/department"
	"github.com/pmn-kao/app/ent/doctor"
	"github.com/pmn-kao/app/ent/nametitle"
)

// Doctor is the model entity for the Doctor schema.
type Doctor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "Email" field.
	Email string `json:"Email,omitempty"`
	// Password holds the value of the "Password" field.
	Password string `json:"Password,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Tel holds the value of the "Tel" field.
	Tel string `json:"Tel,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DoctorQuery when eager-loading is set.
	Edges         DoctorEdges `json:"edges"`
	degree_id     *int
	department_id *int
	doctor_id     *int
}

// DoctorEdges holds the relations/edges for other nodes in the graph.
type DoctorEdges struct {
	// Department holds the value of the department edge.
	Department *Department
	// Degree holds the value of the degree edge.
	Degree *Degree
	// Nametitle holds the value of the nametitle edge.
	Nametitle *Nametitle
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DepartmentOrErr returns the Department value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DoctorEdges) DepartmentOrErr() (*Department, error) {
	if e.loadedTypes[0] {
		if e.Department == nil {
			// The edge department was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.Department, nil
	}
	return nil, &NotLoadedError{edge: "department"}
}

// DegreeOrErr returns the Degree value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DoctorEdges) DegreeOrErr() (*Degree, error) {
	if e.loadedTypes[1] {
		if e.Degree == nil {
			// The edge degree was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: degree.Label}
		}
		return e.Degree, nil
	}
	return nil, &NotLoadedError{edge: "degree"}
}

// NametitleOrErr returns the Nametitle value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DoctorEdges) NametitleOrErr() (*Nametitle, error) {
	if e.loadedTypes[2] {
		if e.Nametitle == nil {
			// The edge nametitle was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: nametitle.Label}
		}
		return e.Nametitle, nil
	}
	return nil, &NotLoadedError{edge: "nametitle"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Doctor) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Email
		&sql.NullString{}, // Password
		&sql.NullString{}, // Name
		&sql.NullString{}, // Tel
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Doctor) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // degree_id
		&sql.NullInt64{}, // department_id
		&sql.NullInt64{}, // doctor_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Doctor fields.
func (d *Doctor) assignValues(values ...interface{}) error {
	if m, n := len(values), len(doctor.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Email", values[0])
	} else if value.Valid {
		d.Email = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Password", values[1])
	} else if value.Valid {
		d.Password = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Name", values[2])
	} else if value.Valid {
		d.Name = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Tel", values[3])
	} else if value.Valid {
		d.Tel = value.String
	}
	values = values[4:]
	if len(values) == len(doctor.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field degree_id", value)
		} else if value.Valid {
			d.degree_id = new(int)
			*d.degree_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field department_id", value)
		} else if value.Valid {
			d.department_id = new(int)
			*d.department_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field doctor_id", value)
		} else if value.Valid {
			d.doctor_id = new(int)
			*d.doctor_id = int(value.Int64)
		}
	}
	return nil
}

// QueryDepartment queries the department edge of the Doctor.
func (d *Doctor) QueryDepartment() *DepartmentQuery {
	return (&DoctorClient{config: d.config}).QueryDepartment(d)
}

// QueryDegree queries the degree edge of the Doctor.
func (d *Doctor) QueryDegree() *DegreeQuery {
	return (&DoctorClient{config: d.config}).QueryDegree(d)
}

// QueryNametitle queries the nametitle edge of the Doctor.
func (d *Doctor) QueryNametitle() *NametitleQuery {
	return (&DoctorClient{config: d.config}).QueryNametitle(d)
}

// Update returns a builder for updating this Doctor.
// Note that, you need to call Doctor.Unwrap() before calling this method, if this Doctor
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Doctor) Update() *DoctorUpdateOne {
	return (&DoctorClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Doctor) Unwrap() *Doctor {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Doctor is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Doctor) String() string {
	var builder strings.Builder
	builder.WriteString("Doctor(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", Email=")
	builder.WriteString(d.Email)
	builder.WriteString(", Password=")
	builder.WriteString(d.Password)
	builder.WriteString(", Name=")
	builder.WriteString(d.Name)
	builder.WriteString(", Tel=")
	builder.WriteString(d.Tel)
	builder.WriteByte(')')
	return builder.String()
}

// Doctors is a parsable slice of Doctor.
type Doctors []*Doctor

func (d Doctors) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
