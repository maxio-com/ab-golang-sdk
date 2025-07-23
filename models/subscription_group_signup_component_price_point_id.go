// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SubscriptionGroupSignupComponentPricePointId represents a SubscriptionGroupSignupComponentPricePointId struct.
// This is a container for one-of cases.
type SubscriptionGroupSignupComponentPricePointId struct {
	value    any
	isString bool
	isNumber bool
}

// String implements the fmt.Stringer interface for SubscriptionGroupSignupComponentPricePointId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupSignupComponentPricePointId) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupComponentPricePointId.
// It customizes the JSON marshaling process for SubscriptionGroupSignupComponentPricePointId objects.
func (s SubscriptionGroupSignupComponentPricePointId) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionGroupSignupComponentPricePointIdContainer.From*` functions to initialize the SubscriptionGroupSignupComponentPricePointId object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupComponentPricePointId object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSignupComponentPricePointId) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupComponentPricePointId.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupComponentPricePointId objects.
func (s *SubscriptionGroupSignupComponentPricePointId) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int), false, &s.isNumber),
	)

	s.value = result
	return err
}

func (s *SubscriptionGroupSignupComponentPricePointId) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *SubscriptionGroupSignupComponentPricePointId) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

// internalSubscriptionGroupSignupComponentPricePointId represents a subscriptionGroupSignupComponentPricePointId struct.
// This is a container for one-of cases.
type internalSubscriptionGroupSignupComponentPricePointId struct{}

var SubscriptionGroupSignupComponentPricePointIdContainer internalSubscriptionGroupSignupComponentPricePointId

// The internalSubscriptionGroupSignupComponentPricePointId instance, wrapping the provided string value.
func (s *internalSubscriptionGroupSignupComponentPricePointId) FromString(val string) SubscriptionGroupSignupComponentPricePointId {
	return SubscriptionGroupSignupComponentPricePointId{value: &val}
}

// The internalSubscriptionGroupSignupComponentPricePointId instance, wrapping the provided int value.
func (s *internalSubscriptionGroupSignupComponentPricePointId) FromNumber(val int) SubscriptionGroupSignupComponentPricePointId {
	return SubscriptionGroupSignupComponentPricePointId{value: &val}
}
