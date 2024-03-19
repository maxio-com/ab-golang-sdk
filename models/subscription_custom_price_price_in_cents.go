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

// SubscriptionCustomPricePriceInCents represents a SubscriptionCustomPricePriceInCents struct.
// This is a container for one-of cases.
type SubscriptionCustomPricePriceInCents struct {
	value    any
	isString bool
	isLong   bool
}

// String converts the SubscriptionCustomPricePriceInCents object to a string representation.
func (s SubscriptionCustomPricePriceInCents) String() string {
	if bytes, err := json.Marshal(s.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionCustomPricePriceInCents.
// It customizes the JSON marshaling process for SubscriptionCustomPricePriceInCents objects.
func (s *SubscriptionCustomPricePriceInCents) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionCustomPricePriceInCentsContainer.From*` functions to initialize the SubscriptionCustomPricePriceInCents object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionCustomPricePriceInCents object to a map representation for JSON marshaling.
func (s *SubscriptionCustomPricePriceInCents) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int64:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionCustomPricePriceInCents.
// It customizes the JSON unmarshaling process for SubscriptionCustomPricePriceInCents objects.
func (s *SubscriptionCustomPricePriceInCents) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int64), false, &s.isLong),
	)

	s.value = result
	return err
}

func (s *SubscriptionCustomPricePriceInCents) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *SubscriptionCustomPricePriceInCents) AsLong() (
	*int64,
	bool) {
	if !s.isLong {
		return nil, false
	}
	return s.value.(*int64), true
}

// internalSubscriptionCustomPricePriceInCents represents a subscriptionCustomPricePriceInCents struct.
// This is a container for one-of cases.
type internalSubscriptionCustomPricePriceInCents struct{}

var SubscriptionCustomPricePriceInCentsContainer internalSubscriptionCustomPricePriceInCents

// The internalSubscriptionCustomPricePriceInCents instance, wrapping the provided string value.
func (s *internalSubscriptionCustomPricePriceInCents) FromString(val string) SubscriptionCustomPricePriceInCents {
	return SubscriptionCustomPricePriceInCents{value: &val}
}

// The internalSubscriptionCustomPricePriceInCents instance, wrapping the provided int64 value.
func (s *internalSubscriptionCustomPricePriceInCents) FromLong(val int64) SubscriptionCustomPricePriceInCents {
	return SubscriptionCustomPricePriceInCents{value: &val}
}
