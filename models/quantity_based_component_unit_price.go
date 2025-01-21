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

// QuantityBasedComponentUnitPrice represents a QuantityBasedComponentUnitPrice struct.
// This is a container for one-of cases.
type QuantityBasedComponentUnitPrice struct {
    value       any
    isString    bool
    isPrecision bool
}

// String implements the fmt.Stringer interface for QuantityBasedComponentUnitPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (q QuantityBasedComponentUnitPrice) String() string {
    return fmt.Sprintf("%v", q.value)
}

// MarshalJSON implements the json.Marshaler interface for QuantityBasedComponentUnitPrice.
// It customizes the JSON marshaling process for QuantityBasedComponentUnitPrice objects.
func (q QuantityBasedComponentUnitPrice) MarshalJSON() (
    []byte,
    error) {
    if q.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.QuantityBasedComponentUnitPriceContainer.From*` functions to initialize the QuantityBasedComponentUnitPrice object.")
    }
    return json.Marshal(q.toMap())
}

// toMap converts the QuantityBasedComponentUnitPrice object to a map representation for JSON marshaling.
func (q *QuantityBasedComponentUnitPrice) toMap() any {
    switch obj := q.value.(type) {
    case *string:
        return *obj
    case *float64:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for QuantityBasedComponentUnitPrice.
// It customizes the JSON unmarshaling process for QuantityBasedComponentUnitPrice objects.
func (q *QuantityBasedComponentUnitPrice) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &q.isString),
        NewTypeHolder(new(float64), false, &q.isPrecision),
    )
    
    q.value = result
    return err
}

func (q *QuantityBasedComponentUnitPrice) AsString() (
    *string,
    bool) {
    if !q.isString {
        return nil, false
    }
    return q.value.(*string), true
}

func (q *QuantityBasedComponentUnitPrice) AsPrecision() (
    *float64,
    bool) {
    if !q.isPrecision {
        return nil, false
    }
    return q.value.(*float64), true
}

// internalQuantityBasedComponentUnitPrice represents a quantityBasedComponentUnitPrice struct.
// This is a container for one-of cases.
type internalQuantityBasedComponentUnitPrice struct {}

var QuantityBasedComponentUnitPriceContainer internalQuantityBasedComponentUnitPrice

// The internalQuantityBasedComponentUnitPrice instance, wrapping the provided string value.
func (q *internalQuantityBasedComponentUnitPrice) FromString(val string) QuantityBasedComponentUnitPrice {
    return QuantityBasedComponentUnitPrice{value: &val}
}

// The internalQuantityBasedComponentUnitPrice instance, wrapping the provided float64 value.
func (q *internalQuantityBasedComponentUnitPrice) FromPrecision(val float64) QuantityBasedComponentUnitPrice {
    return QuantityBasedComponentUnitPrice{value: &val}
}
