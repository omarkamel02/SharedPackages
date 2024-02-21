package SharedPackages

import (
	"encoding/json"
)

// MoneyType represents the type of money object.
type MoneyType string

const (
	MoneyTypeCentPrecision MoneyType = "CentPrecision"
	MoneyTypeHighPrecision MoneyType = "HighPrecision"
)

// Money represents a monetary value with a specific currency.
type Money struct {
	centAmount     int64
	currencyCode   string
	fractionDigits int64
	moneytype      MoneyType
	preciseAmount  int64
}

// UnmarshalJSON unmarshals JSON data into the Money object.
func (m *Money) UnmarshalJSON(data []byte) error {
	var tmp struct {
		CentAmount     *int64    `json:"centAmount,omitempty"`
		CurrencyCode   string    `json:"currencyCode"`
		FractionDigits int64     `json:"fractionDigits,omitempty"`
		Type           MoneyType `json:"type"`
		PreciseAmount  int64     `json:"preciseAmount,omitempty"`
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if _, ok := CurrencyFractions[tmp.CurrencyCode]; !ok {
		return &CurrencyNotSupported{}
	}

	if tmp.Type == MoneyTypeCentPrecision && tmp.CentAmount == nil {
		return &CentAmountMustPresent{}
	}

	if tmp.Type == MoneyTypeCentPrecision {
		tmp.FractionDigits = CurrencyFractions[tmp.CurrencyCode]
	} else {
		if tmp.PreciseAmount > 20 {
			return &TooHighFractionDigits{}
		}
	}

	if tmp.Type == MoneyTypeHighPrecision && tmp.FractionDigits <= CurrencyFractions[tmp.CurrencyCode] {
		return &TooLowFractionDigits{}
	}

	m.centAmount = *tmp.CentAmount
	m.currencyCode = tmp.CurrencyCode
	m.fractionDigits = tmp.FractionDigits
	m.moneytype = tmp.Type
	return nil
}

// NewMoney creates a new Money object with the given cent amount, currency code, fraction digits, and type.
func NewCentPrecisionMoney(centAmount int64, currencyCode string) (money *Money, err error) {

	if _, ok := CurrencyFractions[currencyCode]; !ok {
		return nil, &CurrencyNotSupported{}
	}

	return &Money{
		centAmount:     centAmount,
		currencyCode:   currencyCode,
		fractionDigits: CurrencyFractions[currencyCode],
		moneytype:      MoneyTypeCentPrecision,
		preciseAmount:  0,
	}, nil
}

// NewMoney creates a new Money object with the given cent amount, currency code, fraction digits, and type.
func NewHighPrecisionnMoney(currencyCode string, preciseAmount int64, fractionDigits int64) (money *Money, err error) {

	if _, ok := CurrencyFractions[currencyCode]; !ok {
		return nil, &CurrencyNotSupported{}
	}

	if fractionDigits <= CurrencyFractions[currencyCode] {
		return nil, &TooLowFractionDigits{}
	}

	if fractionDigits > 20 {
		return nil, &TooHighFractionDigits{}
	}
	return &Money{
		centAmount:     0,
		currencyCode:   currencyCode,
		fractionDigits: fractionDigits,
		moneytype:      MoneyTypeHighPrecision,
		preciseAmount:  preciseAmount,
	}, nil
}

// MarshalJSON marshals the Money object to JSON.
func (m *Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"centAmount":     m.centAmount,
		"currencyCode":   m.currencyCode,
		"fractionDigits": m.fractionDigits,
		"type":           m.moneytype,
	})
}
