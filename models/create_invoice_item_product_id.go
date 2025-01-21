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

// CreateInvoiceItemProductId represents a CreateInvoiceItemProductId struct.
// This is a container for one-of cases.
type CreateInvoiceItemProductId struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for CreateInvoiceItemProductId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateInvoiceItemProductId) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceItemProductId.
// It customizes the JSON marshaling process for CreateInvoiceItemProductId objects.
func (c CreateInvoiceItemProductId) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateInvoiceItemProductIdContainer.From*` functions to initialize the CreateInvoiceItemProductId object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceItemProductId object to a map representation for JSON marshaling.
func (c *CreateInvoiceItemProductId) toMap() any {
    switch obj := c.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceItemProductId.
// It customizes the JSON unmarshaling process for CreateInvoiceItemProductId objects.
func (c *CreateInvoiceItemProductId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(int), false, &c.isNumber),
    )
    
    c.value = result
    return err
}

func (c *CreateInvoiceItemProductId) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CreateInvoiceItemProductId) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

// internalCreateInvoiceItemProductId represents a createInvoiceItemProductId struct.
// This is a container for one-of cases.
type internalCreateInvoiceItemProductId struct {}

var CreateInvoiceItemProductIdContainer internalCreateInvoiceItemProductId

// The internalCreateInvoiceItemProductId instance, wrapping the provided string value.
func (c *internalCreateInvoiceItemProductId) FromString(val string) CreateInvoiceItemProductId {
    return CreateInvoiceItemProductId{value: &val}
}

// The internalCreateInvoiceItemProductId instance, wrapping the provided int value.
func (c *internalCreateInvoiceItemProductId) FromNumber(val int) CreateInvoiceItemProductId {
    return CreateInvoiceItemProductId{value: &val}
}
