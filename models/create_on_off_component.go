package models

import (
    "encoding/json"
)

// CreateOnOffComponent represents a CreateOnOffComponent struct.
type CreateOnOffComponent struct {
    OnOffComponent OnOffComponent `json:"on_off_component"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOnOffComponent.
// It customizes the JSON marshaling process for CreateOnOffComponent objects.
func (c *CreateOnOffComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOnOffComponent object to a map representation for JSON marshaling.
func (c *CreateOnOffComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["on_off_component"] = c.OnOffComponent.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOnOffComponent.
// It customizes the JSON unmarshaling process for CreateOnOffComponent objects.
func (c *CreateOnOffComponent) UnmarshalJSON(input []byte) error {
    temp := &struct {
        OnOffComponent OnOffComponent `json:"on_off_component"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.OnOffComponent = temp.OnOffComponent
    return nil
}
