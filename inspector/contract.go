package inspector

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EventHandler interface {
	Signature() string
	Handle(abi abi.ABI, topics []common.Hash, data []byte)
}

type Contract struct {
	ABI           abi.ABI
	Address       common.Address
	Name          string
	EventHandlers []EventHandler
}

func (c *Contract) HandleEvent(log types.Log) {
	for _, e := range c.EventHandlers {
		if e.Signature() == log.Topics[0].String() {
			e.Handle(c.ABI, log.Topics, log.Data)
		}
	}
}

func (c *Contract) RegisterEventHandler(e EventHandler) {
	c.EventHandlers = append(c.EventHandlers, e)
}
