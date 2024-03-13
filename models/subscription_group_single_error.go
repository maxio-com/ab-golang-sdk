package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// SubscriptionGroupSingleError represents a SubscriptionGroupSingleError struct.
type SubscriptionGroupSingleError struct {
	SubscriptionGroup string `json:"subscription_group"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSingleError.
// It customizes the JSON marshaling process for SubscriptionGroupSingleError objects.
func (s *SubscriptionGroupSingleError) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSingleError object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSingleError) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscription_group"] = s.SubscriptionGroup
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSingleError.
// It customizes the JSON unmarshaling process for SubscriptionGroupSingleError objects.
func (s *SubscriptionGroupSingleError) UnmarshalJSON(input []byte) error {
	var temp subscriptionGroupSingleError
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
type subscriptionGroupSingleError struct {
	SubscriptionGroup *string `json:"subscription_group"`
}

func (s *subscriptionGroupSingleError) validate() error {
	var errs []string
	if s.SubscriptionGroup == nil {
		errs = append(errs, "required field `subscription_group` is missing for type `Subscription Group Single Error`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
