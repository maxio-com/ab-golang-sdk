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

// PrepaidUsageComponentUnitPrice represents a PrepaidUsageComponentUnitPrice struct.
// This is a container for one-of cases.
type PrepaidUsageComponentUnitPrice struct {
	value       any
	isString    bool
	isPrecision bool
}

// String converts the PrepaidUsageComponentUnitPrice object to a string representation.
func (p PrepaidUsageComponentUnitPrice) String() string {
	if bytes, err := json.Marshal(p.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for PrepaidUsageComponentUnitPrice.
// It customizes the JSON marshaling process for PrepaidUsageComponentUnitPrice objects.
func (p *PrepaidUsageComponentUnitPrice) MarshalJSON() (
	[]byte,
	error) {
	if p.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.PrepaidUsageComponentUnitPriceContainer.From*` functions to initialize the PrepaidUsageComponentUnitPrice object.")
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PrepaidUsageComponentUnitPrice object to a map representation for JSON marshaling.
func (p *PrepaidUsageComponentUnitPrice) toMap() any {
	switch obj := p.value.(type) {
	case *string:
		return *obj
	case *float64:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidUsageComponentUnitPrice.
// It customizes the JSON unmarshaling process for PrepaidUsageComponentUnitPrice objects.
func (p *PrepaidUsageComponentUnitPrice) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &p.isString),
		NewTypeHolder(new(float64), false, &p.isPrecision),
	)

	p.value = result
	return err
}

func (p *PrepaidUsageComponentUnitPrice) AsString() (
	*string,
	bool) {
	if !p.isString {
		return nil, false
	}
	return p.value.(*string), true
}

func (p *PrepaidUsageComponentUnitPrice) AsPrecision() (
	*float64,
	bool) {
	if !p.isPrecision {
		return nil, false
	}
	return p.value.(*float64), true
}

// internalPrepaidUsageComponentUnitPrice represents a prepaidUsageComponentUnitPrice struct.
// This is a container for one-of cases.
type internalPrepaidUsageComponentUnitPrice struct{}

var PrepaidUsageComponentUnitPriceContainer internalPrepaidUsageComponentUnitPrice

// The internalPrepaidUsageComponentUnitPrice instance, wrapping the provided string value.
func (p *internalPrepaidUsageComponentUnitPrice) FromString(val string) PrepaidUsageComponentUnitPrice {
	return PrepaidUsageComponentUnitPrice{value: &val}
}

// The internalPrepaidUsageComponentUnitPrice instance, wrapping the provided float64 value.
func (p *internalPrepaidUsageComponentUnitPrice) FromPrecision(val float64) PrepaidUsageComponentUnitPrice {
	return PrepaidUsageComponentUnitPrice{value: &val}
}
