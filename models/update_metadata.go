package models

import (
    "encoding/json"
)

// UpdateMetadata represents a UpdateMetadata struct.
type UpdateMetadata struct {
    CurrentName *string `json:"current_name,omitempty"`
    Name        *string `json:"name,omitempty"`
    Value       *string `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateMetadata.
// It customizes the JSON marshaling process for UpdateMetadata objects.
func (u *UpdateMetadata) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateMetadata object to a map representation for JSON marshaling.
func (u *UpdateMetadata) toMap() map[string]any {
    structMap := make(map[string]any)
    if u.CurrentName != nil {
        structMap["current_name"] = u.CurrentName
    }
    if u.Name != nil {
        structMap["name"] = u.Name
    }
    if u.Value != nil {
        structMap["value"] = u.Value
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateMetadata.
// It customizes the JSON unmarshaling process for UpdateMetadata objects.
func (u *UpdateMetadata) UnmarshalJSON(input []byte) error {
    temp := &struct {
        CurrentName *string `json:"current_name,omitempty"`
        Name        *string `json:"name,omitempty"`
        Value       *string `json:"value,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.CurrentName = temp.CurrentName
    u.Name = temp.Name
    u.Value = temp.Value
    return nil
}
