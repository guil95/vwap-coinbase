package usecase

import (
	"github.com/guil95/vwap-coinbase/internal/domain/vwap"

	"github.com/guil95/vwap-coinbase/internal/domain/trader"
	"golang.org/x/net/context"
)

type useCase struct {
	traderChan chan trader.TradeResponse
	vwap       vwap.Vwap
}

const maxMatches = 200

func NewUseCase(traderChan chan trader.TradeResponse, vwap vwap.Vwap) *useCase {
	return &useCase{traderChan: traderChan, vwap: vwap}
}

func (uc *useCase) TradeProducts(ctx context.Context, errorChan chan error) {
	traders := make(map[trader.ProductID][]trader.TradeResponse)

	for trade := range uc.traderChan {
		traders[trader.ProductID(trade.ProductID)] = append(traders[trader.ProductID(trade.ProductID)], trade)

		if len(traders[trader.ProductID(trade.ProductID)]) >= maxMatches {
			uc.vwap.Calc(
				trade.ProductID,
				traders[trader.ProductID(trade.ProductID)][len(traders[trader.ProductID(trade.ProductID)])-maxMatches:],
			)
		}
	}
}
