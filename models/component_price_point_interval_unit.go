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

// ComponentPricePointIntervalUnit represents a ComponentPricePointIntervalUnit struct.
// This is a container for one-of cases.
type ComponentPricePointIntervalUnit struct {
	value          any
	isIntervalUnit bool
}

// String converts the ComponentPricePointIntervalUnit object to a string representation.
func (c ComponentPricePointIntervalUnit) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointIntervalUnit.
// It customizes the JSON marshaling process for ComponentPricePointIntervalUnit objects.
func (c *ComponentPricePointIntervalUnit) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ComponentPricePointIntervalUnitContainer.From*` functions to initialize the ComponentPricePointIntervalUnit object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointIntervalUnit object to a map representation for JSON marshaling.
func (c *ComponentPricePointIntervalUnit) toMap() any {
	switch obj := c.value.(type) {
	case *IntervalUnit:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointIntervalUnit.
// It customizes the JSON unmarshaling process for ComponentPricePointIntervalUnit objects.
func (c *ComponentPricePointIntervalUnit) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(IntervalUnit), false, &c.isIntervalUnit),
	)

	c.value = result
	return err
}

func (c *ComponentPricePointIntervalUnit) AsIntervalUnit() (
	*IntervalUnit,
	bool) {
	if !c.isIntervalUnit {
		return nil, false
	}
	return c.value.(*IntervalUnit), true
}

// internalComponentPricePointIntervalUnit represents a componentPricePointIntervalUnit struct.
// This is a container for one-of cases.
type internalComponentPricePointIntervalUnit struct{}

var ComponentPricePointIntervalUnitContainer internalComponentPricePointIntervalUnit

// The internalComponentPricePointIntervalUnit instance, wrapping the provided IntervalUnit value.
func (c *internalComponentPricePointIntervalUnit) FromIntervalUnit(val IntervalUnit) ComponentPricePointIntervalUnit {
	return ComponentPricePointIntervalUnit{value: &val}
}
