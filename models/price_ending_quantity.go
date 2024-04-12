/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PriceEndingQuantity represents a PriceEndingQuantity struct.
// This is a container for one-of cases.
type PriceEndingQuantity struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the PriceEndingQuantity object to a string representation.
func (p PriceEndingQuantity) String() string {
    if bytes, err := json.Marshal(p.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for PriceEndingQuantity.
// It customizes the JSON marshaling process for PriceEndingQuantity objects.
func (p PriceEndingQuantity) MarshalJSON() (
    []byte,
    error) {
    if p.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.PriceEndingQuantityContainer.From*` functions to initialize the PriceEndingQuantity object.")
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PriceEndingQuantity object to a map representation for JSON marshaling.
func (p *PriceEndingQuantity) toMap() any {
    switch obj := p.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PriceEndingQuantity.
// It customizes the JSON unmarshaling process for PriceEndingQuantity objects.
func (p *PriceEndingQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &p.isNumber),
        NewTypeHolder(new(string), false, &p.isString),
    )
    
    p.value = result
    return err
}

func (p *PriceEndingQuantity) AsNumber() (
    *int,
    bool) {
    if !p.isNumber {
        return nil, false
    }
    return p.value.(*int), true
}

func (p *PriceEndingQuantity) AsString() (
    *string,
    bool) {
    if !p.isString {
        return nil, false
    }
    return p.value.(*string), true
}

// internalPriceEndingQuantity represents a priceEndingQuantity struct.
// This is a container for one-of cases.
type internalPriceEndingQuantity struct {}

var PriceEndingQuantityContainer internalPriceEndingQuantity

// The internalPriceEndingQuantity instance, wrapping the provided int value.
func (p *internalPriceEndingQuantity) FromNumber(val int) PriceEndingQuantity {
    return PriceEndingQuantity{value: &val}
}

// The internalPriceEndingQuantity instance, wrapping the provided string value.
func (p *internalPriceEndingQuantity) FromString(val string) PriceEndingQuantity {
    return PriceEndingQuantity{value: &val}
}
