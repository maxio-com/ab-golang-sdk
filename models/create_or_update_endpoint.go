package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateOrUpdateEndpoint represents a CreateOrUpdateEndpoint struct.
// Used to Create or Update Endpoint
type CreateOrUpdateEndpoint struct {
	Url                  string                `json:"url"`
	WebhookSubscriptions []WebhookSubscription `json:"webhook_subscriptions"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateEndpoint.
// It customizes the JSON marshaling process for CreateOrUpdateEndpoint objects.
func (c *CreateOrUpdateEndpoint) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateEndpoint object to a map representation for JSON marshaling.
func (c *CreateOrUpdateEndpoint) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["url"] = c.Url
	structMap["webhook_subscriptions"] = c.WebhookSubscriptions
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateEndpoint.
// It customizes the JSON unmarshaling process for CreateOrUpdateEndpoint objects.
func (c *CreateOrUpdateEndpoint) UnmarshalJSON(input []byte) error {
	var temp createOrUpdateEndpoint
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.Url = *temp.Url
	c.WebhookSubscriptions = *temp.WebhookSubscriptions
	return nil
}

// TODO
type createOrUpdateEndpoint struct {
	Url                  *string                `json:"url"`
	WebhookSubscriptions *[]WebhookSubscription `json:"webhook_subscriptions"`
}

func (c *createOrUpdateEndpoint) validate() error {
	var errs []string
	if c.Url == nil {
		errs = append(errs, "required field `url` is missing for type `Create or Update Endpoint`")
	}
	if c.WebhookSubscriptions == nil {
		errs = append(errs, "required field `webhook_subscriptions` is missing for type `Create or Update Endpoint`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
