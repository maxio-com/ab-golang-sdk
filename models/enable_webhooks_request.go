package models

import (
    "encoding/json"
)

// EnableWebhooksRequest represents a EnableWebhooksRequest struct.
type EnableWebhooksRequest struct {
    WebhooksEnabled bool `json:"webhooks_enabled"`
}

// MarshalJSON implements the json.Marshaler interface for EnableWebhooksRequest.
// It customizes the JSON marshaling process for EnableWebhooksRequest objects.
func (e *EnableWebhooksRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EnableWebhooksRequest object to a map representation for JSON marshaling.
func (e *EnableWebhooksRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["webhooks_enabled"] = e.WebhooksEnabled
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EnableWebhooksRequest.
// It customizes the JSON unmarshaling process for EnableWebhooksRequest objects.
func (e *EnableWebhooksRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        WebhooksEnabled bool `json:"webhooks_enabled"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    e.WebhooksEnabled = temp.WebhooksEnabled
    return nil
}
