// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ReadProductPricePointProductId represents a ReadProductPricePointProductId struct.
// This is a container for one-of cases.
type ReadProductPricePointProductId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for ReadProductPricePointProductId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ReadProductPricePointProductId) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for ReadProductPricePointProductId.
// It customizes the JSON marshaling process for ReadProductPricePointProductId objects.
func (r ReadProductPricePointProductId) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ReadProductPricePointProductIdContainer.From*` functions to initialize the ReadProductPricePointProductId object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ReadProductPricePointProductId object to a map representation for JSON marshaling.
func (r *ReadProductPricePointProductId) toMap() any {
    switch obj := r.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReadProductPricePointProductId.
// It customizes the JSON unmarshaling process for ReadProductPricePointProductId objects.
func (r *ReadProductPricePointProductId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &r.isNumber),
        NewTypeHolder(new(string), false, &r.isString),
    )
    
    r.value = result
    return err
}

func (r *ReadProductPricePointProductId) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

func (r *ReadProductPricePointProductId) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

// internalReadProductPricePointProductId represents a readProductPricePointProductId struct.
// This is a container for one-of cases.
type internalReadProductPricePointProductId struct {}

var ReadProductPricePointProductIdContainer internalReadProductPricePointProductId

// The internalReadProductPricePointProductId instance, wrapping the provided int value.
func (r *internalReadProductPricePointProductId) FromNumber(val int) ReadProductPricePointProductId {
    return ReadProductPricePointProductId{value: &val}
}

// The internalReadProductPricePointProductId instance, wrapping the provided string value.
func (r *internalReadProductPricePointProductId) FromString(val string) ReadProductPricePointProductId {
    return ReadProductPricePointProductId{value: &val}
}
