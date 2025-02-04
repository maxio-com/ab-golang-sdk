/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateMetadataRequest represents a CreateMetadataRequest struct.
type CreateMetadataRequest struct {
    Metadata             []CreateMetadata       `json:"metadata"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateMetadataRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateMetadataRequest) String() string {
    return fmt.Sprintf(
    	"CreateMetadataRequest[Metadata=%v, AdditionalProperties=%v]",
    	c.Metadata, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateMetadataRequest.
// It customizes the JSON marshaling process for CreateMetadataRequest objects.
func (c CreateMetadataRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "metadata"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMetadataRequest object to a map representation for JSON marshaling.
func (c CreateMetadataRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["metadata"] = c.Metadata
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMetadataRequest.
// It customizes the JSON unmarshaling process for CreateMetadataRequest objects.
func (c *CreateMetadataRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateMetadataRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "metadata")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Metadata = *temp.Metadata
    return nil
}

// tempCreateMetadataRequest is a temporary struct used for validating the fields of CreateMetadataRequest.
type tempCreateMetadataRequest  struct {
    Metadata *[]CreateMetadata `json:"metadata"`
}

func (c *tempCreateMetadataRequest) validate() error {
    var errs []string
    if c.Metadata == nil {
        errs = append(errs, "required field `metadata` is missing for type `Create Metadata Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
