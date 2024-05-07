package models

import (
    "encoding/json"
)

// PauseRequest represents a PauseRequest struct.
// Allows to pause a Subscription
type PauseRequest struct {
    Hold                 *AutoResume    `json:"hold,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PauseRequest.
// It customizes the JSON marshaling process for PauseRequest objects.
func (p PauseRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PauseRequest object to a map representation for JSON marshaling.
func (p PauseRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Hold != nil {
        structMap["hold"] = p.Hold.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PauseRequest.
// It customizes the JSON unmarshaling process for PauseRequest objects.
func (p *PauseRequest) UnmarshalJSON(input []byte) error {
    var temp pauseRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "hold")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Hold = temp.Hold
    return nil
}

// pauseRequest is a temporary struct used for validating the fields of PauseRequest.
type pauseRequest  struct {
    Hold *AutoResume `json:"hold,omitempty"`
}
