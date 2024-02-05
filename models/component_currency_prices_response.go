package models

import (
    "encoding/json"
)

// ComponentCurrencyPricesResponse represents a ComponentCurrencyPricesResponse struct.
type ComponentCurrencyPricesResponse struct {
    CurrencyPrices []ComponentCurrencyPrice `json:"currency_prices"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentCurrencyPricesResponse.
// It customizes the JSON marshaling process for ComponentCurrencyPricesResponse objects.
func (c *ComponentCurrencyPricesResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentCurrencyPricesResponse object to a map representation for JSON marshaling.
func (c *ComponentCurrencyPricesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["currency_prices"] = c.CurrencyPrices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentCurrencyPricesResponse.
// It customizes the JSON unmarshaling process for ComponentCurrencyPricesResponse objects.
func (c *ComponentCurrencyPricesResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        CurrencyPrices []ComponentCurrencyPrice `json:"currency_prices"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.CurrencyPrices = temp.CurrencyPrices
    return nil
}
