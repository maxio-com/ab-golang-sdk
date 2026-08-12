// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// MeteredUsageNewUnitBalance represents a MeteredUsageNewUnitBalance struct.
// This is a container for one-of cases.
type MeteredUsageNewUnitBalance struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for MeteredUsageNewUnitBalance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MeteredUsageNewUnitBalance) String() string {
    return fmt.Sprintf("%v", m.value)
}

// MarshalJSON implements the json.Marshaler interface for MeteredUsageNewUnitBalance.
// It customizes the JSON marshaling process for MeteredUsageNewUnitBalance objects.
func (m MeteredUsageNewUnitBalance) MarshalJSON() (
    []byte,
    error) {
    if m.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.MeteredUsageNewUnitBalanceContainer.From*` functions to initialize the MeteredUsageNewUnitBalance object.")
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MeteredUsageNewUnitBalance object to a map representation for JSON marshaling.
func (m *MeteredUsageNewUnitBalance) toMap() any {
    switch obj := m.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for MeteredUsageNewUnitBalance.
// It customizes the JSON unmarshaling process for MeteredUsageNewUnitBalance objects.
func (m *MeteredUsageNewUnitBalance) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &m.isNumber),
        NewTypeHolder(new(string), false, &m.isString),
    )
    
    m.value = result
    return err
}

func (m *MeteredUsageNewUnitBalance) AsNumber() (
    *int,
    bool) {
    if !m.isNumber {
        return nil, false
    }
    return m.value.(*int), true
}

func (m *MeteredUsageNewUnitBalance) AsString() (
    *string,
    bool) {
    if !m.isString {
        return nil, false
    }
    return m.value.(*string), true
}

// internalMeteredUsageNewUnitBalance represents a meteredUsageNewUnitBalance struct.
// This is a container for one-of cases.
type internalMeteredUsageNewUnitBalance struct {}

var MeteredUsageNewUnitBalanceContainer internalMeteredUsageNewUnitBalance

// The internalMeteredUsageNewUnitBalance instance, wrapping the provided int value.
func (m *internalMeteredUsageNewUnitBalance) FromNumber(val int) MeteredUsageNewUnitBalance {
    return MeteredUsageNewUnitBalance{value: &val}
}

// The internalMeteredUsageNewUnitBalance instance, wrapping the provided string value.
func (m *internalMeteredUsageNewUnitBalance) FromString(val string) MeteredUsageNewUnitBalance {
    return MeteredUsageNewUnitBalance{value: &val}
}
