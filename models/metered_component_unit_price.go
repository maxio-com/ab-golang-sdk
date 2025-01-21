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

// MeteredComponentUnitPrice represents a MeteredComponentUnitPrice struct.
// This is a container for one-of cases.
type MeteredComponentUnitPrice struct {
    value       any
    isString    bool
    isPrecision bool
}

// String implements the fmt.Stringer interface for MeteredComponentUnitPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MeteredComponentUnitPrice) String() string {
    return fmt.Sprintf("%v", m.value)
}

// MarshalJSON implements the json.Marshaler interface for MeteredComponentUnitPrice.
// It customizes the JSON marshaling process for MeteredComponentUnitPrice objects.
func (m MeteredComponentUnitPrice) MarshalJSON() (
    []byte,
    error) {
    if m.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.MeteredComponentUnitPriceContainer.From*` functions to initialize the MeteredComponentUnitPrice object.")
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MeteredComponentUnitPrice object to a map representation for JSON marshaling.
func (m *MeteredComponentUnitPrice) toMap() any {
    switch obj := m.value.(type) {
    case *string:
        return *obj
    case *float64:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for MeteredComponentUnitPrice.
// It customizes the JSON unmarshaling process for MeteredComponentUnitPrice objects.
func (m *MeteredComponentUnitPrice) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &m.isString),
        NewTypeHolder(new(float64), false, &m.isPrecision),
    )
    
    m.value = result
    return err
}

func (m *MeteredComponentUnitPrice) AsString() (
    *string,
    bool) {
    if !m.isString {
        return nil, false
    }
    return m.value.(*string), true
}

func (m *MeteredComponentUnitPrice) AsPrecision() (
    *float64,
    bool) {
    if !m.isPrecision {
        return nil, false
    }
    return m.value.(*float64), true
}

// internalMeteredComponentUnitPrice represents a meteredComponentUnitPrice struct.
// This is a container for one-of cases.
type internalMeteredComponentUnitPrice struct {}

var MeteredComponentUnitPriceContainer internalMeteredComponentUnitPrice

// The internalMeteredComponentUnitPrice instance, wrapping the provided string value.
func (m *internalMeteredComponentUnitPrice) FromString(val string) MeteredComponentUnitPrice {
    return MeteredComponentUnitPrice{value: &val}
}

// The internalMeteredComponentUnitPrice instance, wrapping the provided float64 value.
func (m *internalMeteredComponentUnitPrice) FromPrecision(val float64) MeteredComponentUnitPrice {
    return MeteredComponentUnitPrice{value: &val}
}
