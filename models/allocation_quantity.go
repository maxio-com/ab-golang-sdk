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

// AllocationQuantity represents a AllocationQuantity struct.
// This is a container for one-of cases.
type AllocationQuantity struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the AllocationQuantity object to a string representation.
func (a AllocationQuantity) String() string {
    if bytes, err := json.Marshal(a.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for AllocationQuantity.
// It customizes the JSON marshaling process for AllocationQuantity objects.
func (a AllocationQuantity) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.AllocationQuantityContainer.From*` functions to initialize the AllocationQuantity object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationQuantity object to a map representation for JSON marshaling.
func (a *AllocationQuantity) toMap() any {
    switch obj := a.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationQuantity.
// It customizes the JSON unmarshaling process for AllocationQuantity objects.
func (a *AllocationQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &a.isNumber),
        NewTypeHolder(new(string), false, &a.isString),
    )
    
    a.value = result
    return err
}

func (a *AllocationQuantity) AsNumber() (
    *int,
    bool) {
    if !a.isNumber {
        return nil, false
    }
    return a.value.(*int), true
}

func (a *AllocationQuantity) AsString() (
    *string,
    bool) {
    if !a.isString {
        return nil, false
    }
    return a.value.(*string), true
}

// internalAllocationQuantity represents a allocationQuantity struct.
// This is a container for one-of cases.
type internalAllocationQuantity struct {}

var AllocationQuantityContainer internalAllocationQuantity

// The internalAllocationQuantity instance, wrapping the provided int value.
func (a *internalAllocationQuantity) FromNumber(val int) AllocationQuantity {
    return AllocationQuantity{value: &val}
}

// The internalAllocationQuantity instance, wrapping the provided string value.
func (a *internalAllocationQuantity) FromString(val string) AllocationQuantity {
    return AllocationQuantity{value: &val}
}
