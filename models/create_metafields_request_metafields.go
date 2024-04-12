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

// CreateMetafieldsRequestMetafields represents a CreateMetafieldsRequestMetafields struct.
// This is a container for one-of cases.
type CreateMetafieldsRequestMetafields struct {
    value                    any
    isCreateMetafield        bool
    isArrayOfCreateMetafield bool
}

// String converts the CreateMetafieldsRequestMetafields object to a string representation.
func (c CreateMetafieldsRequestMetafields) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateMetafieldsRequestMetafields.
// It customizes the JSON marshaling process for CreateMetafieldsRequestMetafields objects.
func (c CreateMetafieldsRequestMetafields) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateMetafieldsRequestMetafieldsContainer.From*` functions to initialize the CreateMetafieldsRequestMetafields object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMetafieldsRequestMetafields object to a map representation for JSON marshaling.
func (c *CreateMetafieldsRequestMetafields) toMap() any {
    switch obj := c.value.(type) {
    case *CreateMetafield:
        return obj.toMap()
    case *[]CreateMetafield:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMetafieldsRequestMetafields.
// It customizes the JSON unmarshaling process for CreateMetafieldsRequestMetafields objects.
func (c *CreateMetafieldsRequestMetafields) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&CreateMetafield{}, false, &c.isCreateMetafield),
        NewTypeHolder(&[]CreateMetafield{}, false, &c.isArrayOfCreateMetafield),
    )
    
    c.value = result
    return err
}

func (c *CreateMetafieldsRequestMetafields) AsCreateMetafield() (
    *CreateMetafield,
    bool) {
    if !c.isCreateMetafield {
        return nil, false
    }
    return c.value.(*CreateMetafield), true
}

func (c *CreateMetafieldsRequestMetafields) AsArrayOfCreateMetafield() (
    *[]CreateMetafield,
    bool) {
    if !c.isArrayOfCreateMetafield {
        return nil, false
    }
    return c.value.(*[]CreateMetafield), true
}

// internalCreateMetafieldsRequestMetafields represents a createMetafieldsRequestMetafields struct.
// This is a container for one-of cases.
type internalCreateMetafieldsRequestMetafields struct {}

var CreateMetafieldsRequestMetafieldsContainer internalCreateMetafieldsRequestMetafields

// The internalCreateMetafieldsRequestMetafields instance, wrapping the provided CreateMetafield value.
func (c *internalCreateMetafieldsRequestMetafields) FromCreateMetafield(val CreateMetafield) CreateMetafieldsRequestMetafields {
    return CreateMetafieldsRequestMetafields{value: &val}
}

// The internalCreateMetafieldsRequestMetafields instance, wrapping the provided []CreateMetafield value.
func (c *internalCreateMetafieldsRequestMetafields) FromArrayOfCreateMetafield(val []CreateMetafield) CreateMetafieldsRequestMetafields {
    return CreateMetafieldsRequestMetafields{value: &val}
}
