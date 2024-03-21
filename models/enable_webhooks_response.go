package models

import (
	"encoding/json"
)

// EnableWebhooksResponse represents a EnableWebhooksResponse struct.
type EnableWebhooksResponse struct {
	WebhooksEnabled *bool `json:"webhooks_enabled,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for EnableWebhooksResponse.
// It customizes the JSON marshaling process for EnableWebhooksResponse objects.
func (e *EnableWebhooksResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(e.toMap())
}

// toMap converts the EnableWebhooksResponse object to a map representation for JSON marshaling.
func (e *EnableWebhooksResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if e.WebhooksEnabled != nil {
		structMap["webhooks_enabled"] = e.WebhooksEnabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EnableWebhooksResponse.
// It customizes the JSON unmarshaling process for EnableWebhooksResponse objects.
func (e *EnableWebhooksResponse) UnmarshalJSON(input []byte) error {
	var temp enableWebhooksResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	e.WebhooksEnabled = temp.WebhooksEnabled
	return nil
}

// TODO
type enableWebhooksResponse struct {
	WebhooksEnabled *bool `json:"webhooks_enabled,omitempty"`
}
