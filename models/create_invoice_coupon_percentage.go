/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateInvoiceCouponPercentage represents a CreateInvoiceCouponPercentage struct.
// This is a container for one-of cases.
type CreateInvoiceCouponPercentage struct {
    value       any
    isString    bool
    isPrecision bool
}

// String converts the CreateInvoiceCouponPercentage object to a string representation.
func (c CreateInvoiceCouponPercentage) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceCouponPercentage.
// It customizes the JSON marshaling process for CreateInvoiceCouponPercentage objects.
func (c CreateInvoiceCouponPercentage) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateInvoiceCouponPercentageContainer.From*` functions to initialize the CreateInvoiceCouponPercentage object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceCouponPercentage object to a map representation for JSON marshaling.
func (c *CreateInvoiceCouponPercentage) toMap() any {
    switch obj := c.value.(type) {
    case *string:
        return *obj
    case *float64:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceCouponPercentage.
// It customizes the JSON unmarshaling process for CreateInvoiceCouponPercentage objects.
func (c *CreateInvoiceCouponPercentage) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(float64), false, &c.isPrecision),
    )
    
    c.value = result
    return err
}

func (c *CreateInvoiceCouponPercentage) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CreateInvoiceCouponPercentage) AsPrecision() (
    *float64,
    bool) {
    if !c.isPrecision {
        return nil, false
    }
    return c.value.(*float64), true
}

// internalCreateInvoiceCouponPercentage represents a createInvoiceCouponPercentage struct.
// This is a container for one-of cases.
type internalCreateInvoiceCouponPercentage struct {}

var CreateInvoiceCouponPercentageContainer internalCreateInvoiceCouponPercentage

// The internalCreateInvoiceCouponPercentage instance, wrapping the provided string value.
func (c *internalCreateInvoiceCouponPercentage) FromString(val string) CreateInvoiceCouponPercentage {
    return CreateInvoiceCouponPercentage{value: &val}
}

// The internalCreateInvoiceCouponPercentage instance, wrapping the provided float64 value.
func (c *internalCreateInvoiceCouponPercentage) FromPrecision(val float64) CreateInvoiceCouponPercentage {
    return CreateInvoiceCouponPercentage{value: &val}
}
