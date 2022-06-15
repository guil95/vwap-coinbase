package usecase

import (
	"fmt"
	"github.com/guil95/vwap-coinbase/internal/domain/vwap"

	"github.com/guil95/vwap-coinbase/internal/domain/trader"
)

type useCase struct {
	traderChan chan trader.TradeResponse
	vwap       vwap.Vwap
}

const maxMatches = 200

func NewUseCase(traderChan chan trader.TradeResponse, vwap vwap.Vwap) *useCase {
	return &useCase{traderChan: traderChan, vwap: vwap}
}

func (uc *useCase) TradeProducts(errorChan chan error) {
	traders := make(map[trader.ProductID][]trader.TradeResponse)

	for trade := range uc.traderChan {
		traders[trader.ProductID(trade.ProductID)] = append(traders[trader.ProductID(trade.ProductID)], trade)

		if len(traders[trader.ProductID(trade.ProductID)]) >= maxMatches {
			vwapResult, err := uc.vwap.Calc(
				traders[trader.ProductID(trade.ProductID)][len(traders[trader.ProductID(trade.ProductID)])-maxMatches:],
			)
			if err != nil {
				errorChan <- err
				return
			}

			fmt.Println(fmt.Sprintf("The vwap for product_id %v is %v", trade.ProductID, vwapResult))
		}
	}
}
