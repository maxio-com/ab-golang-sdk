// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// SubscriptionComponentResponse represents a SubscriptionComponentResponse struct.
type SubscriptionComponentResponse struct {
	Component            *SubscriptionComponent `json:"component,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionComponentResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentResponse) String() string {
	return fmt.Sprintf(
		"SubscriptionComponentResponse[Component=%v, AdditionalProperties=%v]",
		s.Component, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentResponse.
// It customizes the JSON marshaling process for SubscriptionComponentResponse objects.
func (s SubscriptionComponentResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"component"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentResponse object to a map representation for JSON marshaling.
func (s SubscriptionComponentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Component != nil {
		structMap["component"] = s.Component.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentResponse.
// It customizes the JSON unmarshaling process for SubscriptionComponentResponse objects.
func (s *SubscriptionComponentResponse) UnmarshalJSON(input []byte) error {
	var temp tempSubscriptionComponentResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "component")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Component = temp.Component
	return nil
}

// tempSubscriptionComponentResponse is a temporary struct used for validating the fields of SubscriptionComponentResponse.
type tempSubscriptionComponentResponse struct {
	Component *SubscriptionComponent `json:"component,omitempty"`
}
