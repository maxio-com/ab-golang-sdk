// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// AllocationPreviewItemQuantity represents a AllocationPreviewItemQuantity struct.
// This is a container for one-of cases.
type AllocationPreviewItemQuantity struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for AllocationPreviewItemQuantity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AllocationPreviewItemQuantity) String() string {
    return fmt.Sprintf("%v", a.value)
}

// MarshalJSON implements the json.Marshaler interface for AllocationPreviewItemQuantity.
// It customizes the JSON marshaling process for AllocationPreviewItemQuantity objects.
func (a AllocationPreviewItemQuantity) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.AllocationPreviewItemQuantityContainer.From*` functions to initialize the AllocationPreviewItemQuantity object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationPreviewItemQuantity object to a map representation for JSON marshaling.
func (a *AllocationPreviewItemQuantity) toMap() any {
    switch obj := a.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviewItemQuantity.
// It customizes the JSON unmarshaling process for AllocationPreviewItemQuantity objects.
func (a *AllocationPreviewItemQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &a.isNumber),
        NewTypeHolder(new(string), false, &a.isString),
    )
    
    a.value = result
    return err
}

func (a *AllocationPreviewItemQuantity) AsNumber() (
    *int,
    bool) {
    if !a.isNumber {
        return nil, false
    }
    return a.value.(*int), true
}

func (a *AllocationPreviewItemQuantity) AsString() (
    *string,
    bool) {
    if !a.isString {
        return nil, false
    }
    return a.value.(*string), true
}

// internalAllocationPreviewItemQuantity represents a allocationPreviewItemQuantity struct.
// This is a container for one-of cases.
type internalAllocationPreviewItemQuantity struct {}

var AllocationPreviewItemQuantityContainer internalAllocationPreviewItemQuantity

// The internalAllocationPreviewItemQuantity instance, wrapping the provided int value.
func (a *internalAllocationPreviewItemQuantity) FromNumber(val int) AllocationPreviewItemQuantity {
    return AllocationPreviewItemQuantity{value: &val}
}

// The internalAllocationPreviewItemQuantity instance, wrapping the provided string value.
func (a *internalAllocationPreviewItemQuantity) FromString(val string) AllocationPreviewItemQuantity {
    return AllocationPreviewItemQuantity{value: &val}
}
