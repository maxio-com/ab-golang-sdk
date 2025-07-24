// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CurrencyPricesResponse represents a CurrencyPricesResponse struct.
type CurrencyPricesResponse struct {
    CurrencyPrices       []CurrencyPrice        `json:"currency_prices"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CurrencyPricesResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CurrencyPricesResponse) String() string {
    return fmt.Sprintf(
    	"CurrencyPricesResponse[CurrencyPrices=%v, AdditionalProperties=%v]",
    	c.CurrencyPrices, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CurrencyPricesResponse.
// It customizes the JSON marshaling process for CurrencyPricesResponse objects.
func (c CurrencyPricesResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "currency_prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CurrencyPricesResponse object to a map representation for JSON marshaling.
func (c CurrencyPricesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["currency_prices"] = c.CurrencyPrices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CurrencyPricesResponse.
// It customizes the JSON unmarshaling process for CurrencyPricesResponse objects.
func (c *CurrencyPricesResponse) UnmarshalJSON(input []byte) error {
    var temp tempCurrencyPricesResponse
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

// tempCurrencyPricesResponse is a temporary struct used for validating the fields of CurrencyPricesResponse.
type tempCurrencyPricesResponse  struct {
    CurrencyPrices *[]CurrencyPrice `json:"currency_prices"`
}

func (c *tempCurrencyPricesResponse) validate() error {
    var errs []string
    if c.CurrencyPrices == nil {
        errs = append(errs, "required field `currency_prices` is missing for type `Currency Prices Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
