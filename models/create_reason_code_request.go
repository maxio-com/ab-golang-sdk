package models

import (
    "encoding/json"
)

// CreateReasonCodeRequest represents a CreateReasonCodeRequest struct.
type CreateReasonCodeRequest struct {
    ReasonCode CreateReasonCode `json:"reason_code"`
}

// MarshalJSON implements the json.Marshaler interface for CreateReasonCodeRequest.
// It customizes the JSON marshaling process for CreateReasonCodeRequest objects.
func (c *CreateReasonCodeRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateReasonCodeRequest object to a map representation for JSON marshaling.
func (c *CreateReasonCodeRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["reason_code"] = c.ReasonCode
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateReasonCodeRequest.
// It customizes the JSON unmarshaling process for CreateReasonCodeRequest objects.
func (c *CreateReasonCodeRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        ReasonCode CreateReasonCode `json:"reason_code"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.ReasonCode = temp.ReasonCode
    return nil
}
