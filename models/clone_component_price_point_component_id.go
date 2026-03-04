// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CloneComponentPricePointComponentId represents a CloneComponentPricePointComponentId struct.
// This is a container for one-of cases.
type CloneComponentPricePointComponentId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for CloneComponentPricePointComponentId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CloneComponentPricePointComponentId) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CloneComponentPricePointComponentId.
// It customizes the JSON marshaling process for CloneComponentPricePointComponentId objects.
func (c CloneComponentPricePointComponentId) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CloneComponentPricePointComponentIdContainer.From*` functions to initialize the CloneComponentPricePointComponentId object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CloneComponentPricePointComponentId object to a map representation for JSON marshaling.
func (c *CloneComponentPricePointComponentId) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CloneComponentPricePointComponentId.
// It customizes the JSON unmarshaling process for CloneComponentPricePointComponentId objects.
func (c *CloneComponentPricePointComponentId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CloneComponentPricePointComponentId) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CloneComponentPricePointComponentId) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCloneComponentPricePointComponentId represents a cloneComponentPricePointComponentId struct.
// This is a container for one-of cases.
type internalCloneComponentPricePointComponentId struct {}

var CloneComponentPricePointComponentIdContainer internalCloneComponentPricePointComponentId

// The internalCloneComponentPricePointComponentId instance, wrapping the provided int value.
func (c *internalCloneComponentPricePointComponentId) FromNumber(val int) CloneComponentPricePointComponentId {
    return CloneComponentPricePointComponentId{value: &val}
}

// The internalCloneComponentPricePointComponentId instance, wrapping the provided string value.
func (c *internalCloneComponentPricePointComponentId) FromString(val string) CloneComponentPricePointComponentId {
    return CloneComponentPricePointComponentId{value: &val}
}
