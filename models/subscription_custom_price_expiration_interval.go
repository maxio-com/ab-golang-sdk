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

// SubscriptionCustomPriceExpirationInterval represents a SubscriptionCustomPriceExpirationInterval struct.
// This is a container for one-of cases.
type SubscriptionCustomPriceExpirationInterval struct {
	value    any
	isString bool
	isNumber bool
}

// String converts the SubscriptionCustomPriceExpirationInterval object to a string representation.
func (s SubscriptionCustomPriceExpirationInterval) String() string {
	if bytes, err := json.Marshal(s.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionCustomPriceExpirationInterval.
// It customizes the JSON marshaling process for SubscriptionCustomPriceExpirationInterval objects.
func (s *SubscriptionCustomPriceExpirationInterval) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionCustomPriceExpirationIntervalContainer.From*` functions to initialize the SubscriptionCustomPriceExpirationInterval object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionCustomPriceExpirationInterval object to a map representation for JSON marshaling.
func (s *SubscriptionCustomPriceExpirationInterval) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionCustomPriceExpirationInterval.
// It customizes the JSON unmarshaling process for SubscriptionCustomPriceExpirationInterval objects.
func (s *SubscriptionCustomPriceExpirationInterval) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int), false, &s.isNumber),
	)

	s.value = result
	return err
}

func (s *SubscriptionCustomPriceExpirationInterval) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *SubscriptionCustomPriceExpirationInterval) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

// internalSubscriptionCustomPriceExpirationInterval represents a subscriptionCustomPriceExpirationInterval struct.
// This is a container for one-of cases.
type internalSubscriptionCustomPriceExpirationInterval struct{}

var SubscriptionCustomPriceExpirationIntervalContainer internalSubscriptionCustomPriceExpirationInterval

// The internalSubscriptionCustomPriceExpirationInterval instance, wrapping the provided string value.
func (s *internalSubscriptionCustomPriceExpirationInterval) FromString(val string) SubscriptionCustomPriceExpirationInterval {
	return SubscriptionCustomPriceExpirationInterval{value: &val}
}

// The internalSubscriptionCustomPriceExpirationInterval instance, wrapping the provided int value.
func (s *internalSubscriptionCustomPriceExpirationInterval) FromNumber(val int) SubscriptionCustomPriceExpirationInterval {
	return SubscriptionCustomPriceExpirationInterval{value: &val}
}
