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

// UpdatePriceStartingQuantity represents a UpdatePriceStartingQuantity struct.
// This is a container for one-of cases.
type UpdatePriceStartingQuantity struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the UpdatePriceStartingQuantity object to a string representation.
func (u UpdatePriceStartingQuantity) String() string {
    if bytes, err := json.Marshal(u.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for UpdatePriceStartingQuantity.
// It customizes the JSON marshaling process for UpdatePriceStartingQuantity objects.
func (u UpdatePriceStartingQuantity) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdatePriceStartingQuantityContainer.From*` functions to initialize the UpdatePriceStartingQuantity object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdatePriceStartingQuantity object to a map representation for JSON marshaling.
func (u *UpdatePriceStartingQuantity) toMap() any {
    switch obj := u.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePriceStartingQuantity.
// It customizes the JSON unmarshaling process for UpdatePriceStartingQuantity objects.
func (u *UpdatePriceStartingQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &u.isNumber),
        NewTypeHolder(new(string), false, &u.isString),
    )
    
    u.value = result
    return err
}

func (u *UpdatePriceStartingQuantity) AsNumber() (
    *int,
    bool) {
    if !u.isNumber {
        return nil, false
    }
    return u.value.(*int), true
}

func (u *UpdatePriceStartingQuantity) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

// internalUpdatePriceStartingQuantity represents a updatePriceStartingQuantity struct.
// This is a container for one-of cases.
type internalUpdatePriceStartingQuantity struct {}

var UpdatePriceStartingQuantityContainer internalUpdatePriceStartingQuantity

// The internalUpdatePriceStartingQuantity instance, wrapping the provided int value.
func (u *internalUpdatePriceStartingQuantity) FromNumber(val int) UpdatePriceStartingQuantity {
    return UpdatePriceStartingQuantity{value: &val}
}

// The internalUpdatePriceStartingQuantity instance, wrapping the provided string value.
func (u *internalUpdatePriceStartingQuantity) FromString(val string) UpdatePriceStartingQuantity {
    return UpdatePriceStartingQuantity{value: &val}
}
