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

// UpdateCouponCurrency represents a UpdateCouponCurrency struct.
type UpdateCouponCurrency struct {
    // ISO code for the site defined currency.
    Currency             string                 `json:"currency"`
    // Price for the given currency.
    Price                int                    `json:"price"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateCouponCurrency,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateCouponCurrency) String() string {
    return fmt.Sprintf(
    	"UpdateCouponCurrency[Currency=%v, Price=%v, AdditionalProperties=%v]",
    	u.Currency, u.Price, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateCouponCurrency.
// It customizes the JSON marshaling process for UpdateCouponCurrency objects.
func (u UpdateCouponCurrency) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "currency", "price"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateCouponCurrency object to a map representation for JSON marshaling.
func (u UpdateCouponCurrency) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["currency"] = u.Currency
    structMap["price"] = u.Price
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCouponCurrency.
// It customizes the JSON unmarshaling process for UpdateCouponCurrency objects.
func (u *UpdateCouponCurrency) UnmarshalJSON(input []byte) error {
    var temp tempUpdateCouponCurrency
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "currency", "price")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Currency = *temp.Currency
    u.Price = *temp.Price
    return nil
}

// tempUpdateCouponCurrency is a temporary struct used for validating the fields of UpdateCouponCurrency.
type tempUpdateCouponCurrency  struct {
    Currency *string `json:"currency"`
    Price    *int    `json:"price"`
}

func (u *tempUpdateCouponCurrency) validate() error {
    var errs []string
    if u.Currency == nil {
        errs = append(errs, "required field `currency` is missing for type `Update Coupon Currency`")
    }
    if u.Price == nil {
        errs = append(errs, "required field `price` is missing for type `Update Coupon Currency`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
