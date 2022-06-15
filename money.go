package moneygo

import "errors"

func (fx *MoneyGo) NewExchanger(base CurrencyCode, rates Rates, defaultSettings FromTo) *MoneyGo {
	return &MoneyGo{
		Base:            base,
		Rates:           rates,
		DefaultSettings: defaultSettings,
	}
}

func (fx *MoneyGo) Convert() (float64, error) {
	if fx.DefaultSettings.From == "" || fx.DefaultSettings.To == "" {
		return 0, errors.New("please set up default config")
	}
	return 0, nil
}

func (fx *MoneyGo) ConvertWithFromTo(fromTo FromTo) (float64, error) {
	return 0, nil
}

func (fx *MoneyGo) From(code CurrencyCode) {
}
