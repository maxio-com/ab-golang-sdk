// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// UpdateInvoiceItemComponentId represents a UpdateInvoiceItemComponentId struct.
// This is a container for one-of cases.
type UpdateInvoiceItemComponentId struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for UpdateInvoiceItemComponentId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateInvoiceItemComponentId) String() string {
    return fmt.Sprintf("%v", u.value)
}

// MarshalJSON implements the json.Marshaler interface for UpdateInvoiceItemComponentId.
// It customizes the JSON marshaling process for UpdateInvoiceItemComponentId objects.
func (u UpdateInvoiceItemComponentId) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateInvoiceItemComponentIdContainer.From*` functions to initialize the UpdateInvoiceItemComponentId object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateInvoiceItemComponentId object to a map representation for JSON marshaling.
func (u *UpdateInvoiceItemComponentId) toMap() any {
    switch obj := u.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateInvoiceItemComponentId.
// It customizes the JSON unmarshaling process for UpdateInvoiceItemComponentId objects.
func (u *UpdateInvoiceItemComponentId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &u.isString),
        NewTypeHolder(new(int), false, &u.isNumber),
    )
    
    u.value = result
    return err
}

func (u *UpdateInvoiceItemComponentId) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

func (u *UpdateInvoiceItemComponentId) AsNumber() (
    *int,
    bool) {
    if !u.isNumber {
        return nil, false
    }
    return u.value.(*int), true
}

// internalUpdateInvoiceItemComponentId represents a updateInvoiceItemComponentId struct.
// This is a container for one-of cases.
type internalUpdateInvoiceItemComponentId struct {}

var UpdateInvoiceItemComponentIdContainer internalUpdateInvoiceItemComponentId

// The internalUpdateInvoiceItemComponentId instance, wrapping the provided string value.
func (u *internalUpdateInvoiceItemComponentId) FromString(val string) UpdateInvoiceItemComponentId {
    return UpdateInvoiceItemComponentId{value: &val}
}

// The internalUpdateInvoiceItemComponentId instance, wrapping the provided int value.
func (u *internalUpdateInvoiceItemComponentId) FromNumber(val int) UpdateInvoiceItemComponentId {
    return UpdateInvoiceItemComponentId{value: &val}
}
