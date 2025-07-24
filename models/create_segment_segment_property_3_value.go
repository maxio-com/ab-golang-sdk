// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CreateSegmentSegmentProperty3Value represents a CreateSegmentSegmentProperty3Value struct.
// This is a container for one-of cases.
type CreateSegmentSegmentProperty3Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
    isBoolean   bool
}

// String implements the fmt.Stringer interface for CreateSegmentSegmentProperty3Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateSegmentSegmentProperty3Value) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateSegmentSegmentProperty3Value.
// It customizes the JSON marshaling process for CreateSegmentSegmentProperty3Value objects.
func (c CreateSegmentSegmentProperty3Value) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateSegmentSegmentProperty3ValueContainer.From*` functions to initialize the CreateSegmentSegmentProperty3Value object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSegmentSegmentProperty3Value object to a map representation for JSON marshaling.
func (c *CreateSegmentSegmentProperty3Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSegmentSegmentProperty3Value.
// It customizes the JSON unmarshaling process for CreateSegmentSegmentProperty3Value objects.
func (c *CreateSegmentSegmentProperty3Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(float64), false, &c.isPrecision),
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(bool), false, &c.isBoolean),
    )
    
    c.value = result
    return err
}

func (c *CreateSegmentSegmentProperty3Value) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CreateSegmentSegmentProperty3Value) AsPrecision() (
    *float64,
    bool) {
    if !c.isPrecision {
        return nil, false
    }
    return c.value.(*float64), true
}

func (c *CreateSegmentSegmentProperty3Value) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CreateSegmentSegmentProperty3Value) AsBoolean() (
    *bool,
    bool) {
    if !c.isBoolean {
        return nil, false
    }
    return c.value.(*bool), true
}

// internalCreateSegmentSegmentProperty3Value represents a createSegmentSegmentProperty3Value struct.
// This is a container for one-of cases.
type internalCreateSegmentSegmentProperty3Value struct {}

var CreateSegmentSegmentProperty3ValueContainer internalCreateSegmentSegmentProperty3Value

// The internalCreateSegmentSegmentProperty3Value instance, wrapping the provided string value.
func (c *internalCreateSegmentSegmentProperty3Value) FromString(val string) CreateSegmentSegmentProperty3Value {
    return CreateSegmentSegmentProperty3Value{value: &val}
}

// The internalCreateSegmentSegmentProperty3Value instance, wrapping the provided float64 value.
func (c *internalCreateSegmentSegmentProperty3Value) FromPrecision(val float64) CreateSegmentSegmentProperty3Value {
    return CreateSegmentSegmentProperty3Value{value: &val}
}

// The internalCreateSegmentSegmentProperty3Value instance, wrapping the provided int value.
func (c *internalCreateSegmentSegmentProperty3Value) FromNumber(val int) CreateSegmentSegmentProperty3Value {
    return CreateSegmentSegmentProperty3Value{value: &val}
}

// The internalCreateSegmentSegmentProperty3Value instance, wrapping the provided bool value.
func (c *internalCreateSegmentSegmentProperty3Value) FromBoolean(val bool) CreateSegmentSegmentProperty3Value {
    return CreateSegmentSegmentProperty3Value{value: &val}
}
