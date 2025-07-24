// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// CouponRequest represents a CouponRequest struct.
type CouponRequest struct {
    Coupon               *CouponPayload         `json:"coupon,omitempty"`
    // An object where the keys are product IDs or handles (prefixed with 'handle:'), and the values are booleans indicating if the coupon should be applicable to the product
    RestrictedProducts   map[string]bool        `json:"restricted_products,omitempty"`
    // An object where the keys are component IDs or handles (prefixed with 'handle:'), and the values are booleans indicating if the coupon should be applicable to the component
    RestrictedComponents map[string]bool        `json:"restricted_components,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CouponRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CouponRequest) String() string {
    return fmt.Sprintf(
    	"CouponRequest[Coupon=%v, RestrictedProducts=%v, RestrictedComponents=%v, AdditionalProperties=%v]",
    	c.Coupon, c.RestrictedProducts, c.RestrictedComponents, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CouponRequest.
// It customizes the JSON marshaling process for CouponRequest objects.
func (c CouponRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "coupon", "restricted_products", "restricted_components"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CouponRequest object to a map representation for JSON marshaling.
func (c CouponRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Coupon != nil {
        structMap["coupon"] = c.Coupon.toMap()
    }
    if c.RestrictedProducts != nil {
        structMap["restricted_products"] = c.RestrictedProducts
    }
    if c.RestrictedComponents != nil {
        structMap["restricted_components"] = c.RestrictedComponents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponRequest.
// It customizes the JSON unmarshaling process for CouponRequest objects.
func (c *CouponRequest) UnmarshalJSON(input []byte) error {
    var temp tempCouponRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "coupon", "restricted_products", "restricted_components")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Coupon = temp.Coupon
    c.RestrictedProducts = temp.RestrictedProducts
    c.RestrictedComponents = temp.RestrictedComponents
    return nil
}

// tempCouponRequest is a temporary struct used for validating the fields of CouponRequest.
type tempCouponRequest  struct {
    Coupon               *CouponPayload  `json:"coupon,omitempty"`
    RestrictedProducts   map[string]bool `json:"restricted_products,omitempty"`
    RestrictedComponents map[string]bool `json:"restricted_components,omitempty"`
}
