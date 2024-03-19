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

// SubscriptionGroupCreditCardFullNumber represents a SubscriptionGroupCreditCardFullNumber struct.
// This is a container for one-of cases.
type SubscriptionGroupCreditCardFullNumber struct {
	value    any
	isString bool
	isNumber bool
}

// String converts the SubscriptionGroupCreditCardFullNumber object to a string representation.
func (s SubscriptionGroupCreditCardFullNumber) String() string {
	if bytes, err := json.Marshal(s.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupCreditCardFullNumber.
// It customizes the JSON marshaling process for SubscriptionGroupCreditCardFullNumber objects.
func (s *SubscriptionGroupCreditCardFullNumber) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionGroupCreditCardFullNumberContainer.From*` functions to initialize the SubscriptionGroupCreditCardFullNumber object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupCreditCardFullNumber object to a map representation for JSON marshaling.
func (s *SubscriptionGroupCreditCardFullNumber) toMap() any {
	switch obj := s.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupCreditCardFullNumber.
// It customizes the JSON unmarshaling process for SubscriptionGroupCreditCardFullNumber objects.
func (s *SubscriptionGroupCreditCardFullNumber) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &s.isString),
		NewTypeHolder(new(int), false, &s.isNumber),
	)

	s.value = result
	return err
}

func (s *SubscriptionGroupCreditCardFullNumber) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

func (s *SubscriptionGroupCreditCardFullNumber) AsNumber() (
	*int,
	bool) {
	if !s.isNumber {
		return nil, false
	}
	return s.value.(*int), true
}

// internalSubscriptionGroupCreditCardFullNumber represents a subscriptionGroupCreditCardFullNumber struct.
// This is a container for one-of cases.
type internalSubscriptionGroupCreditCardFullNumber struct{}

var SubscriptionGroupCreditCardFullNumberContainer internalSubscriptionGroupCreditCardFullNumber

// The internalSubscriptionGroupCreditCardFullNumber instance, wrapping the provided string value.
func (s *internalSubscriptionGroupCreditCardFullNumber) FromString(val string) SubscriptionGroupCreditCardFullNumber {
	return SubscriptionGroupCreditCardFullNumber{value: &val}
}

// The internalSubscriptionGroupCreditCardFullNumber instance, wrapping the provided int value.
func (s *internalSubscriptionGroupCreditCardFullNumber) FromNumber(val int) SubscriptionGroupCreditCardFullNumber {
	return SubscriptionGroupCreditCardFullNumber{value: &val}
}
