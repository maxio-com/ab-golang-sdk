// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CreateComponentPricePointsRequestPricePoints represents a CreateComponentPricePointsRequestPricePoints struct.
// This is Array of a container for any-of cases.
type CreateComponentPricePointsRequestPricePoints struct {
    value                                   any
    isCreateComponentPricePoint             bool
    isCreatePrepaidUsageComponentPricePoint bool
}

// String implements the fmt.Stringer interface for CreateComponentPricePointsRequestPricePoints,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateComponentPricePointsRequestPricePoints) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreateComponentPricePointsRequestPricePoints.
// It customizes the JSON marshaling process for CreateComponentPricePointsRequestPricePoints objects.
func (c CreateComponentPricePointsRequestPricePoints) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateComponentPricePointsRequestPricePointsContainer.From*` functions to initialize the CreateComponentPricePointsRequestPricePoints object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateComponentPricePointsRequestPricePoints object to a map representation for JSON marshaling.
func (c *CreateComponentPricePointsRequestPricePoints) toMap() any {
    switch obj := c.value.(type) {
    case *CreateComponentPricePoint:
        return obj.toMap()
    case *CreatePrepaidUsageComponentPricePoint:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateComponentPricePointsRequestPricePoints.
// It customizes the JSON unmarshaling process for CreateComponentPricePointsRequestPricePoints objects.
func (c *CreateComponentPricePointsRequestPricePoints) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&CreateComponentPricePoint{}, false, &c.isCreateComponentPricePoint),
        NewTypeHolder(&CreatePrepaidUsageComponentPricePoint{}, false, &c.isCreatePrepaidUsageComponentPricePoint),
    )
    
    c.value = result
    return err
}

func (c *CreateComponentPricePointsRequestPricePoints) AsCreateComponentPricePoint() (
    *CreateComponentPricePoint,
    bool) {
    if !c.isCreateComponentPricePoint {
        return nil, false
    }
    return c.value.(*CreateComponentPricePoint), true
}

func (c *CreateComponentPricePointsRequestPricePoints) AsCreatePrepaidUsageComponentPricePoint() (
    *CreatePrepaidUsageComponentPricePoint,
    bool) {
    if !c.isCreatePrepaidUsageComponentPricePoint {
        return nil, false
    }
    return c.value.(*CreatePrepaidUsageComponentPricePoint), true
}

// internalCreateComponentPricePointsRequestPricePoints represents a createComponentPricePointsRequestPricePoints struct.
// This is Array of a container for any-of cases.
type internalCreateComponentPricePointsRequestPricePoints struct {}

var CreateComponentPricePointsRequestPricePointsContainer internalCreateComponentPricePointsRequestPricePoints

// The internalCreateComponentPricePointsRequestPricePoints instance, wrapping the provided CreateComponentPricePoint value.
func (c *internalCreateComponentPricePointsRequestPricePoints) FromCreateComponentPricePoint(val CreateComponentPricePoint) CreateComponentPricePointsRequestPricePoints {
    return CreateComponentPricePointsRequestPricePoints{value: &val}
}

// The internalCreateComponentPricePointsRequestPricePoints instance, wrapping the provided CreatePrepaidUsageComponentPricePoint value.
func (c *internalCreateComponentPricePointsRequestPricePoints) FromCreatePrepaidUsageComponentPricePoint(val CreatePrepaidUsageComponentPricePoint) CreateComponentPricePointsRequestPricePoints {
    return CreateComponentPricePointsRequestPricePoints{value: &val}
}
