// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ReadProductPricePointPricePointId represents a ReadProductPricePointPricePointId struct.
// This is a container for one-of cases.
type ReadProductPricePointPricePointId struct {
	value    any
	isNumber bool
	isString bool
}

// String implements the fmt.Stringer interface for ReadProductPricePointPricePointId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ReadProductPricePointPricePointId) String() string {
	return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for ReadProductPricePointPricePointId.
// It customizes the JSON marshaling process for ReadProductPricePointPricePointId objects.
func (r ReadProductPricePointPricePointId) MarshalJSON() (
	[]byte,
	error) {
	if r.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ReadProductPricePointPricePointIdContainer.From*` functions to initialize the ReadProductPricePointPricePointId object.")
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ReadProductPricePointPricePointId object to a map representation for JSON marshaling.
func (r *ReadProductPricePointPricePointId) toMap() any {
	switch obj := r.value.(type) {
	case *int:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReadProductPricePointPricePointId.
// It customizes the JSON unmarshaling process for ReadProductPricePointPricePointId objects.
func (r *ReadProductPricePointPricePointId) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(int), false, &r.isNumber),
		NewTypeHolder(new(string), false, &r.isString),
	)

	r.value = result
	return err
}

func (r *ReadProductPricePointPricePointId) AsNumber() (
	*int,
	bool) {
	if !r.isNumber {
		return nil, false
	}
	return r.value.(*int), true
}

func (r *ReadProductPricePointPricePointId) AsString() (
	*string,
	bool) {
	if !r.isString {
		return nil, false
	}
	return r.value.(*string), true
}

// internalReadProductPricePointPricePointId represents a readProductPricePointPricePointId struct.
// This is a container for one-of cases.
type internalReadProductPricePointPricePointId struct{}

var ReadProductPricePointPricePointIdContainer internalReadProductPricePointPricePointId

// The internalReadProductPricePointPricePointId instance, wrapping the provided int value.
func (r *internalReadProductPricePointPricePointId) FromNumber(val int) ReadProductPricePointPricePointId {
	return ReadProductPricePointPricePointId{value: &val}
}

// The internalReadProductPricePointPricePointId instance, wrapping the provided string value.
func (r *internalReadProductPricePointPricePointId) FromString(val string) ReadProductPricePointPricePointId {
	return ReadProductPricePointPricePointId{value: &val}
}
