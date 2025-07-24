// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// UsageQuantity represents a UsageQuantity struct.
// This is a container for one-of cases.
type UsageQuantity struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for UsageQuantity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UsageQuantity) String() string {
    return fmt.Sprintf("%v", u.value)
}

// MarshalJSON implements the json.Marshaler interface for UsageQuantity.
// It customizes the JSON marshaling process for UsageQuantity objects.
func (u UsageQuantity) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UsageQuantityContainer.From*` functions to initialize the UsageQuantity object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UsageQuantity object to a map representation for JSON marshaling.
func (u *UsageQuantity) toMap() any {
    switch obj := u.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UsageQuantity.
// It customizes the JSON unmarshaling process for UsageQuantity objects.
func (u *UsageQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &u.isNumber),
        NewTypeHolder(new(string), false, &u.isString),
    )
    
    u.value = result
    return err
}

func (u *UsageQuantity) AsNumber() (
    *int,
    bool) {
    if !u.isNumber {
        return nil, false
    }
    return u.value.(*int), true
}

func (u *UsageQuantity) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

// internalUsageQuantity represents a usageQuantity struct.
// This is a container for one-of cases.
type internalUsageQuantity struct {}

var UsageQuantityContainer internalUsageQuantity

// The internalUsageQuantity instance, wrapping the provided int value.
func (u *internalUsageQuantity) FromNumber(val int) UsageQuantity {
    return UsageQuantity{value: &val}
}

// The internalUsageQuantity instance, wrapping the provided string value.
func (u *internalUsageQuantity) FromString(val string) UsageQuantity {
    return UsageQuantity{value: &val}
}
