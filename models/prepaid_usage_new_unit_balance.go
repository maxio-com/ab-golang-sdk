// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// PrepaidUsageNewUnitBalance represents a PrepaidUsageNewUnitBalance struct.
// This is a container for one-of cases.
type PrepaidUsageNewUnitBalance struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for PrepaidUsageNewUnitBalance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PrepaidUsageNewUnitBalance) String() string {
    return fmt.Sprintf("%v", p.value)
}

// MarshalJSON implements the json.Marshaler interface for PrepaidUsageNewUnitBalance.
// It customizes the JSON marshaling process for PrepaidUsageNewUnitBalance objects.
func (p PrepaidUsageNewUnitBalance) MarshalJSON() (
    []byte,
    error) {
    if p.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.PrepaidUsageNewUnitBalanceContainer.From*` functions to initialize the PrepaidUsageNewUnitBalance object.")
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidUsageNewUnitBalance object to a map representation for JSON marshaling.
func (p *PrepaidUsageNewUnitBalance) toMap() any {
    switch obj := p.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidUsageNewUnitBalance.
// It customizes the JSON unmarshaling process for PrepaidUsageNewUnitBalance objects.
func (p *PrepaidUsageNewUnitBalance) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &p.isNumber),
        NewTypeHolder(new(string), false, &p.isString),
    )
    
    p.value = result
    return err
}

func (p *PrepaidUsageNewUnitBalance) AsNumber() (
    *int,
    bool) {
    if !p.isNumber {
        return nil, false
    }
    return p.value.(*int), true
}

func (p *PrepaidUsageNewUnitBalance) AsString() (
    *string,
    bool) {
    if !p.isString {
        return nil, false
    }
    return p.value.(*string), true
}

// internalPrepaidUsageNewUnitBalance represents a prepaidUsageNewUnitBalance struct.
// This is a container for one-of cases.
type internalPrepaidUsageNewUnitBalance struct {}

var PrepaidUsageNewUnitBalanceContainer internalPrepaidUsageNewUnitBalance

// The internalPrepaidUsageNewUnitBalance instance, wrapping the provided int value.
func (p *internalPrepaidUsageNewUnitBalance) FromNumber(val int) PrepaidUsageNewUnitBalance {
    return PrepaidUsageNewUnitBalance{value: &val}
}

// The internalPrepaidUsageNewUnitBalance instance, wrapping the provided string value.
func (p *internalPrepaidUsageNewUnitBalance) FromString(val string) PrepaidUsageNewUnitBalance {
    return PrepaidUsageNewUnitBalance{value: &val}
}
