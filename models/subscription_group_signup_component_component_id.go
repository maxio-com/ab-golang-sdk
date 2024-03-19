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

// SubscriptionGroupSignupComponentComponentId represents a SubscriptionGroupSignupComponentComponentId struct.
// This is a container for one-of cases.
type SubscriptionGroupSignupComponentComponentId struct {
	value    any
	isString bool
	isNumber bool
}

// String converts the SubscriptionGroupSignupComponentComponentId object to a string representation.
func (s SubscriptionGroupSignupComponentComponentId) String() string {
	if bytes, err := json.Marshal(s.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupComponentComponentId.
// It customizes the JSON marshaling process for SubscriptionGroupSignupComponentComponentId objects.
func (s *SubscriptionGroupSignupComponentComponentId) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionGroupSignupComponentComponentIdContainer.From*` functions to initialize the SubscriptionGroupSignupComponentComponentId object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupComponentComponentId object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSignupComponentComponentId) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupComponentComponentId.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupComponentComponentId objects.
func (s *SubscriptionGroupSignupComponentComponentId) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int), false, &s.isNumber),
	)

	s.value = result
	return err
}

func (s *SubscriptionGroupSignupComponentComponentId) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *SubscriptionGroupSignupComponentComponentId) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

// internalSubscriptionGroupSignupComponentComponentId represents a subscriptionGroupSignupComponentComponentId struct.
// This is a container for one-of cases.
type internalSubscriptionGroupSignupComponentComponentId struct{}

var SubscriptionGroupSignupComponentComponentIdContainer internalSubscriptionGroupSignupComponentComponentId

// The internalSubscriptionGroupSignupComponentComponentId instance, wrapping the provided string value.
func (s *internalSubscriptionGroupSignupComponentComponentId) FromString(val string) SubscriptionGroupSignupComponentComponentId {
	return SubscriptionGroupSignupComponentComponentId{value: &val}
}

// The internalSubscriptionGroupSignupComponentComponentId instance, wrapping the provided int value.
func (s *internalSubscriptionGroupSignupComponentComponentId) FromNumber(val int) SubscriptionGroupSignupComponentComponentId {
	return SubscriptionGroupSignupComponentComponentId{value: &val}
}
