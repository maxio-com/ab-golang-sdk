// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// CreateUsageComponentId represents a CreateUsageComponentId struct.
// This is a container for one-of cases.
type CreateUsageComponentId struct {
	value    any
	isNumber bool
	isString bool
}

// String implements the fmt.Stringer interface for CreateUsageComponentId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateUsageComponentId) String() string {
	return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateUsageComponentId.
// It customizes the JSON marshaling process for CreateUsageComponentId objects.
func (c CreateUsageComponentId) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CreateUsageComponentIdContainer.From*` functions to initialize the CreateUsageComponentId object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateUsageComponentId object to a map representation for JSON marshaling.
func (c *CreateUsageComponentId) toMap() any {
	switch obj := c.value.(type) {
	case *int:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateUsageComponentId.
// It customizes the JSON unmarshaling process for CreateUsageComponentId objects.
func (c *CreateUsageComponentId) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(int), false, &c.isNumber),
		NewTypeHolder(new(string), false, &c.isString),
	)

	c.value = result
	return err
}

func (c *CreateUsageComponentId) AsNumber() (
	*int,
	bool) {
	if !c.isNumber {
		return nil, false
	}
	return c.value.(*int), true
}

func (c *CreateUsageComponentId) AsString() (
	*string,
	bool) {
	if !c.isString {
		return nil, false
	}
	return c.value.(*string), true
}

// internalCreateUsageComponentId represents a createUsageComponentId struct.
// This is a container for one-of cases.
type internalCreateUsageComponentId struct{}

var CreateUsageComponentIdContainer internalCreateUsageComponentId

// The internalCreateUsageComponentId instance, wrapping the provided int value.
func (c *internalCreateUsageComponentId) FromNumber(val int) CreateUsageComponentId {
	return CreateUsageComponentId{value: &val}
}

// The internalCreateUsageComponentId instance, wrapping the provided string value.
func (c *internalCreateUsageComponentId) FromString(val string) CreateUsageComponentId {
	return CreateUsageComponentId{value: &val}
}
