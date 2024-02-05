package models

import (
    "encoding/json"
)

// BaseStringError represents a BaseStringError struct.
// The error is base if it is not directly associated with a single attribute.
type BaseStringError struct {
    Base []string `json:"base,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BaseStringError.
// It customizes the JSON marshaling process for BaseStringError objects.
func (b *BaseStringError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BaseStringError object to a map representation for JSON marshaling.
func (b *BaseStringError) toMap() map[string]any {
    structMap := make(map[string]any)
    if b.Base != nil {
        structMap["base"] = b.Base
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BaseStringError.
// It customizes the JSON unmarshaling process for BaseStringError objects.
func (b *BaseStringError) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Base []string `json:"base,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    b.Base = temp.Base
    return nil
}
