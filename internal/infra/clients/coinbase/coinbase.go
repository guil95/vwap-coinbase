package coinbase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/guil95/vwap-coinbase/internal/domain/client"
	"github.com/guil95/vwap-coinbase/internal/domain/trader"
	"golang.org/x/net/websocket"
	"log"
)

const coinBaseAddress string = "wss://ws-feed.exchange.coinbase.com"

type coinBaseClient struct {
	conn *websocket.Conn
}

type coinBaseRequest struct {
	Type       string    `json:"type"`
	ProductIDs []string  `json:"product_ids"`
	Channels   []channel `json:"channels"`
}

type channel struct {
	Name       string
	ProductIDs []string
}

func NewCoinBaseClient() (client.WsClient, error) {
	conn, err := websocket.Dial(coinBaseAddress, "", "http://localhost")

	if err != nil {
		fmt.Printf("Dial failed: %s\n", err.Error())
		return &coinBaseClient{}, err
	}

	return &coinBaseClient{
		conn: conn,
	}, nil
}

func (cb *coinBaseClient) Subscribe(
	ctx context.Context,
	tradeChan chan trader.TradeResponse,
	errorChan chan error,
) {
	subscription := coinBaseRequest{
		Type:       "subscribe",
		ProductIDs: trader.ProductsIDS,
		Channels: []channel{
			{Name: "matches"},
		},
	}

	payload, err := json.Marshal(subscription)
	if err != nil {
		errorChan <- err
	}

	err = websocket.Message.Send(cb.conn, payload)
	if err != nil {
		errorChan <- err
	}

	go cb.receiveResponseWS(ctx, tradeChan, errorChan)
}

func (cb *coinBaseClient) receiveResponseWS(ctx context.Context, tradeChan chan trader.TradeResponse, errorChan chan error) {
	for {
		select {
		case <-ctx.Done():
			err := cb.conn.Close()
			if err != nil {
				log.Printf("failed closing ws connection: %s", err)
			}
		default:
			var response trader.TradeResponse
			err := websocket.JSON.Receive(cb.conn, &response)
			if err != nil {
				errorChan <- err
			}

			tradeChan <- response
		}
	}
}
