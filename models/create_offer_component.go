package models

import (
    "encoding/json"
)

// CreateOfferComponent represents a CreateOfferComponent struct.
type CreateOfferComponent struct {
    ComponentId      *int `json:"component_id,omitempty"`
    StartingQuantity *int `json:"starting_quantity,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOfferComponent.
// It customizes the JSON marshaling process for CreateOfferComponent objects.
func (c *CreateOfferComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOfferComponent object to a map representation for JSON marshaling.
func (c *CreateOfferComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.ComponentId != nil {
        structMap["component_id"] = c.ComponentId
    }
    if c.StartingQuantity != nil {
        structMap["starting_quantity"] = c.StartingQuantity
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOfferComponent.
// It customizes the JSON unmarshaling process for CreateOfferComponent objects.
func (c *CreateOfferComponent) UnmarshalJSON(input []byte) error {
    temp := &struct {
        ComponentId      *int `json:"component_id,omitempty"`
        StartingQuantity *int `json:"starting_quantity,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.ComponentId = temp.ComponentId
    c.StartingQuantity = temp.StartingQuantity
    return nil
}
