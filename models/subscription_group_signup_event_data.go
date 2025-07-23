// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SubscriptionGroupSignupEventData represents a SubscriptionGroupSignupEventData struct.
type SubscriptionGroupSignupEventData struct {
	SubscriptionGroup    SubscriptionGroupSignupFailureData `json:"subscription_group"`
	Customer             *Customer                          `json:"customer"`
	AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupSignupEventData,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupSignupEventData) String() string {
	return fmt.Sprintf(
		"SubscriptionGroupSignupEventData[SubscriptionGroup=%v, Customer=%v, AdditionalProperties=%v]",
		s.SubscriptionGroup, s.Customer, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupEventData.
// It customizes the JSON marshaling process for SubscriptionGroupSignupEventData objects.
func (s SubscriptionGroupSignupEventData) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"subscription_group", "customer"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupEventData object to a map representation for JSON marshaling.
func (s SubscriptionGroupSignupEventData) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["subscription_group"] = s.SubscriptionGroup.toMap()
	if s.Customer != nil {
		structMap["customer"] = s.Customer.toMap()
	} else {
		structMap["customer"] = nil
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupEventData.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupEventData objects.
func (s *SubscriptionGroupSignupEventData) UnmarshalJSON(input []byte) error {
	var temp tempSubscriptionGroupSignupEventData
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription_group", "customer")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.SubscriptionGroup = *temp.SubscriptionGroup
	s.Customer = temp.Customer
	return nil
}

// tempSubscriptionGroupSignupEventData is a temporary struct used for validating the fields of SubscriptionGroupSignupEventData.
type tempSubscriptionGroupSignupEventData struct {
	SubscriptionGroup *SubscriptionGroupSignupFailureData `json:"subscription_group"`
	Customer          *Customer                           `json:"customer"`
}

func (s *tempSubscriptionGroupSignupEventData) validate() error {
	var errs []string
	if s.SubscriptionGroup == nil {
		errs = append(errs, "required field `subscription_group` is missing for type `Subscription Group Signup Event Data`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
