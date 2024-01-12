package models

import (
    "encoding/json"
)

// ProductPricePointCurrencyPrice represents a ProductPricePointCurrencyPrice struct.
type ProductPricePointCurrencyPrice struct {
    CurrencyPrices []CurrencyPrice `json:"currency_prices"`
}

// MarshalJSON implements the json.Marshaler interface for ProductPricePointCurrencyPrice.
// It customizes the JSON marshaling process for ProductPricePointCurrencyPrice objects.
func (p *ProductPricePointCurrencyPrice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductPricePointCurrencyPrice object to a map representation for JSON marshaling.
func (p *ProductPricePointCurrencyPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["currency_prices"] = p.CurrencyPrices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductPricePointCurrencyPrice.
// It customizes the JSON unmarshaling process for ProductPricePointCurrencyPrice objects.
func (p *ProductPricePointCurrencyPrice) UnmarshalJSON(input []byte) error {
    temp := &struct {
        CurrencyPrices []CurrencyPrice `json:"currency_prices"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.CurrencyPrices = temp.CurrencyPrices
    return nil
}
