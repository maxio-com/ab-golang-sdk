// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// AdjustmentAmount represents a AdjustmentAmount struct.
// This is a container for one-of cases.
type AdjustmentAmount struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for AdjustmentAmount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AdjustmentAmount) String() string {
    return fmt.Sprintf("%v", a.value)
}

// MarshalJSON implements the json.Marshaler interface for AdjustmentAmount.
// It customizes the JSON marshaling process for AdjustmentAmount objects.
func (a AdjustmentAmount) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.AdjustmentAmountContainer.From*` functions to initialize the AdjustmentAmount object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AdjustmentAmount object to a map representation for JSON marshaling.
func (a *AdjustmentAmount) toMap() any {
    switch obj := a.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for AdjustmentAmount.
// It customizes the JSON unmarshaling process for AdjustmentAmount objects.
func (a *AdjustmentAmount) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &a.isString),
        NewTypeHolder(new(int), false, &a.isNumber),
    )
    
    a.value = result
    return err
}

func (a *AdjustmentAmount) AsString() (
    *string,
    bool) {
    if !a.isString {
        return nil, false
    }
    return a.value.(*string), true
}

func (a *AdjustmentAmount) AsNumber() (
    *int,
    bool) {
    if !a.isNumber {
        return nil, false
    }
    return a.value.(*int), true
}

// internalAdjustmentAmount represents a adjustmentAmount struct.
// This is a container for one-of cases.
type internalAdjustmentAmount struct {}

var AdjustmentAmountContainer internalAdjustmentAmount

// The internalAdjustmentAmount instance, wrapping the provided string value.
func (a *internalAdjustmentAmount) FromString(val string) AdjustmentAmount {
    return AdjustmentAmount{value: &val}
}

// The internalAdjustmentAmount instance, wrapping the provided int value.
func (a *internalAdjustmentAmount) FromNumber(val int) AdjustmentAmount {
    return AdjustmentAmount{value: &val}
}
