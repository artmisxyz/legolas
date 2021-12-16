package hashing

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func LogHash(log types.Log) string {
	blockNumber := fmt.Sprintf("%d", log.BlockNumber)
	txIndex := fmt.Sprintf("%d", log.TxIndex)
	logIndex := fmt.Sprintf("%d", log.Index)
	return crypto.Keccak256Hash([]byte(blockNumber), []byte(txIndex), []byte(logIndex)).String()
}
