/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateCurrencyPricesRequest represents a CreateCurrencyPricesRequest struct.
type CreateCurrencyPricesRequest struct {
    CurrencyPrices       []CreateCurrencyPrice  `json:"currency_prices"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateCurrencyPricesRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateCurrencyPricesRequest) String() string {
    return fmt.Sprintf(
    	"CreateCurrencyPricesRequest[CurrencyPrices=%v, AdditionalProperties=%v]",
    	c.CurrencyPrices, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateCurrencyPricesRequest.
// It customizes the JSON marshaling process for CreateCurrencyPricesRequest objects.
func (c CreateCurrencyPricesRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "currency_prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateCurrencyPricesRequest object to a map representation for JSON marshaling.
func (c CreateCurrencyPricesRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["currency_prices"] = c.CurrencyPrices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCurrencyPricesRequest.
// It customizes the JSON unmarshaling process for CreateCurrencyPricesRequest objects.
func (c *CreateCurrencyPricesRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateCurrencyPricesRequest
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

// tempCreateCurrencyPricesRequest is a temporary struct used for validating the fields of CreateCurrencyPricesRequest.
type tempCreateCurrencyPricesRequest  struct {
    CurrencyPrices *[]CreateCurrencyPrice `json:"currency_prices"`
}

func (c *tempCreateCurrencyPricesRequest) validate() error {
    var errs []string
    if c.CurrencyPrices == nil {
        errs = append(errs, "required field `currency_prices` is missing for type `Create Currency Prices Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
