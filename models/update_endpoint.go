package models

import (
    "encoding/json"
)

// UpdateEndpoint represents a UpdateEndpoint struct.
// Used to Create or Update Endpoint
type UpdateEndpoint struct {
    Url                  string                `json:"url"`
    WebhookSubscriptions []WebhookSubscription `json:"webhook_subscriptions"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateEndpoint.
// It customizes the JSON marshaling process for UpdateEndpoint objects.
func (u *UpdateEndpoint) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateEndpoint object to a map representation for JSON marshaling.
func (u *UpdateEndpoint) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["url"] = u.Url
    structMap["webhook_subscriptions"] = u.WebhookSubscriptions
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateEndpoint.
// It customizes the JSON unmarshaling process for UpdateEndpoint objects.
func (u *UpdateEndpoint) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Url                  string                `json:"url"`
        WebhookSubscriptions []WebhookSubscription `json:"webhook_subscriptions"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.Url = temp.Url
    u.WebhookSubscriptions = temp.WebhookSubscriptions
    return nil
}
