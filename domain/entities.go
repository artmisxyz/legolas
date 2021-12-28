package domain

import (
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Time        time.Time `json:"time"`
	Name        string    `json:"name"`
	Signature   string    `json:"signature"`
	Address     string    `json:"address"`
	BlockNumber uint64    `json:"block_number"`
	TxHash      string    `json:"transaction_hash"`
	Index       uint      `json:"index"`
	TxIndex     uint      `json:"transaction_index"`
	BlockHash   string    `json:"block_hash"`
}

type Swap struct {
	Event        *Event `json:"event"`
	Sender       string `json:"sender"`
	Recipient    string `json:"recipient"`
	Amount0      string `json:"amount_0"`
	Amount1      string `json:"amount_1"`
	SqrtPriceX96 string `json:"sqrt_price_x96"`
	Liquidity    string `json:"liquidity"`
	Tick         string `json:"tick"`
}

type Burn struct {
	Event     *Event `json:"event"`
	Owner     string `json:"owner"`
	TickLower string `json:"tick_lower"`
	TickUpper string `json:"tick_upper"`
	Amount    string `json:"amount"`
	Amount0   string `json:"amount_0"`
	Amount1   string `json:"amount_1"`
}

type Collect struct {
	Event     *Event `json:"event"`
	TokenId   string `json:"token_id"`
	Recipient string `json:"recipient"`
	Amount0   string `json:"amount_0"`
	Amount1   string `json:"amount_1"`
}

type DecreaseLiquidity struct {
	Event     *Event `json:"event"`
	TokenId   string `json:"token_id"`
	Liquidity string `json:"liquidity"`
	Amount0   string `json:"amount_0"`
	Amount1   string `json:"amount_1"`
}

type PoolCreated struct {
	Event       *Event `json:"event"`
	Token0      string `json:"token_0"`
	Token1      string `json:"token_1"`
	Fee         string `json:"fee"`
	TickSpacing string `json:"tick_spacing"`
	Pool        string `json:"pool"`
}

type Flash struct {
	Event     *Event `json:"event"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount0   string `json:"amount_0"`
	Amount1   string `json:"amount_1"`
	Paid0     string `json:"paid_0"`
	Paid1     string `json:"paid_1"`
}

type IncreaseLiquidity struct {
	Event     *Event `json:"event"`
	TokenId   string `json:"token_id"`
	Liquidity string `json:"liquidity"`
	Amount0   string `json:"amount_0"`
	Amount1   string `json:"amount_1"`
}

type PoolInitialize struct {
	Event        *Event `json:"event"`
	SqrtPriceX96 string `json:"sqrt_price_x_96"`
	Tick         string `json:"tick"`
}

type Mint struct {
	Event     *Event `json:"event"`
	Sender    string `json:"sender"`
	Owner     string `json:"owner"`
	TickLower string `json:"tick_lower"`
	TickUpper string `json:"tick_upper"`
	Amount    string `json:"amount"`
	Amount0   string `json:"amount_0"`
	Amount1   string `json:"amount_1"`
}

type Transfer struct {
	Event   *Event `json:"event"`
	From    string `json:"from"`
	To      string `json:"to"`
	TokenId string `json:"token_id"`
}
