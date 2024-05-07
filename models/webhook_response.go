package models

import (
    "encoding/json"
)

// WebhookResponse represents a WebhookResponse struct.
type WebhookResponse struct {
    Webhook              *Webhook       `json:"webhook,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookResponse.
// It customizes the JSON marshaling process for WebhookResponse objects.
func (w WebhookResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookResponse object to a map representation for JSON marshaling.
func (w WebhookResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Webhook != nil {
        structMap["webhook"] = w.Webhook.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookResponse.
// It customizes the JSON unmarshaling process for WebhookResponse objects.
func (w *WebhookResponse) UnmarshalJSON(input []byte) error {
    var temp webhookResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "webhook")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Webhook = temp.Webhook
    return nil
}

// webhookResponse is a temporary struct used for validating the fields of WebhookResponse.
type webhookResponse  struct {
    Webhook *Webhook `json:"webhook,omitempty"`
}
