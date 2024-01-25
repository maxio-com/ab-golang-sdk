package models

import (
	"encoding/json"
)

// CreateQuantityBasedComponent represents a CreateQuantityBasedComponent struct.
type CreateQuantityBasedComponent struct {
	QuantityBasedComponent QuantityBasedComponent `json:"quantity_based_component"`
}

// MarshalJSON implements the json.Marshaler interface for CreateQuantityBasedComponent.
// It customizes the JSON marshaling process for CreateQuantityBasedComponent objects.
func (c *CreateQuantityBasedComponent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateQuantityBasedComponent object to a map representation for JSON marshaling.
func (c *CreateQuantityBasedComponent) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["quantity_based_component"] = c.QuantityBasedComponent
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateQuantityBasedComponent.
// It customizes the JSON unmarshaling process for CreateQuantityBasedComponent objects.
func (c *CreateQuantityBasedComponent) UnmarshalJSON(input []byte) error {
	temp := &struct {
		QuantityBasedComponent QuantityBasedComponent `json:"quantity_based_component"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.QuantityBasedComponent = temp.QuantityBasedComponent
	return nil
}
