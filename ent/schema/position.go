package schema

import (
	"database/sql"
	"database/sql/driver"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"fmt"
	"math/big"
)

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

//type Position struct {
//	owner                    common.Address
//	pool                     common.Address
//	token0                   common.Address
//	token1                   common.Address
//	tickLower                *big.Int
//	tickUpper                *big.Int
//	liquidity                *big.Int
//	depositedToken0          *big.Int
//	depositedToken1          *big.Int
//	withdrawnToken0          *big.Int
//	withdrawnToken1          *big.Int
//	collectedToken0          *big.Int
//	collectedToken1          *big.Int
//	collectedFeesToken0      *big.Int
//	collectedFeesToken1      *big.Int
//	feeGrowthInside0LastX128 *big.Int
//	feeGrowthInside1LastX128 *big.Int
//}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.Bytes("owner"),
		field.Bytes("pool"),
		field.Bytes("token0"),
		field.Bytes("token1"),
		field.String("tickLower").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("tickUpper").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("liquidity").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("depositedToken0").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("depositedToken1").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("withdrawnToken0").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("withdrawnToken1").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("collectedToken0").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("collectedToken1").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("collectedFeesToken0").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("collectedFeesToken1").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("feeGrowthInside0LastX128").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("feeGrowthInside1LastX128").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
	}
}

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return nil
}

type BigInt struct {
	big.Int
}

func (b *BigInt) Scan(src interface{}) error {
	var i sql.NullString
	if err := i.Scan(src); err != nil {
		return err
	}
	if !i.Valid {
		return nil
	}
	if _, ok := b.Int.SetString(i.String, 10); ok {
		return nil
	}
	return fmt.Errorf("could not scan type %T with value %v into BigInt", src, src)
}

func (b *BigInt) Value() (driver.Value, error) {
	return b.String(), nil
}
