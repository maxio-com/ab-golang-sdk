/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CouponPayloadPercentage represents a CouponPayloadPercentage struct.
// This is a container for one-of cases.
type CouponPayloadPercentage struct {
    value       any
    isString    bool
    isPrecision bool
}

// String implements the fmt.Stringer interface for CouponPayloadPercentage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CouponPayloadPercentage) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CouponPayloadPercentage.
// It customizes the JSON marshaling process for CouponPayloadPercentage objects.
func (c CouponPayloadPercentage) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CouponPayloadPercentageContainer.From*` functions to initialize the CouponPayloadPercentage object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CouponPayloadPercentage object to a map representation for JSON marshaling.
func (c *CouponPayloadPercentage) toMap() any {
    switch obj := c.value.(type) {
    case *string:
        return *obj
    case *float64:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponPayloadPercentage.
// It customizes the JSON unmarshaling process for CouponPayloadPercentage objects.
func (c *CouponPayloadPercentage) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(float64), false, &c.isPrecision),
    )
    
    c.value = result
    return err
}

func (c *CouponPayloadPercentage) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CouponPayloadPercentage) AsPrecision() (
    *float64,
    bool) {
    if !c.isPrecision {
        return nil, false
    }
    return c.value.(*float64), true
}

// internalCouponPayloadPercentage represents a couponPayloadPercentage struct.
// This is a container for one-of cases.
type internalCouponPayloadPercentage struct {}

var CouponPayloadPercentageContainer internalCouponPayloadPercentage

// The internalCouponPayloadPercentage instance, wrapping the provided string value.
func (c *internalCouponPayloadPercentage) FromString(val string) CouponPayloadPercentage {
    return CouponPayloadPercentage{value: &val}
}

// The internalCouponPayloadPercentage instance, wrapping the provided float64 value.
func (c *internalCouponPayloadPercentage) FromPrecision(val float64) CouponPayloadPercentage {
    return CouponPayloadPercentage{value: &val}
}
