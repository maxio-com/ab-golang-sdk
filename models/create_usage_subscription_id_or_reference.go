// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CreateUsageSubscriptionIdOrReference represents a CreateUsageSubscriptionIdOrReference struct.
// This is a container for one-of cases.
type CreateUsageSubscriptionIdOrReference struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for CreateUsageSubscriptionIdOrReference,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateUsageSubscriptionIdOrReference) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateUsageSubscriptionIdOrReference.
// It customizes the JSON marshaling process for CreateUsageSubscriptionIdOrReference objects.
func (c CreateUsageSubscriptionIdOrReference) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateUsageSubscriptionIdOrReferenceContainer.From*` functions to initialize the CreateUsageSubscriptionIdOrReference object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateUsageSubscriptionIdOrReference object to a map representation for JSON marshaling.
func (c *CreateUsageSubscriptionIdOrReference) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateUsageSubscriptionIdOrReference.
// It customizes the JSON unmarshaling process for CreateUsageSubscriptionIdOrReference objects.
func (c *CreateUsageSubscriptionIdOrReference) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CreateUsageSubscriptionIdOrReference) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CreateUsageSubscriptionIdOrReference) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCreateUsageSubscriptionIdOrReference represents a createUsageSubscriptionIdOrReference struct.
// This is a container for one-of cases.
type internalCreateUsageSubscriptionIdOrReference struct {}

var CreateUsageSubscriptionIdOrReferenceContainer internalCreateUsageSubscriptionIdOrReference

// The internalCreateUsageSubscriptionIdOrReference instance, wrapping the provided int value.
func (c *internalCreateUsageSubscriptionIdOrReference) FromNumber(val int) CreateUsageSubscriptionIdOrReference {
    return CreateUsageSubscriptionIdOrReference{value: &val}
}

// The internalCreateUsageSubscriptionIdOrReference instance, wrapping the provided string value.
func (c *internalCreateUsageSubscriptionIdOrReference) FromString(val string) CreateUsageSubscriptionIdOrReference {
    return CreateUsageSubscriptionIdOrReference{value: &val}
}
