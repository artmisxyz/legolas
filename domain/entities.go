package domain

type Event struct {
	Name        string `json:"name"`
	Signature   string `json:"signature"`
	Address     string `json:"address"`
	BlockNumber uint64 `json:"block_number"`
	TxHash      string `json:"transaction_hash"`
	Index       uint   `json:"index"`
	TxIndex     uint   `json:"transaction_index"`
	BlockHash   string `json:"block_hash"`
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
