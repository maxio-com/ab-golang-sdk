// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// CouponCurrency represents a CouponCurrency struct.
type CouponCurrency struct {
    Id                   Optional[int]          `json:"id"`
    Currency             *string                `json:"currency,omitempty"`
    Price                Optional[float64]      `json:"price"`
    CouponId             *int                   `json:"coupon_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CouponCurrency,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CouponCurrency) String() string {
    return fmt.Sprintf(
    	"CouponCurrency[Id=%v, Currency=%v, Price=%v, CouponId=%v, AdditionalProperties=%v]",
    	c.Id, c.Currency, c.Price, c.CouponId, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CouponCurrency.
// It customizes the JSON marshaling process for CouponCurrency objects.
func (c CouponCurrency) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "id", "currency", "price", "coupon_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CouponCurrency object to a map representation for JSON marshaling.
func (c CouponCurrency) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id.IsValueSet() {
        if c.Id.Value() != nil {
            structMap["id"] = c.Id.Value()
        } else {
            structMap["id"] = nil
        }
    }
    if c.Currency != nil {
        structMap["currency"] = c.Currency
    }
    if c.Price.IsValueSet() {
        if c.Price.Value() != nil {
            structMap["price"] = c.Price.Value()
        } else {
            structMap["price"] = nil
        }
    }
    if c.CouponId != nil {
        structMap["coupon_id"] = c.CouponId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponCurrency.
// It customizes the JSON unmarshaling process for CouponCurrency objects.
func (c *CouponCurrency) UnmarshalJSON(input []byte) error {
    var temp tempCouponCurrency
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "currency", "price", "coupon_id")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Id = temp.Id
    c.Currency = temp.Currency
    c.Price = temp.Price
    c.CouponId = temp.CouponId
    return nil
}

// tempCouponCurrency is a temporary struct used for validating the fields of CouponCurrency.
type tempCouponCurrency  struct {
    Id       Optional[int]     `json:"id"`
    Currency *string           `json:"currency,omitempty"`
    Price    Optional[float64] `json:"price"`
    CouponId *int              `json:"coupon_id,omitempty"`
}
