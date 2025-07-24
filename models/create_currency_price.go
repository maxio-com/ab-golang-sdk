// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// CreateCurrencyPrice represents a CreateCurrencyPrice struct.
type CreateCurrencyPrice struct {
    // ISO code for a currency defined on the site level
    Currency             *string                `json:"currency,omitempty"`
    // Price for the price level in this currency
    Price                *float64               `json:"price,omitempty"`
    // ID of the price that this corresponds with
    PriceId              *int                   `json:"price_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateCurrencyPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateCurrencyPrice) String() string {
    return fmt.Sprintf(
    	"CreateCurrencyPrice[Currency=%v, Price=%v, PriceId=%v, AdditionalProperties=%v]",
    	c.Currency, c.Price, c.PriceId, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateCurrencyPrice.
// It customizes the JSON marshaling process for CreateCurrencyPrice objects.
func (c CreateCurrencyPrice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "currency", "price", "price_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateCurrencyPrice object to a map representation for JSON marshaling.
func (c CreateCurrencyPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Currency != nil {
        structMap["currency"] = c.Currency
    }
    if c.Price != nil {
        structMap["price"] = c.Price
    }
    if c.PriceId != nil {
        structMap["price_id"] = c.PriceId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCurrencyPrice.
// It customizes the JSON unmarshaling process for CreateCurrencyPrice objects.
func (c *CreateCurrencyPrice) UnmarshalJSON(input []byte) error {
    var temp tempCreateCurrencyPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "currency", "price", "price_id")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Currency = temp.Currency
    c.Price = temp.Price
    c.PriceId = temp.PriceId
    return nil
}

// tempCreateCurrencyPrice is a temporary struct used for validating the fields of CreateCurrencyPrice.
type tempCreateCurrencyPrice  struct {
    Currency *string  `json:"currency,omitempty"`
    Price    *float64 `json:"price,omitempty"`
    PriceId  *int     `json:"price_id,omitempty"`
}
