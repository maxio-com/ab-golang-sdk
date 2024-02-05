package models

import (
    "encoding/json"
)

// AttributeError represents a AttributeError struct.
type AttributeError struct {
    Attribute []string `json:"attribute"`
}

// MarshalJSON implements the json.Marshaler interface for AttributeError.
// It customizes the JSON marshaling process for AttributeError objects.
func (a *AttributeError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AttributeError object to a map representation for JSON marshaling.
func (a *AttributeError) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["attribute"] = a.Attribute
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AttributeError.
// It customizes the JSON unmarshaling process for AttributeError objects.
func (a *AttributeError) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Attribute []string `json:"attribute"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    a.Attribute = temp.Attribute
    return nil
}
