// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ComponentPricePointAssignmentComponentId represents a ComponentPricePointAssignmentComponentId struct.
// This is a container for one-of cases.
type ComponentPricePointAssignmentComponentId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for ComponentPricePointAssignmentComponentId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ComponentPricePointAssignmentComponentId) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointAssignmentComponentId.
// It customizes the JSON marshaling process for ComponentPricePointAssignmentComponentId objects.
func (c ComponentPricePointAssignmentComponentId) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ComponentPricePointAssignmentComponentIdContainer.From*` functions to initialize the ComponentPricePointAssignmentComponentId object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointAssignmentComponentId object to a map representation for JSON marshaling.
func (c *ComponentPricePointAssignmentComponentId) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointAssignmentComponentId.
// It customizes the JSON unmarshaling process for ComponentPricePointAssignmentComponentId objects.
func (c *ComponentPricePointAssignmentComponentId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *ComponentPricePointAssignmentComponentId) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *ComponentPricePointAssignmentComponentId) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalComponentPricePointAssignmentComponentId represents a componentPricePointAssignmentComponentId struct.
// This is a container for one-of cases.
type internalComponentPricePointAssignmentComponentId struct {}

var ComponentPricePointAssignmentComponentIdContainer internalComponentPricePointAssignmentComponentId

// The internalComponentPricePointAssignmentComponentId instance, wrapping the provided int value.
func (c *internalComponentPricePointAssignmentComponentId) FromNumber(val int) ComponentPricePointAssignmentComponentId {
    return ComponentPricePointAssignmentComponentId{value: &val}
}

// The internalComponentPricePointAssignmentComponentId instance, wrapping the provided string value.
func (c *internalComponentPricePointAssignmentComponentId) FromString(val string) ComponentPricePointAssignmentComponentId {
    return ComponentPricePointAssignmentComponentId{value: &val}
}
