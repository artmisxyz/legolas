package uniswapv3

import (
	"gotest.tools/assert"
	"testing"
)

func TestPoolCreatedEventHandler_Signature(t *testing.T) {
	p := NewPoolCreatedEventHandler()
	assert.Equal(t, p.Signature(), "0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118")
}
