package moneygo

type FromTo struct {
	From CurrencyCode
	To   CurrencyCode
}

type Rates map[CurrencyCode]float64

type MoneyGo struct {
	base            CurrencyCode `json:"base"`
	rates           Rates        `json:"rates"`
	DefaultSettings FromTo       `json:"defaultSettings"`
}
