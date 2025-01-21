/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// CouponCurrencyResponse represents a CouponCurrencyResponse struct.
type CouponCurrencyResponse struct {
    CurrencyPrices       []CouponCurrency       `json:"currency_prices,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CouponCurrencyResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CouponCurrencyResponse) String() string {
    return fmt.Sprintf(
    	"CouponCurrencyResponse[CurrencyPrices=%v, AdditionalProperties=%v]",
    	c.CurrencyPrices, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CouponCurrencyResponse.
// It customizes the JSON marshaling process for CouponCurrencyResponse objects.
func (c CouponCurrencyResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "currency_prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CouponCurrencyResponse object to a map representation for JSON marshaling.
func (c CouponCurrencyResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.CurrencyPrices != nil {
        structMap["currency_prices"] = c.CurrencyPrices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponCurrencyResponse.
// It customizes the JSON unmarshaling process for CouponCurrencyResponse objects.
func (c *CouponCurrencyResponse) UnmarshalJSON(input []byte) error {
    var temp tempCouponCurrencyResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "currency_prices")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.CurrencyPrices = temp.CurrencyPrices
    return nil
}

// tempCouponCurrencyResponse is a temporary struct used for validating the fields of CouponCurrencyResponse.
type tempCouponCurrencyResponse  struct {
    CurrencyPrices []CouponCurrency `json:"currency_prices,omitempty"`
}
