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

// ComponentPricePointAssignmentPricePoint represents a ComponentPricePointAssignmentPricePoint struct.
// This is a container for one-of cases.
type ComponentPricePointAssignmentPricePoint struct {
	value    any
	isString bool
	isNumber bool
}

// String converts the ComponentPricePointAssignmentPricePoint object to a string representation.
func (c ComponentPricePointAssignmentPricePoint) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointAssignmentPricePoint.
// It customizes the JSON marshaling process for ComponentPricePointAssignmentPricePoint objects.
func (c *ComponentPricePointAssignmentPricePoint) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ComponentPricePointAssignmentPricePointContainer.From*` functions to initialize the ComponentPricePointAssignmentPricePoint object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointAssignmentPricePoint object to a map representation for JSON marshaling.
func (c *ComponentPricePointAssignmentPricePoint) toMap() any {
	switch obj := c.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointAssignmentPricePoint.
// It customizes the JSON unmarshaling process for ComponentPricePointAssignmentPricePoint objects.
func (c *ComponentPricePointAssignmentPricePoint) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &c.isString),
		NewTypeHolder(new(int), false, &c.isNumber),
	)

	c.value = result
	return err
}

func (c *ComponentPricePointAssignmentPricePoint) AsString() (
	*string,
	bool) {
	if !c.isString {
		return nil, false
	}
	return c.value.(*string), true
}

func (c *ComponentPricePointAssignmentPricePoint) AsNumber() (
	*int,
	bool) {
	if !c.isNumber {
		return nil, false
	}
	return c.value.(*int), true
}

// internalComponentPricePointAssignmentPricePoint represents a componentPricePointAssignmentPricePoint struct.
// This is a container for one-of cases.
type internalComponentPricePointAssignmentPricePoint struct{}

var ComponentPricePointAssignmentPricePointContainer internalComponentPricePointAssignmentPricePoint

func (c *internalComponentPricePointAssignmentPricePoint) FromString(val string) ComponentPricePointAssignmentPricePoint {
	return ComponentPricePointAssignmentPricePoint{value: &val}
}

func (c *internalComponentPricePointAssignmentPricePoint) FromNumber(val int) ComponentPricePointAssignmentPricePoint {
	return ComponentPricePointAssignmentPricePoint{value: &val}
}
