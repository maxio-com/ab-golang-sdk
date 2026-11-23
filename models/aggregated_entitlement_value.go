// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// AggregatedEntitlementValue represents a AggregatedEntitlementValue struct.
// This is a container for one-of cases.
type AggregatedEntitlementValue struct {
    value       any
    isBoolean   bool
    isPrecision bool
    isString    bool
}

// String implements the fmt.Stringer interface for AggregatedEntitlementValue,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AggregatedEntitlementValue) String() string {
    return fmt.Sprintf("%v", a.value)
}

// MarshalJSON implements the json.Marshaler interface for AggregatedEntitlementValue.
// It customizes the JSON marshaling process for AggregatedEntitlementValue objects.
func (a AggregatedEntitlementValue) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.AggregatedEntitlementValueContainer.From*` functions to initialize the AggregatedEntitlementValue object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AggregatedEntitlementValue object to a map representation for JSON marshaling.
func (a *AggregatedEntitlementValue) toMap() any {
    switch obj := a.value.(type) {
    case *bool:
        return *obj
    case *float64:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for AggregatedEntitlementValue.
// It customizes the JSON unmarshaling process for AggregatedEntitlementValue objects.
func (a *AggregatedEntitlementValue) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(bool), false, &a.isBoolean),
        NewTypeHolder(new(float64), false, &a.isPrecision),
        NewTypeHolder(new(string), false, &a.isString),
    )
    
    a.value = result
    return err
}

func (a *AggregatedEntitlementValue) AsBoolean() (
    *bool,
    bool) {
    if !a.isBoolean {
        return nil, false
    }
    return a.value.(*bool), true
}

func (a *AggregatedEntitlementValue) AsPrecision() (
    *float64,
    bool) {
    if !a.isPrecision {
        return nil, false
    }
    return a.value.(*float64), true
}

func (a *AggregatedEntitlementValue) AsString() (
    *string,
    bool) {
    if !a.isString {
        return nil, false
    }
    return a.value.(*string), true
}

// internalAggregatedEntitlementValue represents a aggregatedEntitlementValue struct.
// This is a container for one-of cases.
type internalAggregatedEntitlementValue struct {}

var AggregatedEntitlementValueContainer internalAggregatedEntitlementValue

// The internalAggregatedEntitlementValue instance, wrapping the provided bool value.
func (a *internalAggregatedEntitlementValue) FromBoolean(val bool) AggregatedEntitlementValue {
    return AggregatedEntitlementValue{value: &val}
}

// The internalAggregatedEntitlementValue instance, wrapping the provided float64 value.
func (a *internalAggregatedEntitlementValue) FromPrecision(val float64) AggregatedEntitlementValue {
    return AggregatedEntitlementValue{value: &val}
}

// The internalAggregatedEntitlementValue instance, wrapping the provided string value.
func (a *internalAggregatedEntitlementValue) FromString(val string) AggregatedEntitlementValue {
    return AggregatedEntitlementValue{value: &val}
}
