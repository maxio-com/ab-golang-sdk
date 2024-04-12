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

// CreateSubscriptionOfferId represents a CreateSubscriptionOfferId struct.
// This is a container for one-of cases.
type CreateSubscriptionOfferId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the CreateSubscriptionOfferId object to a string representation.
func (c CreateSubscriptionOfferId) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionOfferId.
// It customizes the JSON marshaling process for CreateSubscriptionOfferId objects.
func (c CreateSubscriptionOfferId) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateSubscriptionOfferIdContainer.From*` functions to initialize the CreateSubscriptionOfferId object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionOfferId object to a map representation for JSON marshaling.
func (c *CreateSubscriptionOfferId) toMap() any {
    switch obj := c.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionOfferId.
// It customizes the JSON unmarshaling process for CreateSubscriptionOfferId objects.
func (c *CreateSubscriptionOfferId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(int), false, &c.isNumber),
    )
    
    c.value = result
    return err
}

func (c *CreateSubscriptionOfferId) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CreateSubscriptionOfferId) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

// internalCreateSubscriptionOfferId represents a createSubscriptionOfferId struct.
// This is a container for one-of cases.
type internalCreateSubscriptionOfferId struct {}

var CreateSubscriptionOfferIdContainer internalCreateSubscriptionOfferId

// The internalCreateSubscriptionOfferId instance, wrapping the provided string value.
func (c *internalCreateSubscriptionOfferId) FromString(val string) CreateSubscriptionOfferId {
    return CreateSubscriptionOfferId{value: &val}
}

// The internalCreateSubscriptionOfferId instance, wrapping the provided int value.
func (c *internalCreateSubscriptionOfferId) FromNumber(val int) CreateSubscriptionOfferId {
    return CreateSubscriptionOfferId{value: &val}
}
