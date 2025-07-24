// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ReadComponentPricePointComponentId represents a ReadComponentPricePointComponentId struct.
// This is a container for one-of cases.
type ReadComponentPricePointComponentId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for ReadComponentPricePointComponentId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ReadComponentPricePointComponentId) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for ReadComponentPricePointComponentId.
// It customizes the JSON marshaling process for ReadComponentPricePointComponentId objects.
func (r ReadComponentPricePointComponentId) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ReadComponentPricePointComponentIdContainer.From*` functions to initialize the ReadComponentPricePointComponentId object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ReadComponentPricePointComponentId object to a map representation for JSON marshaling.
func (r *ReadComponentPricePointComponentId) toMap() any {
    switch obj := r.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReadComponentPricePointComponentId.
// It customizes the JSON unmarshaling process for ReadComponentPricePointComponentId objects.
func (r *ReadComponentPricePointComponentId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &r.isNumber),
        NewTypeHolder(new(string), false, &r.isString),
    )
    
    r.value = result
    return err
}

func (r *ReadComponentPricePointComponentId) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

func (r *ReadComponentPricePointComponentId) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

// internalReadComponentPricePointComponentId represents a readComponentPricePointComponentId struct.
// This is a container for one-of cases.
type internalReadComponentPricePointComponentId struct {}

var ReadComponentPricePointComponentIdContainer internalReadComponentPricePointComponentId

// The internalReadComponentPricePointComponentId instance, wrapping the provided int value.
func (r *internalReadComponentPricePointComponentId) FromNumber(val int) ReadComponentPricePointComponentId {
    return ReadComponentPricePointComponentId{value: &val}
}

// The internalReadComponentPricePointComponentId instance, wrapping the provided string value.
func (r *internalReadComponentPricePointComponentId) FromString(val string) ReadComponentPricePointComponentId {
    return ReadComponentPricePointComponentId{value: &val}
}
