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

// ArchiveProductPricePointPricePointId represents a ArchiveProductPricePointPricePointId struct.
// This is a container for one-of cases.
type ArchiveProductPricePointPricePointId struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the ArchiveProductPricePointPricePointId object to a string representation.
func (a ArchiveProductPricePointPricePointId) String() string {
    if bytes, err := json.Marshal(a.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ArchiveProductPricePointPricePointId.
// It customizes the JSON marshaling process for ArchiveProductPricePointPricePointId objects.
func (a ArchiveProductPricePointPricePointId) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ArchiveProductPricePointPricePointIdContainer.From*` functions to initialize the ArchiveProductPricePointPricePointId object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ArchiveProductPricePointPricePointId object to a map representation for JSON marshaling.
func (a *ArchiveProductPricePointPricePointId) toMap() any {
    switch obj := a.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ArchiveProductPricePointPricePointId.
// It customizes the JSON unmarshaling process for ArchiveProductPricePointPricePointId objects.
func (a *ArchiveProductPricePointPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &a.isNumber),
        NewTypeHolder(new(string), false, &a.isString),
    )
    
    a.value = result
    return err
}

func (a *ArchiveProductPricePointPricePointId) AsNumber() (
    *int,
    bool) {
    if !a.isNumber {
        return nil, false
    }
    return a.value.(*int), true
}

func (a *ArchiveProductPricePointPricePointId) AsString() (
    *string,
    bool) {
    if !a.isString {
        return nil, false
    }
    return a.value.(*string), true
}

// internalArchiveProductPricePointPricePointId represents a archiveProductPricePointPricePointId struct.
// This is a container for one-of cases.
type internalArchiveProductPricePointPricePointId struct {}

var ArchiveProductPricePointPricePointIdContainer internalArchiveProductPricePointPricePointId

// The internalArchiveProductPricePointPricePointId instance, wrapping the provided int value.
func (a *internalArchiveProductPricePointPricePointId) FromNumber(val int) ArchiveProductPricePointPricePointId {
    return ArchiveProductPricePointPricePointId{value: &val}
}

// The internalArchiveProductPricePointPricePointId instance, wrapping the provided string value.
func (a *internalArchiveProductPricePointPricePointId) FromString(val string) ArchiveProductPricePointPricePointId {
    return ArchiveProductPricePointPricePointId{value: &val}
}
