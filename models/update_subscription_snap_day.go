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

// UpdateSubscriptionSnapDay represents a UpdateSubscriptionSnapDay struct.
// This is a container for one-of cases.
type UpdateSubscriptionSnapDay struct {
	value     any
	isSnapDay bool
	isNumber  bool
}

// String converts the UpdateSubscriptionSnapDay object to a string representation.
func (u UpdateSubscriptionSnapDay) String() string {
	if bytes, err := json.Marshal(u.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionSnapDay.
// It customizes the JSON marshaling process for UpdateSubscriptionSnapDay objects.
func (u *UpdateSubscriptionSnapDay) MarshalJSON() (
	[]byte,
	error) {
	if u.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateSubscriptionSnapDayContainer.From*` functions to initialize the UpdateSubscriptionSnapDay object.")
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionSnapDay object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionSnapDay) toMap() any {
	switch obj := u.value.(type) {
	case *SnapDay:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionSnapDay.
// It customizes the JSON unmarshaling process for UpdateSubscriptionSnapDay objects.
func (u *UpdateSubscriptionSnapDay) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(SnapDay), false, &u.isSnapDay),
		NewTypeHolder(new(int), false, &u.isNumber),
	)

	u.value = result
	return err
}

func (u *UpdateSubscriptionSnapDay) AsSnapDay() (
	*SnapDay,
	bool) {
	if !u.isSnapDay {
		return nil, false
	}
	return u.value.(*SnapDay), true
}

func (u *UpdateSubscriptionSnapDay) AsNumber() (
	*int,
	bool) {
	if !u.isNumber {
		return nil, false
	}
	return u.value.(*int), true
}

// internalUpdateSubscriptionSnapDay represents a updateSubscriptionSnapDay struct.
// This is a container for one-of cases.
type internalUpdateSubscriptionSnapDay struct{}

var UpdateSubscriptionSnapDayContainer internalUpdateSubscriptionSnapDay

func (u *internalUpdateSubscriptionSnapDay) FromSnapDay(val SnapDay) UpdateSubscriptionSnapDay {
	return UpdateSubscriptionSnapDay{value: &val}
}

func (u *internalUpdateSubscriptionSnapDay) FromNumber(val int) UpdateSubscriptionSnapDay {
	return UpdateSubscriptionSnapDay{value: &val}
}
