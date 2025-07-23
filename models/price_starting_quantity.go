// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// PriceStartingQuantity represents a PriceStartingQuantity struct.
// This is a container for one-of cases.
type PriceStartingQuantity struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for PriceStartingQuantity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PriceStartingQuantity) String() string {
    return fmt.Sprintf("%v", p.value)
}

// MarshalJSON implements the json.Marshaler interface for PriceStartingQuantity.
// It customizes the JSON marshaling process for PriceStartingQuantity objects.
func (p PriceStartingQuantity) MarshalJSON() (
    []byte,
    error) {
    if p.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.PriceStartingQuantityContainer.From*` functions to initialize the PriceStartingQuantity object.")
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PriceStartingQuantity object to a map representation for JSON marshaling.
func (p *PriceStartingQuantity) toMap() any {
    switch obj := p.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PriceStartingQuantity.
// It customizes the JSON unmarshaling process for PriceStartingQuantity objects.
func (p *PriceStartingQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &p.isNumber),
        NewTypeHolder(new(string), false, &p.isString),
    )
    
    p.value = result
    return err
}

func (p *PriceStartingQuantity) AsNumber() (
    *int,
    bool) {
    if !p.isNumber {
        return nil, false
    }
    return p.value.(*int), true
}

func (p *PriceStartingQuantity) AsString() (
    *string,
    bool) {
    if !p.isString {
        return nil, false
    }
    return p.value.(*string), true
}

// internalPriceStartingQuantity represents a priceStartingQuantity struct.
// This is a container for one-of cases.
type internalPriceStartingQuantity struct {}

var PriceStartingQuantityContainer internalPriceStartingQuantity

// The internalPriceStartingQuantity instance, wrapping the provided int value.
func (p *internalPriceStartingQuantity) FromNumber(val int) PriceStartingQuantity {
    return PriceStartingQuantity{value: &val}
}

// The internalPriceStartingQuantity instance, wrapping the provided string value.
func (p *internalPriceStartingQuantity) FromString(val string) PriceStartingQuantity {
    return PriceStartingQuantity{value: &val}
}
