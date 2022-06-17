package vwap

import (
	"github.com/shopspring/decimal"
	"testing"

	"github.com/guil95/vwap-coinbase/internal/domain/trader"
	"github.com/stretchr/testify/assert"
)

func TestVwapCalc(t *testing.T) {
	t.Run("Test calc with success", func(t *testing.T) {
		vwapStruct := NewVwap()
		traders := []trader.TradeResponse{
			{
				ProductID: string(trader.EthBtc),
				Size:      "0.20",
				Price:     "0.10",
				Type:      "match",
			},
			{
				ProductID: string(trader.EthBtc),
				Size:      "0.20",
				Price:     "0.20",
				Type:      "match",
			},
		}

		calc, err := vwapStruct.Calc(traders)
		expected := float64(0.15)
		actual, _ := calc.Float64()

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Test calc with only one trade should return error", func(t *testing.T) {
		vwapStruct := NewVwap()
		traders := []trader.TradeResponse{
			{
				ProductID: string(trader.EthBtc),
				Size:      "0.20",
				Price:     "0.20",
				Type:      "match",
			},
		}

		calc, err := vwapStruct.Calc(traders)

		assert.Error(t, err)
		assert.ErrorIs(t, err, VwapInvalidQtdTraderError{})
		assert.Equal(t, errorMessage, err.Error())
		assert.Equal(t, decimal.Decimal{}, calc)
	})

	t.Run("Test calc with invalid size should return error", func(t *testing.T) {
		vwapStruct := NewVwap()
		traders := []trader.TradeResponse{
			{
				ProductID: string(trader.EthBtc),
				Size:      "invalid_size",
				Price:     "0.10",
				Type:      "match",
			},
			{
				ProductID: string(trader.EthBtc),
				Size:      "0.20",
				Price:     "0.20",
				Type:      "match",
			},
		}

		calc, err := vwapStruct.Calc(traders)

		assert.Error(t, err)
		assert.ErrorIs(t, err, VwapInvalidSizeError{})
		assert.Equal(t, errorInvalidSizeMessage, err.Error())
		assert.Equal(t, decimal.Decimal{}, calc)
	})

	t.Run("Test calc with invalid price should return error", func(t *testing.T) {
		vwapStruct := NewVwap()
		traders := []trader.TradeResponse{
			{
				ProductID: string(trader.EthBtc),
				Size:      "0.10",
				Price:     "invalid price",
				Type:      "match",
			},
			{
				ProductID: string(trader.EthBtc),
				Size:      "0.20",
				Price:     "0.20",
				Type:      "match",
			},
		}

		calc, err := vwapStruct.Calc(traders)

		assert.Error(t, err)
		assert.ErrorIs(t, err, VwapInvalidPriceError{})
		assert.Equal(t, errorInvalidPriceMessage, err.Error())
		assert.Equal(t, decimal.Decimal{}, calc)
	})
}
