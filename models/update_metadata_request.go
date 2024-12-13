/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// UpdateMetadataRequest represents a UpdateMetadataRequest struct.
type UpdateMetadataRequest struct {
    Metadata             *UpdateMetadata        `json:"metadata,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateMetadataRequest.
// It customizes the JSON marshaling process for UpdateMetadataRequest objects.
func (u UpdateMetadataRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "metadata"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateMetadataRequest object to a map representation for JSON marshaling.
func (u UpdateMetadataRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Metadata != nil {
        structMap["metadata"] = u.Metadata.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateMetadataRequest.
// It customizes the JSON unmarshaling process for UpdateMetadataRequest objects.
func (u *UpdateMetadataRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateMetadataRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "metadata")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Metadata = temp.Metadata
    return nil
}

// tempUpdateMetadataRequest is a temporary struct used for validating the fields of UpdateMetadataRequest.
type tempUpdateMetadataRequest  struct {
    Metadata *UpdateMetadata `json:"metadata,omitempty"`
}
