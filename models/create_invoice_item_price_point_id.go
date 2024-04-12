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

// CreateInvoiceItemPricePointId represents a CreateInvoiceItemPricePointId struct.
// This is a container for one-of cases.
type CreateInvoiceItemPricePointId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the CreateInvoiceItemPricePointId object to a string representation.
func (c CreateInvoiceItemPricePointId) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceItemPricePointId.
// It customizes the JSON marshaling process for CreateInvoiceItemPricePointId objects.
func (c CreateInvoiceItemPricePointId) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateInvoiceItemPricePointIdContainer.From*` functions to initialize the CreateInvoiceItemPricePointId object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceItemPricePointId object to a map representation for JSON marshaling.
func (c *CreateInvoiceItemPricePointId) toMap() any {
    switch obj := c.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceItemPricePointId.
// It customizes the JSON unmarshaling process for CreateInvoiceItemPricePointId objects.
func (c *CreateInvoiceItemPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(int), false, &c.isNumber),
    )
    
    c.value = result
    return err
}

func (c *CreateInvoiceItemPricePointId) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CreateInvoiceItemPricePointId) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

// internalCreateInvoiceItemPricePointId represents a createInvoiceItemPricePointId struct.
// This is a container for one-of cases.
type internalCreateInvoiceItemPricePointId struct {}

var CreateInvoiceItemPricePointIdContainer internalCreateInvoiceItemPricePointId

// The internalCreateInvoiceItemPricePointId instance, wrapping the provided string value.
func (c *internalCreateInvoiceItemPricePointId) FromString(val string) CreateInvoiceItemPricePointId {
    return CreateInvoiceItemPricePointId{value: &val}
}

// The internalCreateInvoiceItemPricePointId instance, wrapping the provided int value.
func (c *internalCreateInvoiceItemPricePointId) FromNumber(val int) CreateInvoiceItemPricePointId {
    return CreateInvoiceItemPricePointId{value: &val}
}
