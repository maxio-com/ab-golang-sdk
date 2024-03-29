package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// UpdateSubscriptionRequest represents a UpdateSubscriptionRequest struct.
type UpdateSubscriptionRequest struct {
	Subscription UpdateSubscription `json:"subscription"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionRequest objects.
func (u *UpdateSubscriptionRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscription"] = u.Subscription.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionRequest objects.
func (u *UpdateSubscriptionRequest) UnmarshalJSON(input []byte) error {
	var temp updateSubscriptionRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	u.Subscription = *temp.Subscription
	return nil
}

// TODO
type updateSubscriptionRequest struct {
	Subscription *UpdateSubscription `json:"subscription"`
}

func (u *updateSubscriptionRequest) validate() error {
	var errs []string
	if u.Subscription == nil {
		errs = append(errs, "required field `subscription` is missing for type `Update Subscription Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
