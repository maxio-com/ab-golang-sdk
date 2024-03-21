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

// ProductTrialIntervalUnit represents a ProductTrialIntervalUnit struct.
// This is a container for one-of cases.
type ProductTrialIntervalUnit struct {
	value          any
	isIntervalUnit bool
}

// String converts the ProductTrialIntervalUnit object to a string representation.
func (p ProductTrialIntervalUnit) String() string {
	if bytes, err := json.Marshal(p.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for ProductTrialIntervalUnit.
// It customizes the JSON marshaling process for ProductTrialIntervalUnit objects.
func (p *ProductTrialIntervalUnit) MarshalJSON() (
	[]byte,
	error) {
	if p.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ProductTrialIntervalUnitContainer.From*` functions to initialize the ProductTrialIntervalUnit object.")
	}
	return json.Marshal(p.toMap())
}

// toMap converts the ProductTrialIntervalUnit object to a map representation for JSON marshaling.
func (p *ProductTrialIntervalUnit) toMap() any {
	switch obj := p.value.(type) {
	case *IntervalUnit:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductTrialIntervalUnit.
// It customizes the JSON unmarshaling process for ProductTrialIntervalUnit objects.
func (p *ProductTrialIntervalUnit) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(IntervalUnit), false, &p.isIntervalUnit),
	)

	p.value = result
	return err
}

func (p *ProductTrialIntervalUnit) AsIntervalUnit() (
	*IntervalUnit,
	bool) {
	if !p.isIntervalUnit {
		return nil, false
	}
	return p.value.(*IntervalUnit), true
}

// internalProductTrialIntervalUnit represents a productTrialIntervalUnit struct.
// This is a container for one-of cases.
type internalProductTrialIntervalUnit struct{}

var ProductTrialIntervalUnitContainer internalProductTrialIntervalUnit

// The internalProductTrialIntervalUnit instance, wrapping the provided IntervalUnit value.
func (p *internalProductTrialIntervalUnit) FromIntervalUnit(val IntervalUnit) ProductTrialIntervalUnit {
	return ProductTrialIntervalUnit{value: &val}
}
