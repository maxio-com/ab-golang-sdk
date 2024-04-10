package models

import (
    "encoding/json"
)

// UpdateReasonCode represents a UpdateReasonCode struct.
type UpdateReasonCode struct {
    // The unique identifier for the ReasonCode
    Code                 *string        `json:"code,omitempty"`
    // The friendly summary of what the code signifies
    Description          *string        `json:"description,omitempty"`
    // The order that code appears in lists
    Position             *int           `json:"position,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateReasonCode.
// It customizes the JSON marshaling process for UpdateReasonCode objects.
func (u UpdateReasonCode) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateReasonCode object to a map representation for JSON marshaling.
func (u UpdateReasonCode) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Code != nil {
        structMap["code"] = u.Code
    }
    if u.Description != nil {
        structMap["description"] = u.Description
    }
    if u.Position != nil {
        structMap["position"] = u.Position
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateReasonCode.
// It customizes the JSON unmarshaling process for UpdateReasonCode objects.
func (u *UpdateReasonCode) UnmarshalJSON(input []byte) error {
    var temp updateReasonCode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "code", "description", "position")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Code = temp.Code
    u.Description = temp.Description
    u.Position = temp.Position
    return nil
}

// TODO
type updateReasonCode  struct {
    Code        *string `json:"code,omitempty"`
    Description *string `json:"description,omitempty"`
    Position    *int    `json:"position,omitempty"`
}
