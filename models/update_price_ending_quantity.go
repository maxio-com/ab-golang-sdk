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

// UpdatePriceEndingQuantity represents a UpdatePriceEndingQuantity struct.
// This is a container for one-of cases.
type UpdatePriceEndingQuantity struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for UpdatePriceEndingQuantity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdatePriceEndingQuantity) String() string {
    return fmt.Sprintf("%v", u.value)
}

// MarshalJSON implements the json.Marshaler interface for UpdatePriceEndingQuantity.
// It customizes the JSON marshaling process for UpdatePriceEndingQuantity objects.
func (u UpdatePriceEndingQuantity) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdatePriceEndingQuantityContainer.From*` functions to initialize the UpdatePriceEndingQuantity object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdatePriceEndingQuantity object to a map representation for JSON marshaling.
func (u *UpdatePriceEndingQuantity) toMap() any {
    switch obj := u.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePriceEndingQuantity.
// It customizes the JSON unmarshaling process for UpdatePriceEndingQuantity objects.
func (u *UpdatePriceEndingQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &u.isNumber),
        NewTypeHolder(new(string), false, &u.isString),
    )
    
    u.value = result
    return err
}

func (u *UpdatePriceEndingQuantity) AsNumber() (
    *int,
    bool) {
    if !u.isNumber {
        return nil, false
    }
    return u.value.(*int), true
}

func (u *UpdatePriceEndingQuantity) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

// internalUpdatePriceEndingQuantity represents a updatePriceEndingQuantity struct.
// This is a container for one-of cases.
type internalUpdatePriceEndingQuantity struct {}

var UpdatePriceEndingQuantityContainer internalUpdatePriceEndingQuantity

// The internalUpdatePriceEndingQuantity instance, wrapping the provided int value.
func (u *internalUpdatePriceEndingQuantity) FromNumber(val int) UpdatePriceEndingQuantity {
    return UpdatePriceEndingQuantity{value: &val}
}

// The internalUpdatePriceEndingQuantity instance, wrapping the provided string value.
func (u *internalUpdatePriceEndingQuantity) FromString(val string) UpdatePriceEndingQuantity {
    return UpdatePriceEndingQuantity{value: &val}
}
