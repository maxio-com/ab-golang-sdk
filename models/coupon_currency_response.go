package models

import (
	"encoding/json"
)

// CouponCurrencyResponse represents a CouponCurrencyResponse struct.
type CouponCurrencyResponse struct {
	CurrencyPrices []CouponCurrency `json:"currency_prices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CouponCurrencyResponse.
// It customizes the JSON marshaling process for CouponCurrencyResponse objects.
func (c *CouponCurrencyResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CouponCurrencyResponse object to a map representation for JSON marshaling.
func (c *CouponCurrencyResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.CurrencyPrices != nil {
		structMap["currency_prices"] = c.CurrencyPrices
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponCurrencyResponse.
// It customizes the JSON unmarshaling process for CouponCurrencyResponse objects.
func (c *CouponCurrencyResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CurrencyPrices []CouponCurrency `json:"currency_prices,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CurrencyPrices = temp.CurrencyPrices
	return nil
}
