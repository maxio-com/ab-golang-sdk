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

// CreateSubscriptionComponentPricePointId represents a CreateSubscriptionComponentPricePointId struct.
// This is a container for one-of cases.
type CreateSubscriptionComponentPricePointId struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the CreateSubscriptionComponentPricePointId object to a string representation.
func (c CreateSubscriptionComponentPricePointId) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionComponentPricePointId.
// It customizes the JSON marshaling process for CreateSubscriptionComponentPricePointId objects.
func (c CreateSubscriptionComponentPricePointId) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateSubscriptionComponentPricePointIdContainer.From*` functions to initialize the CreateSubscriptionComponentPricePointId object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionComponentPricePointId object to a map representation for JSON marshaling.
func (c *CreateSubscriptionComponentPricePointId) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionComponentPricePointId.
// It customizes the JSON unmarshaling process for CreateSubscriptionComponentPricePointId objects.
func (c *CreateSubscriptionComponentPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CreateSubscriptionComponentPricePointId) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CreateSubscriptionComponentPricePointId) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCreateSubscriptionComponentPricePointId represents a createSubscriptionComponentPricePointId struct.
// This is a container for one-of cases.
type internalCreateSubscriptionComponentPricePointId struct {}

var CreateSubscriptionComponentPricePointIdContainer internalCreateSubscriptionComponentPricePointId

// The internalCreateSubscriptionComponentPricePointId instance, wrapping the provided int value.
func (c *internalCreateSubscriptionComponentPricePointId) FromNumber(val int) CreateSubscriptionComponentPricePointId {
    return CreateSubscriptionComponentPricePointId{value: &val}
}

// The internalCreateSubscriptionComponentPricePointId instance, wrapping the provided string value.
func (c *internalCreateSubscriptionComponentPricePointId) FromString(val string) CreateSubscriptionComponentPricePointId {
    return CreateSubscriptionComponentPricePointId{value: &val}
}
