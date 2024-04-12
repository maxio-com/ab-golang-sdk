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

// ArchiveProductPricePointProductId represents a ArchiveProductPricePointProductId struct.
// This is a container for one-of cases.
type ArchiveProductPricePointProductId struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the ArchiveProductPricePointProductId object to a string representation.
func (a ArchiveProductPricePointProductId) String() string {
    if bytes, err := json.Marshal(a.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ArchiveProductPricePointProductId.
// It customizes the JSON marshaling process for ArchiveProductPricePointProductId objects.
func (a ArchiveProductPricePointProductId) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ArchiveProductPricePointProductIdContainer.From*` functions to initialize the ArchiveProductPricePointProductId object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ArchiveProductPricePointProductId object to a map representation for JSON marshaling.
func (a *ArchiveProductPricePointProductId) toMap() any {
    switch obj := a.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ArchiveProductPricePointProductId.
// It customizes the JSON unmarshaling process for ArchiveProductPricePointProductId objects.
func (a *ArchiveProductPricePointProductId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &a.isNumber),
        NewTypeHolder(new(string), false, &a.isString),
    )
    
    a.value = result
    return err
}

func (a *ArchiveProductPricePointProductId) AsNumber() (
    *int,
    bool) {
    if !a.isNumber {
        return nil, false
    }
    return a.value.(*int), true
}

func (a *ArchiveProductPricePointProductId) AsString() (
    *string,
    bool) {
    if !a.isString {
        return nil, false
    }
    return a.value.(*string), true
}

// internalArchiveProductPricePointProductId represents a archiveProductPricePointProductId struct.
// This is a container for one-of cases.
type internalArchiveProductPricePointProductId struct {}

var ArchiveProductPricePointProductIdContainer internalArchiveProductPricePointProductId

// The internalArchiveProductPricePointProductId instance, wrapping the provided int value.
func (a *internalArchiveProductPricePointProductId) FromNumber(val int) ArchiveProductPricePointProductId {
    return ArchiveProductPricePointProductId{value: &val}
}

// The internalArchiveProductPricePointProductId instance, wrapping the provided string value.
func (a *internalArchiveProductPricePointProductId) FromString(val string) ArchiveProductPricePointProductId {
    return ArchiveProductPricePointProductId{value: &val}
}
