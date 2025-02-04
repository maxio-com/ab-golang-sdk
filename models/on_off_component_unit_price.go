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

// OnOffComponentUnitPrice represents a OnOffComponentUnitPrice struct.
// This is a container for one-of cases.
type OnOffComponentUnitPrice struct {
    value       any
    isString    bool
    isPrecision bool
}

// String implements the fmt.Stringer interface for OnOffComponentUnitPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OnOffComponentUnitPrice) String() string {
    return fmt.Sprintf("%v", o.value)
}

// MarshalJSON implements the json.Marshaler interface for OnOffComponentUnitPrice.
// It customizes the JSON marshaling process for OnOffComponentUnitPrice objects.
func (o OnOffComponentUnitPrice) MarshalJSON() (
    []byte,
    error) {
    if o.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.OnOffComponentUnitPriceContainer.From*` functions to initialize the OnOffComponentUnitPrice object.")
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OnOffComponentUnitPrice object to a map representation for JSON marshaling.
func (o *OnOffComponentUnitPrice) toMap() any {
    switch obj := o.value.(type) {
    case *string:
        return *obj
    case *float64:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for OnOffComponentUnitPrice.
// It customizes the JSON unmarshaling process for OnOffComponentUnitPrice objects.
func (o *OnOffComponentUnitPrice) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &o.isString),
        NewTypeHolder(new(float64), false, &o.isPrecision),
    )
    
    o.value = result
    return err
}

func (o *OnOffComponentUnitPrice) AsString() (
    *string,
    bool) {
    if !o.isString {
        return nil, false
    }
    return o.value.(*string), true
}

func (o *OnOffComponentUnitPrice) AsPrecision() (
    *float64,
    bool) {
    if !o.isPrecision {
        return nil, false
    }
    return o.value.(*float64), true
}

// internalOnOffComponentUnitPrice represents a onOffComponentUnitPrice struct.
// This is a container for one-of cases.
type internalOnOffComponentUnitPrice struct {}

var OnOffComponentUnitPriceContainer internalOnOffComponentUnitPrice

// The internalOnOffComponentUnitPrice instance, wrapping the provided string value.
func (o *internalOnOffComponentUnitPrice) FromString(val string) OnOffComponentUnitPrice {
    return OnOffComponentUnitPrice{value: &val}
}

// The internalOnOffComponentUnitPrice instance, wrapping the provided float64 value.
func (o *internalOnOffComponentUnitPrice) FromPrecision(val float64) OnOffComponentUnitPrice {
    return OnOffComponentUnitPrice{value: &val}
}
