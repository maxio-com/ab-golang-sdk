package models

import (
    "encoding/json"
)

// UpdateMetadataRequest represents a UpdateMetadataRequest struct.
type UpdateMetadataRequest struct {
    Metadata             *UpdateMetadata `json:"metadata,omitempty"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateMetadataRequest.
// It customizes the JSON marshaling process for UpdateMetadataRequest objects.
func (u UpdateMetadataRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateMetadataRequest object to a map representation for JSON marshaling.
func (u UpdateMetadataRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Metadata != nil {
        structMap["metadata"] = u.Metadata.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateMetadataRequest.
// It customizes the JSON unmarshaling process for UpdateMetadataRequest objects.
func (u *UpdateMetadataRequest) UnmarshalJSON(input []byte) error {
    var temp updateMetadataRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "metadata")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Metadata = temp.Metadata
    return nil
}

// updateMetadataRequest is a temporary struct used for validating the fields of UpdateMetadataRequest.
type updateMetadataRequest  struct {
    Metadata *UpdateMetadata `json:"metadata,omitempty"`
}
