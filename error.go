package moneygo

import "errors"

var ErrorConfig = errors.New("please set up default settings")
var ErrorUnknown = errors.New("fx error")
var ErrorCurrencyCode = errors.New("not found currency code to exchange")
