package uniswapv3

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Position struct {
	owner                    common.Address
	pool                     common.Address
	token0                   common.Address
	token1                   common.Address
	tickLower                *big.Int
	tickUpper                *big.Int
	liquidity                *big.Int
	depositedToken0          *big.Int
	depositedToken1          *big.Int
	withdrawnToken0          *big.Int
	withdrawnToken1          *big.Int
	collectedToken0          *big.Int
	collectedToken1          *big.Int
	collectedFeesToken0      *big.Int
	collectedFeesToken1      *big.Int
	feeGrowthInside0LastX128 *big.Int
	feeGrowthInside1LastX128 *big.Int
}
