package vwap

import (
	"github.com/guil95/vwap-coinbase/internal/domain/trader"
	"github.com/shopspring/decimal"
	"log"
)

type Vwap struct {
}

func NewVwap() Vwap {
	return Vwap{}
}

const minTradersToCalc = 1

func (v *Vwap) Calc(traders []trader.TradeResponse) (decimal.Decimal, error) {
	sumVolPrice := decimal.NewFromInt(0)
	sumVol := decimal.NewFromInt(0)

	if len(traders) <= minTradersToCalc {
		return decimal.Decimal{}, VwapInvalidQtdTraderError{}
	}

	for _, trade := range traders {
		tradeVol, err := decimal.NewFromString(trade.Size)
		if err != nil {
			log.Print(err)
			return decimal.Decimal{}, VwapInvalidSizeError{}
		}
		tradePrice, err := decimal.NewFromString(trade.Price)
		if err != nil {
			log.Print(err)
			return decimal.Decimal{}, VwapInvalidPriceError{}
		}
		sumVol = sumVol.Add(tradeVol)
		sumVolPrice = sumVolPrice.Add(tradePrice.Mul(tradeVol))
	}

	return sumVolPrice.Div(sumVol), nil
}
