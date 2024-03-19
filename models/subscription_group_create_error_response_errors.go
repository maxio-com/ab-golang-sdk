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

// SubscriptionGroupCreateErrorResponseErrors represents a SubscriptionGroupCreateErrorResponseErrors struct.
// This is a container for one-of cases.
type SubscriptionGroupCreateErrorResponseErrors struct {
	value                                any
	isSubscriptionGroupMembersArrayError bool
	isSubscriptionGroupSingleError       bool
	isString                             bool
}

// String converts the SubscriptionGroupCreateErrorResponseErrors object to a string representation.
func (s SubscriptionGroupCreateErrorResponseErrors) String() string {
	if bytes, err := json.Marshal(s.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupCreateErrorResponseErrors.
// It customizes the JSON marshaling process for SubscriptionGroupCreateErrorResponseErrors objects.
func (s *SubscriptionGroupCreateErrorResponseErrors) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionGroupCreateErrorResponseErrorsContainer.From*` functions to initialize the SubscriptionGroupCreateErrorResponseErrors object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupCreateErrorResponseErrors object to a map representation for JSON marshaling.
func (s *SubscriptionGroupCreateErrorResponseErrors) toMap() any {
	switch obj := s.value.(type) {
	case *SubscriptionGroupMembersArrayError:
		return obj.toMap()
	case *SubscriptionGroupSingleError:
		return obj.toMap()
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupCreateErrorResponseErrors.
// It customizes the JSON unmarshaling process for SubscriptionGroupCreateErrorResponseErrors objects.
func (s *SubscriptionGroupCreateErrorResponseErrors) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&SubscriptionGroupMembersArrayError{}, false, &s.isSubscriptionGroupMembersArrayError),
		NewTypeHolder(&SubscriptionGroupSingleError{}, false, &s.isSubscriptionGroupSingleError),
		NewTypeHolder(new(string), false, &s.isString),
	)

	s.value = result
	return err
}

func (s *SubscriptionGroupCreateErrorResponseErrors) AsSubscriptionGroupMembersArrayError() (
	*SubscriptionGroupMembersArrayError,
	bool) {
	if !s.isSubscriptionGroupMembersArrayError {
		return nil, false
	}
	return s.value.(*SubscriptionGroupMembersArrayError), true
}

func (s *SubscriptionGroupCreateErrorResponseErrors) AsSubscriptionGroupSingleError() (
	*SubscriptionGroupSingleError,
	bool) {
	if !s.isSubscriptionGroupSingleError {
		return nil, false
	}
	return s.value.(*SubscriptionGroupSingleError), true
}

func (s *SubscriptionGroupCreateErrorResponseErrors) AsString() (
	*string,
	bool) {
	if !s.isString {
		return nil, false
	}
	return s.value.(*string), true
}

// internalSubscriptionGroupCreateErrorResponseErrors represents a subscriptionGroupCreateErrorResponseErrors struct.
// This is a container for one-of cases.
type internalSubscriptionGroupCreateErrorResponseErrors struct{}

var SubscriptionGroupCreateErrorResponseErrorsContainer internalSubscriptionGroupCreateErrorResponseErrors

// The internalSubscriptionGroupCreateErrorResponseErrors instance, wrapping the provided SubscriptionGroupMembersArrayError value.
func (s *internalSubscriptionGroupCreateErrorResponseErrors) FromSubscriptionGroupMembersArrayError(val SubscriptionGroupMembersArrayError) SubscriptionGroupCreateErrorResponseErrors {
	return SubscriptionGroupCreateErrorResponseErrors{value: &val}
}

// The internalSubscriptionGroupCreateErrorResponseErrors instance, wrapping the provided SubscriptionGroupSingleError value.
func (s *internalSubscriptionGroupCreateErrorResponseErrors) FromSubscriptionGroupSingleError(val SubscriptionGroupSingleError) SubscriptionGroupCreateErrorResponseErrors {
	return SubscriptionGroupCreateErrorResponseErrors{value: &val}
}

// The internalSubscriptionGroupCreateErrorResponseErrors instance, wrapping the provided string value.
func (s *internalSubscriptionGroupCreateErrorResponseErrors) FromString(val string) SubscriptionGroupCreateErrorResponseErrors {
	return SubscriptionGroupCreateErrorResponseErrors{value: &val}
}
