// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// UpdateInvoiceItemUnitPrice represents a UpdateInvoiceItemUnitPrice struct.
// This is a container for one-of cases.
type UpdateInvoiceItemUnitPrice struct {
    value       any
    isPrecision bool
    isString    bool
}

// String implements the fmt.Stringer interface for UpdateInvoiceItemUnitPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateInvoiceItemUnitPrice) String() string {
    return fmt.Sprintf("%v", u.value)
}

// MarshalJSON implements the json.Marshaler interface for UpdateInvoiceItemUnitPrice.
// It customizes the JSON marshaling process for UpdateInvoiceItemUnitPrice objects.
func (u UpdateInvoiceItemUnitPrice) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateInvoiceItemUnitPriceContainer.From*` functions to initialize the UpdateInvoiceItemUnitPrice object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateInvoiceItemUnitPrice object to a map representation for JSON marshaling.
func (u *UpdateInvoiceItemUnitPrice) toMap() any {
    switch obj := u.value.(type) {
    case *float64:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateInvoiceItemUnitPrice.
// It customizes the JSON unmarshaling process for UpdateInvoiceItemUnitPrice objects.
func (u *UpdateInvoiceItemUnitPrice) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(float64), false, &u.isPrecision),
        NewTypeHolder(new(string), false, &u.isString),
    )
    
    u.value = result
    return err
}

func (u *UpdateInvoiceItemUnitPrice) AsPrecision() (
    *float64,
    bool) {
    if !u.isPrecision {
        return nil, false
    }
    return u.value.(*float64), true
}

func (u *UpdateInvoiceItemUnitPrice) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

// internalUpdateInvoiceItemUnitPrice represents a updateInvoiceItemUnitPrice struct.
// This is a container for one-of cases.
type internalUpdateInvoiceItemUnitPrice struct {}

var UpdateInvoiceItemUnitPriceContainer internalUpdateInvoiceItemUnitPrice

// The internalUpdateInvoiceItemUnitPrice instance, wrapping the provided float64 value.
func (u *internalUpdateInvoiceItemUnitPrice) FromPrecision(val float64) UpdateInvoiceItemUnitPrice {
    return UpdateInvoiceItemUnitPrice{value: &val}
}

// The internalUpdateInvoiceItemUnitPrice instance, wrapping the provided string value.
func (u *internalUpdateInvoiceItemUnitPrice) FromString(val string) UpdateInvoiceItemUnitPrice {
    return UpdateInvoiceItemUnitPrice{value: &val}
}
