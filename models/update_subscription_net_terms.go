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

// UpdateSubscriptionNetTerms represents a UpdateSubscriptionNetTerms struct.
// This is a container for one-of cases.
type UpdateSubscriptionNetTerms struct {
	value    any
	isString bool
	isNumber bool
}

// String converts the UpdateSubscriptionNetTerms object to a string representation.
func (u UpdateSubscriptionNetTerms) String() string {
	if bytes, err := json.Marshal(u.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionNetTerms.
// It customizes the JSON marshaling process for UpdateSubscriptionNetTerms objects.
func (u *UpdateSubscriptionNetTerms) MarshalJSON() (
	[]byte,
	error) {
	if u.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateSubscriptionNetTermsContainer.From*` functions to initialize the UpdateSubscriptionNetTerms object.")
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionNetTerms object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionNetTerms) toMap() any {
	switch obj := u.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionNetTerms.
// It customizes the JSON unmarshaling process for UpdateSubscriptionNetTerms objects.
func (u *UpdateSubscriptionNetTerms) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &u.isString),
		NewTypeHolder(new(int), false, &u.isNumber),
	)

	u.value = result
	return err
}

func (u *UpdateSubscriptionNetTerms) AsString() (
	*string,
	bool) {
	if !u.isString {
		return nil, false
	}
	return u.value.(*string), true
}

func (u *UpdateSubscriptionNetTerms) AsNumber() (
	*int,
	bool) {
	if !u.isNumber {
		return nil, false
	}
	return u.value.(*int), true
}

// internalUpdateSubscriptionNetTerms represents a updateSubscriptionNetTerms struct.
// This is a container for one-of cases.
type internalUpdateSubscriptionNetTerms struct{}

var UpdateSubscriptionNetTermsContainer internalUpdateSubscriptionNetTerms

// The internalUpdateSubscriptionNetTerms instance, wrapping the provided string value.
func (u *internalUpdateSubscriptionNetTerms) FromString(val string) UpdateSubscriptionNetTerms {
	return UpdateSubscriptionNetTerms{value: &val}
}

// The internalUpdateSubscriptionNetTerms instance, wrapping the provided int value.
func (u *internalUpdateSubscriptionNetTerms) FromNumber(val int) UpdateSubscriptionNetTerms {
	return UpdateSubscriptionNetTerms{value: &val}
}
