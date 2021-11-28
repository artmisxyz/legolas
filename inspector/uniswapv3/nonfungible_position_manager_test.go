package uniswapv3

import (
	"gotest.tools/assert"
	"testing"
)

func TestIncreaseLiquidityEventHandler_Signature(t *testing.T) {
	i := NewIncreaseLiquidityEventHandler()
	assert.Equal(t, i.Signature(), "0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f")
}
