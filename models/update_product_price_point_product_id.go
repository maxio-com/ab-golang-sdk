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

// UpdateProductPricePointProductId represents a UpdateProductPricePointProductId struct.
// This is a container for one-of cases.
type UpdateProductPricePointProductId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for UpdateProductPricePointProductId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateProductPricePointProductId) String() string {
    return fmt.Sprintf("%v", u.value)
}

// MarshalJSON implements the json.Marshaler interface for UpdateProductPricePointProductId.
// It customizes the JSON marshaling process for UpdateProductPricePointProductId objects.
func (u UpdateProductPricePointProductId) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateProductPricePointProductIdContainer.From*` functions to initialize the UpdateProductPricePointProductId object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateProductPricePointProductId object to a map representation for JSON marshaling.
func (u *UpdateProductPricePointProductId) toMap() any {
    switch obj := u.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateProductPricePointProductId.
// It customizes the JSON unmarshaling process for UpdateProductPricePointProductId objects.
func (u *UpdateProductPricePointProductId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &u.isNumber),
        NewTypeHolder(new(string), false, &u.isString),
    )
    
    u.value = result
    return err
}

func (u *UpdateProductPricePointProductId) AsNumber() (
    *int,
    bool) {
    if !u.isNumber {
        return nil, false
    }
    return u.value.(*int), true
}

func (u *UpdateProductPricePointProductId) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

// internalUpdateProductPricePointProductId represents a updateProductPricePointProductId struct.
// This is a container for one-of cases.
type internalUpdateProductPricePointProductId struct {}

var UpdateProductPricePointProductIdContainer internalUpdateProductPricePointProductId

// The internalUpdateProductPricePointProductId instance, wrapping the provided int value.
func (u *internalUpdateProductPricePointProductId) FromNumber(val int) UpdateProductPricePointProductId {
    return UpdateProductPricePointProductId{value: &val}
}

// The internalUpdateProductPricePointProductId instance, wrapping the provided string value.
func (u *internalUpdateProductPricePointProductId) FromString(val string) UpdateProductPricePointProductId {
    return UpdateProductPricePointProductId{value: &val}
}
