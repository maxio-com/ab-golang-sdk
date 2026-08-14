// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// UpdateInvoiceItemQuantity represents a UpdateInvoiceItemQuantity struct.
// This is a container for one-of cases.
type UpdateInvoiceItemQuantity struct {
    value       any
    isPrecision bool
    isString    bool
}

// String implements the fmt.Stringer interface for UpdateInvoiceItemQuantity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateInvoiceItemQuantity) String() string {
    return fmt.Sprintf("%v", u.value)
}

// MarshalJSON implements the json.Marshaler interface for UpdateInvoiceItemQuantity.
// It customizes the JSON marshaling process for UpdateInvoiceItemQuantity objects.
func (u UpdateInvoiceItemQuantity) MarshalJSON() (
    []byte,
    error) {
    if u.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.UpdateInvoiceItemQuantityContainer.From*` functions to initialize the UpdateInvoiceItemQuantity object.")
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateInvoiceItemQuantity object to a map representation for JSON marshaling.
func (u *UpdateInvoiceItemQuantity) toMap() any {
    switch obj := u.value.(type) {
    case *float64:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateInvoiceItemQuantity.
// It customizes the JSON unmarshaling process for UpdateInvoiceItemQuantity objects.
func (u *UpdateInvoiceItemQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(float64), false, &u.isPrecision),
        NewTypeHolder(new(string), false, &u.isString),
    )
    
    u.value = result
    return err
}

func (u *UpdateInvoiceItemQuantity) AsPrecision() (
    *float64,
    bool) {
    if !u.isPrecision {
        return nil, false
    }
    return u.value.(*float64), true
}

func (u *UpdateInvoiceItemQuantity) AsString() (
    *string,
    bool) {
    if !u.isString {
        return nil, false
    }
    return u.value.(*string), true
}

// internalUpdateInvoiceItemQuantity represents a updateInvoiceItemQuantity struct.
// This is a container for one-of cases.
type internalUpdateInvoiceItemQuantity struct {}

var UpdateInvoiceItemQuantityContainer internalUpdateInvoiceItemQuantity

// The internalUpdateInvoiceItemQuantity instance, wrapping the provided float64 value.
func (u *internalUpdateInvoiceItemQuantity) FromPrecision(val float64) UpdateInvoiceItemQuantity {
    return UpdateInvoiceItemQuantity{value: &val}
}

// The internalUpdateInvoiceItemQuantity instance, wrapping the provided string value.
func (u *internalUpdateInvoiceItemQuantity) FromString(val string) UpdateInvoiceItemQuantity {
    return UpdateInvoiceItemQuantity{value: &val}
}
