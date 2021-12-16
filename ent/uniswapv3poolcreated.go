// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/artmisxyz/blockinspector/ent/event"
	"github.com/artmisxyz/blockinspector/ent/schema"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3poolcreated"
)

// UniswapV3PoolCreated is the model entity for the UniswapV3PoolCreated schema.
type UniswapV3PoolCreated struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Token0 holds the value of the "token0" field.
	Token0 string `json:"token0,omitempty"`
	// Token1 holds the value of the "token1" field.
	Token1 string `json:"token1,omitempty"`
	// Fee holds the value of the "fee" field.
	Fee *schema.BigInt `json:"fee,omitempty"`
	// TickSpacing holds the value of the "tick_spacing" field.
	TickSpacing *schema.BigInt `json:"tick_spacing,omitempty"`
	// Pool holds the value of the "pool" field.
	Pool string `json:"pool,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UniswapV3PoolCreatedQuery when eager-loading is set.
	Edges              UniswapV3PoolCreatedEdges `json:"edges"`
	event_pool_created *int
}

// UniswapV3PoolCreatedEdges holds the relations/edges for other nodes in the graph.
type UniswapV3PoolCreatedEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UniswapV3PoolCreatedEdges) EventOrErr() (*Event, error) {
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
func (*UniswapV3PoolCreated) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case uniswapv3poolcreated.FieldFee, uniswapv3poolcreated.FieldTickSpacing:
			values[i] = new(schema.BigInt)
		case uniswapv3poolcreated.FieldID:
			values[i] = new(sql.NullInt64)
		case uniswapv3poolcreated.FieldToken0, uniswapv3poolcreated.FieldToken1, uniswapv3poolcreated.FieldPool:
			values[i] = new(sql.NullString)
		case uniswapv3poolcreated.ForeignKeys[0]: // event_pool_created
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UniswapV3PoolCreated", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UniswapV3PoolCreated fields.
func (uvc *UniswapV3PoolCreated) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case uniswapv3poolcreated.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uvc.ID = int(value.Int64)
		case uniswapv3poolcreated.FieldToken0:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token0", values[i])
			} else if value.Valid {
				uvc.Token0 = value.String
			}
		case uniswapv3poolcreated.FieldToken1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token1", values[i])
			} else if value.Valid {
				uvc.Token1 = value.String
			}
		case uniswapv3poolcreated.FieldFee:
			if value, ok := values[i].(*schema.BigInt); !ok {
				return fmt.Errorf("unexpected type %T for field fee", values[i])
			} else if value != nil {
				uvc.Fee = value
			}
		case uniswapv3poolcreated.FieldTickSpacing:
			if value, ok := values[i].(*schema.BigInt); !ok {
				return fmt.Errorf("unexpected type %T for field tick_spacing", values[i])
			} else if value != nil {
				uvc.TickSpacing = value
			}
		case uniswapv3poolcreated.FieldPool:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pool", values[i])
			} else if value.Valid {
				uvc.Pool = value.String
			}
		case uniswapv3poolcreated.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field event_pool_created", value)
			} else if value.Valid {
				uvc.event_pool_created = new(int)
				*uvc.event_pool_created = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryEvent queries the "event" edge of the UniswapV3PoolCreated entity.
func (uvc *UniswapV3PoolCreated) QueryEvent() *EventQuery {
	return (&UniswapV3PoolCreatedClient{config: uvc.config}).QueryEvent(uvc)
}

// Update returns a builder for updating this UniswapV3PoolCreated.
// Note that you need to call UniswapV3PoolCreated.Unwrap() before calling this method if this UniswapV3PoolCreated
// was returned from a transaction, and the transaction was committed or rolled back.
func (uvc *UniswapV3PoolCreated) Update() *UniswapV3PoolCreatedUpdateOne {
	return (&UniswapV3PoolCreatedClient{config: uvc.config}).UpdateOne(uvc)
}

// Unwrap unwraps the UniswapV3PoolCreated entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uvc *UniswapV3PoolCreated) Unwrap() *UniswapV3PoolCreated {
	tx, ok := uvc.config.driver.(*txDriver)
	if !ok {
		panic("ent: UniswapV3PoolCreated is not a transactional entity")
	}
	uvc.config.driver = tx.drv
	return uvc
}

// String implements the fmt.Stringer.
func (uvc *UniswapV3PoolCreated) String() string {
	var builder strings.Builder
	builder.WriteString("UniswapV3PoolCreated(")
	builder.WriteString(fmt.Sprintf("id=%v", uvc.ID))
	builder.WriteString(", token0=")
	builder.WriteString(uvc.Token0)
	builder.WriteString(", token1=")
	builder.WriteString(uvc.Token1)
	builder.WriteString(", fee=")
	builder.WriteString(fmt.Sprintf("%v", uvc.Fee))
	builder.WriteString(", tick_spacing=")
	builder.WriteString(fmt.Sprintf("%v", uvc.TickSpacing))
	builder.WriteString(", pool=")
	builder.WriteString(uvc.Pool)
	builder.WriteByte(')')
	return builder.String()
}

// UniswapV3PoolCreateds is a parsable slice of UniswapV3PoolCreated.
type UniswapV3PoolCreateds []*UniswapV3PoolCreated

func (uvc UniswapV3PoolCreateds) config(cfg config) {
	for _i := range uvc {
		uvc[_i].config = cfg
	}
}
