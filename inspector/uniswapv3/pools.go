package uniswapv3

import (
	"github.com/artmisxyz/blockinspector/inspector"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type InitializeEvent struct {
}

type initializeEventHandler struct {

}

func NewInitializeEventHandler() inspector.EventHandler {
	return &initializeEventHandler{}
}

func (i initializeEventHandler) Signature() string {
	return ""
}

func (i initializeEventHandler) Handle(abi abi.ABI, topics []common.Hash, data []byte) error {
	return nil
}