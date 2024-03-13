package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// SubscriptionGroupSignupSuccess represents a SubscriptionGroupSignupSuccess struct.
type SubscriptionGroupSignupSuccess struct {
	SubscriptionGroup SubscriptionGroupSignupSuccessData `json:"subscription_group"`
	Customer          Customer                           `json:"customer"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupSuccess.
// It customizes the JSON marshaling process for SubscriptionGroupSignupSuccess objects.
func (s *SubscriptionGroupSignupSuccess) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupSuccess object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSignupSuccess) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscription_group"] = s.SubscriptionGroup.toMap()
	structMap["customer"] = s.Customer.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupSuccess.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupSuccess objects.
func (s *SubscriptionGroupSignupSuccess) UnmarshalJSON(input []byte) error {
	var temp subscriptionGroupSignupSuccess
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	s.SubscriptionGroup = *temp.SubscriptionGroup
	s.Customer = *temp.Customer
	return nil
}

// TODO
type subscriptionGroupSignupSuccess struct {
	SubscriptionGroup *SubscriptionGroupSignupSuccessData `json:"subscription_group"`
	Customer          *Customer                           `json:"customer"`
}

func (s *subscriptionGroupSignupSuccess) validate() error {
	var errs []string
	if s.SubscriptionGroup == nil {
		errs = append(errs, "required field `subscription_group` is missing for type `Subscription Group Signup Success`")
	}
	if s.Customer == nil {
		errs = append(errs, "required field `customer` is missing for type `Subscription Group Signup Success`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
