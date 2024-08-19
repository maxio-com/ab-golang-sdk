/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateProductCurrencyPrice represents a CreateProductCurrencyPrice struct.
type CreateProductCurrencyPrice struct {
    // ISO code for one of the site level currencies.
    Currency             string            `json:"currency"`
    // Price for the given role.
    Price                int               `json:"price"`
    // Role for the price.
    Role                 CurrencyPriceRole `json:"role"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateProductCurrencyPrice.
// It customizes the JSON marshaling process for CreateProductCurrencyPrice objects.
func (c CreateProductCurrencyPrice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateProductCurrencyPrice object to a map representation for JSON marshaling.
func (c CreateProductCurrencyPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["currency"] = c.Currency
    structMap["price"] = c.Price
    structMap["role"] = c.Role
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductCurrencyPrice.
// It customizes the JSON unmarshaling process for CreateProductCurrencyPrice objects.
func (c *CreateProductCurrencyPrice) UnmarshalJSON(input []byte) error {
    var temp tempCreateProductCurrencyPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "currency", "price", "role")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Currency = *temp.Currency
    c.Price = *temp.Price
    c.Role = *temp.Role
    return nil
}

// tempCreateProductCurrencyPrice is a temporary struct used for validating the fields of CreateProductCurrencyPrice.
type tempCreateProductCurrencyPrice  struct {
    Currency *string            `json:"currency"`
    Price    *int               `json:"price"`
    Role     *CurrencyPriceRole `json:"role"`
}

func (c *tempCreateProductCurrencyPrice) validate() error {
    var errs []string
    if c.Currency == nil {
        errs = append(errs, "required field `currency` is missing for type `Create Product Currency Price`")
    }
    if c.Price == nil {
        errs = append(errs, "required field `price` is missing for type `Create Product Currency Price`")
    }
    if c.Role == nil {
        errs = append(errs, "required field `role` is missing for type `Create Product Currency Price`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
