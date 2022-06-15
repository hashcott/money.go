package moneygo

func NewExchanger(base CurrencyCode, rates Rates) *MoneyGo {
	return &MoneyGo{
		base:  base,
		rates: rates,
	}
}

// Returns the exchange rate to `target` currency from `base` currency
func (fx *MoneyGo) getRate(fromTo FromTo) (float64, error) {
	// Make sure the base rate is in the rates object:
	fx.rates[fx.base] = 1

	// Throw an error if either rate isn't in the rates array
	if fx.rates[fromTo.To] == 0 || fx.rates[fromTo.From] == 0 {
		return 0, ErrorCurrencyCode
	}

	// If `from` currency === fx.base, return the basic exchange rate for the `to` currency
	if fromTo.From == fx.base {
		return fx.rates[fromTo.To], nil
	}

	// If `to` currency === fx.base, return the basic inverse rate of the `from` currency
	if fromTo.To == fx.base {
		return 1 / fx.rates[fromTo.From], nil
	}

	// Otherwise, return the `to` rate multiply by the inverse of the `from` rate to get the
	// relative exchange rate between the two currencies
	return fx.rates[fromTo.To] * 1 / fx.rates[fromTo.From], nil
}

func (fx *MoneyGo) Convert(val float64) (float64, error) {

	// We need to know the `from` and `to` currencies
	if fx.DefaultSettings.From == "" || fx.DefaultSettings.To == "" {
		return 0, ErrorConfig
	}

	// Multiply the value by the exchange rate
	rate, err := fx.getRate(fx.DefaultSettings)
	if err != nil {
		return 0, err
	}

	return rate * val, nil
}

func (fx *MoneyGo) ConvertWithFromTo(val float64, fromTo FromTo) (float64, error) {
	// We need to know the `from` and `to` currencies
	if fromTo.From == "" || fromTo.To == "" {
		return 0, ErrorConfig
	}

	// Multiply the value by the exchange rate
	rate, err := fx.getRate(fromTo)
	if err != nil {
		return 0, ErrorUnknown
	}

	return rate * val, nil
}
