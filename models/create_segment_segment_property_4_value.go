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

// CreateSegmentSegmentProperty4Value represents a CreateSegmentSegmentProperty4Value struct.
// This is a container for one-of cases.
type CreateSegmentSegmentProperty4Value struct {
	value       any
	isString    bool
	isPrecision bool
	isNumber    bool
	isBoolean   bool
}

// String converts the CreateSegmentSegmentProperty4Value object to a string representation.
func (c CreateSegmentSegmentProperty4Value) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateSegmentSegmentProperty4Value.
// It customizes the JSON marshaling process for CreateSegmentSegmentProperty4Value objects.
func (c *CreateSegmentSegmentProperty4Value) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CreateSegmentSegmentProperty4ValueContainer.From*` functions to initialize the CreateSegmentSegmentProperty4Value object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSegmentSegmentProperty4Value object to a map representation for JSON marshaling.
func (c *CreateSegmentSegmentProperty4Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSegmentSegmentProperty4Value.
// It customizes the JSON unmarshaling process for CreateSegmentSegmentProperty4Value objects.
func (c *CreateSegmentSegmentProperty4Value) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &c.isString),
		NewTypeHolder(new(float64), false, &c.isPrecision),
		NewTypeHolder(new(int), false, &c.isNumber),
		NewTypeHolder(new(bool), false, &c.isBoolean),
	)

	c.value = result
	return err
}

func (c *CreateSegmentSegmentProperty4Value) AsString() (
	*string,
	bool) {
	if !c.isString {
		return nil, false
	}
	return c.value.(*string), true
}

func (c *CreateSegmentSegmentProperty4Value) AsPrecision() (
	*float64,
	bool) {
	if !c.isPrecision {
		return nil, false
	}
	return c.value.(*float64), true
}

func (c *CreateSegmentSegmentProperty4Value) AsNumber() (
	*int,
	bool) {
	if !c.isNumber {
		return nil, false
	}
	return c.value.(*int), true
}

func (c *CreateSegmentSegmentProperty4Value) AsBoolean() (
	*bool,
	bool) {
	if !c.isBoolean {
		return nil, false
	}
	return c.value.(*bool), true
}

// internalCreateSegmentSegmentProperty4Value represents a createSegmentSegmentProperty4Value struct.
// This is a container for one-of cases.
type internalCreateSegmentSegmentProperty4Value struct{}

var CreateSegmentSegmentProperty4ValueContainer internalCreateSegmentSegmentProperty4Value

func (c *internalCreateSegmentSegmentProperty4Value) FromString(val string) CreateSegmentSegmentProperty4Value {
	return CreateSegmentSegmentProperty4Value{value: &val}
}

func (c *internalCreateSegmentSegmentProperty4Value) FromPrecision(val float64) CreateSegmentSegmentProperty4Value {
	return CreateSegmentSegmentProperty4Value{value: &val}
}

func (c *internalCreateSegmentSegmentProperty4Value) FromNumber(val int) CreateSegmentSegmentProperty4Value {
	return CreateSegmentSegmentProperty4Value{value: &val}
}

func (c *internalCreateSegmentSegmentProperty4Value) FromBoolean(val bool) CreateSegmentSegmentProperty4Value {
	return CreateSegmentSegmentProperty4Value{value: &val}
}
