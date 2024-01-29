package models

import (
	"encoding/json"
)

// CreateProductCurrencyPricesRequest represents a CreateProductCurrencyPricesRequest struct.
type CreateProductCurrencyPricesRequest struct {
	CurrencyPrices []CreateProductCurrencyPrice `json:"currency_prices"`
}

// MarshalJSON implements the json.Marshaler interface for CreateProductCurrencyPricesRequest.
// It customizes the JSON marshaling process for CreateProductCurrencyPricesRequest objects.
func (c *CreateProductCurrencyPricesRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateProductCurrencyPricesRequest object to a map representation for JSON marshaling.
func (c *CreateProductCurrencyPricesRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["currency_prices"] = c.CurrencyPrices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductCurrencyPricesRequest.
// It customizes the JSON unmarshaling process for CreateProductCurrencyPricesRequest objects.
func (c *CreateProductCurrencyPricesRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CurrencyPrices []CreateProductCurrencyPrice `json:"currency_prices"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CurrencyPrices = temp.CurrencyPrices
	return nil
}
