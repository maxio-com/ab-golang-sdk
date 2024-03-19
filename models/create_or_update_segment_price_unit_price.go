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

// CreateOrUpdateSegmentPriceUnitPrice represents a CreateOrUpdateSegmentPriceUnitPrice struct.
// This is a container for one-of cases.
type CreateOrUpdateSegmentPriceUnitPrice struct {
	value       any
	isString    bool
	isPrecision bool
}

// String converts the CreateOrUpdateSegmentPriceUnitPrice object to a string representation.
func (c CreateOrUpdateSegmentPriceUnitPrice) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateSegmentPriceUnitPrice.
// It customizes the JSON marshaling process for CreateOrUpdateSegmentPriceUnitPrice objects.
func (c *CreateOrUpdateSegmentPriceUnitPrice) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CreateOrUpdateSegmentPriceUnitPriceContainer.From*` functions to initialize the CreateOrUpdateSegmentPriceUnitPrice object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateSegmentPriceUnitPrice object to a map representation for JSON marshaling.
func (c *CreateOrUpdateSegmentPriceUnitPrice) toMap() any {
	switch obj := c.value.(type) {
	case *string:
		return *obj
	case *float64:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateSegmentPriceUnitPrice.
// It customizes the JSON unmarshaling process for CreateOrUpdateSegmentPriceUnitPrice objects.
func (c *CreateOrUpdateSegmentPriceUnitPrice) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &c.isString),
		NewTypeHolder(new(float64), false, &c.isPrecision),
	)

	c.value = result
	return err
}

func (c *CreateOrUpdateSegmentPriceUnitPrice) AsString() (
	*string,
	bool) {
	if !c.isString {
		return nil, false
	}
	return c.value.(*string), true
}

func (c *CreateOrUpdateSegmentPriceUnitPrice) AsPrecision() (
	*float64,
	bool) {
	if !c.isPrecision {
		return nil, false
	}
	return c.value.(*float64), true
}

// internalCreateOrUpdateSegmentPriceUnitPrice represents a createOrUpdateSegmentPriceUnitPrice struct.
// This is a container for one-of cases.
type internalCreateOrUpdateSegmentPriceUnitPrice struct{}

var CreateOrUpdateSegmentPriceUnitPriceContainer internalCreateOrUpdateSegmentPriceUnitPrice

// The internalCreateOrUpdateSegmentPriceUnitPrice instance, wrapping the provided string value.
func (c *internalCreateOrUpdateSegmentPriceUnitPrice) FromString(val string) CreateOrUpdateSegmentPriceUnitPrice {
	return CreateOrUpdateSegmentPriceUnitPrice{value: &val}
}

// The internalCreateOrUpdateSegmentPriceUnitPrice instance, wrapping the provided float64 value.
func (c *internalCreateOrUpdateSegmentPriceUnitPrice) FromPrecision(val float64) CreateOrUpdateSegmentPriceUnitPrice {
	return CreateOrUpdateSegmentPriceUnitPrice{value: &val}
}
