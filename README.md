# Money. <img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Aqua.svg" alt="go" width="56"/>

### Simple JavaScript currency conversion library with no dependencies.


#### Install
```go
go get -u github.com/hashcott/money.go
```

You'll need to do one more thing before you can use it, which is:

#### Setting up exchange rates with rates
To use money.go to convert currencies, you'll need to feed it with some exchange rate data and provide a base currency. (As long as you have exchange rates for every currency relative to one single other ('base') currency, money.go can convert between any other two)

The library doesn't specify a format for currency names/codes (we recommend sticking to the standard three-letter codes). It also does not mind how accurate they are, or which currency is your base rate.

They should like this:

#### NewExchanger(base CurrencyCode, rates Rates)

```go
package main

import (
	moneygo "github.com/hashcott/money.go"
)

func main()  {
	fx := moneygo.NewExchanger(moneygo.USD, moneygo.Rates{"USD": 1, "EUR": 0.745101, "GBP": 0.647710,
		"HKD": 7.781919})
}
```

#### .Convert(val float64) 
The basic function of the library - converts a value from one currency to another. Uses the default from and to currencies in fx.DefaultSettings:

```go
        fx.DefaultSettings = moneygo.FromTo{From: moneygo.USD, To: moneygo.GBP}
	val, err := fx.Convert(1000); // convert from USD to GBP
```

#### .ConvertWithFromTo(val float64, fromTo FromTo)

```go
        fx.ConvertWithFromTo(1000, moneygo.FromTo{From: moneygo.USD, To: moneygo.GBP})
```

