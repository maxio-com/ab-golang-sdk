package models

import (
	"encoding/json"
)

// OverrideSubscriptionRequest represents a OverrideSubscriptionRequest struct.
type OverrideSubscriptionRequest struct {
	Subscription OverrideSubscription `json:"subscription"`
}

// MarshalJSON implements the json.Marshaler interface for OverrideSubscriptionRequest.
// It customizes the JSON marshaling process for OverrideSubscriptionRequest objects.
func (o *OverrideSubscriptionRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OverrideSubscriptionRequest object to a map representation for JSON marshaling.
func (o *OverrideSubscriptionRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscription"] = o.Subscription
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OverrideSubscriptionRequest.
// It customizes the JSON unmarshaling process for OverrideSubscriptionRequest objects.
func (o *OverrideSubscriptionRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Subscription OverrideSubscription `json:"subscription"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	o.Subscription = temp.Subscription
	return nil
}
