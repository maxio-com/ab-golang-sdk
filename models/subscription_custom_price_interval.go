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

// SubscriptionCustomPriceInterval represents a SubscriptionCustomPriceInterval struct.
// This is a container for one-of cases.
type SubscriptionCustomPriceInterval struct {
	value    any
	isString bool
	isNumber bool
}

// String converts the SubscriptionCustomPriceInterval object to a string representation.
func (s SubscriptionCustomPriceInterval) String() string {
	if bytes, err := json.Marshal(s.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionCustomPriceInterval.
// It customizes the JSON marshaling process for SubscriptionCustomPriceInterval objects.
func (s *SubscriptionCustomPriceInterval) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionCustomPriceIntervalContainer.From*` functions to initialize the SubscriptionCustomPriceInterval object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionCustomPriceInterval object to a map representation for JSON marshaling.
func (s *SubscriptionCustomPriceInterval) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionCustomPriceInterval.
// It customizes the JSON unmarshaling process for SubscriptionCustomPriceInterval objects.
func (s *SubscriptionCustomPriceInterval) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int), false, &s.isNumber),
	)

	s.value = result
	return err
}

func (s *SubscriptionCustomPriceInterval) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *SubscriptionCustomPriceInterval) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

// internalSubscriptionCustomPriceInterval represents a subscriptionCustomPriceInterval struct.
// This is a container for one-of cases.
type internalSubscriptionCustomPriceInterval struct{}

var SubscriptionCustomPriceIntervalContainer internalSubscriptionCustomPriceInterval

func (s *internalSubscriptionCustomPriceInterval) FromString(val string) SubscriptionCustomPriceInterval {
	return SubscriptionCustomPriceInterval{value: &val}
}

func (s *internalSubscriptionCustomPriceInterval) FromNumber(val int) SubscriptionCustomPriceInterval {
	return SubscriptionCustomPriceInterval{value: &val}
}
