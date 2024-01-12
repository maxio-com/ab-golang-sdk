package models

import (
    "encoding/json"
)

// CreateReasonCode represents a CreateReasonCode struct.
type CreateReasonCode struct {
    // The unique identifier for the ReasonCode
    Code        string `json:"code"`
    // The friendly summary of what the code signifies
    Description string `json:"description"`
    // The order that code appears in lists
    Position    *int   `json:"position,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateReasonCode.
// It customizes the JSON marshaling process for CreateReasonCode objects.
func (c *CreateReasonCode) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateReasonCode object to a map representation for JSON marshaling.
func (c *CreateReasonCode) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["code"] = c.Code
    structMap["description"] = c.Description
    if c.Position != nil {
        structMap["position"] = c.Position
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateReasonCode.
// It customizes the JSON unmarshaling process for CreateReasonCode objects.
func (c *CreateReasonCode) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Code        string `json:"code"`
        Description string `json:"description"`
        Position    *int   `json:"position,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Code = temp.Code
    c.Description = temp.Description
    c.Position = temp.Position
    return nil
}
