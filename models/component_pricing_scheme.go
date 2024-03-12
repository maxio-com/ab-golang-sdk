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

// ComponentPricingScheme represents a ComponentPricingScheme struct.
// This is a container for one-of cases.
type ComponentPricingScheme struct {
	value           any
	isPricingScheme bool
}

// String converts the ComponentPricingScheme object to a string representation.
func (c ComponentPricingScheme) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricingScheme.
// It customizes the JSON marshaling process for ComponentPricingScheme objects.
func (c *ComponentPricingScheme) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ComponentPricingSchemeContainer.From*` functions to initialize the ComponentPricingScheme object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricingScheme object to a map representation for JSON marshaling.
func (c *ComponentPricingScheme) toMap() any {
	switch obj := c.value.(type) {
	case *PricingScheme:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricingScheme.
// It customizes the JSON unmarshaling process for ComponentPricingScheme objects.
func (c *ComponentPricingScheme) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(PricingScheme), false, &c.isPricingScheme),
	)

	c.value = result
	return err
}

func (c *ComponentPricingScheme) AsPricingScheme() (
	*PricingScheme,
	bool) {
	if !c.isPricingScheme {
		return nil, false
	}
	return c.value.(*PricingScheme), true
}

// internalComponentPricingScheme represents a componentPricingScheme struct.
// This is a container for one-of cases.
type internalComponentPricingScheme struct{}

var ComponentPricingSchemeContainer internalComponentPricingScheme

func (c *internalComponentPricingScheme) FromPricingScheme(val PricingScheme) ComponentPricingScheme {
	return ComponentPricingScheme{value: &val}
}
