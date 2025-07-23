// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// PauseRequest represents a PauseRequest struct.
// Allows to pause a Subscription
type PauseRequest struct {
	Hold                 *AutoResume            `json:"hold,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PauseRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PauseRequest) String() string {
	return fmt.Sprintf(
		"PauseRequest[Hold=%v, AdditionalProperties=%v]",
		p.Hold, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PauseRequest.
// It customizes the JSON marshaling process for PauseRequest objects.
func (p PauseRequest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"hold"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PauseRequest object to a map representation for JSON marshaling.
func (p PauseRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, p.AdditionalProperties)
	if p.Hold != nil {
		structMap["hold"] = p.Hold.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PauseRequest.
// It customizes the JSON unmarshaling process for PauseRequest objects.
func (p *PauseRequest) UnmarshalJSON(input []byte) error {
	var temp tempPauseRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "hold")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.Hold = temp.Hold
	return nil
}

// tempPauseRequest is a temporary struct used for validating the fields of PauseRequest.
type tempPauseRequest struct {
	Hold *AutoResume `json:"hold,omitempty"`
}
