// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationship"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationshipinfo"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
)

// Relationship is the model entity for the Relationship schema.
type Relationship struct {
	config `json:"-"`
	// Weight holds the value of the "weight" field.
	Weight int `json:"weight,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// RelativeID holds the value of the "relative_id" field.
	RelativeID int `json:"relative_id,omitempty"`
	// InfoID holds the value of the "info_id" field.
	InfoID int `json:"info_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RelationshipQuery when eager-loading is set.
	Edges RelationshipEdges `json:"edges"`
}

// RelationshipEdges holds the relations/edges for other nodes in the graph.
type RelationshipEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Relative holds the value of the relative edge.
	Relative *User `json:"relative,omitempty"`
	// Info holds the value of the info edge.
	Info *RelationshipInfo `json:"info,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RelationshipEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// RelativeOrErr returns the Relative value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RelationshipEdges) RelativeOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Relative == nil {
			// The edge relative was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Relative, nil
	}
	return nil, &NotLoadedError{edge: "relative"}
}

// InfoOrErr returns the Info value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RelationshipEdges) InfoOrErr() (*RelationshipInfo, error) {
	if e.loadedTypes[2] {
		if e.Info == nil {
			// The edge info was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: relationshipinfo.Label}
		}
		return e.Info, nil
	}
	return nil, &NotLoadedError{edge: "info"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Relationship) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case relationship.FieldWeight, relationship.FieldUserID, relationship.FieldRelativeID, relationship.FieldInfoID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Relationship", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Relationship fields.
func (r *Relationship) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case relationship.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				r.Weight = int(value.Int64)
			}
		case relationship.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				r.UserID = int(value.Int64)
			}
		case relationship.FieldRelativeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relative_id", values[i])
			} else if value.Valid {
				r.RelativeID = int(value.Int64)
			}
		case relationship.FieldInfoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field info_id", values[i])
			} else if value.Valid {
				r.InfoID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Relationship entity.
func (r *Relationship) QueryUser() *UserQuery {
	return (&RelationshipClient{config: r.config}).QueryUser(r)
}

// QueryRelative queries the "relative" edge of the Relationship entity.
func (r *Relationship) QueryRelative() *UserQuery {
	return (&RelationshipClient{config: r.config}).QueryRelative(r)
}

// QueryInfo queries the "info" edge of the Relationship entity.
func (r *Relationship) QueryInfo() *RelationshipInfoQuery {
	return (&RelationshipClient{config: r.config}).QueryInfo(r)
}

// Update returns a builder for updating this Relationship.
// Note that you need to call Relationship.Unwrap() before calling this method if this Relationship
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Relationship) Update() *RelationshipUpdateOne {
	return (&RelationshipClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Relationship entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Relationship) Unwrap() *Relationship {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Relationship is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Relationship) String() string {
	var builder strings.Builder
	builder.WriteString("Relationship(")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", r.Weight))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", r.UserID))
	builder.WriteString(", ")
	builder.WriteString("relative_id=")
	builder.WriteString(fmt.Sprintf("%v", r.RelativeID))
	builder.WriteString(", ")
	builder.WriteString("info_id=")
	builder.WriteString(fmt.Sprintf("%v", r.InfoID))
	builder.WriteByte(')')
	return builder.String()
}

// Relationships is a parsable slice of Relationship.
type Relationships []*Relationship

func (r Relationships) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
