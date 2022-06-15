package moneygo

type MoneyInterface interface {
	getRate(fromTo FromTo) (float64, error)
	Convert(val float64) (float64, error)
	ConvertWithFromTo(val float64, fromTo FromTo) (float64, error)
}
