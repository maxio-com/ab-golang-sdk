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

// CreateOrUpdatePercentageCouponPercentage represents a CreateOrUpdatePercentageCouponPercentage struct.
// This is a container for one-of cases.
type CreateOrUpdatePercentageCouponPercentage struct {
	value       any
	isString    bool
	isPrecision bool
}

// String converts the CreateOrUpdatePercentageCouponPercentage object to a string representation.
func (c CreateOrUpdatePercentageCouponPercentage) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdatePercentageCouponPercentage.
// It customizes the JSON marshaling process for CreateOrUpdatePercentageCouponPercentage objects.
func (c *CreateOrUpdatePercentageCouponPercentage) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CreateOrUpdatePercentageCouponPercentageContainer.From*` functions to initialize the CreateOrUpdatePercentageCouponPercentage object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdatePercentageCouponPercentage object to a map representation for JSON marshaling.
func (c *CreateOrUpdatePercentageCouponPercentage) toMap() any {
	switch obj := c.value.(type) {
	case *string:
		return *obj
	case *float64:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdatePercentageCouponPercentage.
// It customizes the JSON unmarshaling process for CreateOrUpdatePercentageCouponPercentage objects.
func (c *CreateOrUpdatePercentageCouponPercentage) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &c.isString),
		NewTypeHolder(new(float64), false, &c.isPrecision),
	)

	c.value = result
	return err
}

func (c *CreateOrUpdatePercentageCouponPercentage) AsString() (
	*string,
	bool) {
	if !c.isString {
		return nil, false
	}
	return c.value.(*string), true
}

func (c *CreateOrUpdatePercentageCouponPercentage) AsPrecision() (
	*float64,
	bool) {
	if !c.isPrecision {
		return nil, false
	}
	return c.value.(*float64), true
}

// internalCreateOrUpdatePercentageCouponPercentage represents a createOrUpdatePercentageCouponPercentage struct.
// This is a container for one-of cases.
type internalCreateOrUpdatePercentageCouponPercentage struct{}

var CreateOrUpdatePercentageCouponPercentageContainer internalCreateOrUpdatePercentageCouponPercentage

func (c *internalCreateOrUpdatePercentageCouponPercentage) FromString(val string) CreateOrUpdatePercentageCouponPercentage {
	return CreateOrUpdatePercentageCouponPercentage{value: &val}
}

func (c *internalCreateOrUpdatePercentageCouponPercentage) FromPrecision(val float64) CreateOrUpdatePercentageCouponPercentage {
	return CreateOrUpdatePercentageCouponPercentage{value: &val}
}
