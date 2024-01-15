package models

import (
    "encoding/json"
)

// CreateCurrencyPricesRequest represents a CreateCurrencyPricesRequest struct.
type CreateCurrencyPricesRequest struct {
    CurrencyPrices []CreateCurrencyPrice `json:"currency_prices"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCurrencyPricesRequest.
// It customizes the JSON marshaling process for CreateCurrencyPricesRequest objects.
func (c *CreateCurrencyPricesRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateCurrencyPricesRequest object to a map representation for JSON marshaling.
func (c *CreateCurrencyPricesRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["currency_prices"] = c.CurrencyPrices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCurrencyPricesRequest.
// It customizes the JSON unmarshaling process for CreateCurrencyPricesRequest objects.
func (c *CreateCurrencyPricesRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        CurrencyPrices []CreateCurrencyPrice `json:"currency_prices"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.CurrencyPrices = temp.CurrencyPrices
    return nil
}
