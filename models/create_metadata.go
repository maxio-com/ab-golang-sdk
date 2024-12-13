/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// CreateMetadata represents a CreateMetadata struct.
type CreateMetadata struct {
    Name                 *string                `json:"name,omitempty"`
    Value                *string                `json:"value,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateMetadata.
// It customizes the JSON marshaling process for CreateMetadata objects.
func (c CreateMetadata) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "name", "value"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMetadata object to a map representation for JSON marshaling.
func (c CreateMetadata) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    var temp tempCreateMetadata
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "value")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Name = temp.Name
    c.Value = temp.Value
    return nil
}

// tempCreateMetadata is a temporary struct used for validating the fields of CreateMetadata.
type tempCreateMetadata  struct {
    Name  *string `json:"name,omitempty"`
    Value *string `json:"value,omitempty"`
}
