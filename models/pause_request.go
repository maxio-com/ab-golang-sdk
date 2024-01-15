package models

import (
    "encoding/json"
)

// PauseRequest represents a PauseRequest struct.
// Allows to pause a Subscription
type PauseRequest struct {
    Hold *AutoResume `json:"hold,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PauseRequest.
// It customizes the JSON marshaling process for PauseRequest objects.
func (p *PauseRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PauseRequest object to a map representation for JSON marshaling.
func (p *PauseRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.Hold != nil {
        structMap["hold"] = p.Hold
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PauseRequest.
// It customizes the JSON unmarshaling process for PauseRequest objects.
func (p *PauseRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Hold *AutoResume `json:"hold,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Hold = temp.Hold
    return nil
}
