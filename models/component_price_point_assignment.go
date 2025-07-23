// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// ComponentPricePointAssignment represents a ComponentPricePointAssignment struct.
type ComponentPricePointAssignment struct {
	ComponentId          *int                                     `json:"component_id,omitempty"`
	PricePoint           *ComponentPricePointAssignmentPricePoint `json:"price_point,omitempty"`
	AdditionalProperties map[string]interface{}                   `json:"_"`
}

// String implements the fmt.Stringer interface for ComponentPricePointAssignment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ComponentPricePointAssignment) String() string {
	return fmt.Sprintf(
		"ComponentPricePointAssignment[ComponentId=%v, PricePoint=%v, AdditionalProperties=%v]",
		c.ComponentId, c.PricePoint, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointAssignment.
// It customizes the JSON marshaling process for ComponentPricePointAssignment objects.
func (c ComponentPricePointAssignment) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"component_id", "price_point"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointAssignment object to a map representation for JSON marshaling.
func (c ComponentPricePointAssignment) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.ComponentId != nil {
		structMap["component_id"] = c.ComponentId
	}
	if c.PricePoint != nil {
		structMap["price_point"] = c.PricePoint.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointAssignment.
// It customizes the JSON unmarshaling process for ComponentPricePointAssignment objects.
func (c *ComponentPricePointAssignment) UnmarshalJSON(input []byte) error {
	var temp tempComponentPricePointAssignment
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "component_id", "price_point")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.ComponentId = temp.ComponentId
	c.PricePoint = temp.PricePoint
	return nil
}

// tempComponentPricePointAssignment is a temporary struct used for validating the fields of ComponentPricePointAssignment.
type tempComponentPricePointAssignment struct {
	ComponentId *int                                     `json:"component_id,omitempty"`
	PricePoint  *ComponentPricePointAssignmentPricePoint `json:"price_point,omitempty"`
}
