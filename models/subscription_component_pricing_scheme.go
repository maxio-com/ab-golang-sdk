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

// SubscriptionComponentPricingScheme represents a SubscriptionComponentPricingScheme struct.
// This is a container for one-of cases.
type SubscriptionComponentPricingScheme struct {
	value           any
	isPricingScheme bool
}

// String converts the SubscriptionComponentPricingScheme object to a string representation.
func (s SubscriptionComponentPricingScheme) String() string {
	if bytes, err := json.Marshal(s.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentPricingScheme.
// It customizes the JSON marshaling process for SubscriptionComponentPricingScheme objects.
func (s *SubscriptionComponentPricingScheme) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentPricingSchemeContainer.From*` functions to initialize the SubscriptionComponentPricingScheme object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentPricingScheme object to a map representation for JSON marshaling.
func (s *SubscriptionComponentPricingScheme) toMap() any {
	switch obj := s.value.(type) {
	case *PricingScheme:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentPricingScheme.
// It customizes the JSON unmarshaling process for SubscriptionComponentPricingScheme objects.
func (s *SubscriptionComponentPricingScheme) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(PricingScheme), false, &s.isPricingScheme),
	)

	s.value = result
	return err
}

func (s *SubscriptionComponentPricingScheme) AsPricingScheme() (
	*PricingScheme,
	bool) {
	if !s.isPricingScheme {
		return nil, false
	}
	return s.value.(*PricingScheme), true
}

// internalSubscriptionComponentPricingScheme represents a subscriptionComponentPricingScheme struct.
// This is a container for one-of cases.
type internalSubscriptionComponentPricingScheme struct{}

var SubscriptionComponentPricingSchemeContainer internalSubscriptionComponentPricingScheme

// The internalSubscriptionComponentPricingScheme instance, wrapping the provided PricingScheme value.
func (s *internalSubscriptionComponentPricingScheme) FromPricingScheme(val PricingScheme) SubscriptionComponentPricingScheme {
	return SubscriptionComponentPricingScheme{value: &val}
}
