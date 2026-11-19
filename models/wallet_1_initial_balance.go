// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// Wallet1InitialBalance represents a Wallet1InitialBalance struct.
// This is a container for one-of cases.
type Wallet1InitialBalance struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for Wallet1InitialBalance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w Wallet1InitialBalance) String() string {
    return fmt.Sprintf("%v", w.value)
}

// MarshalJSON implements the json.Marshaler interface for Wallet1InitialBalance.
// It customizes the JSON marshaling process for Wallet1InitialBalance objects.
func (w Wallet1InitialBalance) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.Wallet1InitialBalanceContainer.From*` functions to initialize the Wallet1InitialBalance object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the Wallet1InitialBalance object to a map representation for JSON marshaling.
func (w *Wallet1InitialBalance) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Wallet1InitialBalance.
// It customizes the JSON unmarshaling process for Wallet1InitialBalance objects.
func (w *Wallet1InitialBalance) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *Wallet1InitialBalance) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *Wallet1InitialBalance) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWallet1InitialBalance represents a wallet1InitialBalance struct.
// This is a container for one-of cases.
type internalWallet1InitialBalance struct {}

var Wallet1InitialBalanceContainer internalWallet1InitialBalance

// The internalWallet1InitialBalance instance, wrapping the provided string value.
func (w *internalWallet1InitialBalance) FromString(val string) Wallet1InitialBalance {
    return Wallet1InitialBalance{value: &val}
}

// The internalWallet1InitialBalance instance, wrapping the provided int value.
func (w *internalWallet1InitialBalance) FromNumber(val int) Wallet1InitialBalance {
    return Wallet1InitialBalance{value: &val}
}
