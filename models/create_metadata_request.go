package models

import (
    "encoding/json"
)

// CreateMetadataRequest represents a CreateMetadataRequest struct.
type CreateMetadataRequest struct {
    Metadata []CreateMetadata `json:"metadata"`
}

// MarshalJSON implements the json.Marshaler interface for CreateMetadataRequest.
// It customizes the JSON marshaling process for CreateMetadataRequest objects.
func (c *CreateMetadataRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMetadataRequest object to a map representation for JSON marshaling.
func (c *CreateMetadataRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["metadata"] = c.Metadata
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMetadataRequest.
// It customizes the JSON unmarshaling process for CreateMetadataRequest objects.
func (c *CreateMetadataRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Metadata []CreateMetadata `json:"metadata"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Metadata = temp.Metadata
    return nil
}
