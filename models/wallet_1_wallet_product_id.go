// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// Wallet1WalletProductId represents a Wallet1WalletProductId struct.
// This is a container for one-of cases.
type Wallet1WalletProductId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for Wallet1WalletProductId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w Wallet1WalletProductId) String() string {
    return fmt.Sprintf("%v", w.value)
}

// MarshalJSON implements the json.Marshaler interface for Wallet1WalletProductId.
// It customizes the JSON marshaling process for Wallet1WalletProductId objects.
func (w Wallet1WalletProductId) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.Wallet1WalletProductIdContainer.From*` functions to initialize the Wallet1WalletProductId object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the Wallet1WalletProductId object to a map representation for JSON marshaling.
func (w *Wallet1WalletProductId) toMap() any {
    switch obj := w.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Wallet1WalletProductId.
// It customizes the JSON unmarshaling process for Wallet1WalletProductId objects.
func (w *Wallet1WalletProductId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &w.isNumber),
        NewTypeHolder(new(string), false, &w.isString),
    )
    
    w.value = result
    return err
}

func (w *Wallet1WalletProductId) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

func (w *Wallet1WalletProductId) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

// internalWallet1WalletProductId represents a wallet1WalletProductId struct.
// This is a container for one-of cases.
type internalWallet1WalletProductId struct {}

var Wallet1WalletProductIdContainer internalWallet1WalletProductId

// The internalWallet1WalletProductId instance, wrapping the provided int value.
func (w *internalWallet1WalletProductId) FromNumber(val int) Wallet1WalletProductId {
    return Wallet1WalletProductId{value: &val}
}

// The internalWallet1WalletProductId instance, wrapping the provided string value.
func (w *internalWallet1WalletProductId) FromString(val string) Wallet1WalletProductId {
    return Wallet1WalletProductId{value: &val}
}
