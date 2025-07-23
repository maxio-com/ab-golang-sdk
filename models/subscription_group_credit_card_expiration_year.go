// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SubscriptionGroupCreditCardExpirationYear represents a SubscriptionGroupCreditCardExpirationYear struct.
// This is a container for one-of cases.
type SubscriptionGroupCreditCardExpirationYear struct {
	value    any
	isString bool
	isNumber bool
}

// String implements the fmt.Stringer interface for SubscriptionGroupCreditCardExpirationYear,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupCreditCardExpirationYear) String() string {
	return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupCreditCardExpirationYear.
// It customizes the JSON marshaling process for SubscriptionGroupCreditCardExpirationYear objects.
func (s SubscriptionGroupCreditCardExpirationYear) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionGroupCreditCardExpirationYearContainer.From*` functions to initialize the SubscriptionGroupCreditCardExpirationYear object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupCreditCardExpirationYear object to a map representation for JSON marshaling.
func (s *SubscriptionGroupCreditCardExpirationYear) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupCreditCardExpirationYear.
// It customizes the JSON unmarshaling process for SubscriptionGroupCreditCardExpirationYear objects.
func (s *SubscriptionGroupCreditCardExpirationYear) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int), false, &s.isNumber),
	)

	s.value = result
	return err
}

func (s *SubscriptionGroupCreditCardExpirationYear) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *SubscriptionGroupCreditCardExpirationYear) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

// internalSubscriptionGroupCreditCardExpirationYear represents a subscriptionGroupCreditCardExpirationYear struct.
// This is a container for one-of cases.
type internalSubscriptionGroupCreditCardExpirationYear struct{}

var SubscriptionGroupCreditCardExpirationYearContainer internalSubscriptionGroupCreditCardExpirationYear

// The internalSubscriptionGroupCreditCardExpirationYear instance, wrapping the provided string value.
func (s *internalSubscriptionGroupCreditCardExpirationYear) FromString(val string) SubscriptionGroupCreditCardExpirationYear {
	return SubscriptionGroupCreditCardExpirationYear{value: &val}
}

// The internalSubscriptionGroupCreditCardExpirationYear instance, wrapping the provided int value.
func (s *internalSubscriptionGroupCreditCardExpirationYear) FromNumber(val int) SubscriptionGroupCreditCardExpirationYear {
	return SubscriptionGroupCreditCardExpirationYear{value: &val}
}
