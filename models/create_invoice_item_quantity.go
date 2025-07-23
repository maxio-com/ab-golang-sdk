// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CreateInvoiceItemQuantity represents a CreateInvoiceItemQuantity struct.
// This is a container for one-of cases.
type CreateInvoiceItemQuantity struct {
    value       any
    isPrecision bool
    isString    bool
}

// String implements the fmt.Stringer interface for CreateInvoiceItemQuantity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateInvoiceItemQuantity) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceItemQuantity.
// It customizes the JSON marshaling process for CreateInvoiceItemQuantity objects.
func (c CreateInvoiceItemQuantity) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateInvoiceItemQuantityContainer.From*` functions to initialize the CreateInvoiceItemQuantity object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceItemQuantity object to a map representation for JSON marshaling.
func (c *CreateInvoiceItemQuantity) toMap() any {
    switch obj := c.value.(type) {
    case *float64:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceItemQuantity.
// It customizes the JSON unmarshaling process for CreateInvoiceItemQuantity objects.
func (c *CreateInvoiceItemQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(float64), false, &c.isPrecision),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CreateInvoiceItemQuantity) AsPrecision() (
    *float64,
    bool) {
    if !c.isPrecision {
        return nil, false
    }
    return c.value.(*float64), true
}

func (c *CreateInvoiceItemQuantity) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCreateInvoiceItemQuantity represents a createInvoiceItemQuantity struct.
// This is a container for one-of cases.
type internalCreateInvoiceItemQuantity struct {}

var CreateInvoiceItemQuantityContainer internalCreateInvoiceItemQuantity

// The internalCreateInvoiceItemQuantity instance, wrapping the provided float64 value.
func (c *internalCreateInvoiceItemQuantity) FromPrecision(val float64) CreateInvoiceItemQuantity {
    return CreateInvoiceItemQuantity{value: &val}
}

// The internalCreateInvoiceItemQuantity instance, wrapping the provided string value.
func (c *internalCreateInvoiceItemQuantity) FromString(val string) CreateInvoiceItemQuantity {
    return CreateInvoiceItemQuantity{value: &val}
}
