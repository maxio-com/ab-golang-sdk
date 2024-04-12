package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateMetadataRequest represents a CreateMetadataRequest struct.
type CreateMetadataRequest struct {
    Metadata             []CreateMetadata `json:"metadata"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateMetadataRequest.
// It customizes the JSON marshaling process for CreateMetadataRequest objects.
func (c CreateMetadataRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMetadataRequest object to a map representation for JSON marshaling.
func (c CreateMetadataRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["metadata"] = c.Metadata
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMetadataRequest.
// It customizes the JSON unmarshaling process for CreateMetadataRequest objects.
func (c *CreateMetadataRequest) UnmarshalJSON(input []byte) error {
    var temp createMetadataRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "metadata")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Metadata = *temp.Metadata
    return nil
}

// TODO
type createMetadataRequest  struct {
    Metadata *[]CreateMetadata `json:"metadata"`
}

func (c *createMetadataRequest) validate() error {
    var errs []string
    if c.Metadata == nil {
        errs = append(errs, "required field `metadata` is missing for type `Create Metadata Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
