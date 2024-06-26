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

// ArchiveComponentPricePointComponentId represents a ArchiveComponentPricePointComponentId struct.
// This is a container for one-of cases.
type ArchiveComponentPricePointComponentId struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the ArchiveComponentPricePointComponentId object to a string representation.
func (a ArchiveComponentPricePointComponentId) String() string {
    if bytes, err := json.Marshal(a.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ArchiveComponentPricePointComponentId.
// It customizes the JSON marshaling process for ArchiveComponentPricePointComponentId objects.
func (a ArchiveComponentPricePointComponentId) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ArchiveComponentPricePointComponentIdContainer.From*` functions to initialize the ArchiveComponentPricePointComponentId object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ArchiveComponentPricePointComponentId object to a map representation for JSON marshaling.
func (a *ArchiveComponentPricePointComponentId) toMap() any {
    switch obj := a.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ArchiveComponentPricePointComponentId.
// It customizes the JSON unmarshaling process for ArchiveComponentPricePointComponentId objects.
func (a *ArchiveComponentPricePointComponentId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &a.isNumber),
        NewTypeHolder(new(string), false, &a.isString),
    )
    
    a.value = result
    return err
}

func (a *ArchiveComponentPricePointComponentId) AsNumber() (
    *int,
    bool) {
    if !a.isNumber {
        return nil, false
    }
    return a.value.(*int), true
}

func (a *ArchiveComponentPricePointComponentId) AsString() (
    *string,
    bool) {
    if !a.isString {
        return nil, false
    }
    return a.value.(*string), true
}

// internalArchiveComponentPricePointComponentId represents a archiveComponentPricePointComponentId struct.
// This is a container for one-of cases.
type internalArchiveComponentPricePointComponentId struct {}

var ArchiveComponentPricePointComponentIdContainer internalArchiveComponentPricePointComponentId

// The internalArchiveComponentPricePointComponentId instance, wrapping the provided int value.
func (a *internalArchiveComponentPricePointComponentId) FromNumber(val int) ArchiveComponentPricePointComponentId {
    return ArchiveComponentPricePointComponentId{value: &val}
}

// The internalArchiveComponentPricePointComponentId instance, wrapping the provided string value.
func (a *internalArchiveComponentPricePointComponentId) FromString(val string) ArchiveComponentPricePointComponentId {
    return ArchiveComponentPricePointComponentId{value: &val}
}
