package SharedPackages

import (
	"encoding/json"
	"math"

	"github.com/go-playground/validator"
)

// MoneyType represents the type of money object.
type MoneyType string

const ()

// Money represents a monetary value with a specific currency.
type Money struct {
	currencyCode   string
	fractionDigits int64
	preciseAmount  int64
}

func (m *Money) CentAmount() float64 {

	if m.fractionDigits == CurrencyFractions[m.currencyCode] {
		return float64(m.preciseAmount)
	} else {
		return float64(m.preciseAmount) / (math.Pow(10, float64(m.fractionDigits-CurrencyFractions[m.currencyCode])))
	}

}
func (m *Money) CurrencyCode() string {
	return m.currencyCode
}

func (m *Money) FractionDigits() int64 {
	return m.fractionDigits
}

func (m *Money) PreciseAmount() int64 {
	return m.preciseAmount
}

// UnmarshalJSON unmarshals JSON data into the Money object.
func (m *Money) UnmarshalJSON(data []byte) error {

	validate := validator.New()
	var tmp struct {
		CurrencyCode   string `json:"currencyCode" validate:"required" `
		FractionDigits int64  `json:"fractionDigits" validate:"lte=20"`
		PreciseAmount  int64  `json:"preciseAmount" validate:"required"`
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	err := validate.Struct(tmp)

	if err != nil {
		return &BadMoneyValue{}
	}

	if _, ok := CurrencyFractions[tmp.CurrencyCode]; !ok {
		return &CurrencyNotSupported{}
	}

	m.currencyCode = tmp.CurrencyCode

	if tmp.FractionDigits == 0 {
		m.fractionDigits = CurrencyFractions[tmp.CurrencyCode]
	} else if tmp.FractionDigits < CurrencyFractions[tmp.CurrencyCode] {
		return &BadMoneyValue{}
	} else {
		m.fractionDigits = tmp.FractionDigits
	}

	m.preciseAmount = tmp.PreciseAmount

	return nil
}

// NewMoney creates a new Money object with the given  currency code, fraction digits, and type.
func NewMoney(preciseAmount int64, fractionDigits int64, currencyCode string) (money *Money, err error) {

	if _, ok := CurrencyFractions[currencyCode]; !ok {
		return nil, &CurrencyNotSupported{}
	}

	if fractionDigits > 20 {
		return nil, &BadMoneyValue{}
	}
	if fractionDigits == 0 {
		fractionDigits = CurrencyFractions[currencyCode]
	} else if fractionDigits < CurrencyFractions[currencyCode] {
		return nil, &BadMoneyValue{}
	}
	return &Money{

		currencyCode:   currencyCode,
		fractionDigits: fractionDigits,
		preciseAmount:  preciseAmount,
	}, nil
}

// MarshalJSON marshals the Money object to JSON.
func (m *Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{

		"currencyCode":   m.currencyCode,
		"fractionDigits": m.fractionDigits,
		"preciseAmount":  m.preciseAmount,
	})
}
