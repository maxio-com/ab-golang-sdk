// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CreateSubscriptionComponentUnitBalance represents a CreateSubscriptionComponentUnitBalance struct.
// This is a container for one-of cases.
type CreateSubscriptionComponentUnitBalance struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for CreateSubscriptionComponentUnitBalance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateSubscriptionComponentUnitBalance) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionComponentUnitBalance.
// It customizes the JSON marshaling process for CreateSubscriptionComponentUnitBalance objects.
func (c CreateSubscriptionComponentUnitBalance) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateSubscriptionComponentUnitBalanceContainer.From*` functions to initialize the CreateSubscriptionComponentUnitBalance object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionComponentUnitBalance object to a map representation for JSON marshaling.
func (c *CreateSubscriptionComponentUnitBalance) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionComponentUnitBalance.
// It customizes the JSON unmarshaling process for CreateSubscriptionComponentUnitBalance objects.
func (c *CreateSubscriptionComponentUnitBalance) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CreateSubscriptionComponentUnitBalance) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CreateSubscriptionComponentUnitBalance) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCreateSubscriptionComponentUnitBalance represents a createSubscriptionComponentUnitBalance struct.
// This is a container for one-of cases.
type internalCreateSubscriptionComponentUnitBalance struct {}

var CreateSubscriptionComponentUnitBalanceContainer internalCreateSubscriptionComponentUnitBalance

// The internalCreateSubscriptionComponentUnitBalance instance, wrapping the provided int value.
func (c *internalCreateSubscriptionComponentUnitBalance) FromNumber(val int) CreateSubscriptionComponentUnitBalance {
    return CreateSubscriptionComponentUnitBalance{value: &val}
}

// The internalCreateSubscriptionComponentUnitBalance instance, wrapping the provided string value.
func (c *internalCreateSubscriptionComponentUnitBalance) FromString(val string) CreateSubscriptionComponentUnitBalance {
    return CreateSubscriptionComponentUnitBalance{value: &val}
}
