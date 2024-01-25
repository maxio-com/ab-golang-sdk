package models

import (
	"encoding/json"
)

// CouponCurrencyRequest represents a CouponCurrencyRequest struct.
type CouponCurrencyRequest struct {
	CurrencyPrices []UpdateCouponCurrency `json:"currency_prices"`
}

// MarshalJSON implements the json.Marshaler interface for CouponCurrencyRequest.
// It customizes the JSON marshaling process for CouponCurrencyRequest objects.
func (c *CouponCurrencyRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CouponCurrencyRequest object to a map representation for JSON marshaling.
func (c *CouponCurrencyRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["currency_prices"] = c.CurrencyPrices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponCurrencyRequest.
// It customizes the JSON unmarshaling process for CouponCurrencyRequest objects.
func (c *CouponCurrencyRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CurrencyPrices []UpdateCouponCurrency `json:"currency_prices"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CurrencyPrices = temp.CurrencyPrices
	return nil
}
