// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ListUsagesInputSubscriptionIdOrReference represents a ListUsagesInputSubscriptionIdOrReference struct.
// This is a container for one-of cases.
type ListUsagesInputSubscriptionIdOrReference struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for ListUsagesInputSubscriptionIdOrReference,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListUsagesInputSubscriptionIdOrReference) String() string {
    return fmt.Sprintf("%v", l.value)
}

// MarshalJSON implements the json.Marshaler interface for ListUsagesInputSubscriptionIdOrReference.
// It customizes the JSON marshaling process for ListUsagesInputSubscriptionIdOrReference objects.
func (l ListUsagesInputSubscriptionIdOrReference) MarshalJSON() (
    []byte,
    error) {
    if l.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ListUsagesInputSubscriptionIdOrReferenceContainer.From*` functions to initialize the ListUsagesInputSubscriptionIdOrReference object.")
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListUsagesInputSubscriptionIdOrReference object to a map representation for JSON marshaling.
func (l *ListUsagesInputSubscriptionIdOrReference) toMap() any {
    switch obj := l.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListUsagesInputSubscriptionIdOrReference.
// It customizes the JSON unmarshaling process for ListUsagesInputSubscriptionIdOrReference objects.
func (l *ListUsagesInputSubscriptionIdOrReference) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &l.isNumber),
        NewTypeHolder(new(string), false, &l.isString),
    )
    
    l.value = result
    return err
}

func (l *ListUsagesInputSubscriptionIdOrReference) AsNumber() (
    *int,
    bool) {
    if !l.isNumber {
        return nil, false
    }
    return l.value.(*int), true
}

func (l *ListUsagesInputSubscriptionIdOrReference) AsString() (
    *string,
    bool) {
    if !l.isString {
        return nil, false
    }
    return l.value.(*string), true
}

// internalListUsagesInputSubscriptionIdOrReference represents a listUsagesInputSubscriptionIdOrReference struct.
// This is a container for one-of cases.
type internalListUsagesInputSubscriptionIdOrReference struct {}

var ListUsagesInputSubscriptionIdOrReferenceContainer internalListUsagesInputSubscriptionIdOrReference

// The internalListUsagesInputSubscriptionIdOrReference instance, wrapping the provided int value.
func (l *internalListUsagesInputSubscriptionIdOrReference) FromNumber(val int) ListUsagesInputSubscriptionIdOrReference {
    return ListUsagesInputSubscriptionIdOrReference{value: &val}
}

// The internalListUsagesInputSubscriptionIdOrReference instance, wrapping the provided string value.
func (l *internalListUsagesInputSubscriptionIdOrReference) FromString(val string) ListUsagesInputSubscriptionIdOrReference {
    return ListUsagesInputSubscriptionIdOrReference{value: &val}
}
