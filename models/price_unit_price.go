// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// PriceUnitPrice represents a PriceUnitPrice struct.
// This is a container for one-of cases.
type PriceUnitPrice struct {
	value       any
	isPrecision bool
	isString    bool
}

// String implements the fmt.Stringer interface for PriceUnitPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PriceUnitPrice) String() string {
	return fmt.Sprintf("%v", p.value)
}

// MarshalJSON implements the json.Marshaler interface for PriceUnitPrice.
// It customizes the JSON marshaling process for PriceUnitPrice objects.
func (p PriceUnitPrice) MarshalJSON() (
	[]byte,
	error) {
	if p.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.PriceUnitPriceContainer.From*` functions to initialize the PriceUnitPrice object.")
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PriceUnitPrice object to a map representation for JSON marshaling.
func (p *PriceUnitPrice) toMap() any {
	switch obj := p.value.(type) {
	case *float64:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PriceUnitPrice.
// It customizes the JSON unmarshaling process for PriceUnitPrice objects.
func (p *PriceUnitPrice) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(float64), false, &p.isPrecision),
		NewTypeHolder(new(string), false, &p.isString),
	)

	p.value = result
	return err
}

func (p *PriceUnitPrice) AsPrecision() (
	*float64,
	bool) {
	if !p.isPrecision {
		return nil, false
	}
	return p.value.(*float64), true
}

func (p *PriceUnitPrice) AsString() (
	*string,
	bool) {
	if !p.isString {
		return nil, false
	}
	return p.value.(*string), true
}

// internalPriceUnitPrice represents a priceUnitPrice struct.
// This is a container for one-of cases.
type internalPriceUnitPrice struct{}

var PriceUnitPriceContainer internalPriceUnitPrice

// The internalPriceUnitPrice instance, wrapping the provided float64 value.
func (p *internalPriceUnitPrice) FromPrecision(val float64) PriceUnitPrice {
	return PriceUnitPrice{value: &val}
}

// The internalPriceUnitPrice instance, wrapping the provided string value.
func (p *internalPriceUnitPrice) FromString(val string) PriceUnitPrice {
	return PriceUnitPrice{value: &val}
}
