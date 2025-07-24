// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CreateSubscriptionComponentAllocatedQuantity represents a CreateSubscriptionComponentAllocatedQuantity struct.
// This is a container for one-of cases.
type CreateSubscriptionComponentAllocatedQuantity struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for CreateSubscriptionComponentAllocatedQuantity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateSubscriptionComponentAllocatedQuantity) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionComponentAllocatedQuantity.
// It customizes the JSON marshaling process for CreateSubscriptionComponentAllocatedQuantity objects.
func (c CreateSubscriptionComponentAllocatedQuantity) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateSubscriptionComponentAllocatedQuantityContainer.From*` functions to initialize the CreateSubscriptionComponentAllocatedQuantity object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionComponentAllocatedQuantity object to a map representation for JSON marshaling.
func (c *CreateSubscriptionComponentAllocatedQuantity) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionComponentAllocatedQuantity.
// It customizes the JSON unmarshaling process for CreateSubscriptionComponentAllocatedQuantity objects.
func (c *CreateSubscriptionComponentAllocatedQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CreateSubscriptionComponentAllocatedQuantity) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CreateSubscriptionComponentAllocatedQuantity) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCreateSubscriptionComponentAllocatedQuantity represents a createSubscriptionComponentAllocatedQuantity struct.
// This is a container for one-of cases.
type internalCreateSubscriptionComponentAllocatedQuantity struct {}

var CreateSubscriptionComponentAllocatedQuantityContainer internalCreateSubscriptionComponentAllocatedQuantity

// The internalCreateSubscriptionComponentAllocatedQuantity instance, wrapping the provided int value.
func (c *internalCreateSubscriptionComponentAllocatedQuantity) FromNumber(val int) CreateSubscriptionComponentAllocatedQuantity {
    return CreateSubscriptionComponentAllocatedQuantity{value: &val}
}

// The internalCreateSubscriptionComponentAllocatedQuantity instance, wrapping the provided string value.
func (c *internalCreateSubscriptionComponentAllocatedQuantity) FromString(val string) CreateSubscriptionComponentAllocatedQuantity {
    return CreateSubscriptionComponentAllocatedQuantity{value: &val}
}
