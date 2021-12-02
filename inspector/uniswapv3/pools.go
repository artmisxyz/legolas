package uniswapv3

import (
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type InitializeEvent struct {
	sqrtPriceX96 *big.Int
	tick         int
}

type initializeEventHandler struct {
}

func NewInitializeEventHandler() inspector.EventHandler {
	return &initializeEventHandler{}
}

func (i *initializeEventHandler) Signature() string {
	return crypto.Keccak256Hash([]byte("Initialize(uint160,int24)")).String()
}

func (i *initializeEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
	var v InitializeEvent
	err := abi.UnpackIntoInterface(&v, "Initialize", data)
	if err != nil {
		return err
	}
	return nil
}

type swapEvent struct {
	sender       common.Address
	recipient    common.Address
	amount0      *big.Int
	amount1      *big.Int
	sqrtPriceX96 *big.Int
	liquidity    *big.Int
	tick         int
}

type swapEventHandler struct{}

func NewSwapEventHandler() inspector.EventHandler {
	return &swapEventHandler{}
}

func (s swapEventHandler) Signature() string {
	return crypto.Keccak256Hash([]byte("Swap(address,address,int256,int256,uint160,uint128,int24)")).String()
}

func (s swapEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
	var v swapEvent
	err := abi.UnpackIntoInterface(&v, "Swap", data)
	if err != nil {
		return err
	}
	v.sender = common.BytesToAddress(topics[1].Bytes())
	v.recipient = common.BytesToAddress(topics[2].Bytes())
	return nil
}

type MintEvent struct {
	sender    common.Address
	owner     common.Address
	tickLower int64
	tickUpper int64
	amount    *big.Int
	amount0   *big.Int
	amount1   *big.Int
}

type mintEventHandler struct{}

func NewMintEventHandler() inspector.EventHandler {
	return &mintEventHandler{}
}

func (m mintEventHandler) Signature() string {
	return crypto.Keccak256Hash([]byte("Mint(address,address,int24,int24,uint128,uint256,uint256)")).String()
}

func (m mintEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
	var v MintEvent
	err := abi.UnpackIntoInterface(&v, "Mint", data)
	if err != nil {
		return err
	}
	v.owner = common.BytesToAddress(topics[2].Bytes())
	v.tickLower = topics[3].Big().Int64()
	v.tickUpper = topics[4].Big().Int64()
	return nil
}

type BurnEvent struct {
	owner     common.Address
	tickLower int64
	tickUpper int64
	amount    *big.Int
	amount0   *big.Int
	amount1   *big.Int
}

type burnEventHandler struct{}

func NewBurnEventHanlder() inspector.EventHandler {
	return &burnEventHandler{}
}

func (b burnEventHandler) Signature() string {
	return crypto.Keccak256Hash([]byte("Burn(address,int24,int24,uint128,uint256,uint256)")).String()
}

func (b burnEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
	var v BurnEvent
	err := abi.UnpackIntoInterface(&v, "Burn", data)
	if err != nil {
		return err
	}
	v.owner = common.BytesToAddress(topics[1].Bytes())
	v.tickLower = topics[2].Big().Int64()
	v.tickUpper = topics[3].Big().Int64()
	return nil
}

type FlashEvent struct {
	sender    common.Address
	recipient common.Address
	amount0   *big.Int
	amount1   *big.Int
	paid0     *big.Int
	paid1     *big.Int
}

type flashEventHandler struct {
}

func NewFlashEventHandler() inspector.EventHandler {
	return &flashEventHandler{}
}

func (f flashEventHandler) Signature() string {
	return crypto.Keccak256Hash([]byte("Burn(address,int24,int24,uint128,uint256,uint256)")).String()
}

func (f flashEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
	var v FlashEvent
	err := abi.UnpackIntoInterface(&v, "Flash", data)
	if err != nil {
		return err
	}
	v.sender = common.BytesToAddress(topics[1].Bytes())
	v.recipient = common.BytesToAddress(topics[2].Bytes())
	return nil
}
