// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CreateProductPricePointProductId represents a CreateProductPricePointProductId struct.
// This is a container for one-of cases.
type CreateProductPricePointProductId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for CreateProductPricePointProductId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateProductPricePointProductId) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateProductPricePointProductId.
// It customizes the JSON marshaling process for CreateProductPricePointProductId objects.
func (c CreateProductPricePointProductId) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateProductPricePointProductIdContainer.From*` functions to initialize the CreateProductPricePointProductId object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateProductPricePointProductId object to a map representation for JSON marshaling.
func (c *CreateProductPricePointProductId) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductPricePointProductId.
// It customizes the JSON unmarshaling process for CreateProductPricePointProductId objects.
func (c *CreateProductPricePointProductId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CreateProductPricePointProductId) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CreateProductPricePointProductId) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCreateProductPricePointProductId represents a createProductPricePointProductId struct.
// This is a container for one-of cases.
type internalCreateProductPricePointProductId struct {}

var CreateProductPricePointProductIdContainer internalCreateProductPricePointProductId

// The internalCreateProductPricePointProductId instance, wrapping the provided int value.
func (c *internalCreateProductPricePointProductId) FromNumber(val int) CreateProductPricePointProductId {
    return CreateProductPricePointProductId{value: &val}
}

// The internalCreateProductPricePointProductId instance, wrapping the provided string value.
func (c *internalCreateProductPricePointProductId) FromString(val string) CreateProductPricePointProductId {
    return CreateProductPricePointProductId{value: &val}
}
