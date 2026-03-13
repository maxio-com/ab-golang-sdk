// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// AvailableActions represents a AvailableActions struct.
type AvailableActions struct {
    SendEmail            *SendEmail             `json:"send_email,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AvailableActions,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AvailableActions) String() string {
    return fmt.Sprintf(
    	"AvailableActions[SendEmail=%v, AdditionalProperties=%v]",
    	a.SendEmail, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AvailableActions.
// It customizes the JSON marshaling process for AvailableActions objects.
func (a AvailableActions) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "send_email"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AvailableActions object to a map representation for JSON marshaling.
func (a AvailableActions) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.SendEmail != nil {
        structMap["send_email"] = a.SendEmail.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AvailableActions.
// It customizes the JSON unmarshaling process for AvailableActions objects.
func (a *AvailableActions) UnmarshalJSON(input []byte) error {
    var temp tempAvailableActions
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "send_email")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.SendEmail = temp.SendEmail
    return nil
}

// tempAvailableActions is a temporary struct used for validating the fields of AvailableActions.
type tempAvailableActions  struct {
    SendEmail *SendEmail `json:"send_email,omitempty"`
}
