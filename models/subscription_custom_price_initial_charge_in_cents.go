// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SubscriptionCustomPriceInitialChargeInCents represents a SubscriptionCustomPriceInitialChargeInCents struct.
// This is a container for one-of cases.
type SubscriptionCustomPriceInitialChargeInCents struct {
	value    any
	isString bool
	isLong   bool
}

// String implements the fmt.Stringer interface for SubscriptionCustomPriceInitialChargeInCents,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionCustomPriceInitialChargeInCents) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionCustomPriceInitialChargeInCents.
// It customizes the JSON marshaling process for SubscriptionCustomPriceInitialChargeInCents objects.
func (s SubscriptionCustomPriceInitialChargeInCents) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionCustomPriceInitialChargeInCentsContainer.From*` functions to initialize the SubscriptionCustomPriceInitialChargeInCents object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionCustomPriceInitialChargeInCents object to a map representation for JSON marshaling.
func (s *SubscriptionCustomPriceInitialChargeInCents) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int64:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionCustomPriceInitialChargeInCents.
// It customizes the JSON unmarshaling process for SubscriptionCustomPriceInitialChargeInCents objects.
func (s *SubscriptionCustomPriceInitialChargeInCents) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int64), false, &s.isLong),
	)

	s.value = result
	return err
}

func (s *SubscriptionCustomPriceInitialChargeInCents) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *SubscriptionCustomPriceInitialChargeInCents) AsLong() (
	*int64,
	bool) {
	if !s.isLong {
		return nil, false
	}
	return s.value.(*int64), true
}

// internalSubscriptionCustomPriceInitialChargeInCents represents a subscriptionCustomPriceInitialChargeInCents struct.
// This is a container for one-of cases.
type internalSubscriptionCustomPriceInitialChargeInCents struct{}

var SubscriptionCustomPriceInitialChargeInCentsContainer internalSubscriptionCustomPriceInitialChargeInCents

// The internalSubscriptionCustomPriceInitialChargeInCents instance, wrapping the provided string value.
func (s *internalSubscriptionCustomPriceInitialChargeInCents) FromString(val string) SubscriptionCustomPriceInitialChargeInCents {
	return SubscriptionCustomPriceInitialChargeInCents{value: &val}
}

// The internalSubscriptionCustomPriceInitialChargeInCents instance, wrapping the provided int64 value.
func (s *internalSubscriptionCustomPriceInitialChargeInCents) FromLong(val int64) SubscriptionCustomPriceInitialChargeInCents {
	return SubscriptionCustomPriceInitialChargeInCents{value: &val}
}
