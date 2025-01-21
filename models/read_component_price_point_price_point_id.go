/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ReadComponentPricePointPricePointId represents a ReadComponentPricePointPricePointId struct.
// This is a container for one-of cases.
type ReadComponentPricePointPricePointId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for ReadComponentPricePointPricePointId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ReadComponentPricePointPricePointId) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for ReadComponentPricePointPricePointId.
// It customizes the JSON marshaling process for ReadComponentPricePointPricePointId objects.
func (r ReadComponentPricePointPricePointId) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ReadComponentPricePointPricePointIdContainer.From*` functions to initialize the ReadComponentPricePointPricePointId object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ReadComponentPricePointPricePointId object to a map representation for JSON marshaling.
func (r *ReadComponentPricePointPricePointId) toMap() any {
    switch obj := r.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReadComponentPricePointPricePointId.
// It customizes the JSON unmarshaling process for ReadComponentPricePointPricePointId objects.
func (r *ReadComponentPricePointPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &r.isNumber),
        NewTypeHolder(new(string), false, &r.isString),
    )
    
    r.value = result
    return err
}

func (r *ReadComponentPricePointPricePointId) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

func (r *ReadComponentPricePointPricePointId) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

// internalReadComponentPricePointPricePointId represents a readComponentPricePointPricePointId struct.
// This is a container for one-of cases.
type internalReadComponentPricePointPricePointId struct {}

var ReadComponentPricePointPricePointIdContainer internalReadComponentPricePointPricePointId

// The internalReadComponentPricePointPricePointId instance, wrapping the provided int value.
func (r *internalReadComponentPricePointPricePointId) FromNumber(val int) ReadComponentPricePointPricePointId {
    return ReadComponentPricePointPricePointId{value: &val}
}

// The internalReadComponentPricePointPricePointId instance, wrapping the provided string value.
func (r *internalReadComponentPricePointPricePointId) FromString(val string) ReadComponentPricePointPricePointId {
    return ReadComponentPricePointPricePointId{value: &val}
}
