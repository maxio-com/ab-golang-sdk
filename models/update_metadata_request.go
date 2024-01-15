package models

import (
    "encoding/json"
)

// UpdateMetadataRequest represents a UpdateMetadataRequest struct.
type UpdateMetadataRequest struct {
    Metadata *UpdateMetadata `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateMetadataRequest.
// It customizes the JSON marshaling process for UpdateMetadataRequest objects.
func (u *UpdateMetadataRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateMetadataRequest object to a map representation for JSON marshaling.
func (u *UpdateMetadataRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    if u.Metadata != nil {
        structMap["metadata"] = u.Metadata
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateMetadataRequest.
// It customizes the JSON unmarshaling process for UpdateMetadataRequest objects.
func (u *UpdateMetadataRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Metadata *UpdateMetadata `json:"metadata,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.Metadata = temp.Metadata
    return nil
}
