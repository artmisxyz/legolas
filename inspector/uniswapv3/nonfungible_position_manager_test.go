package uniswapv3

import (
	"gotest.tools/assert"
	"testing"
)

func TestIncreaseLiquidityEventHandler_Signature(t *testing.T) {
	i := NewIncreaseLiquidityEventHandler()
	assert.Equal(t, i.Signature(), "0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f")
}

func TestDecreaseLiquidityEventHandler_Signature(t *testing.T) {
	i := NewDecreaseLiquidityEventHandler()
	assert.Equal(t, i.Signature(), "0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4")
}