// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionComponentUnitBalance represents a SubscriptionComponentUnitBalance struct.
// This is a container for one-of cases.
type SubscriptionComponentUnitBalance struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentUnitBalance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentUnitBalance) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentUnitBalance.
// It customizes the JSON marshaling process for SubscriptionComponentUnitBalance objects.
func (s SubscriptionComponentUnitBalance) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentUnitBalanceContainer.From*` functions to initialize the SubscriptionComponentUnitBalance object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentUnitBalance object to a map representation for JSON marshaling.
func (s *SubscriptionComponentUnitBalance) toMap() any {
    switch obj := s.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentUnitBalance.
// It customizes the JSON unmarshaling process for SubscriptionComponentUnitBalance objects.
func (s *SubscriptionComponentUnitBalance) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &s.isNumber),
        NewTypeHolder(new(string), false, &s.isString),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentUnitBalance) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

func (s *SubscriptionComponentUnitBalance) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

// internalSubscriptionComponentUnitBalance represents a subscriptionComponentUnitBalance struct.
// This is a container for one-of cases.
type internalSubscriptionComponentUnitBalance struct {}

var SubscriptionComponentUnitBalanceContainer internalSubscriptionComponentUnitBalance

// The internalSubscriptionComponentUnitBalance instance, wrapping the provided int value.
func (s *internalSubscriptionComponentUnitBalance) FromNumber(val int) SubscriptionComponentUnitBalance {
    return SubscriptionComponentUnitBalance{value: &val}
}

// The internalSubscriptionComponentUnitBalance instance, wrapping the provided string value.
func (s *internalSubscriptionComponentUnitBalance) FromString(val string) SubscriptionComponentUnitBalance {
    return SubscriptionComponentUnitBalance{value: &val}
}
