package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/guil95/vwap-coinbase/internal/domain/trader"
	"github.com/guil95/vwap-coinbase/internal/domain/vwap"
	"github.com/guil95/vwap-coinbase/internal/infra/clients/coinbase"
	"github.com/guil95/vwap-coinbase/internal/usecase"
	"golang.org/x/net/context"
)

func main() {
	wsClient, _ := coinbase.NewCoinBaseClient()
	traderChan := make(chan trader.TradeResponse)
	uc := usecase.NewUseCase(traderChan, vwap.NewVwap())
	errorChan := make(chan error)
	interrupt := make(chan os.Signal, 1)
	ctx := context.Background()

	signal.Notify(interrupt, os.Interrupt)
	go uc.TradeProducts(errorChan)
	go wsClient.Subscribe(ctx, traderChan, errorChan)

	select {
	case <-interrupt:
		closeCtx(ctx)
		log.Println("Finish process")
		return
	case err := <-errorChan:
		closeCtx(ctx)
		log.Fatal("Error: ", err)
	}
}

func closeCtx(ctx context.Context) {
	done := make(chan struct{})
	go func() {
		defer close(done)
		ctx.Done()
	}()
}
