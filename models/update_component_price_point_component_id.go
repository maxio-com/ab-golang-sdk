// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// UpdateComponentPricePointComponentId represents a UpdateComponentPricePointComponentId struct.
// This is a container for one-of cases.
type UpdateComponentPricePointComponentId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for UpdateComponentPricePointComponentId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateComponentPricePointComponentId) String() string {
    return fmt.Sprintf("%v", u.value)
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponentPricePointComponentId.
// It customizes the JSON marshaling process for UpdateComponentPricePointComponentId objects.
func (u UpdateComponentPricePointComponentId) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateComponentPricePointComponentIdContainer.From*` functions to initialize the UpdateComponentPricePointComponentId object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponentPricePointComponentId object to a map representation for JSON marshaling.
func (u *UpdateComponentPricePointComponentId) toMap() any {
    switch obj := u.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponentPricePointComponentId.
// It customizes the JSON unmarshaling process for UpdateComponentPricePointComponentId objects.
func (u *UpdateComponentPricePointComponentId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &u.isNumber),
        NewTypeHolder(new(string), false, &u.isString),
    )
    
    u.value = result
    return err
}

func (u *UpdateComponentPricePointComponentId) AsNumber() (
    *int,
    bool) {
    if !u.isNumber {
        return nil, false
    }
    return u.value.(*int), true
}

func (u *UpdateComponentPricePointComponentId) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

// internalUpdateComponentPricePointComponentId represents a updateComponentPricePointComponentId struct.
// This is a container for one-of cases.
type internalUpdateComponentPricePointComponentId struct {}

var UpdateComponentPricePointComponentIdContainer internalUpdateComponentPricePointComponentId

// The internalUpdateComponentPricePointComponentId instance, wrapping the provided int value.
func (u *internalUpdateComponentPricePointComponentId) FromNumber(val int) UpdateComponentPricePointComponentId {
    return UpdateComponentPricePointComponentId{value: &val}
}

// The internalUpdateComponentPricePointComponentId instance, wrapping the provided string value.
func (u *internalUpdateComponentPricePointComponentId) FromString(val string) UpdateComponentPricePointComponentId {
    return UpdateComponentPricePointComponentId{value: &val}
}
