package vwap

import (
	"github.com/guil95/vwap-coinbase/internal/domain/trader"
	"github.com/shopspring/decimal"
)

type Vwap struct {
}

func NewVwap() Vwap {
	return Vwap{}
}

func (v *Vwap) Calc(traders []trader.TradeResponse) (decimal.Decimal, error) {
	sumVolPrice := decimal.NewFromInt(0)
	sumVol := decimal.NewFromInt(0)

	for _, trade := range traders {
		tradeVol, err := decimal.NewFromString(trade.Size)
		if err != nil {
			return decimal.Decimal{}, err
		}
		tradePrice, err := decimal.NewFromString(trade.Price)
		if err != nil {
			return decimal.Decimal{}, err
		}
		sumVol = sumVol.Add(tradeVol)
		sumVolPrice = sumVolPrice.Add(tradePrice.Mul(tradeVol))
	}

	return sumVolPrice.Div(sumVol), nil
}
