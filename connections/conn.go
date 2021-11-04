package connections

import (
	"context"
	"fmt"
	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RPC(baseUrl string) *ethclient.Client {
	c, err := ethclient.Dial(baseUrl)
	if err != nil {
		panic("cannot connect to client: " + err.Error())
	}
	return c
}

func Websocket(baseUrl string) *ethclient.Client {
	c, err := ethclient.Dial(baseUrl)
	if err != nil {
		panic(fmt.Sprintln("cannot connect to client: ", err.Error()))
	}
	return c
}

func CurrentBlockChan(logger *zap.Logger, c *ethclient.Client) (chan *types.Header, error) {
	headerChan := make(chan *types.Header, 12)
	sub, err := c.SubscribeNewHead(context.Background(), headerChan)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			err := <-sub.Err()
			logger.Error("subscription error", zap.Error(err))
		}
	}()
	return headerChan, nil
}
