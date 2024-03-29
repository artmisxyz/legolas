// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/artmisxyz/legolas/ent/event"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolflash"
)

// UniswapV3PoolFlash is the model entity for the UniswapV3PoolFlash schema.
type UniswapV3PoolFlash struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Sender holds the value of the "sender" field.
	Sender string `json:"sender,omitempty"`
	// Recipient holds the value of the "recipient" field.
	Recipient string `json:"recipient,omitempty"`
	// Amount0 holds the value of the "amount0" field.
	Amount0 string `json:"amount0,omitempty"`
	// Amount1 holds the value of the "amount1" field.
	Amount1 string `json:"amount1,omitempty"`
	// Paid0 holds the value of the "paid0" field.
	Paid0 string `json:"paid0,omitempty"`
	// Paid1 holds the value of the "paid1" field.
	Paid1 string `json:"paid1,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UniswapV3PoolFlashQuery when eager-loading is set.
	Edges    UniswapV3PoolFlashEdges `json:"edges"`
	event_id *int
}

// UniswapV3PoolFlashEdges holds the relations/edges for other nodes in the graph.
type UniswapV3PoolFlashEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UniswapV3PoolFlashEdges) EventOrErr() (*Event, error) {
	if e.loadedTypes[0] {
		if e.Event == nil {
			// The edge event was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: event.Label}
		}
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UniswapV3PoolFlash) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case uniswapv3poolflash.FieldID:
			values[i] = new(sql.NullInt64)
		case uniswapv3poolflash.FieldSender, uniswapv3poolflash.FieldRecipient, uniswapv3poolflash.FieldAmount0, uniswapv3poolflash.FieldAmount1, uniswapv3poolflash.FieldPaid0, uniswapv3poolflash.FieldPaid1:
			values[i] = new(sql.NullString)
		case uniswapv3poolflash.ForeignKeys[0]: // event_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UniswapV3PoolFlash", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UniswapV3PoolFlash fields.
func (uvf *UniswapV3PoolFlash) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case uniswapv3poolflash.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uvf.ID = int(value.Int64)
		case uniswapv3poolflash.FieldSender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender", values[i])
			} else if value.Valid {
				uvf.Sender = value.String
			}
		case uniswapv3poolflash.FieldRecipient:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field recipient", values[i])
			} else if value.Valid {
				uvf.Recipient = value.String
			}
		case uniswapv3poolflash.FieldAmount0:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field amount0", values[i])
			} else if value.Valid {
				uvf.Amount0 = value.String
			}
		case uniswapv3poolflash.FieldAmount1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field amount1", values[i])
			} else if value.Valid {
				uvf.Amount1 = value.String
			}
		case uniswapv3poolflash.FieldPaid0:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paid0", values[i])
			} else if value.Valid {
				uvf.Paid0 = value.String
			}
		case uniswapv3poolflash.FieldPaid1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paid1", values[i])
			} else if value.Valid {
				uvf.Paid1 = value.String
			}
		case uniswapv3poolflash.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field event_id", value)
			} else if value.Valid {
				uvf.event_id = new(int)
				*uvf.event_id = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryEvent queries the "event" edge of the UniswapV3PoolFlash entity.
func (uvf *UniswapV3PoolFlash) QueryEvent() *EventQuery {
	return (&UniswapV3PoolFlashClient{config: uvf.config}).QueryEvent(uvf)
}

// Update returns a builder for updating this UniswapV3PoolFlash.
// Note that you need to call UniswapV3PoolFlash.Unwrap() before calling this method if this UniswapV3PoolFlash
// was returned from a transaction, and the transaction was committed or rolled back.
func (uvf *UniswapV3PoolFlash) Update() *UniswapV3PoolFlashUpdateOne {
	return (&UniswapV3PoolFlashClient{config: uvf.config}).UpdateOne(uvf)
}

// Unwrap unwraps the UniswapV3PoolFlash entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uvf *UniswapV3PoolFlash) Unwrap() *UniswapV3PoolFlash {
	tx, ok := uvf.config.driver.(*txDriver)
	if !ok {
		panic("ent: UniswapV3PoolFlash is not a transactional entity")
	}
	uvf.config.driver = tx.drv
	return uvf
}

// String implements the fmt.Stringer.
func (uvf *UniswapV3PoolFlash) String() string {
	var builder strings.Builder
	builder.WriteString("UniswapV3PoolFlash(")
	builder.WriteString(fmt.Sprintf("id=%v", uvf.ID))
	builder.WriteString(", sender=")
	builder.WriteString(uvf.Sender)
	builder.WriteString(", recipient=")
	builder.WriteString(uvf.Recipient)
	builder.WriteString(", amount0=")
	builder.WriteString(uvf.Amount0)
	builder.WriteString(", amount1=")
	builder.WriteString(uvf.Amount1)
	builder.WriteString(", paid0=")
	builder.WriteString(uvf.Paid0)
	builder.WriteString(", paid1=")
	builder.WriteString(uvf.Paid1)
	builder.WriteByte(')')
	return builder.String()
}

// UniswapV3PoolFlashes is a parsable slice of UniswapV3PoolFlash.
type UniswapV3PoolFlashes []*UniswapV3PoolFlash

func (uvf UniswapV3PoolFlashes) config(cfg config) {
	for _i := range uvf {
		uvf[_i].config = cfg
	}
}
