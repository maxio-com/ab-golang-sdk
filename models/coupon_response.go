/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// CouponResponse represents a CouponResponse struct.
type CouponResponse struct {
    Coupon               *Coupon                `json:"coupon,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CouponResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CouponResponse) String() string {
    return fmt.Sprintf(
    	"CouponResponse[Coupon=%v, AdditionalProperties=%v]",
    	c.Coupon, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CouponResponse.
// It customizes the JSON marshaling process for CouponResponse objects.
func (c CouponResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "coupon"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CouponResponse object to a map representation for JSON marshaling.
func (c CouponResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Coupon != nil {
        structMap["coupon"] = c.Coupon.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponResponse.
// It customizes the JSON unmarshaling process for CouponResponse objects.
func (c *CouponResponse) UnmarshalJSON(input []byte) error {
    var temp tempCouponResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "coupon")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Coupon = temp.Coupon
    return nil
}

// tempCouponResponse is a temporary struct used for validating the fields of CouponResponse.
type tempCouponResponse  struct {
    Coupon *Coupon `json:"coupon,omitempty"`
}
