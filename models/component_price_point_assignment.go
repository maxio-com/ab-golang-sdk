package models

import (
    "encoding/json"
)

// ComponentPricePointAssignment represents a ComponentPricePointAssignment struct.
type ComponentPricePointAssignment struct {
    ComponentId          *int                                     `json:"component_id,omitempty"`
    PricePoint           *ComponentPricePointAssignmentPricePoint `json:"price_point,omitempty"`
    AdditionalProperties map[string]any                           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointAssignment.
// It customizes the JSON marshaling process for ComponentPricePointAssignment objects.
func (c ComponentPricePointAssignment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointAssignment object to a map representation for JSON marshaling.
func (c ComponentPricePointAssignment) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
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
    var temp componentPricePointAssignment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "component_id", "price_point")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.ComponentId = temp.ComponentId
    c.PricePoint = temp.PricePoint
    return nil
}

// componentPricePointAssignment is a temporary struct used for validating the fields of ComponentPricePointAssignment.
type componentPricePointAssignment  struct {
    ComponentId *int                                     `json:"component_id,omitempty"`
    PricePoint  *ComponentPricePointAssignmentPricePoint `json:"price_point,omitempty"`
}
