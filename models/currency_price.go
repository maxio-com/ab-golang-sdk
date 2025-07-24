// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// CurrencyPrice represents a CurrencyPrice struct.
type CurrencyPrice struct {
    Id                   *int                   `json:"id,omitempty"`
    Currency             *string                `json:"currency,omitempty"`
    Price                *float64               `json:"price,omitempty"`
    FormattedPrice       *string                `json:"formatted_price,omitempty"`
    PriceId              *int                   `json:"price_id,omitempty"`
    PricePointId         *int                   `json:"price_point_id,omitempty"`
    ProductPricePointId  *int                   `json:"product_price_point_id,omitempty"`
    // Role for the price.
    Role                 *CurrencyPriceRole     `json:"role,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CurrencyPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CurrencyPrice) String() string {
    return fmt.Sprintf(
    	"CurrencyPrice[Id=%v, Currency=%v, Price=%v, FormattedPrice=%v, PriceId=%v, PricePointId=%v, ProductPricePointId=%v, Role=%v, AdditionalProperties=%v]",
    	c.Id, c.Currency, c.Price, c.FormattedPrice, c.PriceId, c.PricePointId, c.ProductPricePointId, c.Role, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CurrencyPrice.
// It customizes the JSON marshaling process for CurrencyPrice objects.
func (c CurrencyPrice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "id", "currency", "price", "formatted_price", "price_id", "price_point_id", "product_price_point_id", "role"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CurrencyPrice object to a map representation for JSON marshaling.
func (c CurrencyPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    if c.PriceId != nil {
        structMap["price_id"] = c.PriceId
    }
    if c.PricePointId != nil {
        structMap["price_point_id"] = c.PricePointId
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
    var temp tempCurrencyPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "currency", "price", "formatted_price", "price_id", "price_point_id", "product_price_point_id", "role")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Id = temp.Id
    c.Currency = temp.Currency
    c.Price = temp.Price
    c.FormattedPrice = temp.FormattedPrice
    c.PriceId = temp.PriceId
    c.PricePointId = temp.PricePointId
    c.ProductPricePointId = temp.ProductPricePointId
    c.Role = temp.Role
    return nil
}

// tempCurrencyPrice is a temporary struct used for validating the fields of CurrencyPrice.
type tempCurrencyPrice  struct {
    Id                  *int               `json:"id,omitempty"`
    Currency            *string            `json:"currency,omitempty"`
    Price               *float64           `json:"price,omitempty"`
    FormattedPrice      *string            `json:"formatted_price,omitempty"`
    PriceId             *int               `json:"price_id,omitempty"`
    PricePointId        *int               `json:"price_point_id,omitempty"`
    ProductPricePointId *int               `json:"product_price_point_id,omitempty"`
    Role                *CurrencyPriceRole `json:"role,omitempty"`
}
