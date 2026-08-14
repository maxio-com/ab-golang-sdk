// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ListSubscriptionsInputProduct represents a ListSubscriptionsInputProduct struct.
// This is a container for one-of cases.
type ListSubscriptionsInputProduct struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for ListSubscriptionsInputProduct,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListSubscriptionsInputProduct) String() string {
    return fmt.Sprintf("%v", l.value)
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionsInputProduct.
// It customizes the JSON marshaling process for ListSubscriptionsInputProduct objects.
func (l ListSubscriptionsInputProduct) MarshalJSON() (
    []byte,
    error) {
    if l.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ListSubscriptionsInputProductContainer.From*` functions to initialize the ListSubscriptionsInputProduct object.")
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionsInputProduct object to a map representation for JSON marshaling.
func (l *ListSubscriptionsInputProduct) toMap() any {
    switch obj := l.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionsInputProduct.
// It customizes the JSON unmarshaling process for ListSubscriptionsInputProduct objects.
func (l *ListSubscriptionsInputProduct) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &l.isNumber),
        NewTypeHolder(new(string), false, &l.isString),
    )
    
    l.value = result
    return err
}

func (l *ListSubscriptionsInputProduct) AsNumber() (
    *int,
    bool) {
    if !l.isNumber {
        return nil, false
    }
    return l.value.(*int), true
}

func (l *ListSubscriptionsInputProduct) AsString() (
    *string,
    bool) {
    if !l.isString {
        return nil, false
    }
    return l.value.(*string), true
}

// internalListSubscriptionsInputProduct represents a listSubscriptionsInputProduct struct.
// This is a container for one-of cases.
type internalListSubscriptionsInputProduct struct {}

var ListSubscriptionsInputProductContainer internalListSubscriptionsInputProduct

// The internalListSubscriptionsInputProduct instance, wrapping the provided int value.
func (l *internalListSubscriptionsInputProduct) FromNumber(val int) ListSubscriptionsInputProduct {
    return ListSubscriptionsInputProduct{value: &val}
}

// The internalListSubscriptionsInputProduct instance, wrapping the provided string value.
func (l *internalListSubscriptionsInputProduct) FromString(val string) ListSubscriptionsInputProduct {
    return ListSubscriptionsInputProduct{value: &val}
}
