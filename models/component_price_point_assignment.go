package models

import (
    "encoding/json"
)

// ComponentPricePointAssignment represents a ComponentPricePointAssignment struct.
type ComponentPricePointAssignment struct {
    ComponentId *int         `json:"component_id,omitempty"`
    PricePoint  *interface{} `json:"price_point,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointAssignment.
// It customizes the JSON marshaling process for ComponentPricePointAssignment objects.
func (c *ComponentPricePointAssignment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointAssignment object to a map representation for JSON marshaling.
func (c *ComponentPricePointAssignment) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.ComponentId != nil {
        structMap["component_id"] = c.ComponentId
    }
    if c.PricePoint != nil {
        structMap["price_point"] = c.PricePoint
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointAssignment.
// It customizes the JSON unmarshaling process for ComponentPricePointAssignment objects.
func (c *ComponentPricePointAssignment) UnmarshalJSON(input []byte) error {
    temp := &struct {
        ComponentId *int         `json:"component_id,omitempty"`
        PricePoint  *interface{} `json:"price_point,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.ComponentId = temp.ComponentId
    c.PricePoint = temp.PricePoint
    return nil
}
