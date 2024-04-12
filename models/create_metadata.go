package models

import (
    "encoding/json"
)

// CreateMetadata represents a CreateMetadata struct.
type CreateMetadata struct {
    Name                 *string        `json:"name,omitempty"`
    Value                *string        `json:"value,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateMetadata.
// It customizes the JSON marshaling process for CreateMetadata objects.
func (c CreateMetadata) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMetadata object to a map representation for JSON marshaling.
func (c CreateMetadata) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
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
    var temp createMetadata
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "value")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Name = temp.Name
    c.Value = temp.Value
    return nil
}

// TODO
type createMetadata  struct {
    Name  *string `json:"name,omitempty"`
    Value *string `json:"value,omitempty"`
}
