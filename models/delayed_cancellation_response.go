package models

import (
    "encoding/json"
)

// DelayedCancellationResponse represents a DelayedCancellationResponse struct.
type DelayedCancellationResponse struct {
    Message              *string        `json:"message,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DelayedCancellationResponse.
// It customizes the JSON marshaling process for DelayedCancellationResponse objects.
func (d DelayedCancellationResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DelayedCancellationResponse object to a map representation for JSON marshaling.
func (d DelayedCancellationResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Message != nil {
        structMap["message"] = d.Message
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DelayedCancellationResponse.
// It customizes the JSON unmarshaling process for DelayedCancellationResponse objects.
func (d *DelayedCancellationResponse) UnmarshalJSON(input []byte) error {
    var temp delayedCancellationResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "message")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Message = temp.Message
    return nil
}

// TODO
type delayedCancellationResponse  struct {
    Message *string `json:"message,omitempty"`
}
