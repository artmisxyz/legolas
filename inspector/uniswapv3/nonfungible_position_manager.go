package uniswapv3

import (
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
	return "0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f"
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

type decreaseLiquidityEventHandler struct{}

func (d *decreaseLiquidityEventHandler) Signature() string {
	return "0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4"
}

func (d *decreaseLiquidityEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
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

type CollectEvent struct {
	tokenId   *big.Int
	recipient *big.Int
	amount0   *big.Int
	amount1   *big.Int
}

type collectEventHandler struct {
}

func NewCollectEventHandler() inspector.EventHandler {
	return &collectEventHandler{}
}

func (c *collectEventHandler) Signature() string {
	return "0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01"
}

func (c *collectEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
	var v CollectEvent
	err := abi.UnpackIntoInterface(&v, "Collect", data)
	if err != nil {
		return err
	}
	v.tokenId = topics[1].Big()
	return nil
}

type TransferEvent struct {
	from    common.Address
	to      common.Address
	tokenId *big.Int
}

type transferEventHandler struct{}

func NewTransferEventHandler() inspector.EventHandler {
	return &transferEventHandler{}
}

func (t transferEventHandler) Signature() string {
	return "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
}

func (t transferEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
	var v TransferEvent
	v.from = common.BytesToAddress(topics[1].Bytes())
	v.to = common.BytesToAddress(topics[2].Bytes())
	v.tokenId = topics[3].Big()
	return nil
}
