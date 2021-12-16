package tokens

import (
	"context"
	"github.com/artmisxyz/uniswap-go/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ERC20 struct {
	binding     *erc20.Erc20
	Name        string
	Symbol      string
	Decimals    uint8
	TotalSupply *big.Int
}

func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	binding, err := erc20.NewErc20(address, backend)
	if err != nil {
		return nil, err
	}

	instance := &ERC20{binding: binding}
	opts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}

	if err = instance.LoadName(opts); err != nil {
		return nil, err
	}
	if err = instance.LoadSymbol(opts); err != nil {
		return nil, err
	}
	if err = instance.LoadDecimals(opts); err != nil {
		return nil, err
	}
	if err = instance.LoadTotalSupply(opts); err != nil {
		return nil, err
	}

	return instance, nil
}

func (t *ERC20) LoadName(opts *bind.CallOpts) error {
	name, err := t.binding.Name(opts)
	if err != nil {
		return err
	}
	t.Name = name
	return nil
}

func (t *ERC20) LoadSymbol(opts *bind.CallOpts) error {
	symbol, err := t.binding.Symbol(opts)
	if err != nil {
		return err
	}
	t.Symbol = symbol
	return nil
}

func (t *ERC20) LoadDecimals(opts *bind.CallOpts) error {
	decimals, err := t.binding.Decimals(opts)
	if err != nil {
		return err
	}
	t.Decimals = decimals
	return nil
}

func (t *ERC20) LoadTotalSupply(opts *bind.CallOpts) error {
	totalSupply, err := t.binding.TotalSupply(opts)
	if err != nil {
		return err
	}
	t.TotalSupply = totalSupply
	return nil
}

func (t *ERC20) BalanceOf(addr common.Address) (*big.Int, error) {
	opts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}
	return t.binding.BalanceOf(opts, addr)
}
