package models

import (
	"encoding/json"
)

// CreateSubscriptionRequest represents a CreateSubscriptionRequest struct.
type CreateSubscriptionRequest struct {
	Subscription CreateSubscription `json:"subscription"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionRequest.
// It customizes the JSON marshaling process for CreateSubscriptionRequest objects.
func (c *CreateSubscriptionRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionRequest object to a map representation for JSON marshaling.
func (c *CreateSubscriptionRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscription"] = c.Subscription
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionRequest.
// It customizes the JSON unmarshaling process for CreateSubscriptionRequest objects.
func (c *CreateSubscriptionRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Subscription CreateSubscription `json:"subscription"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Subscription = temp.Subscription
	return nil
}
