// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// CouponCurrencyRequest represents a CouponCurrencyRequest struct.
type CouponCurrencyRequest struct {
	CurrencyPrices       []UpdateCouponCurrency `json:"currency_prices"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CouponCurrencyRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CouponCurrencyRequest) String() string {
	return fmt.Sprintf(
		"CouponCurrencyRequest[CurrencyPrices=%v, AdditionalProperties=%v]",
		c.CurrencyPrices, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CouponCurrencyRequest.
// It customizes the JSON marshaling process for CouponCurrencyRequest objects.
func (c CouponCurrencyRequest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"currency_prices"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CouponCurrencyRequest object to a map representation for JSON marshaling.
func (c CouponCurrencyRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["currency_prices"] = c.CurrencyPrices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponCurrencyRequest.
// It customizes the JSON unmarshaling process for CouponCurrencyRequest objects.
func (c *CouponCurrencyRequest) UnmarshalJSON(input []byte) error {
	var temp tempCouponCurrencyRequest
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

// tempCouponCurrencyRequest is a temporary struct used for validating the fields of CouponCurrencyRequest.
type tempCouponCurrencyRequest struct {
	CurrencyPrices *[]UpdateCouponCurrency `json:"currency_prices"`
}

func (c *tempCouponCurrencyRequest) validate() error {
	var errs []string
	if c.CurrencyPrices == nil {
		errs = append(errs, "required field `currency_prices` is missing for type `Coupon Currency Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
