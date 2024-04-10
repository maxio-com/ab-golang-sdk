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

// CreateSegmentSegmentProperty2Value represents a CreateSegmentSegmentProperty2Value struct.
// This is a container for one-of cases.
type CreateSegmentSegmentProperty2Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
    isBoolean   bool
}

// String converts the CreateSegmentSegmentProperty2Value object to a string representation.
func (c CreateSegmentSegmentProperty2Value) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateSegmentSegmentProperty2Value.
// It customizes the JSON marshaling process for CreateSegmentSegmentProperty2Value objects.
func (c CreateSegmentSegmentProperty2Value) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateSegmentSegmentProperty2ValueContainer.From*` functions to initialize the CreateSegmentSegmentProperty2Value object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSegmentSegmentProperty2Value object to a map representation for JSON marshaling.
func (c *CreateSegmentSegmentProperty2Value) toMap() any {
    switch obj := c.value.(type) {
    case *string:
        return *obj
    case *float64:
        return *obj
    case *int:
        return *obj
    case *bool:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSegmentSegmentProperty2Value.
// It customizes the JSON unmarshaling process for CreateSegmentSegmentProperty2Value objects.
func (c *CreateSegmentSegmentProperty2Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(float64), false, &c.isPrecision),
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(bool), false, &c.isBoolean),
    )
    
    c.value = result
    return err
}

func (c *CreateSegmentSegmentProperty2Value) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CreateSegmentSegmentProperty2Value) AsPrecision() (
    *float64,
    bool) {
    if !c.isPrecision {
        return nil, false
    }
    return c.value.(*float64), true
}

func (c *CreateSegmentSegmentProperty2Value) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CreateSegmentSegmentProperty2Value) AsBoolean() (
    *bool,
    bool) {
    if !c.isBoolean {
        return nil, false
    }
    return c.value.(*bool), true
}

// internalCreateSegmentSegmentProperty2Value represents a createSegmentSegmentProperty2Value struct.
// This is a container for one-of cases.
type internalCreateSegmentSegmentProperty2Value struct {}

var CreateSegmentSegmentProperty2ValueContainer internalCreateSegmentSegmentProperty2Value

// The internalCreateSegmentSegmentProperty2Value instance, wrapping the provided string value.
func (c *internalCreateSegmentSegmentProperty2Value) FromString(val string) CreateSegmentSegmentProperty2Value {
    return CreateSegmentSegmentProperty2Value{value: &val}
}

// The internalCreateSegmentSegmentProperty2Value instance, wrapping the provided float64 value.
func (c *internalCreateSegmentSegmentProperty2Value) FromPrecision(val float64) CreateSegmentSegmentProperty2Value {
    return CreateSegmentSegmentProperty2Value{value: &val}
}

// The internalCreateSegmentSegmentProperty2Value instance, wrapping the provided int value.
func (c *internalCreateSegmentSegmentProperty2Value) FromNumber(val int) CreateSegmentSegmentProperty2Value {
    return CreateSegmentSegmentProperty2Value{value: &val}
}

// The internalCreateSegmentSegmentProperty2Value instance, wrapping the provided bool value.
func (c *internalCreateSegmentSegmentProperty2Value) FromBoolean(val bool) CreateSegmentSegmentProperty2Value {
    return CreateSegmentSegmentProperty2Value{value: &val}
}
