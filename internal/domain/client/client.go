package client

import (
	"github.com/guil95/vwap-coinbase/internal/domain/trader"
	"golang.org/x/net/context"
)

type WsClient interface {
	Subscribe(ctx context.Context, responseChan chan trader.TradeResponse, errorChan chan error)
}
