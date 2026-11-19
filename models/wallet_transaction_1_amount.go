// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// WalletTransaction1Amount represents a WalletTransaction1Amount struct.
// This is a container for one-of cases.
type WalletTransaction1Amount struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for WalletTransaction1Amount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletTransaction1Amount) String() string {
    return fmt.Sprintf("%v", w.value)
}

// MarshalJSON implements the json.Marshaler interface for WalletTransaction1Amount.
// It customizes the JSON marshaling process for WalletTransaction1Amount objects.
func (w WalletTransaction1Amount) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.WalletTransaction1AmountContainer.From*` functions to initialize the WalletTransaction1Amount object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletTransaction1Amount object to a map representation for JSON marshaling.
func (w *WalletTransaction1Amount) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletTransaction1Amount.
// It customizes the JSON unmarshaling process for WalletTransaction1Amount objects.
func (w *WalletTransaction1Amount) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *WalletTransaction1Amount) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *WalletTransaction1Amount) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWalletTransaction1Amount represents a walletTransaction1Amount struct.
// This is a container for one-of cases.
type internalWalletTransaction1Amount struct {}

var WalletTransaction1AmountContainer internalWalletTransaction1Amount

// The internalWalletTransaction1Amount instance, wrapping the provided string value.
func (w *internalWalletTransaction1Amount) FromString(val string) WalletTransaction1Amount {
    return WalletTransaction1Amount{value: &val}
}

// The internalWalletTransaction1Amount instance, wrapping the provided int value.
func (w *internalWalletTransaction1Amount) FromNumber(val int) WalletTransaction1Amount {
    return WalletTransaction1Amount{value: &val}
}
