// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// Wallet1ConversionRate represents a Wallet1ConversionRate struct.
// This is a container for one-of cases.
type Wallet1ConversionRate struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for Wallet1ConversionRate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w Wallet1ConversionRate) String() string {
    return fmt.Sprintf("%v", w.value)
}

// MarshalJSON implements the json.Marshaler interface for Wallet1ConversionRate.
// It customizes the JSON marshaling process for Wallet1ConversionRate objects.
func (w Wallet1ConversionRate) MarshalJSON() (
    []byte,
    error) {
    if w.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.Wallet1ConversionRateContainer.From*` functions to initialize the Wallet1ConversionRate object.")
    }
    return json.Marshal(w.toMap())
}

// toMap converts the Wallet1ConversionRate object to a map representation for JSON marshaling.
func (w *Wallet1ConversionRate) toMap() any {
    switch obj := w.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Wallet1ConversionRate.
// It customizes the JSON unmarshaling process for Wallet1ConversionRate objects.
func (w *Wallet1ConversionRate) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &w.isString),
        NewTypeHolder(new(int), false, &w.isNumber),
    )
    
    w.value = result
    return err
}

func (w *Wallet1ConversionRate) AsString() (
    *string,
    bool) {
    if !w.isString {
        return nil, false
    }
    return w.value.(*string), true
}

func (w *Wallet1ConversionRate) AsNumber() (
    *int,
    bool) {
    if !w.isNumber {
        return nil, false
    }
    return w.value.(*int), true
}

// internalWallet1ConversionRate represents a wallet1ConversionRate struct.
// This is a container for one-of cases.
type internalWallet1ConversionRate struct {}

var Wallet1ConversionRateContainer internalWallet1ConversionRate

// The internalWallet1ConversionRate instance, wrapping the provided string value.
func (w *internalWallet1ConversionRate) FromString(val string) Wallet1ConversionRate {
    return Wallet1ConversionRate{value: &val}
}

// The internalWallet1ConversionRate instance, wrapping the provided int value.
func (w *internalWallet1ConversionRate) FromNumber(val int) Wallet1ConversionRate {
    return Wallet1ConversionRate{value: &val}
}
