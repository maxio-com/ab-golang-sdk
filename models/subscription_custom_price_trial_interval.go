// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SubscriptionCustomPriceTrialInterval represents a SubscriptionCustomPriceTrialInterval struct.
// This is a container for one-of cases.
type SubscriptionCustomPriceTrialInterval struct {
	value    any
	isString bool
	isNumber bool
}

// String implements the fmt.Stringer interface for SubscriptionCustomPriceTrialInterval,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionCustomPriceTrialInterval) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionCustomPriceTrialInterval.
// It customizes the JSON marshaling process for SubscriptionCustomPriceTrialInterval objects.
func (s SubscriptionCustomPriceTrialInterval) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionCustomPriceTrialIntervalContainer.From*` functions to initialize the SubscriptionCustomPriceTrialInterval object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionCustomPriceTrialInterval object to a map representation for JSON marshaling.
func (s *SubscriptionCustomPriceTrialInterval) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionCustomPriceTrialInterval.
// It customizes the JSON unmarshaling process for SubscriptionCustomPriceTrialInterval objects.
func (s *SubscriptionCustomPriceTrialInterval) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int), false, &s.isNumber),
	)

	s.value = result
	return err
}

func (s *SubscriptionCustomPriceTrialInterval) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *SubscriptionCustomPriceTrialInterval) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

// internalSubscriptionCustomPriceTrialInterval represents a subscriptionCustomPriceTrialInterval struct.
// This is a container for one-of cases.
type internalSubscriptionCustomPriceTrialInterval struct{}

var SubscriptionCustomPriceTrialIntervalContainer internalSubscriptionCustomPriceTrialInterval

// The internalSubscriptionCustomPriceTrialInterval instance, wrapping the provided string value.
func (s *internalSubscriptionCustomPriceTrialInterval) FromString(val string) SubscriptionCustomPriceTrialInterval {
	return SubscriptionCustomPriceTrialInterval{value: &val}
}

// The internalSubscriptionCustomPriceTrialInterval instance, wrapping the provided int value.
func (s *internalSubscriptionCustomPriceTrialInterval) FromNumber(val int) SubscriptionCustomPriceTrialInterval {
	return SubscriptionCustomPriceTrialInterval{value: &val}
}
