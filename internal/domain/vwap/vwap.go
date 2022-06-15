package vwap

import (
	"fmt"

	"github.com/guil95/vwap-coinbase/internal/domain/trader"
)

type Vwap struct {
}

func NewVwap() Vwap {
	return Vwap{}
}

func (v *Vwap) Calc(tradeName string, traders []trader.TradeResponse) {
	fmt.Println(fmt.Sprintf("Trade Name: %s trade len %v first %v last %v", tradeName, len(traders), traders[0], traders[len(traders)-1]))
}
