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

// SubscriptionGroupSignupComponentUnitBalance represents a SubscriptionGroupSignupComponentUnitBalance struct.
// This is a container for one-of cases.
type SubscriptionGroupSignupComponentUnitBalance struct {
	value    any
	isString bool
	isNumber bool
}

// String converts the SubscriptionGroupSignupComponentUnitBalance object to a string representation.
func (s SubscriptionGroupSignupComponentUnitBalance) String() string {
	if bytes, err := json.Marshal(s.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupComponentUnitBalance.
// It customizes the JSON marshaling process for SubscriptionGroupSignupComponentUnitBalance objects.
func (s *SubscriptionGroupSignupComponentUnitBalance) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionGroupSignupComponentUnitBalanceContainer.From*` functions to initialize the SubscriptionGroupSignupComponentUnitBalance object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupComponentUnitBalance object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSignupComponentUnitBalance) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupComponentUnitBalance.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupComponentUnitBalance objects.
func (s *SubscriptionGroupSignupComponentUnitBalance) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int), false, &s.isNumber),
	)

	s.value = result
	return err
}

func (s *SubscriptionGroupSignupComponentUnitBalance) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *SubscriptionGroupSignupComponentUnitBalance) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

// internalSubscriptionGroupSignupComponentUnitBalance represents a subscriptionGroupSignupComponentUnitBalance struct.
// This is a container for one-of cases.
type internalSubscriptionGroupSignupComponentUnitBalance struct{}

var SubscriptionGroupSignupComponentUnitBalanceContainer internalSubscriptionGroupSignupComponentUnitBalance

// The internalSubscriptionGroupSignupComponentUnitBalance instance, wrapping the provided string value.
func (s *internalSubscriptionGroupSignupComponentUnitBalance) FromString(val string) SubscriptionGroupSignupComponentUnitBalance {
	return SubscriptionGroupSignupComponentUnitBalance{value: &val}
}

// The internalSubscriptionGroupSignupComponentUnitBalance instance, wrapping the provided int value.
func (s *internalSubscriptionGroupSignupComponentUnitBalance) FromNumber(val int) SubscriptionGroupSignupComponentUnitBalance {
	return SubscriptionGroupSignupComponentUnitBalance{value: &val}
}
