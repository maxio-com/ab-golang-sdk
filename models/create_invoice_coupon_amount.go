// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CreateInvoiceCouponAmount represents a CreateInvoiceCouponAmount struct.
// This is a container for one-of cases.
type CreateInvoiceCouponAmount struct {
    value       any
    isString    bool
    isPrecision bool
}

// String implements the fmt.Stringer interface for CreateInvoiceCouponAmount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateInvoiceCouponAmount) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceCouponAmount.
// It customizes the JSON marshaling process for CreateInvoiceCouponAmount objects.
func (c CreateInvoiceCouponAmount) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateInvoiceCouponAmountContainer.From*` functions to initialize the CreateInvoiceCouponAmount object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceCouponAmount object to a map representation for JSON marshaling.
func (c *CreateInvoiceCouponAmount) toMap() any {
    switch obj := c.value.(type) {
    case *string:
        return *obj
    case *float64:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceCouponAmount.
// It customizes the JSON unmarshaling process for CreateInvoiceCouponAmount objects.
func (c *CreateInvoiceCouponAmount) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(float64), false, &c.isPrecision),
    )
    
    c.value = result
    return err
}

func (c *CreateInvoiceCouponAmount) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CreateInvoiceCouponAmount) AsPrecision() (
    *float64,
    bool) {
    if !c.isPrecision {
        return nil, false
    }
    return c.value.(*float64), true
}

// internalCreateInvoiceCouponAmount represents a createInvoiceCouponAmount struct.
// This is a container for one-of cases.
type internalCreateInvoiceCouponAmount struct {}

var CreateInvoiceCouponAmountContainer internalCreateInvoiceCouponAmount

// The internalCreateInvoiceCouponAmount instance, wrapping the provided string value.
func (c *internalCreateInvoiceCouponAmount) FromString(val string) CreateInvoiceCouponAmount {
    return CreateInvoiceCouponAmount{value: &val}
}

// The internalCreateInvoiceCouponAmount instance, wrapping the provided float64 value.
func (c *internalCreateInvoiceCouponAmount) FromPrecision(val float64) CreateInvoiceCouponAmount {
    return CreateInvoiceCouponAmount{value: &val}
}
