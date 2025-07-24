// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CreateInvoiceItemProductPricePointId represents a CreateInvoiceItemProductPricePointId struct.
// This is a container for one-of cases.
type CreateInvoiceItemProductPricePointId struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for CreateInvoiceItemProductPricePointId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateInvoiceItemProductPricePointId) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceItemProductPricePointId.
// It customizes the JSON marshaling process for CreateInvoiceItemProductPricePointId objects.
func (c CreateInvoiceItemProductPricePointId) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateInvoiceItemProductPricePointIdContainer.From*` functions to initialize the CreateInvoiceItemProductPricePointId object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceItemProductPricePointId object to a map representation for JSON marshaling.
func (c *CreateInvoiceItemProductPricePointId) toMap() any {
    switch obj := c.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceItemProductPricePointId.
// It customizes the JSON unmarshaling process for CreateInvoiceItemProductPricePointId objects.
func (c *CreateInvoiceItemProductPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(int), false, &c.isNumber),
    )
    
    c.value = result
    return err
}

func (c *CreateInvoiceItemProductPricePointId) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CreateInvoiceItemProductPricePointId) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

// internalCreateInvoiceItemProductPricePointId represents a createInvoiceItemProductPricePointId struct.
// This is a container for one-of cases.
type internalCreateInvoiceItemProductPricePointId struct {}

var CreateInvoiceItemProductPricePointIdContainer internalCreateInvoiceItemProductPricePointId

// The internalCreateInvoiceItemProductPricePointId instance, wrapping the provided string value.
func (c *internalCreateInvoiceItemProductPricePointId) FromString(val string) CreateInvoiceItemProductPricePointId {
    return CreateInvoiceItemProductPricePointId{value: &val}
}

// The internalCreateInvoiceItemProductPricePointId instance, wrapping the provided int value.
func (c *internalCreateInvoiceItemProductPricePointId) FromNumber(val int) CreateInvoiceItemProductPricePointId {
    return CreateInvoiceItemProductPricePointId{value: &val}
}
