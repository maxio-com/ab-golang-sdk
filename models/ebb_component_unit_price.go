/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// EBBComponentUnitPrice represents a EBBComponentUnitPrice struct.
// This is a container for one-of cases.
type EBBComponentUnitPrice struct {
    value       any
    isString    bool
    isPrecision bool
}

// String implements the fmt.Stringer interface for EBBComponentUnitPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EBBComponentUnitPrice) String() string {
    return fmt.Sprintf("%v", e.value)
}

// MarshalJSON implements the json.Marshaler interface for EBBComponentUnitPrice.
// It customizes the JSON marshaling process for EBBComponentUnitPrice objects.
func (e EBBComponentUnitPrice) MarshalJSON() (
    []byte,
    error) {
    if e.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.EBBComponentUnitPriceContainer.From*` functions to initialize the EBBComponentUnitPrice object.")
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EBBComponentUnitPrice object to a map representation for JSON marshaling.
func (e *EBBComponentUnitPrice) toMap() any {
    switch obj := e.value.(type) {
    case *string:
        return *obj
    case *float64:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for EBBComponentUnitPrice.
// It customizes the JSON unmarshaling process for EBBComponentUnitPrice objects.
func (e *EBBComponentUnitPrice) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &e.isString),
        NewTypeHolder(new(float64), false, &e.isPrecision),
    )
    
    e.value = result
    return err
}

func (e *EBBComponentUnitPrice) AsString() (
    *string,
    bool) {
    if !e.isString {
        return nil, false
    }
    return e.value.(*string), true
}

func (e *EBBComponentUnitPrice) AsPrecision() (
    *float64,
    bool) {
    if !e.isPrecision {
        return nil, false
    }
    return e.value.(*float64), true
}

// internalEBBComponentUnitPrice represents a eBBComponentUnitPrice struct.
// This is a container for one-of cases.
type internalEBBComponentUnitPrice struct {}

var EBBComponentUnitPriceContainer internalEBBComponentUnitPrice

// The internalEBBComponentUnitPrice instance, wrapping the provided string value.
func (e *internalEBBComponentUnitPrice) FromString(val string) EBBComponentUnitPrice {
    return EBBComponentUnitPrice{value: &val}
}

// The internalEBBComponentUnitPrice instance, wrapping the provided float64 value.
func (e *internalEBBComponentUnitPrice) FromPrecision(val float64) EBBComponentUnitPrice {
    return EBBComponentUnitPrice{value: &val}
}
