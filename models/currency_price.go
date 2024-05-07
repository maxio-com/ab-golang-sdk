package models

import (
    "encoding/json"
)

// CurrencyPrice represents a CurrencyPrice struct.
type CurrencyPrice struct {
    Id                   *int               `json:"id,omitempty"`
    Currency             *string            `json:"currency,omitempty"`
    Price                *float64           `json:"price,omitempty"`
    FormattedPrice       *string            `json:"formatted_price,omitempty"`
    ProductPricePointId  *int               `json:"product_price_point_id,omitempty"`
    // Role for the price.
    Role                 *CurrencyPriceRole `json:"role,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CurrencyPrice.
// It customizes the JSON marshaling process for CurrencyPrice objects.
func (c CurrencyPrice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CurrencyPrice object to a map representation for JSON marshaling.
func (c CurrencyPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.Currency != nil {
        structMap["currency"] = c.Currency
    }
    if c.Price != nil {
        structMap["price"] = c.Price
    }
    if c.FormattedPrice != nil {
        structMap["formatted_price"] = c.FormattedPrice
    }
    if c.ProductPricePointId != nil {
        structMap["product_price_point_id"] = c.ProductPricePointId
    }
    if c.Role != nil {
        structMap["role"] = c.Role
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CurrencyPrice.
// It customizes the JSON unmarshaling process for CurrencyPrice objects.
func (c *CurrencyPrice) UnmarshalJSON(input []byte) error {
    var temp currencyPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "currency", "price", "formatted_price", "product_price_point_id", "role")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Id = temp.Id
    c.Currency = temp.Currency
    c.Price = temp.Price
    c.FormattedPrice = temp.FormattedPrice
    c.ProductPricePointId = temp.ProductPricePointId
    c.Role = temp.Role
    return nil
}

// currencyPrice is a temporary struct used for validating the fields of CurrencyPrice.
type currencyPrice  struct {
    Id                  *int               `json:"id,omitempty"`
    Currency            *string            `json:"currency,omitempty"`
    Price               *float64           `json:"price,omitempty"`
    FormattedPrice      *string            `json:"formatted_price,omitempty"`
    ProductPricePointId *int               `json:"product_price_point_id,omitempty"`
    Role                *CurrencyPriceRole `json:"role,omitempty"`
}
