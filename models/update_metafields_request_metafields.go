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

// UpdateMetafieldsRequestMetafields represents a UpdateMetafieldsRequestMetafields struct.
// This is a container for one-of cases.
type UpdateMetafieldsRequestMetafields struct {
	value                    any
	isUpdateMetafield        bool
	isArrayOfUpdateMetafield bool
}

// String converts the UpdateMetafieldsRequestMetafields object to a string representation.
func (u UpdateMetafieldsRequestMetafields) String() string {
	if bytes, err := json.Marshal(u.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for UpdateMetafieldsRequestMetafields.
// It customizes the JSON marshaling process for UpdateMetafieldsRequestMetafields objects.
func (u *UpdateMetafieldsRequestMetafields) MarshalJSON() (
	[]byte,
	error) {
	if u.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateMetafieldsRequestMetafieldsContainer.From*` functions to initialize the UpdateMetafieldsRequestMetafields object.")
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateMetafieldsRequestMetafields object to a map representation for JSON marshaling.
func (u *UpdateMetafieldsRequestMetafields) toMap() any {
	switch obj := u.value.(type) {
	case *UpdateMetafield:
		return obj.toMap()
	case *[]UpdateMetafield:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateMetafieldsRequestMetafields.
// It customizes the JSON unmarshaling process for UpdateMetafieldsRequestMetafields objects.
func (u *UpdateMetafieldsRequestMetafields) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&UpdateMetafield{}, false, &u.isUpdateMetafield),
		NewTypeHolder(&[]UpdateMetafield{}, false, &u.isArrayOfUpdateMetafield),
	)

	u.value = result
	return err
}

func (u *UpdateMetafieldsRequestMetafields) AsUpdateMetafield() (
	*UpdateMetafield,
	bool) {
	if !u.isUpdateMetafield {
		return nil, false
	}
	return u.value.(*UpdateMetafield), true
}

func (u *UpdateMetafieldsRequestMetafields) AsArrayOfUpdateMetafield() (
	*[]UpdateMetafield,
	bool) {
	if !u.isArrayOfUpdateMetafield {
		return nil, false
	}
	return u.value.(*[]UpdateMetafield), true
}

// internalUpdateMetafieldsRequestMetafields represents a updateMetafieldsRequestMetafields struct.
// This is a container for one-of cases.
type internalUpdateMetafieldsRequestMetafields struct{}

var UpdateMetafieldsRequestMetafieldsContainer internalUpdateMetafieldsRequestMetafields

// The internalUpdateMetafieldsRequestMetafields instance, wrapping the provided UpdateMetafield value.
func (u *internalUpdateMetafieldsRequestMetafields) FromUpdateMetafield(val UpdateMetafield) UpdateMetafieldsRequestMetafields {
	return UpdateMetafieldsRequestMetafields{value: &val}
}

// The internalUpdateMetafieldsRequestMetafields instance, wrapping the provided []UpdateMetafield value.
func (u *internalUpdateMetafieldsRequestMetafields) FromArrayOfUpdateMetafield(val []UpdateMetafield) UpdateMetafieldsRequestMetafields {
	return UpdateMetafieldsRequestMetafields{value: &val}
}
