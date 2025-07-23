// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// ReplayWebhooksResponse represents a ReplayWebhooksResponse struct.
type ReplayWebhooksResponse struct {
	Status               *string                `json:"status,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ReplayWebhooksResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ReplayWebhooksResponse) String() string {
	return fmt.Sprintf(
		"ReplayWebhooksResponse[Status=%v, AdditionalProperties=%v]",
		r.Status, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ReplayWebhooksResponse.
// It customizes the JSON marshaling process for ReplayWebhooksResponse objects.
func (r ReplayWebhooksResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"status"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ReplayWebhooksResponse object to a map representation for JSON marshaling.
func (r ReplayWebhooksResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Status != nil {
		structMap["status"] = r.Status
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReplayWebhooksResponse.
// It customizes the JSON unmarshaling process for ReplayWebhooksResponse objects.
func (r *ReplayWebhooksResponse) UnmarshalJSON(input []byte) error {
	var temp tempReplayWebhooksResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "status")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Status = temp.Status
	return nil
}

// tempReplayWebhooksResponse is a temporary struct used for validating the fields of ReplayWebhooksResponse.
type tempReplayWebhooksResponse struct {
	Status *string `json:"status,omitempty"`
}
