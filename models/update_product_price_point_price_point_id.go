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

// UpdateProductPricePointPricePointId represents a UpdateProductPricePointPricePointId struct.
// This is a container for one-of cases.
type UpdateProductPricePointPricePointId struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the UpdateProductPricePointPricePointId object to a string representation.
func (u UpdateProductPricePointPricePointId) String() string {
    if bytes, err := json.Marshal(u.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for UpdateProductPricePointPricePointId.
// It customizes the JSON marshaling process for UpdateProductPricePointPricePointId objects.
func (u UpdateProductPricePointPricePointId) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateProductPricePointPricePointIdContainer.From*` functions to initialize the UpdateProductPricePointPricePointId object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateProductPricePointPricePointId object to a map representation for JSON marshaling.
func (u *UpdateProductPricePointPricePointId) toMap() any {
    switch obj := u.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateProductPricePointPricePointId.
// It customizes the JSON unmarshaling process for UpdateProductPricePointPricePointId objects.
func (u *UpdateProductPricePointPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &u.isNumber),
        NewTypeHolder(new(string), false, &u.isString),
    )
    
    u.value = result
    return err
}

func (u *UpdateProductPricePointPricePointId) AsNumber() (
    *int,
    bool) {
    if !u.isNumber {
        return nil, false
    }
    return u.value.(*int), true
}

func (u *UpdateProductPricePointPricePointId) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

// internalUpdateProductPricePointPricePointId represents a updateProductPricePointPricePointId struct.
// This is a container for one-of cases.
type internalUpdateProductPricePointPricePointId struct {}

var UpdateProductPricePointPricePointIdContainer internalUpdateProductPricePointPricePointId

// The internalUpdateProductPricePointPricePointId instance, wrapping the provided int value.
func (u *internalUpdateProductPricePointPricePointId) FromNumber(val int) UpdateProductPricePointPricePointId {
    return UpdateProductPricePointPricePointId{value: &val}
}

// The internalUpdateProductPricePointPricePointId instance, wrapping the provided string value.
func (u *internalUpdateProductPricePointPricePointId) FromString(val string) UpdateProductPricePointPricePointId {
    return UpdateProductPricePointPricePointId{value: &val}
}
