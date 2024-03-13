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

// UpdatePriceUnitPrice represents a UpdatePriceUnitPrice struct.
// This is a container for one-of cases.
type UpdatePriceUnitPrice struct {
	value       any
	isPrecision bool
	isString    bool
}

// String converts the UpdatePriceUnitPrice object to a string representation.
func (u UpdatePriceUnitPrice) String() string {
	if bytes, err := json.Marshal(u.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for UpdatePriceUnitPrice.
// It customizes the JSON marshaling process for UpdatePriceUnitPrice objects.
func (u *UpdatePriceUnitPrice) MarshalJSON() (
	[]byte,
	error) {
	if u.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.UpdatePriceUnitPriceContainer.From*` functions to initialize the UpdatePriceUnitPrice object.")
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpdatePriceUnitPrice object to a map representation for JSON marshaling.
func (u *UpdatePriceUnitPrice) toMap() any {
	switch obj := u.value.(type) {
	case *float64:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePriceUnitPrice.
// It customizes the JSON unmarshaling process for UpdatePriceUnitPrice objects.
func (u *UpdatePriceUnitPrice) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(float64), false, &u.isPrecision),
		NewTypeHolder(new(string), false, &u.isString),
	)

	u.value = result
	return err
}

func (u *UpdatePriceUnitPrice) AsPrecision() (
	*float64,
	bool) {
	if !u.isPrecision {
		return nil, false
	}
	return u.value.(*float64), true
}

func (u *UpdatePriceUnitPrice) AsString() (
	*string,
	bool) {
	if !u.isString {
		return nil, false
	}
	return u.value.(*string), true
}

// internalUpdatePriceUnitPrice represents a updatePriceUnitPrice struct.
// This is a container for one-of cases.
type internalUpdatePriceUnitPrice struct{}

var UpdatePriceUnitPriceContainer internalUpdatePriceUnitPrice

func (u *internalUpdatePriceUnitPrice) FromPrecision(val float64) UpdatePriceUnitPrice {
	return UpdatePriceUnitPrice{value: &val}
}

func (u *internalUpdatePriceUnitPrice) FromString(val string) UpdatePriceUnitPrice {
	return UpdatePriceUnitPrice{value: &val}
}
