// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/artmisxyz/legolas/ent/event"
	"github.com/artmisxyz/legolas/ent/schema"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolswap"
)

// UniswapV3PoolSwap is the model entity for the UniswapV3PoolSwap schema.
type UniswapV3PoolSwap struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Sender holds the value of the "sender" field.
	Sender string `json:"sender,omitempty"`
	// Recipient holds the value of the "recipient" field.
	Recipient string `json:"recipient,omitempty"`
	// Amount0 holds the value of the "amount0" field.
	Amount0 *schema.BigInt `json:"amount0,omitempty"`
	// Amount1 holds the value of the "amount1" field.
	Amount1 *schema.BigInt `json:"amount1,omitempty"`
	// SqrtPriceX96 holds the value of the "sqrt_price_x96" field.
	SqrtPriceX96 *schema.BigInt `json:"sqrt_price_x96,omitempty"`
	// Liquidity holds the value of the "liquidity" field.
	Liquidity *schema.BigInt `json:"liquidity,omitempty"`
	// Tick holds the value of the "tick" field.
	Tick *schema.BigInt `json:"tick,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UniswapV3PoolSwapQuery when eager-loading is set.
	Edges    UniswapV3PoolSwapEdges `json:"edges"`
	event_id *int
}

// UniswapV3PoolSwapEdges holds the relations/edges for other nodes in the graph.
type UniswapV3PoolSwapEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UniswapV3PoolSwapEdges) EventOrErr() (*Event, error) {
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
func (*UniswapV3PoolSwap) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case uniswapv3poolswap.FieldAmount0, uniswapv3poolswap.FieldAmount1, uniswapv3poolswap.FieldSqrtPriceX96, uniswapv3poolswap.FieldLiquidity, uniswapv3poolswap.FieldTick:
			values[i] = new(schema.BigInt)
		case uniswapv3poolswap.FieldID:
			values[i] = new(sql.NullInt64)
		case uniswapv3poolswap.FieldSender, uniswapv3poolswap.FieldRecipient:
			values[i] = new(sql.NullString)
		case uniswapv3poolswap.ForeignKeys[0]: // event_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UniswapV3PoolSwap", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UniswapV3PoolSwap fields.
func (uvs *UniswapV3PoolSwap) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case uniswapv3poolswap.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uvs.ID = int(value.Int64)
		case uniswapv3poolswap.FieldSender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender", values[i])
			} else if value.Valid {
				uvs.Sender = value.String
			}
		case uniswapv3poolswap.FieldRecipient:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field recipient", values[i])
			} else if value.Valid {
				uvs.Recipient = value.String
			}
		case uniswapv3poolswap.FieldAmount0:
			if value, ok := values[i].(*schema.BigInt); !ok {
				return fmt.Errorf("unexpected type %T for field amount0", values[i])
			} else if value != nil {
				uvs.Amount0 = value
			}
		case uniswapv3poolswap.FieldAmount1:
			if value, ok := values[i].(*schema.BigInt); !ok {
				return fmt.Errorf("unexpected type %T for field amount1", values[i])
			} else if value != nil {
				uvs.Amount1 = value
			}
		case uniswapv3poolswap.FieldSqrtPriceX96:
			if value, ok := values[i].(*schema.BigInt); !ok {
				return fmt.Errorf("unexpected type %T for field sqrt_price_x96", values[i])
			} else if value != nil {
				uvs.SqrtPriceX96 = value
			}
		case uniswapv3poolswap.FieldLiquidity:
			if value, ok := values[i].(*schema.BigInt); !ok {
				return fmt.Errorf("unexpected type %T for field liquidity", values[i])
			} else if value != nil {
				uvs.Liquidity = value
			}
		case uniswapv3poolswap.FieldTick:
			if value, ok := values[i].(*schema.BigInt); !ok {
				return fmt.Errorf("unexpected type %T for field tick", values[i])
			} else if value != nil {
				uvs.Tick = value
			}
		case uniswapv3poolswap.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field event_id", value)
			} else if value.Valid {
				uvs.event_id = new(int)
				*uvs.event_id = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryEvent queries the "event" edge of the UniswapV3PoolSwap entity.
func (uvs *UniswapV3PoolSwap) QueryEvent() *EventQuery {
	return (&UniswapV3PoolSwapClient{config: uvs.config}).QueryEvent(uvs)
}

// Update returns a builder for updating this UniswapV3PoolSwap.
// Note that you need to call UniswapV3PoolSwap.Unwrap() before calling this method if this UniswapV3PoolSwap
// was returned from a transaction, and the transaction was committed or rolled back.
func (uvs *UniswapV3PoolSwap) Update() *UniswapV3PoolSwapUpdateOne {
	return (&UniswapV3PoolSwapClient{config: uvs.config}).UpdateOne(uvs)
}

// Unwrap unwraps the UniswapV3PoolSwap entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uvs *UniswapV3PoolSwap) Unwrap() *UniswapV3PoolSwap {
	tx, ok := uvs.config.driver.(*txDriver)
	if !ok {
		panic("ent: UniswapV3PoolSwap is not a transactional entity")
	}
	uvs.config.driver = tx.drv
	return uvs
}

// String implements the fmt.Stringer.
func (uvs *UniswapV3PoolSwap) String() string {
	var builder strings.Builder
	builder.WriteString("UniswapV3PoolSwap(")
	builder.WriteString(fmt.Sprintf("id=%v", uvs.ID))
	builder.WriteString(", sender=")
	builder.WriteString(uvs.Sender)
	builder.WriteString(", recipient=")
	builder.WriteString(uvs.Recipient)
	builder.WriteString(", amount0=")
	builder.WriteString(fmt.Sprintf("%v", uvs.Amount0))
	builder.WriteString(", amount1=")
	builder.WriteString(fmt.Sprintf("%v", uvs.Amount1))
	builder.WriteString(", sqrt_price_x96=")
	builder.WriteString(fmt.Sprintf("%v", uvs.SqrtPriceX96))
	builder.WriteString(", liquidity=")
	builder.WriteString(fmt.Sprintf("%v", uvs.Liquidity))
	builder.WriteString(", tick=")
	builder.WriteString(fmt.Sprintf("%v", uvs.Tick))
	builder.WriteByte(')')
	return builder.String()
}

// UniswapV3PoolSwaps is a parsable slice of UniswapV3PoolSwap.
type UniswapV3PoolSwaps []*UniswapV3PoolSwap

func (uvs UniswapV3PoolSwaps) config(cfg config) {
	for _i := range uvs {
		uvs[_i].config = cfg
	}
}
