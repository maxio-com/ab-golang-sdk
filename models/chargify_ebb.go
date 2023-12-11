package models

import (
	"encoding/json"
)

// ChargifyEBB represents a ChargifyEBB struct.
type ChargifyEBB struct {
	// This timestamp determines what billing period the event will be billed in. If your request payload does not include it, Chargify will add `chargify.timestamp` to the event payload and set the value to `now`.
	Timestamp *string `json:"timestamp,omitempty"`
	// A unique ID set by Chargify. Please note that this field is reserved. If `chargify.id` is present in the request payload, it will be overwritten.
	Id *string `json:"id,omitempty"`
	// An ISO-8601 timestamp, set by Chargify at the time each event is recorded. Please note that this field is reserved. If `chargify.created_at` is present in the request payload, it will be overwritten.
	CreatedAt *string `json:"created_at,omitempty"`
	// User-defined string scoped per-stream. Duplicate events within a stream will be silently ignored. Tokens expire after 31 days.
	UniquenessToken *string `json:"uniqueness_token,omitempty"`
	// Id of Maxio Advanced Billing Subscription which is connected to this event.
	// Provide `subscription_id` if you configured `chargify.subscription_id` as Subscription Identifier in your Event Stream.
	SubscriptionId *int `json:"subscription_id,omitempty"`
	// Reference of Maxio Advanced Billing Subscription which is connected to this event.
	// Provide `subscription_reference` if you configured `chargify.subscription_reference` as Subscription Identifier in your Event Stream.
	SubscriptionReference *string `json:"subscription_reference,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ChargifyEBB.
// It customizes the JSON marshaling process for ChargifyEBB objects.
func (c *ChargifyEBB) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ChargifyEBB object to a map representation for JSON marshaling.
func (c *ChargifyEBB) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Timestamp != nil {
		structMap["timestamp"] = c.Timestamp
	}
	if c.Id != nil {
		structMap["id"] = c.Id
	}
	if c.CreatedAt != nil {
		structMap["created_at"] = c.CreatedAt
	}
	if c.UniquenessToken != nil {
		structMap["uniqueness_token"] = c.UniquenessToken
	}
	if c.SubscriptionId != nil {
		structMap["subscription_id"] = c.SubscriptionId
	}
	if c.SubscriptionReference != nil {
		structMap["subscription_reference"] = c.SubscriptionReference
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChargifyEBB.
// It customizes the JSON unmarshaling process for ChargifyEBB objects.
func (c *ChargifyEBB) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Timestamp             *string `json:"timestamp,omitempty"`
		Id                    *string `json:"id,omitempty"`
		CreatedAt             *string `json:"created_at,omitempty"`
		UniquenessToken       *string `json:"uniqueness_token,omitempty"`
		SubscriptionId        *int    `json:"subscription_id,omitempty"`
		SubscriptionReference *string `json:"subscription_reference,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Timestamp = temp.Timestamp
	c.Id = temp.Id
	c.CreatedAt = temp.CreatedAt
	c.UniquenessToken = temp.UniquenessToken
	c.SubscriptionId = temp.SubscriptionId
	c.SubscriptionReference = temp.SubscriptionReference
	return nil
}
