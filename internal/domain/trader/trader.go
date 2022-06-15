package trader

type TradeResponse struct {
	Type      string `json:"type"`
	Size      string `json:"size"`
	Price     string `json:"price"`
	ProductID string `json:"product_id"`
}
type ProductID string

const (
	BtcUsd ProductID = "BTC-USD"
	EthUsd ProductID = "ETH-USD"
	EthBtc ProductID = "ETH-BTC"
)

var ProductsIDS = []string{string(BtcUsd), string(EthUsd), string(EthBtc)}
