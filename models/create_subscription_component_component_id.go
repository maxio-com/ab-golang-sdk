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

// CreateSubscriptionComponentComponentId represents a CreateSubscriptionComponentComponentId struct.
// This is a container for one-of cases.
type CreateSubscriptionComponentComponentId struct {
	value    any
	isNumber bool
	isString bool
}

// String converts the CreateSubscriptionComponentComponentId object to a string representation.
func (c CreateSubscriptionComponentComponentId) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionComponentComponentId.
// It customizes the JSON marshaling process for CreateSubscriptionComponentComponentId objects.
func (c *CreateSubscriptionComponentComponentId) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CreateSubscriptionComponentComponentIdContainer.From*` functions to initialize the CreateSubscriptionComponentComponentId object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionComponentComponentId object to a map representation for JSON marshaling.
func (c *CreateSubscriptionComponentComponentId) toMap() any {
	switch obj := c.value.(type) {
	case *int:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionComponentComponentId.
// It customizes the JSON unmarshaling process for CreateSubscriptionComponentComponentId objects.
func (c *CreateSubscriptionComponentComponentId) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(int), false, &c.isNumber),
		NewTypeHolder(new(string), false, &c.isString),
	)

	c.value = result
	return err
}

func (c *CreateSubscriptionComponentComponentId) AsNumber() (
	*int,
	bool) {
	if !c.isNumber {
		return nil, false
	}
	return c.value.(*int), true
}

func (c *CreateSubscriptionComponentComponentId) AsString() (
	*string,
	bool) {
	if !c.isString {
		return nil, false
	}
	return c.value.(*string), true
}

// internalCreateSubscriptionComponentComponentId represents a createSubscriptionComponentComponentId struct.
// This is a container for one-of cases.
type internalCreateSubscriptionComponentComponentId struct{}

var CreateSubscriptionComponentComponentIdContainer internalCreateSubscriptionComponentComponentId

func (c *internalCreateSubscriptionComponentComponentId) FromNumber(val int) CreateSubscriptionComponentComponentId {
	return CreateSubscriptionComponentComponentId{value: &val}
}

func (c *internalCreateSubscriptionComponentComponentId) FromString(val string) CreateSubscriptionComponentComponentId {
	return CreateSubscriptionComponentComponentId{value: &val}
}
