package models

import (
	"encoding/json"
)

// CurrencyPricesResponse represents a CurrencyPricesResponse struct.
type CurrencyPricesResponse struct {
	CurrencyPrices []CurrencyPrice `json:"currency_prices"`
}

// MarshalJSON implements the json.Marshaler interface for CurrencyPricesResponse.
// It customizes the JSON marshaling process for CurrencyPricesResponse objects.
func (c *CurrencyPricesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CurrencyPricesResponse object to a map representation for JSON marshaling.
func (c *CurrencyPricesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["currency_prices"] = c.CurrencyPrices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CurrencyPricesResponse.
// It customizes the JSON unmarshaling process for CurrencyPricesResponse objects.
func (c *CurrencyPricesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CurrencyPrices []CurrencyPrice `json:"currency_prices"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CurrencyPrices = temp.CurrencyPrices
	return nil
}
