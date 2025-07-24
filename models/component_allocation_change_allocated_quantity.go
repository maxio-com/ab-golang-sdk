// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ComponentAllocationChangeAllocatedQuantity represents a ComponentAllocationChangeAllocatedQuantity struct.
// This is a container for one-of cases.
type ComponentAllocationChangeAllocatedQuantity struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for ComponentAllocationChangeAllocatedQuantity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ComponentAllocationChangeAllocatedQuantity) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for ComponentAllocationChangeAllocatedQuantity.
// It customizes the JSON marshaling process for ComponentAllocationChangeAllocatedQuantity objects.
func (c ComponentAllocationChangeAllocatedQuantity) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ComponentAllocationChangeAllocatedQuantityContainer.From*` functions to initialize the ComponentAllocationChangeAllocatedQuantity object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentAllocationChangeAllocatedQuantity object to a map representation for JSON marshaling.
func (c *ComponentAllocationChangeAllocatedQuantity) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentAllocationChangeAllocatedQuantity.
// It customizes the JSON unmarshaling process for ComponentAllocationChangeAllocatedQuantity objects.
func (c *ComponentAllocationChangeAllocatedQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *ComponentAllocationChangeAllocatedQuantity) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *ComponentAllocationChangeAllocatedQuantity) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalComponentAllocationChangeAllocatedQuantity represents a componentAllocationChangeAllocatedQuantity struct.
// This is a container for one-of cases.
type internalComponentAllocationChangeAllocatedQuantity struct {}

var ComponentAllocationChangeAllocatedQuantityContainer internalComponentAllocationChangeAllocatedQuantity

// The internalComponentAllocationChangeAllocatedQuantity instance, wrapping the provided int value.
func (c *internalComponentAllocationChangeAllocatedQuantity) FromNumber(val int) ComponentAllocationChangeAllocatedQuantity {
    return ComponentAllocationChangeAllocatedQuantity{value: &val}
}

// The internalComponentAllocationChangeAllocatedQuantity instance, wrapping the provided string value.
func (c *internalComponentAllocationChangeAllocatedQuantity) FromString(val string) ComponentAllocationChangeAllocatedQuantity {
    return ComponentAllocationChangeAllocatedQuantity{value: &val}
}
