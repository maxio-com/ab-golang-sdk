// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// EnableWebhooksResponse represents a EnableWebhooksResponse struct.
type EnableWebhooksResponse struct {
    WebhooksEnabled      *bool                  `json:"webhooks_enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EnableWebhooksResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EnableWebhooksResponse) String() string {
    return fmt.Sprintf(
    	"EnableWebhooksResponse[WebhooksEnabled=%v, AdditionalProperties=%v]",
    	e.WebhooksEnabled, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EnableWebhooksResponse.
// It customizes the JSON marshaling process for EnableWebhooksResponse objects.
func (e EnableWebhooksResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "webhooks_enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EnableWebhooksResponse object to a map representation for JSON marshaling.
func (e EnableWebhooksResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.WebhooksEnabled != nil {
        structMap["webhooks_enabled"] = e.WebhooksEnabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EnableWebhooksResponse.
// It customizes the JSON unmarshaling process for EnableWebhooksResponse objects.
func (e *EnableWebhooksResponse) UnmarshalJSON(input []byte) error {
    var temp tempEnableWebhooksResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "webhooks_enabled")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.WebhooksEnabled = temp.WebhooksEnabled
    return nil
}

// tempEnableWebhooksResponse is a temporary struct used for validating the fields of EnableWebhooksResponse.
type tempEnableWebhooksResponse  struct {
    WebhooksEnabled *bool `json:"webhooks_enabled,omitempty"`
}
