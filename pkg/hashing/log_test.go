package hashing

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogHash(t *testing.T) {
	log := types.Log{
		BlockNumber: 3817381,
		TxIndex:     31,
		Index:       4,
	}
	assert.Equal(t, "0xdf2265702f50b14171f1311e44cf10d5a89c56d4a8a1f0269874230fa188360a", LogHash(log))
}
