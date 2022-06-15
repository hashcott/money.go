package moneygo

type MoneyGoInterface interface {
	from(val float64) MoneyGoInterface
	to(val float64) int
}
