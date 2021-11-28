package uniswapv3

import (
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

type IncreaseLiquidityEvent struct {
	tokenId   *big.Int
	liquidity *big.Int
	amount0   *big.Int
	amount1   *big.Int
}

type increaseLiquidityEventHandler struct {
}

func NewIncreaseLiquidityEventHandler() inspector.EventHandler {
	return &increaseLiquidityEventHandler{}
}

func (i *increaseLiquidityEventHandler) Signature() string {
	return crypto.Keccak256Hash([]byte("IncreaseLiquidity(uint256,uint128,uint256,uint256)")).String()
}

func (i *increaseLiquidityEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
	var v IncreaseLiquidityEvent
	err := abi.UnpackIntoInterface(&v, "IncreaseLiquidity", data)
	if err != nil {
		return err
	}
	v.tokenId = topics[1].Big()
	return nil
}



type DecreaseLiquidityEvent struct {
	tokenId   *big.Int
	liquidity *big.Int
	amount0   *big.Int
	amount1   *big.Int
}

type decreaseLiquidityEventHandler struct {}

func (d decreaseLiquidityEventHandler) Signature() string {
	return crypto.Keccak256Hash([]byte("DecreaseLiquidity(uint256,uint128,uint256,uint256)")).String()
}

func (d decreaseLiquidityEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
	var v DecreaseLiquidityEvent
	err := abi.UnpackIntoInterface(&v, "DecreaseLiquidity", data)
	if err != nil {
		return err
	}
	v.tokenId = topics[1].Big()
	return nil
}

func NewDecreaseLiquidityEventHandler() inspector.EventHandler {
	return &decreaseLiquidityEventHandler{}
}
