package moneygo_test

import (
	moneygo "github.com/hashcott/money.go"
	"testing"
)

func TestMoneyGo_Convert(t *testing.T) {
	fx := moneygo.NewExchanger(moneygo.USD, moneygo.Rates{"USD": 1, "EUR": 0.745101, "GBP": 0.647710,
		"HKD": 7.781919})
	fx.DefaultSettings = moneygo.FromTo{From: moneygo.USD, To: moneygo.GBP}

	if val, err := fx.Convert(1000); val != 647.71 {
		t.Errorf("fx error: %s", err)
	}
}

func TestMoneyGo_ConvertWithFromTo(t *testing.T) {
	fx := moneygo.NewExchanger(moneygo.USD, moneygo.Rates{"USD": 1, "EUR": 0.745101, "GBP": 0.647710,
		"HKD": 7.781919})

	if val, err := fx.ConvertWithFromTo(1000, moneygo.FromTo{From: moneygo.USD, To: moneygo.GBP}); val != 647.71 {
		t.Errorf("fx error: %s", err)
	}
}
