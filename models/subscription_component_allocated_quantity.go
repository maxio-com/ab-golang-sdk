// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SubscriptionComponentAllocatedQuantity represents a SubscriptionComponentAllocatedQuantity struct.
// This is a container for one-of cases.
type SubscriptionComponentAllocatedQuantity struct {
	value    any
	isNumber bool
	isString bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentAllocatedQuantity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentAllocatedQuantity) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentAllocatedQuantity.
// It customizes the JSON marshaling process for SubscriptionComponentAllocatedQuantity objects.
func (s SubscriptionComponentAllocatedQuantity) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentAllocatedQuantityContainer.From*` functions to initialize the SubscriptionComponentAllocatedQuantity object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentAllocatedQuantity object to a map representation for JSON marshaling.
func (s *SubscriptionComponentAllocatedQuantity) toMap() any {
	switch obj := s.value.(type) {
	case *int:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentAllocatedQuantity.
// It customizes the JSON unmarshaling process for SubscriptionComponentAllocatedQuantity objects.
func (s *SubscriptionComponentAllocatedQuantity) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(int), false, &s.isNumber),
		NewTypeHolder(new(string), false, &s.isString),
	)

	s.value = result
	return err
}

func (s *SubscriptionComponentAllocatedQuantity) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

func (s *SubscriptionComponentAllocatedQuantity) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

// internalSubscriptionComponentAllocatedQuantity represents a subscriptionComponentAllocatedQuantity struct.
// This is a container for one-of cases.
type internalSubscriptionComponentAllocatedQuantity struct{}

var SubscriptionComponentAllocatedQuantityContainer internalSubscriptionComponentAllocatedQuantity

// The internalSubscriptionComponentAllocatedQuantity instance, wrapping the provided int value.
func (s *internalSubscriptionComponentAllocatedQuantity) FromNumber(val int) SubscriptionComponentAllocatedQuantity {
	return SubscriptionComponentAllocatedQuantity{value: &val}
}

// The internalSubscriptionComponentAllocatedQuantity instance, wrapping the provided string value.
func (s *internalSubscriptionComponentAllocatedQuantity) FromString(val string) SubscriptionComponentAllocatedQuantity {
	return SubscriptionComponentAllocatedQuantity{value: &val}
}
