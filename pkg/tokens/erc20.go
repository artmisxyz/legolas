package tokens

import (
	"context"
	"github.com/artmisxyz/uniswap-go/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ERC20Token struct {
	binding     *erc20.Erc20
	name        string
	symbol      string
	decimals    uint8
	totalSupply *big.Int
}

func NewToken(address common.Address, backend bind.ContractBackend) *ERC20Token {
	binding, err := erc20.NewErc20(address, backend)
	if err != nil {
		panic(err)
	}
	return &ERC20Token{binding: binding}
}

func (t *ERC20Token) Name() (string, error) {
	if t.name != "" {
		return t.name, nil
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}
	name, err := t.binding.Name(opts)
	t.name = name
	if err != nil {
		return "", err
	}
	return t.name, nil
}

func (t *ERC20Token) Symbol() (string, error) {
	if t.symbol != "" {
		return t.symbol, nil
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}
	symbol, err := t.binding.Symbol(opts)
	t.symbol = symbol
	if err != nil {
		return "", err
	}
	return t.symbol, nil
}

func (t *ERC20Token) Decimals() (uint8, error) {
	if t.decimals != 0 {
		return t.decimals, nil
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}
	decimals, err := t.binding.Decimals(opts)
	t.decimals = decimals
	if err != nil {
		return 0, err
	}
	return t.decimals, err
}

func (t *ERC20Token) TotalSupply() (*big.Int, error) {
	if t.totalSupply != nil {
		return t.totalSupply, nil
	}

	opts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}
	totalSupply, err := t.binding.TotalSupply(opts)
	t.totalSupply = totalSupply
	if err != nil {
		return nil, err
	}
	return t.totalSupply, nil
}

func (t *ERC20Token) BalanceOf(addr common.Address) (*big.Int, error) {
	opts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}
	return t.binding.BalanceOf(opts, addr)
}
