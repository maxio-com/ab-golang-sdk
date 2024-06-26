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

// UpdateComponentPricePointPricePointId represents a UpdateComponentPricePointPricePointId struct.
// This is a container for one-of cases.
type UpdateComponentPricePointPricePointId struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the UpdateComponentPricePointPricePointId object to a string representation.
func (u UpdateComponentPricePointPricePointId) String() string {
    if bytes, err := json.Marshal(u.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponentPricePointPricePointId.
// It customizes the JSON marshaling process for UpdateComponentPricePointPricePointId objects.
func (u UpdateComponentPricePointPricePointId) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateComponentPricePointPricePointIdContainer.From*` functions to initialize the UpdateComponentPricePointPricePointId object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponentPricePointPricePointId object to a map representation for JSON marshaling.
func (u *UpdateComponentPricePointPricePointId) toMap() any {
    switch obj := u.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponentPricePointPricePointId.
// It customizes the JSON unmarshaling process for UpdateComponentPricePointPricePointId objects.
func (u *UpdateComponentPricePointPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &u.isNumber),
        NewTypeHolder(new(string), false, &u.isString),
    )
    
    u.value = result
    return err
}

func (u *UpdateComponentPricePointPricePointId) AsNumber() (
    *int,
    bool) {
    if !u.isNumber {
        return nil, false
    }
    return u.value.(*int), true
}

func (u *UpdateComponentPricePointPricePointId) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

// internalUpdateComponentPricePointPricePointId represents a updateComponentPricePointPricePointId struct.
// This is a container for one-of cases.
type internalUpdateComponentPricePointPricePointId struct {}

var UpdateComponentPricePointPricePointIdContainer internalUpdateComponentPricePointPricePointId

// The internalUpdateComponentPricePointPricePointId instance, wrapping the provided int value.
func (u *internalUpdateComponentPricePointPricePointId) FromNumber(val int) UpdateComponentPricePointPricePointId {
    return UpdateComponentPricePointPricePointId{value: &val}
}

// The internalUpdateComponentPricePointPricePointId instance, wrapping the provided string value.
func (u *internalUpdateComponentPricePointPricePointId) FromString(val string) UpdateComponentPricePointPricePointId {
    return UpdateComponentPricePointPricePointId{value: &val}
}
