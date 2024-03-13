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

// ProductExpirationIntervalUnit represents a ProductExpirationIntervalUnit struct.
// This is a container for one-of cases.
type ProductExpirationIntervalUnit struct {
	value                  any
	isExtendedIntervalUnit bool
}

// String converts the ProductExpirationIntervalUnit object to a string representation.
func (p ProductExpirationIntervalUnit) String() string {
	if bytes, err := json.Marshal(p.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for ProductExpirationIntervalUnit.
// It customizes the JSON marshaling process for ProductExpirationIntervalUnit objects.
func (p *ProductExpirationIntervalUnit) MarshalJSON() (
	[]byte,
	error) {
	if p.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ProductExpirationIntervalUnitContainer.From*` functions to initialize the ProductExpirationIntervalUnit object.")
	}
	return json.Marshal(p.toMap())
}

// toMap converts the ProductExpirationIntervalUnit object to a map representation for JSON marshaling.
func (p *ProductExpirationIntervalUnit) toMap() any {
	switch obj := p.value.(type) {
	case *ExtendedIntervalUnit:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductExpirationIntervalUnit.
// It customizes the JSON unmarshaling process for ProductExpirationIntervalUnit objects.
func (p *ProductExpirationIntervalUnit) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(ExtendedIntervalUnit), false, &p.isExtendedIntervalUnit),
	)

	p.value = result
	return err
}

func (p *ProductExpirationIntervalUnit) AsExtendedIntervalUnit() (
	*ExtendedIntervalUnit,
	bool) {
	if !p.isExtendedIntervalUnit {
		return nil, false
	}
	return p.value.(*ExtendedIntervalUnit), true
}

// internalProductExpirationIntervalUnit represents a productExpirationIntervalUnit struct.
// This is a container for one-of cases.
type internalProductExpirationIntervalUnit struct{}

var ProductExpirationIntervalUnitContainer internalProductExpirationIntervalUnit

func (p *internalProductExpirationIntervalUnit) FromExtendedIntervalUnit(val ExtendedIntervalUnit) ProductExpirationIntervalUnit {
	return ProductExpirationIntervalUnit{value: &val}
}
