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

// CreateInvoicePaymentAmount represents a CreateInvoicePaymentAmount struct.
// This is a container for one-of cases.
type CreateInvoicePaymentAmount struct {
    value       any
    isString    bool
    isPrecision bool
}

// String implements the fmt.Stringer interface for CreateInvoicePaymentAmount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateInvoicePaymentAmount) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoicePaymentAmount.
// It customizes the JSON marshaling process for CreateInvoicePaymentAmount objects.
func (c CreateInvoicePaymentAmount) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateInvoicePaymentAmountContainer.From*` functions to initialize the CreateInvoicePaymentAmount object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoicePaymentAmount object to a map representation for JSON marshaling.
func (c *CreateInvoicePaymentAmount) toMap() any {
    switch obj := c.value.(type) {
    case *string:
        return *obj
    case *float64:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoicePaymentAmount.
// It customizes the JSON unmarshaling process for CreateInvoicePaymentAmount objects.
func (c *CreateInvoicePaymentAmount) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(float64), false, &c.isPrecision),
    )
    
    c.value = result
    return err
}

func (c *CreateInvoicePaymentAmount) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CreateInvoicePaymentAmount) AsPrecision() (
    *float64,
    bool) {
    if !c.isPrecision {
        return nil, false
    }
    return c.value.(*float64), true
}

// internalCreateInvoicePaymentAmount represents a createInvoicePaymentAmount struct.
// This is a container for one-of cases.
type internalCreateInvoicePaymentAmount struct {}

var CreateInvoicePaymentAmountContainer internalCreateInvoicePaymentAmount

// The internalCreateInvoicePaymentAmount instance, wrapping the provided string value.
func (c *internalCreateInvoicePaymentAmount) FromString(val string) CreateInvoicePaymentAmount {
    return CreateInvoicePaymentAmount{value: &val}
}

// The internalCreateInvoicePaymentAmount instance, wrapping the provided float64 value.
func (c *internalCreateInvoicePaymentAmount) FromPrecision(val float64) CreateInvoicePaymentAmount {
    return CreateInvoicePaymentAmount{value: &val}
}
