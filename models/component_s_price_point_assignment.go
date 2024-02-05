package models

import (
    "encoding/json"
)

// ComponentSPricePointAssignment represents a ComponentSPricePointAssignment struct.
type ComponentSPricePointAssignment struct {
    ComponentId *int         `json:"component_id,omitempty"`
    PricePoint  *interface{} `json:"price_point,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentSPricePointAssignment.
// It customizes the JSON marshaling process for ComponentSPricePointAssignment objects.
func (c *ComponentSPricePointAssignment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentSPricePointAssignment object to a map representation for JSON marshaling.
func (c *ComponentSPricePointAssignment) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.ComponentId != nil {
        structMap["component_id"] = c.ComponentId
    }
    if c.PricePoint != nil {
        structMap["price_point"] = c.PricePoint
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentSPricePointAssignment.
// It customizes the JSON unmarshaling process for ComponentSPricePointAssignment objects.
func (c *ComponentSPricePointAssignment) UnmarshalJSON(input []byte) error {
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
