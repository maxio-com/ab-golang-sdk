// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UpdateCurrencyPricesRequest represents a UpdateCurrencyPricesRequest struct.
type UpdateCurrencyPricesRequest struct {
    CurrencyPrices       []UpdateCurrencyPrice  `json:"currency_prices"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateCurrencyPricesRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateCurrencyPricesRequest) String() string {
    return fmt.Sprintf(
    	"UpdateCurrencyPricesRequest[CurrencyPrices=%v, AdditionalProperties=%v]",
    	u.CurrencyPrices, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateCurrencyPricesRequest.
// It customizes the JSON marshaling process for UpdateCurrencyPricesRequest objects.
func (u UpdateCurrencyPricesRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "currency_prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateCurrencyPricesRequest object to a map representation for JSON marshaling.
func (u UpdateCurrencyPricesRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["currency_prices"] = u.CurrencyPrices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCurrencyPricesRequest.
// It customizes the JSON unmarshaling process for UpdateCurrencyPricesRequest objects.
func (u *UpdateCurrencyPricesRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateCurrencyPricesRequest
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
    u.AdditionalProperties = additionalProperties
    
    u.CurrencyPrices = *temp.CurrencyPrices
    return nil
}

// tempUpdateCurrencyPricesRequest is a temporary struct used for validating the fields of UpdateCurrencyPricesRequest.
type tempUpdateCurrencyPricesRequest  struct {
    CurrencyPrices *[]UpdateCurrencyPrice `json:"currency_prices"`
}

func (u *tempUpdateCurrencyPricesRequest) validate() error {
    var errs []string
    if u.CurrencyPrices == nil {
        errs = append(errs, "required field `currency_prices` is missing for type `Update Currency Prices Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
