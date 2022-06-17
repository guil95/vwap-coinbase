package usecase

import (
	"fmt"
	"log"

	"github.com/guil95/vwap-coinbase/internal/domain/trader"
	"github.com/guil95/vwap-coinbase/internal/domain/vwap"
)

type useCase struct {
	traderChan chan trader.TradeResponse
	vwap       vwap.Vwap
}

const (
	maxMatches = 200
	minMatches = 2
)

func NewUseCase(traderChan chan trader.TradeResponse, vwap vwap.Vwap) *useCase {
	return &useCase{traderChan: traderChan, vwap: vwap}
}

func (uc *useCase) TradeProducts(errorChan chan error) {
	traders := make(map[trader.ProductID][]trader.TradeResponse)

	for trade := range uc.traderChan {
		traders[trader.ProductID(trade.ProductID)] = append(traders[trader.ProductID(trade.ProductID)], trade)
		tradersToCalc := traders[trader.ProductID(trade.ProductID)]

		if len(traders[trader.ProductID(trade.ProductID)]) >= maxMatches {
			tradersToCalc = traders[trader.ProductID(trade.ProductID)][len(traders[trader.ProductID(trade.ProductID)])-maxMatches:]
		}

		if len(tradersToCalc) >= minMatches {
			vwapResult, err := uc.vwap.Calc(tradersToCalc[0:])
			if err != nil {
				errorChan <- err
				return
			}

			log.Print(fmt.Sprintf("The vwap for product_id %v is %v", trade.ProductID, vwapResult))
		}
	}
}
