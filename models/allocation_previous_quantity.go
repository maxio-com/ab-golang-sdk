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

// AllocationPreviousQuantity represents a AllocationPreviousQuantity struct.
// This is a container for one-of cases.
type AllocationPreviousQuantity struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the AllocationPreviousQuantity object to a string representation.
func (a AllocationPreviousQuantity) String() string {
    if bytes, err := json.Marshal(a.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for AllocationPreviousQuantity.
// It customizes the JSON marshaling process for AllocationPreviousQuantity objects.
func (a AllocationPreviousQuantity) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.AllocationPreviousQuantityContainer.From*` functions to initialize the AllocationPreviousQuantity object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationPreviousQuantity object to a map representation for JSON marshaling.
func (a *AllocationPreviousQuantity) toMap() any {
    switch obj := a.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviousQuantity.
// It customizes the JSON unmarshaling process for AllocationPreviousQuantity objects.
func (a *AllocationPreviousQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &a.isNumber),
        NewTypeHolder(new(string), false, &a.isString),
    )
    
    a.value = result
    return err
}

func (a *AllocationPreviousQuantity) AsNumber() (
    *int,
    bool) {
    if !a.isNumber {
        return nil, false
    }
    return a.value.(*int), true
}

func (a *AllocationPreviousQuantity) AsString() (
    *string,
    bool) {
    if !a.isString {
        return nil, false
    }
    return a.value.(*string), true
}

// internalAllocationPreviousQuantity represents a allocationPreviousQuantity struct.
// This is a container for one-of cases.
type internalAllocationPreviousQuantity struct {}

var AllocationPreviousQuantityContainer internalAllocationPreviousQuantity

// The internalAllocationPreviousQuantity instance, wrapping the provided int value.
func (a *internalAllocationPreviousQuantity) FromNumber(val int) AllocationPreviousQuantity {
    return AllocationPreviousQuantity{value: &val}
}

// The internalAllocationPreviousQuantity instance, wrapping the provided string value.
func (a *internalAllocationPreviousQuantity) FromString(val string) AllocationPreviousQuantity {
    return AllocationPreviousQuantity{value: &val}
}
