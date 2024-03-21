package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// SubscriptionGroupResponse represents a SubscriptionGroupResponse struct.
type SubscriptionGroupResponse struct {
	SubscriptionGroup SubscriptionGroup `json:"subscription_group"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupResponse.
// It customizes the JSON marshaling process for SubscriptionGroupResponse objects.
func (s *SubscriptionGroupResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupResponse object to a map representation for JSON marshaling.
func (s *SubscriptionGroupResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscription_group"] = s.SubscriptionGroup.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupResponse.
// It customizes the JSON unmarshaling process for SubscriptionGroupResponse objects.
func (s *SubscriptionGroupResponse) UnmarshalJSON(input []byte) error {
	var temp subscriptionGroupResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	s.SubscriptionGroup = *temp.SubscriptionGroup
	return nil
}

// TODO
type subscriptionGroupResponse struct {
	SubscriptionGroup *SubscriptionGroup `json:"subscription_group"`
}

func (s *subscriptionGroupResponse) validate() error {
	var errs []string
	if s.SubscriptionGroup == nil {
		errs = append(errs, "required field `subscription_group` is missing for type `Subscription Group Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
