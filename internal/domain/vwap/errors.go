package vwap

const (
	errorMessage             = "Internal error calc vwap"
	errorInvalidSizeMessage  = "Internal error invalid size to calc vwap"
	errorInvalidPriceMessage = "Internal error invalid price to calc vwap"
)

type VwapInvalidQtdTraderError struct{}

func (v VwapInvalidQtdTraderError) Error() string {
	return errorMessage
}

type VwapInvalidSizeError struct{}

func (v VwapInvalidSizeError) Error() string {
	return errorInvalidSizeMessage
}

type VwapInvalidPriceError struct{}

func (v VwapInvalidPriceError) Error() string {
	return errorInvalidPriceMessage
}
