/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
	"encoding/json"
)

// WebhookResponse represents a WebhookResponse struct.
type WebhookResponse struct {
	Webhook *Webhook `json:"webhook,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookResponse.
// It customizes the JSON marshaling process for WebhookResponse objects.
func (w *WebhookResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookResponse object to a map representation for JSON marshaling.
func (w *WebhookResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if w.Webhook != nil {
		structMap["webhook"] = w.Webhook
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookResponse.
// It customizes the JSON unmarshaling process for WebhookResponse objects.
func (w *WebhookResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Webhook *Webhook `json:"webhook,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	w.Webhook = temp.Webhook
	return nil
}
