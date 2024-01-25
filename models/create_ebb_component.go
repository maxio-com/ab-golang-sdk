package models

import (
	"encoding/json"
)

// CreateEBBComponent represents a CreateEBBComponent struct.
type CreateEBBComponent struct {
	EventBasedComponent EBBComponent `json:"event_based_component"`
}

// MarshalJSON implements the json.Marshaler interface for CreateEBBComponent.
// It customizes the JSON marshaling process for CreateEBBComponent objects.
func (c *CreateEBBComponent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateEBBComponent object to a map representation for JSON marshaling.
func (c *CreateEBBComponent) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["event_based_component"] = c.EventBasedComponent
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateEBBComponent.
// It customizes the JSON unmarshaling process for CreateEBBComponent objects.
func (c *CreateEBBComponent) UnmarshalJSON(input []byte) error {
	temp := &struct {
		EventBasedComponent EBBComponent `json:"event_based_component"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.EventBasedComponent = temp.EventBasedComponent
	return nil
}
