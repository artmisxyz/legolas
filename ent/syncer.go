// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/artmisxyz/legolas/ent/syncer"
)

// Syncer is the model entity for the Syncer schema.
type Syncer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Start holds the value of the "start" field.
	Start uint64 `json:"start,omitempty"`
	// Finish holds the value of the "finish" field.
	Finish uint64 `json:"finish,omitempty"`
	// Current holds the value of the "current" field.
	Current uint64 `json:"current,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Syncer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case syncer.FieldID, syncer.FieldStart, syncer.FieldFinish, syncer.FieldCurrent:
			values[i] = new(sql.NullInt64)
		case syncer.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Syncer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Syncer fields.
func (s *Syncer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case syncer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case syncer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case syncer.FieldStart:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start", values[i])
			} else if value.Valid {
				s.Start = uint64(value.Int64)
			}
		case syncer.FieldFinish:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field finish", values[i])
			} else if value.Valid {
				s.Finish = uint64(value.Int64)
			}
		case syncer.FieldCurrent:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field current", values[i])
			} else if value.Valid {
				s.Current = uint64(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Syncer.
// Note that you need to call Syncer.Unwrap() before calling this method if this Syncer
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Syncer) Update() *SyncerUpdateOne {
	return (&SyncerClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Syncer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Syncer) Unwrap() *Syncer {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Syncer is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Syncer) String() string {
	var builder strings.Builder
	builder.WriteString("Syncer(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteString(", start=")
	builder.WriteString(fmt.Sprintf("%v", s.Start))
	builder.WriteString(", finish=")
	builder.WriteString(fmt.Sprintf("%v", s.Finish))
	builder.WriteString(", current=")
	builder.WriteString(fmt.Sprintf("%v", s.Current))
	builder.WriteByte(')')
	return builder.String()
}

// Syncers is a parsable slice of Syncer.
type Syncers []*Syncer

func (s Syncers) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
