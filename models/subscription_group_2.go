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

// SubscriptionGroup2 represents a SubscriptionGroup2 struct.
// This is a container for one-of cases.
type SubscriptionGroup2 struct {
	value                     any
	isNestedSubscriptionGroup bool
}

// String converts the SubscriptionGroup2 object to a string representation.
func (s SubscriptionGroup2) String() string {
	if bytes, err := json.Marshal(s.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroup2.
// It customizes the JSON marshaling process for SubscriptionGroup2 objects.
func (s *SubscriptionGroup2) MarshalJSON() (
	[]byte,
	error) {
	if s.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionGroup2Container.From*` functions to initialize the SubscriptionGroup2 object.")
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroup2 object to a map representation for JSON marshaling.
func (s *SubscriptionGroup2) toMap() any {
	switch obj := s.value.(type) {
	case *NestedSubscriptionGroup:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroup2.
// It customizes the JSON unmarshaling process for SubscriptionGroup2 objects.
func (s *SubscriptionGroup2) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&NestedSubscriptionGroup{}, false, &s.isNestedSubscriptionGroup),
	)

	s.value = result
	return err
}

func (s *SubscriptionGroup2) AsNestedSubscriptionGroup() (
	*NestedSubscriptionGroup,
	bool) {
	if !s.isNestedSubscriptionGroup {
		return nil, false
	}
	return s.value.(*NestedSubscriptionGroup), true
}

// internalSubscriptionGroup2 represents a subscriptionGroup2 struct.
// This is a container for one-of cases.
type internalSubscriptionGroup2 struct{}

var SubscriptionGroup2Container internalSubscriptionGroup2

func (s *internalSubscriptionGroup2) FromNestedSubscriptionGroup(val NestedSubscriptionGroup) SubscriptionGroup2 {
	return SubscriptionGroup2{value: &val}
}
