package models

import (
	"encoding/json"
)

// CreateMetadata represents a CreateMetadata struct.
type CreateMetadata struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateMetadata.
// It customizes the JSON marshaling process for CreateMetadata objects.
func (c *CreateMetadata) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateMetadata object to a map representation for JSON marshaling.
func (c *CreateMetadata) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Name != nil {
		structMap["name"] = c.Name
	}
	if c.Value != nil {
		structMap["value"] = c.Value
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMetadata.
// It customizes the JSON unmarshaling process for CreateMetadata objects.
func (c *CreateMetadata) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name  *string `json:"name,omitempty"`
		Value *string `json:"value,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Name = temp.Name
	c.Value = temp.Value
	return nil
}
