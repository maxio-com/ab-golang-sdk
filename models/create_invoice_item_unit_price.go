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

// CreateInvoiceItemUnitPrice represents a CreateInvoiceItemUnitPrice struct.
// This is a container for one-of cases.
type CreateInvoiceItemUnitPrice struct {
    value       any
    isPrecision bool
    isString    bool
}

// String converts the CreateInvoiceItemUnitPrice object to a string representation.
func (c CreateInvoiceItemUnitPrice) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceItemUnitPrice.
// It customizes the JSON marshaling process for CreateInvoiceItemUnitPrice objects.
func (c CreateInvoiceItemUnitPrice) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateInvoiceItemUnitPriceContainer.From*` functions to initialize the CreateInvoiceItemUnitPrice object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceItemUnitPrice object to a map representation for JSON marshaling.
func (c *CreateInvoiceItemUnitPrice) toMap() any {
    switch obj := c.value.(type) {
    case *float64:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceItemUnitPrice.
// It customizes the JSON unmarshaling process for CreateInvoiceItemUnitPrice objects.
func (c *CreateInvoiceItemUnitPrice) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(float64), false, &c.isPrecision),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CreateInvoiceItemUnitPrice) AsPrecision() (
    *float64,
    bool) {
    if !c.isPrecision {
        return nil, false
    }
    return c.value.(*float64), true
}

func (c *CreateInvoiceItemUnitPrice) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCreateInvoiceItemUnitPrice represents a createInvoiceItemUnitPrice struct.
// This is a container for one-of cases.
type internalCreateInvoiceItemUnitPrice struct {}

var CreateInvoiceItemUnitPriceContainer internalCreateInvoiceItemUnitPrice

// The internalCreateInvoiceItemUnitPrice instance, wrapping the provided float64 value.
func (c *internalCreateInvoiceItemUnitPrice) FromPrecision(val float64) CreateInvoiceItemUnitPrice {
    return CreateInvoiceItemUnitPrice{value: &val}
}

// The internalCreateInvoiceItemUnitPrice instance, wrapping the provided string value.
func (c *internalCreateInvoiceItemUnitPrice) FromString(val string) CreateInvoiceItemUnitPrice {
    return CreateInvoiceItemUnitPrice{value: &val}
}
