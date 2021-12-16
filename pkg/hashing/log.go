package hashing

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func LogHash(log types.Log) string {
	logIndex := fmt.Sprintf("%d", log.Index)
	return crypto.Keccak256Hash([]byte(log.TxHash.String()), []byte(logIndex)).String()
}
