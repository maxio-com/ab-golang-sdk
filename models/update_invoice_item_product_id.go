// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// UpdateInvoiceItemProductId represents a UpdateInvoiceItemProductId struct.
// This is a container for one-of cases.
type UpdateInvoiceItemProductId struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for UpdateInvoiceItemProductId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateInvoiceItemProductId) String() string {
    return fmt.Sprintf("%v", u.value)
}

// MarshalJSON implements the json.Marshaler interface for UpdateInvoiceItemProductId.
// It customizes the JSON marshaling process for UpdateInvoiceItemProductId objects.
func (u UpdateInvoiceItemProductId) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateInvoiceItemProductIdContainer.From*` functions to initialize the UpdateInvoiceItemProductId object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateInvoiceItemProductId object to a map representation for JSON marshaling.
func (u *UpdateInvoiceItemProductId) toMap() any {
    switch obj := u.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateInvoiceItemProductId.
// It customizes the JSON unmarshaling process for UpdateInvoiceItemProductId objects.
func (u *UpdateInvoiceItemProductId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &u.isString),
        NewTypeHolder(new(int), false, &u.isNumber),
    )
    
    u.value = result
    return err
}

func (u *UpdateInvoiceItemProductId) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

func (u *UpdateInvoiceItemProductId) AsNumber() (
    *int,
    bool) {
    if !u.isNumber {
        return nil, false
    }
    return u.value.(*int), true
}

// internalUpdateInvoiceItemProductId represents a updateInvoiceItemProductId struct.
// This is a container for one-of cases.
type internalUpdateInvoiceItemProductId struct {}

var UpdateInvoiceItemProductIdContainer internalUpdateInvoiceItemProductId

// The internalUpdateInvoiceItemProductId instance, wrapping the provided string value.
func (u *internalUpdateInvoiceItemProductId) FromString(val string) UpdateInvoiceItemProductId {
    return UpdateInvoiceItemProductId{value: &val}
}

// The internalUpdateInvoiceItemProductId instance, wrapping the provided int value.
func (u *internalUpdateInvoiceItemProductId) FromNumber(val int) UpdateInvoiceItemProductId {
    return UpdateInvoiceItemProductId{value: &val}
}
