package uniswapv3

import (
	"fmt"
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"os"
)

type uniswapV3 struct {
	logger    *zap.Logger
	contracts []contract
}

type contract struct {
	abi     abi.ABI
	address common.Address
	name    string
}

const Inspector = "uniswapV3:inspector"

const (
	UniswapV3Factory                   = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	Multicall2                         = "0x5BA1e12693Dc8F9c48aAD8770482f4739bEeD696"
	ProxyAdmin                         = "0xB753548F6E010e7e680BA186F9Ca1BdAB2E90cf2"
	TickLens                           = "0xbfd8137f7d1516D3ea5cA83523914859ec47F573"
	Quoter                             = "0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6"
	SwapRouter                         = "0xE592427A0AEce92De3Edee1F18E0157C05861564"
	NFTDescriptor                      = "0x42B24A95702b9986e82d421cC3568932790A48Ec"
	NonfungibleTokenPositionDescriptor = "0x91ae842A5Ffd8d12023116943e72A606179294f3"
	TransparentUpgradeableProxy        = "0xEe6A57eC80ea46401049E92587E52f5Ec1c24785"
	NonfungiblePositionManager         = "0xC36442b4a4522E871399CD717aBDD847Ab11FE88"
	V3Migrator                         = "0xA5644E29708357803b5A882D272c41cC0dF92B34"
)

func NewUniswapV3(logger *zap.Logger) inspector.Inspector {
	v := &uniswapV3{
		logger:    logger.Named(Inspector),
		contracts: make([]contract, 11),
	}
	base := "inspector/uniswapV3/abis/"

	v.mustRegisterContract(UniswapV3Factory, base+"UniswapV3Factory.json", "UniswapV3Factory")
	v.mustRegisterContract(Multicall2, base+"Multicall2.json", "Multicall2")
	v.mustRegisterContract(ProxyAdmin, base+"ProxyAdmin.json", "ProxyAdmin")
	v.mustRegisterContract(TickLens, base+"TickLens.json", "TickLens")
	v.mustRegisterContract(Quoter, base+"Quoter.json", "Quoter")
	v.mustRegisterContract(SwapRouter, base+"SwapRouter.json", "SwapRouter")
	v.mustRegisterContract(NFTDescriptor, base+"NFTDescriptor.json", "NFTDescriptor")
	v.mustRegisterContract(NonfungibleTokenPositionDescriptor, base+"NonfungibleTokenPositionDescriptor.json", "NonfungibleTokenPositionDescriptor")
	v.mustRegisterContract(TransparentUpgradeableProxy, base+"TransparentUpgradeableProxy.json", "TransparentUpgradeableProxy")
	v.mustRegisterContract(NonfungiblePositionManager, base+"NonfungiblePositionManager.json", "NonfungiblePositionManager")
	v.mustRegisterContract(V3Migrator, base+"V3Migrator.json", "V3Migrator")

	return v
}

func (v *uniswapV3) Name() string {
	return Inspector
}

func (v *uniswapV3) InspectBlock(block *types.Block) error {
	txs := v.Filter(block)
	for _, tx := range txs {
		err := v.InspectTransaction(tx)
		if err != nil {
			return fmt.Errorf("error inspecting tx. %w", err)
		}
	}
	return nil
}

func (v *uniswapV3) InspectTransaction(tx *types.Transaction) error {
	v.logger.Info("inspecting transaction", zap.Any("tx", tx.Hash()))
	fmt.Println(tx.Hash())
	return nil
}

func (v *uniswapV3) Filter(block *types.Block) []*types.Transaction {
	var filtered []*types.Transaction
	txs := block.Transactions()
	for _, tx := range txs {
		if tx.To() != nil {
			if v.isMine(tx.To().String()) {
				filtered = append(filtered, tx)
			}
		}
	}
	return filtered
}

func (v *uniswapV3) mustRegisterContract(address, abiPath, name string) {
	f, err := os.Open(abiPath)
	if err != nil {
		panic(err)
	}
	a, err := abi.JSON(f)
	if err != nil {
		panic(err)
	}
	c := contract{
		name:    name,
		abi:     a,
		address: common.HexToAddress(address),
	}
	v.contracts = append(v.contracts, c)
}

func (v *uniswapV3) isMine(address string) bool {
	for _, c := range v.contracts {
		if c.address.String() == address {
			return true
		}
	}
	return false
}
