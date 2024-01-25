package models

import (
	"encoding/json"
)

// Endpoint represents a Endpoint struct.
type Endpoint struct {
	Id                   *int     `json:"id,omitempty"`
	Url                  *string  `json:"url,omitempty"`
	SiteId               *int     `json:"site_id,omitempty"`
	Status               *string  `json:"status,omitempty"`
	WebhookSubscriptions []string `json:"webhook_subscriptions,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Endpoint.
// It customizes the JSON marshaling process for Endpoint objects.
func (e *Endpoint) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(e.toMap())
}

// toMap converts the Endpoint object to a map representation for JSON marshaling.
func (e *Endpoint) toMap() map[string]any {
	structMap := make(map[string]any)
	if e.Id != nil {
		structMap["id"] = e.Id
	}
	if e.Url != nil {
		structMap["url"] = e.Url
	}
	if e.SiteId != nil {
		structMap["site_id"] = e.SiteId
	}
	if e.Status != nil {
		structMap["status"] = e.Status
	}
	if e.WebhookSubscriptions != nil {
		structMap["webhook_subscriptions"] = e.WebhookSubscriptions
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Endpoint.
// It customizes the JSON unmarshaling process for Endpoint objects.
func (e *Endpoint) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                   *int     `json:"id,omitempty"`
		Url                  *string  `json:"url,omitempty"`
		SiteId               *int     `json:"site_id,omitempty"`
		Status               *string  `json:"status,omitempty"`
		WebhookSubscriptions []string `json:"webhook_subscriptions,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	e.Id = temp.Id
	e.Url = temp.Url
	e.SiteId = temp.SiteId
	e.Status = temp.Status
	e.WebhookSubscriptions = temp.WebhookSubscriptions
	return nil
}
