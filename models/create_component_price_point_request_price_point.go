/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CreateComponentPricePointRequestPricePoint represents a CreateComponentPricePointRequestPricePoint struct.
// This is a container for any-of cases.
type CreateComponentPricePointRequestPricePoint struct {
    value                                   any
    isCreateComponentPricePoint             bool
    isCreatePrepaidUsageComponentPricePoint bool
}

// String implements the fmt.Stringer interface for CreateComponentPricePointRequestPricePoint,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateComponentPricePointRequestPricePoint) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateComponentPricePointRequestPricePoint.
// It customizes the JSON marshaling process for CreateComponentPricePointRequestPricePoint objects.
func (c CreateComponentPricePointRequestPricePoint) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateComponentPricePointRequestPricePointContainer.From*` functions to initialize the CreateComponentPricePointRequestPricePoint object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateComponentPricePointRequestPricePoint object to a map representation for JSON marshaling.
func (c *CreateComponentPricePointRequestPricePoint) toMap() any {
    switch obj := c.value.(type) {
    case *CreateComponentPricePoint:
        return obj.toMap()
    case *CreatePrepaidUsageComponentPricePoint:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateComponentPricePointRequestPricePoint.
// It customizes the JSON unmarshaling process for CreateComponentPricePointRequestPricePoint objects.
func (c *CreateComponentPricePointRequestPricePoint) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&CreateComponentPricePoint{}, false, &c.isCreateComponentPricePoint),
        NewTypeHolder(&CreatePrepaidUsageComponentPricePoint{}, false, &c.isCreatePrepaidUsageComponentPricePoint),
    )
    
    c.value = result
    return err
}

func (c *CreateComponentPricePointRequestPricePoint) AsCreateComponentPricePoint() (
    *CreateComponentPricePoint,
    bool) {
    if !c.isCreateComponentPricePoint {
        return nil, false
    }
    return c.value.(*CreateComponentPricePoint), true
}

func (c *CreateComponentPricePointRequestPricePoint) AsCreatePrepaidUsageComponentPricePoint() (
    *CreatePrepaidUsageComponentPricePoint,
    bool) {
    if !c.isCreatePrepaidUsageComponentPricePoint {
        return nil, false
    }
    return c.value.(*CreatePrepaidUsageComponentPricePoint), true
}

// internalCreateComponentPricePointRequestPricePoint represents a createComponentPricePointRequestPricePoint struct.
// This is a container for any-of cases.
type internalCreateComponentPricePointRequestPricePoint struct {}

var CreateComponentPricePointRequestPricePointContainer internalCreateComponentPricePointRequestPricePoint

// The internalCreateComponentPricePointRequestPricePoint instance, wrapping the provided CreateComponentPricePoint value.
func (c *internalCreateComponentPricePointRequestPricePoint) FromCreateComponentPricePoint(val CreateComponentPricePoint) CreateComponentPricePointRequestPricePoint {
    return CreateComponentPricePointRequestPricePoint{value: &val}
}

// The internalCreateComponentPricePointRequestPricePoint instance, wrapping the provided CreatePrepaidUsageComponentPricePoint value.
func (c *internalCreateComponentPricePointRequestPricePoint) FromCreatePrepaidUsageComponentPricePoint(val CreatePrepaidUsageComponentPricePoint) CreateComponentPricePointRequestPricePoint {
    return CreateComponentPricePointRequestPricePoint{value: &val}
}
