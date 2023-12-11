package models

import (
	"encoding/json"
)

// CreateMetafieldsRequest represents a CreateMetafieldsRequest struct.
type CreateMetafieldsRequest struct {
	Metafields Metafields `json:"metafields"`
}

// MarshalJSON implements the json.Marshaler interface for CreateMetafieldsRequest.
// It customizes the JSON marshaling process for CreateMetafieldsRequest objects.
func (c *CreateMetafieldsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateMetafieldsRequest object to a map representation for JSON marshaling.
func (c *CreateMetafieldsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["metafields"] = c.Metafields
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMetafieldsRequest.
// It customizes the JSON unmarshaling process for CreateMetafieldsRequest objects.
func (c *CreateMetafieldsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Metafields Metafields `json:"metafields"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Metafields = temp.Metafields
	return nil
}
