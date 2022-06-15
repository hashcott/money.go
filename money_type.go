package moneygo

type FromTo struct {
	From CurrencyCode
	To   CurrencyCode
}

type Rates map[CurrencyCode]float64

type MoneyGo struct {
	Base            CurrencyCode `json:"base"`
	Rates           Rates        `json:"rates"`
	DefaultSettings FromTo       `json:"defaultSettings"`
	MoneyInterface
}
