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

// ListUsagesInputComponentId represents a ListUsagesInputComponentId struct.
// This is a container for one-of cases.
type ListUsagesInputComponentId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for ListUsagesInputComponentId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListUsagesInputComponentId) String() string {
    return fmt.Sprintf("%v", l.value)
}

// MarshalJSON implements the json.Marshaler interface for ListUsagesInputComponentId.
// It customizes the JSON marshaling process for ListUsagesInputComponentId objects.
func (l ListUsagesInputComponentId) MarshalJSON() (
    []byte,
    error) {
    if l.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ListUsagesInputComponentIdContainer.From*` functions to initialize the ListUsagesInputComponentId object.")
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListUsagesInputComponentId object to a map representation for JSON marshaling.
func (l *ListUsagesInputComponentId) toMap() any {
    switch obj := l.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListUsagesInputComponentId.
// It customizes the JSON unmarshaling process for ListUsagesInputComponentId objects.
func (l *ListUsagesInputComponentId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &l.isNumber),
        NewTypeHolder(new(string), false, &l.isString),
    )
    
    l.value = result
    return err
}

func (l *ListUsagesInputComponentId) AsNumber() (
    *int,
    bool) {
    if !l.isNumber {
        return nil, false
    }
    return l.value.(*int), true
}

func (l *ListUsagesInputComponentId) AsString() (
    *string,
    bool) {
    if !l.isString {
        return nil, false
    }
    return l.value.(*string), true
}

// internalListUsagesInputComponentId represents a listUsagesInputComponentId struct.
// This is a container for one-of cases.
type internalListUsagesInputComponentId struct {}

var ListUsagesInputComponentIdContainer internalListUsagesInputComponentId

// The internalListUsagesInputComponentId instance, wrapping the provided int value.
func (l *internalListUsagesInputComponentId) FromNumber(val int) ListUsagesInputComponentId {
    return ListUsagesInputComponentId{value: &val}
}

// The internalListUsagesInputComponentId instance, wrapping the provided string value.
func (l *internalListUsagesInputComponentId) FromString(val string) ListUsagesInputComponentId {
    return ListUsagesInputComponentId{value: &val}
}
