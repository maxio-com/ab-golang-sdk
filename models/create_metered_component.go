package models

import (
    "encoding/json"
)

// CreateMeteredComponent represents a CreateMeteredComponent struct.
type CreateMeteredComponent struct {
    MeteredComponent MeteredComponent `json:"metered_component"`
}

// MarshalJSON implements the json.Marshaler interface for CreateMeteredComponent.
// It customizes the JSON marshaling process for CreateMeteredComponent objects.
func (c *CreateMeteredComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMeteredComponent object to a map representation for JSON marshaling.
func (c *CreateMeteredComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["metered_component"] = c.MeteredComponent.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMeteredComponent.
// It customizes the JSON unmarshaling process for CreateMeteredComponent objects.
func (c *CreateMeteredComponent) UnmarshalJSON(input []byte) error {
    temp := &struct {
        MeteredComponent MeteredComponent `json:"metered_component"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.MeteredComponent = temp.MeteredComponent
    return nil
}
