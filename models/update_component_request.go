package models

import (
    "encoding/json"
)

// UpdateComponentRequest represents a UpdateComponentRequest struct.
type UpdateComponentRequest struct {
    Component UpdateComponent `json:"component"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponentRequest.
// It customizes the JSON marshaling process for UpdateComponentRequest objects.
func (u *UpdateComponentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponentRequest object to a map representation for JSON marshaling.
func (u *UpdateComponentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["component"] = u.Component
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponentRequest.
// It customizes the JSON unmarshaling process for UpdateComponentRequest objects.
func (u *UpdateComponentRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Component UpdateComponent `json:"component"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.Component = temp.Component
    return nil
}
