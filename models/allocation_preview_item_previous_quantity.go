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

// AllocationPreviewItemPreviousQuantity represents a AllocationPreviewItemPreviousQuantity struct.
// This is a container for one-of cases.
type AllocationPreviewItemPreviousQuantity struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for AllocationPreviewItemPreviousQuantity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AllocationPreviewItemPreviousQuantity) String() string {
    return fmt.Sprintf("%v", a.value)
}

// MarshalJSON implements the json.Marshaler interface for AllocationPreviewItemPreviousQuantity.
// It customizes the JSON marshaling process for AllocationPreviewItemPreviousQuantity objects.
func (a AllocationPreviewItemPreviousQuantity) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.AllocationPreviewItemPreviousQuantityContainer.From*` functions to initialize the AllocationPreviewItemPreviousQuantity object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationPreviewItemPreviousQuantity object to a map representation for JSON marshaling.
func (a *AllocationPreviewItemPreviousQuantity) toMap() any {
    switch obj := a.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviewItemPreviousQuantity.
// It customizes the JSON unmarshaling process for AllocationPreviewItemPreviousQuantity objects.
func (a *AllocationPreviewItemPreviousQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &a.isNumber),
        NewTypeHolder(new(string), false, &a.isString),
    )
    
    a.value = result
    return err
}

func (a *AllocationPreviewItemPreviousQuantity) AsNumber() (
    *int,
    bool) {
    if !a.isNumber {
        return nil, false
    }
    return a.value.(*int), true
}

func (a *AllocationPreviewItemPreviousQuantity) AsString() (
    *string,
    bool) {
    if !a.isString {
        return nil, false
    }
    return a.value.(*string), true
}

// internalAllocationPreviewItemPreviousQuantity represents a allocationPreviewItemPreviousQuantity struct.
// This is a container for one-of cases.
type internalAllocationPreviewItemPreviousQuantity struct {}

var AllocationPreviewItemPreviousQuantityContainer internalAllocationPreviewItemPreviousQuantity

// The internalAllocationPreviewItemPreviousQuantity instance, wrapping the provided int value.
func (a *internalAllocationPreviewItemPreviousQuantity) FromNumber(val int) AllocationPreviewItemPreviousQuantity {
    return AllocationPreviewItemPreviousQuantity{value: &val}
}

// The internalAllocationPreviewItemPreviousQuantity instance, wrapping the provided string value.
func (a *internalAllocationPreviewItemPreviousQuantity) FromString(val string) AllocationPreviewItemPreviousQuantity {
    return AllocationPreviewItemPreviousQuantity{value: &val}
}
