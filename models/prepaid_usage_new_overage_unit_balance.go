// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// PrepaidUsageNewOverageUnitBalance represents a PrepaidUsageNewOverageUnitBalance struct.
// This is a container for one-of cases.
type PrepaidUsageNewOverageUnitBalance struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for PrepaidUsageNewOverageUnitBalance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PrepaidUsageNewOverageUnitBalance) String() string {
    return fmt.Sprintf("%v", p.value)
}

// MarshalJSON implements the json.Marshaler interface for PrepaidUsageNewOverageUnitBalance.
// It customizes the JSON marshaling process for PrepaidUsageNewOverageUnitBalance objects.
func (p PrepaidUsageNewOverageUnitBalance) MarshalJSON() (
    []byte,
    error) {
    if p.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.PrepaidUsageNewOverageUnitBalanceContainer.From*` functions to initialize the PrepaidUsageNewOverageUnitBalance object.")
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidUsageNewOverageUnitBalance object to a map representation for JSON marshaling.
func (p *PrepaidUsageNewOverageUnitBalance) toMap() any {
    switch obj := p.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidUsageNewOverageUnitBalance.
// It customizes the JSON unmarshaling process for PrepaidUsageNewOverageUnitBalance objects.
func (p *PrepaidUsageNewOverageUnitBalance) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &p.isNumber),
        NewTypeHolder(new(string), false, &p.isString),
    )
    
    p.value = result
    return err
}

func (p *PrepaidUsageNewOverageUnitBalance) AsNumber() (
    *int,
    bool) {
    if !p.isNumber {
        return nil, false
    }
    return p.value.(*int), true
}

func (p *PrepaidUsageNewOverageUnitBalance) AsString() (
    *string,
    bool) {
    if !p.isString {
        return nil, false
    }
    return p.value.(*string), true
}

// internalPrepaidUsageNewOverageUnitBalance represents a prepaidUsageNewOverageUnitBalance struct.
// This is a container for one-of cases.
type internalPrepaidUsageNewOverageUnitBalance struct {}

var PrepaidUsageNewOverageUnitBalanceContainer internalPrepaidUsageNewOverageUnitBalance

// The internalPrepaidUsageNewOverageUnitBalance instance, wrapping the provided int value.
func (p *internalPrepaidUsageNewOverageUnitBalance) FromNumber(val int) PrepaidUsageNewOverageUnitBalance {
    return PrepaidUsageNewOverageUnitBalance{value: &val}
}

// The internalPrepaidUsageNewOverageUnitBalance instance, wrapping the provided string value.
func (p *internalPrepaidUsageNewOverageUnitBalance) FromString(val string) PrepaidUsageNewOverageUnitBalance {
    return PrepaidUsageNewOverageUnitBalance{value: &val}
}
