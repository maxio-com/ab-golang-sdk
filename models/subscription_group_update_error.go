// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// SubscriptionGroupUpdateError represents a SubscriptionGroupUpdateError struct.
type SubscriptionGroupUpdateError struct {
	Members              []string               `json:"members,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupUpdateError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupUpdateError) String() string {
	return fmt.Sprintf(
		"SubscriptionGroupUpdateError[Members=%v, AdditionalProperties=%v]",
		s.Members, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupUpdateError.
// It customizes the JSON marshaling process for SubscriptionGroupUpdateError objects.
func (s SubscriptionGroupUpdateError) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"members"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupUpdateError object to a map representation for JSON marshaling.
func (s SubscriptionGroupUpdateError) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Members != nil {
		structMap["members"] = s.Members
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupUpdateError.
// It customizes the JSON unmarshaling process for SubscriptionGroupUpdateError objects.
func (s *SubscriptionGroupUpdateError) UnmarshalJSON(input []byte) error {
	var temp tempSubscriptionGroupUpdateError
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "members")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Members = temp.Members
	return nil
}

// tempSubscriptionGroupUpdateError is a temporary struct used for validating the fields of SubscriptionGroupUpdateError.
type tempSubscriptionGroupUpdateError struct {
	Members []string `json:"members,omitempty"`
}
