// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ArchiveComponentPricePointPricePointId represents a ArchiveComponentPricePointPricePointId struct.
// This is a container for one-of cases.
type ArchiveComponentPricePointPricePointId struct {
	value    any
	isNumber bool
	isString bool
}

// String implements the fmt.Stringer interface for ArchiveComponentPricePointPricePointId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ArchiveComponentPricePointPricePointId) String() string {
	return fmt.Sprintf("%v", a.value)
}

// MarshalJSON implements the json.Marshaler interface for ArchiveComponentPricePointPricePointId.
// It customizes the JSON marshaling process for ArchiveComponentPricePointPricePointId objects.
func (a ArchiveComponentPricePointPricePointId) MarshalJSON() (
	[]byte,
	error) {
	if a.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ArchiveComponentPricePointPricePointIdContainer.From*` functions to initialize the ArchiveComponentPricePointPricePointId object.")
	}
	return json.Marshal(a.toMap())
}

// toMap converts the ArchiveComponentPricePointPricePointId object to a map representation for JSON marshaling.
func (a *ArchiveComponentPricePointPricePointId) toMap() any {
	switch obj := a.value.(type) {
	case *int:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ArchiveComponentPricePointPricePointId.
// It customizes the JSON unmarshaling process for ArchiveComponentPricePointPricePointId objects.
func (a *ArchiveComponentPricePointPricePointId) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(int), false, &a.isNumber),
		NewTypeHolder(new(string), false, &a.isString),
	)

	a.value = result
	return err
}

func (a *ArchiveComponentPricePointPricePointId) AsNumber() (
	*int,
	bool) {
	if !a.isNumber {
		return nil, false
	}
	return a.value.(*int), true
}

func (a *ArchiveComponentPricePointPricePointId) AsString() (
	*string,
	bool) {
	if !a.isString {
		return nil, false
	}
	return a.value.(*string), true
}

// internalArchiveComponentPricePointPricePointId represents a archiveComponentPricePointPricePointId struct.
// This is a container for one-of cases.
type internalArchiveComponentPricePointPricePointId struct{}

var ArchiveComponentPricePointPricePointIdContainer internalArchiveComponentPricePointPricePointId

// The internalArchiveComponentPricePointPricePointId instance, wrapping the provided int value.
func (a *internalArchiveComponentPricePointPricePointId) FromNumber(val int) ArchiveComponentPricePointPricePointId {
	return ArchiveComponentPricePointPricePointId{value: &val}
}

// The internalArchiveComponentPricePointPricePointId instance, wrapping the provided string value.
func (a *internalArchiveComponentPricePointPricePointId) FromString(val string) ArchiveComponentPricePointPricePointId {
	return ArchiveComponentPricePointPricePointId{value: &val}
}
