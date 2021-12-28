// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/artmisxyz/legolas/ent/event"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolmint"
)

// UniswapV3PoolMint is the model entity for the UniswapV3PoolMint schema.
type UniswapV3PoolMint struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Sender holds the value of the "sender" field.
	Sender string `json:"sender,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// TickLower holds the value of the "tick_lower" field.
	TickLower string `json:"tick_lower,omitempty"`
	// TickUpper holds the value of the "tick_upper" field.
	TickUpper string `json:"tick_upper,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount string `json:"amount,omitempty"`
	// Amount0 holds the value of the "amount0" field.
	Amount0 string `json:"amount0,omitempty"`
	// Amount1 holds the value of the "amount1" field.
	Amount1 string `json:"amount1,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UniswapV3PoolMintQuery when eager-loading is set.
	Edges    UniswapV3PoolMintEdges `json:"edges"`
	event_id *int
}

// UniswapV3PoolMintEdges holds the relations/edges for other nodes in the graph.
type UniswapV3PoolMintEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UniswapV3PoolMintEdges) EventOrErr() (*Event, error) {
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
func (*UniswapV3PoolMint) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case uniswapv3poolmint.FieldID:
			values[i] = new(sql.NullInt64)
		case uniswapv3poolmint.FieldSender, uniswapv3poolmint.FieldOwner, uniswapv3poolmint.FieldTickLower, uniswapv3poolmint.FieldTickUpper, uniswapv3poolmint.FieldAmount, uniswapv3poolmint.FieldAmount0, uniswapv3poolmint.FieldAmount1:
			values[i] = new(sql.NullString)
		case uniswapv3poolmint.ForeignKeys[0]: // event_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UniswapV3PoolMint", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UniswapV3PoolMint fields.
func (uvm *UniswapV3PoolMint) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case uniswapv3poolmint.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uvm.ID = int(value.Int64)
		case uniswapv3poolmint.FieldSender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender", values[i])
			} else if value.Valid {
				uvm.Sender = value.String
			}
		case uniswapv3poolmint.FieldOwner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner", values[i])
			} else if value.Valid {
				uvm.Owner = value.String
			}
		case uniswapv3poolmint.FieldTickLower:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tick_lower", values[i])
			} else if value.Valid {
				uvm.TickLower = value.String
			}
		case uniswapv3poolmint.FieldTickUpper:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tick_upper", values[i])
			} else if value.Valid {
				uvm.TickUpper = value.String
			}
		case uniswapv3poolmint.FieldAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				uvm.Amount = value.String
			}
		case uniswapv3poolmint.FieldAmount0:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field amount0", values[i])
			} else if value.Valid {
				uvm.Amount0 = value.String
			}
		case uniswapv3poolmint.FieldAmount1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field amount1", values[i])
			} else if value.Valid {
				uvm.Amount1 = value.String
			}
		case uniswapv3poolmint.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field event_id", value)
			} else if value.Valid {
				uvm.event_id = new(int)
				*uvm.event_id = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryEvent queries the "event" edge of the UniswapV3PoolMint entity.
func (uvm *UniswapV3PoolMint) QueryEvent() *EventQuery {
	return (&UniswapV3PoolMintClient{config: uvm.config}).QueryEvent(uvm)
}

// Update returns a builder for updating this UniswapV3PoolMint.
// Note that you need to call UniswapV3PoolMint.Unwrap() before calling this method if this UniswapV3PoolMint
// was returned from a transaction, and the transaction was committed or rolled back.
func (uvm *UniswapV3PoolMint) Update() *UniswapV3PoolMintUpdateOne {
	return (&UniswapV3PoolMintClient{config: uvm.config}).UpdateOne(uvm)
}

// Unwrap unwraps the UniswapV3PoolMint entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uvm *UniswapV3PoolMint) Unwrap() *UniswapV3PoolMint {
	tx, ok := uvm.config.driver.(*txDriver)
	if !ok {
		panic("ent: UniswapV3PoolMint is not a transactional entity")
	}
	uvm.config.driver = tx.drv
	return uvm
}

// String implements the fmt.Stringer.
func (uvm *UniswapV3PoolMint) String() string {
	var builder strings.Builder
	builder.WriteString("UniswapV3PoolMint(")
	builder.WriteString(fmt.Sprintf("id=%v", uvm.ID))
	builder.WriteString(", sender=")
	builder.WriteString(uvm.Sender)
	builder.WriteString(", owner=")
	builder.WriteString(uvm.Owner)
	builder.WriteString(", tick_lower=")
	builder.WriteString(uvm.TickLower)
	builder.WriteString(", tick_upper=")
	builder.WriteString(uvm.TickUpper)
	builder.WriteString(", amount=")
	builder.WriteString(uvm.Amount)
	builder.WriteString(", amount0=")
	builder.WriteString(uvm.Amount0)
	builder.WriteString(", amount1=")
	builder.WriteString(uvm.Amount1)
	builder.WriteByte(')')
	return builder.String()
}

// UniswapV3PoolMints is a parsable slice of UniswapV3PoolMint.
type UniswapV3PoolMints []*UniswapV3PoolMint

func (uvm UniswapV3PoolMints) config(cfg config) {
	for _i := range uvm {
		uvm[_i].config = cfg
	}
}
