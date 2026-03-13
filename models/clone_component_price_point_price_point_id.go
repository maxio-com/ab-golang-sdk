// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CloneComponentPricePointPricePointId represents a CloneComponentPricePointPricePointId struct.
// This is a container for one-of cases.
type CloneComponentPricePointPricePointId struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for CloneComponentPricePointPricePointId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CloneComponentPricePointPricePointId) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CloneComponentPricePointPricePointId.
// It customizes the JSON marshaling process for CloneComponentPricePointPricePointId objects.
func (c CloneComponentPricePointPricePointId) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CloneComponentPricePointPricePointIdContainer.From*` functions to initialize the CloneComponentPricePointPricePointId object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CloneComponentPricePointPricePointId object to a map representation for JSON marshaling.
func (c *CloneComponentPricePointPricePointId) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CloneComponentPricePointPricePointId.
// It customizes the JSON unmarshaling process for CloneComponentPricePointPricePointId objects.
func (c *CloneComponentPricePointPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CloneComponentPricePointPricePointId) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CloneComponentPricePointPricePointId) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCloneComponentPricePointPricePointId represents a cloneComponentPricePointPricePointId struct.
// This is a container for one-of cases.
type internalCloneComponentPricePointPricePointId struct {}

var CloneComponentPricePointPricePointIdContainer internalCloneComponentPricePointPricePointId

// The internalCloneComponentPricePointPricePointId instance, wrapping the provided int value.
func (c *internalCloneComponentPricePointPricePointId) FromNumber(val int) CloneComponentPricePointPricePointId {
    return CloneComponentPricePointPricePointId{value: &val}
}

// The internalCloneComponentPricePointPricePointId instance, wrapping the provided string value.
func (c *internalCloneComponentPricePointPricePointId) FromString(val string) CloneComponentPricePointPricePointId {
    return CloneComponentPricePointPricePointId{value: &val}
}
