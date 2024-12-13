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

// ComponentCurrencyPricesResponse represents a ComponentCurrencyPricesResponse struct.
type ComponentCurrencyPricesResponse struct {
    CurrencyPrices       []ComponentCurrencyPrice `json:"currency_prices"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentCurrencyPricesResponse.
// It customizes the JSON marshaling process for ComponentCurrencyPricesResponse objects.
func (c ComponentCurrencyPricesResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "currency_prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentCurrencyPricesResponse object to a map representation for JSON marshaling.
func (c ComponentCurrencyPricesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["currency_prices"] = c.CurrencyPrices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentCurrencyPricesResponse.
// It customizes the JSON unmarshaling process for ComponentCurrencyPricesResponse objects.
func (c *ComponentCurrencyPricesResponse) UnmarshalJSON(input []byte) error {
    var temp tempComponentCurrencyPricesResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "currency_prices")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.CurrencyPrices = *temp.CurrencyPrices
    return nil
}

// tempComponentCurrencyPricesResponse is a temporary struct used for validating the fields of ComponentCurrencyPricesResponse.
type tempComponentCurrencyPricesResponse  struct {
    CurrencyPrices *[]ComponentCurrencyPrice `json:"currency_prices"`
}

func (c *tempComponentCurrencyPricesResponse) validate() error {
    var errs []string
    if c.CurrencyPrices == nil {
        errs = append(errs, "required field `currency_prices` is missing for type `Component Currency Prices Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
