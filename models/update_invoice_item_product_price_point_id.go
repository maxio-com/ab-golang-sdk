// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// UpdateInvoiceItemProductPricePointId represents a UpdateInvoiceItemProductPricePointId struct.
// This is a container for one-of cases.
type UpdateInvoiceItemProductPricePointId struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for UpdateInvoiceItemProductPricePointId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateInvoiceItemProductPricePointId) String() string {
    return fmt.Sprintf("%v", u.value)
}

// MarshalJSON implements the json.Marshaler interface for UpdateInvoiceItemProductPricePointId.
// It customizes the JSON marshaling process for UpdateInvoiceItemProductPricePointId objects.
func (u UpdateInvoiceItemProductPricePointId) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateInvoiceItemProductPricePointIdContainer.From*` functions to initialize the UpdateInvoiceItemProductPricePointId object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateInvoiceItemProductPricePointId object to a map representation for JSON marshaling.
func (u *UpdateInvoiceItemProductPricePointId) toMap() any {
    switch obj := u.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateInvoiceItemProductPricePointId.
// It customizes the JSON unmarshaling process for UpdateInvoiceItemProductPricePointId objects.
func (u *UpdateInvoiceItemProductPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &u.isString),
        NewTypeHolder(new(int), false, &u.isNumber),
    )
    
    u.value = result
    return err
}

func (u *UpdateInvoiceItemProductPricePointId) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

func (u *UpdateInvoiceItemProductPricePointId) AsNumber() (
    *int,
    bool) {
    if !u.isNumber {
        return nil, false
    }
    return u.value.(*int), true
}

// internalUpdateInvoiceItemProductPricePointId represents a updateInvoiceItemProductPricePointId struct.
// This is a container for one-of cases.
type internalUpdateInvoiceItemProductPricePointId struct {}

var UpdateInvoiceItemProductPricePointIdContainer internalUpdateInvoiceItemProductPricePointId

// The internalUpdateInvoiceItemProductPricePointId instance, wrapping the provided string value.
func (u *internalUpdateInvoiceItemProductPricePointId) FromString(val string) UpdateInvoiceItemProductPricePointId {
    return UpdateInvoiceItemProductPricePointId{value: &val}
}

// The internalUpdateInvoiceItemProductPricePointId instance, wrapping the provided int value.
func (u *internalUpdateInvoiceItemProductPricePointId) FromNumber(val int) UpdateInvoiceItemProductPricePointId {
    return UpdateInvoiceItemProductPricePointId{value: &val}
}
