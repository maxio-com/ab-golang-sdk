// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// UpdateInvoiceItemPricePointId represents a UpdateInvoiceItemPricePointId struct.
// This is a container for one-of cases.
type UpdateInvoiceItemPricePointId struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for UpdateInvoiceItemPricePointId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateInvoiceItemPricePointId) String() string {
    return fmt.Sprintf("%v", u.value)
}

// MarshalJSON implements the json.Marshaler interface for UpdateInvoiceItemPricePointId.
// It customizes the JSON marshaling process for UpdateInvoiceItemPricePointId objects.
func (u UpdateInvoiceItemPricePointId) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateInvoiceItemPricePointIdContainer.From*` functions to initialize the UpdateInvoiceItemPricePointId object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateInvoiceItemPricePointId object to a map representation for JSON marshaling.
func (u *UpdateInvoiceItemPricePointId) toMap() any {
    switch obj := u.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateInvoiceItemPricePointId.
// It customizes the JSON unmarshaling process for UpdateInvoiceItemPricePointId objects.
func (u *UpdateInvoiceItemPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &u.isString),
        NewTypeHolder(new(int), false, &u.isNumber),
    )
    
    u.value = result
    return err
}

func (u *UpdateInvoiceItemPricePointId) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

func (u *UpdateInvoiceItemPricePointId) AsNumber() (
    *int,
    bool) {
    if !u.isNumber {
        return nil, false
    }
    return u.value.(*int), true
}

// internalUpdateInvoiceItemPricePointId represents a updateInvoiceItemPricePointId struct.
// This is a container for one-of cases.
type internalUpdateInvoiceItemPricePointId struct {}

var UpdateInvoiceItemPricePointIdContainer internalUpdateInvoiceItemPricePointId

// The internalUpdateInvoiceItemPricePointId instance, wrapping the provided string value.
func (u *internalUpdateInvoiceItemPricePointId) FromString(val string) UpdateInvoiceItemPricePointId {
    return UpdateInvoiceItemPricePointId{value: &val}
}

// The internalUpdateInvoiceItemPricePointId instance, wrapping the provided int value.
func (u *internalUpdateInvoiceItemPricePointId) FromNumber(val int) UpdateInvoiceItemPricePointId {
    return UpdateInvoiceItemPricePointId{value: &val}
}
