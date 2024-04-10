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

// CreateAllocationPricePointId represents a CreateAllocationPricePointId struct.
// This is a container for one-of cases.
type CreateAllocationPricePointId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the CreateAllocationPricePointId object to a string representation.
func (c CreateAllocationPricePointId) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateAllocationPricePointId.
// It customizes the JSON marshaling process for CreateAllocationPricePointId objects.
func (c CreateAllocationPricePointId) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateAllocationPricePointIdContainer.From*` functions to initialize the CreateAllocationPricePointId object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateAllocationPricePointId object to a map representation for JSON marshaling.
func (c *CreateAllocationPricePointId) toMap() any {
    switch obj := c.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateAllocationPricePointId.
// It customizes the JSON unmarshaling process for CreateAllocationPricePointId objects.
func (c *CreateAllocationPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(int), false, &c.isNumber),
    )
    
    c.value = result
    return err
}

func (c *CreateAllocationPricePointId) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CreateAllocationPricePointId) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

// internalCreateAllocationPricePointId represents a createAllocationPricePointId struct.
// This is a container for one-of cases.
type internalCreateAllocationPricePointId struct {}

var CreateAllocationPricePointIdContainer internalCreateAllocationPricePointId

// The internalCreateAllocationPricePointId instance, wrapping the provided string value.
func (c *internalCreateAllocationPricePointId) FromString(val string) CreateAllocationPricePointId {
    return CreateAllocationPricePointId{value: &val}
}

// The internalCreateAllocationPricePointId instance, wrapping the provided int value.
func (c *internalCreateAllocationPricePointId) FromNumber(val int) CreateAllocationPricePointId {
    return CreateAllocationPricePointId{value: &val}
}
