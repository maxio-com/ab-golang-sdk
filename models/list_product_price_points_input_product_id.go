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

// ListProductPricePointsInputProductId represents a ListProductPricePointsInputProductId struct.
// This is a container for one-of cases.
type ListProductPricePointsInputProductId struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the ListProductPricePointsInputProductId object to a string representation.
func (l ListProductPricePointsInputProductId) String() string {
    if bytes, err := json.Marshal(l.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ListProductPricePointsInputProductId.
// It customizes the JSON marshaling process for ListProductPricePointsInputProductId objects.
func (l ListProductPricePointsInputProductId) MarshalJSON() (
    []byte,
    error) {
    if l.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ListProductPricePointsInputProductIdContainer.From*` functions to initialize the ListProductPricePointsInputProductId object.")
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListProductPricePointsInputProductId object to a map representation for JSON marshaling.
func (l *ListProductPricePointsInputProductId) toMap() any {
    switch obj := l.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListProductPricePointsInputProductId.
// It customizes the JSON unmarshaling process for ListProductPricePointsInputProductId objects.
func (l *ListProductPricePointsInputProductId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &l.isNumber),
        NewTypeHolder(new(string), false, &l.isString),
    )
    
    l.value = result
    return err
}

func (l *ListProductPricePointsInputProductId) AsNumber() (
    *int,
    bool) {
    if !l.isNumber {
        return nil, false
    }
    return l.value.(*int), true
}

func (l *ListProductPricePointsInputProductId) AsString() (
    *string,
    bool) {
    if !l.isString {
        return nil, false
    }
    return l.value.(*string), true
}

// internalListProductPricePointsInputProductId represents a listProductPricePointsInputProductId struct.
// This is a container for one-of cases.
type internalListProductPricePointsInputProductId struct {}

var ListProductPricePointsInputProductIdContainer internalListProductPricePointsInputProductId

// The internalListProductPricePointsInputProductId instance, wrapping the provided int value.
func (l *internalListProductPricePointsInputProductId) FromNumber(val int) ListProductPricePointsInputProductId {
    return ListProductPricePointsInputProductId{value: &val}
}

// The internalListProductPricePointsInputProductId instance, wrapping the provided string value.
func (l *internalListProductPricePointsInputProductId) FromString(val string) ListProductPricePointsInputProductId {
    return ListProductPricePointsInputProductId{value: &val}
}
