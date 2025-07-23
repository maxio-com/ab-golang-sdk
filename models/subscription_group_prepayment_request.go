// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SubscriptionGroupPrepaymentRequest represents a SubscriptionGroupPrepaymentRequest struct.
type SubscriptionGroupPrepaymentRequest struct {
	Prepayment           SubscriptionGroupPrepayment `json:"prepayment"`
	AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupPrepaymentRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupPrepaymentRequest) String() string {
	return fmt.Sprintf(
		"SubscriptionGroupPrepaymentRequest[Prepayment=%v, AdditionalProperties=%v]",
		s.Prepayment, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupPrepaymentRequest.
// It customizes the JSON marshaling process for SubscriptionGroupPrepaymentRequest objects.
func (s SubscriptionGroupPrepaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"prepayment"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupPrepaymentRequest object to a map representation for JSON marshaling.
func (s SubscriptionGroupPrepaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["prepayment"] = s.Prepayment.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupPrepaymentRequest.
// It customizes the JSON unmarshaling process for SubscriptionGroupPrepaymentRequest objects.
func (s *SubscriptionGroupPrepaymentRequest) UnmarshalJSON(input []byte) error {
	var temp tempSubscriptionGroupPrepaymentRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "prepayment")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Prepayment = *temp.Prepayment
	return nil
}

// tempSubscriptionGroupPrepaymentRequest is a temporary struct used for validating the fields of SubscriptionGroupPrepaymentRequest.
type tempSubscriptionGroupPrepaymentRequest struct {
	Prepayment *SubscriptionGroupPrepayment `json:"prepayment"`
}

func (s *tempSubscriptionGroupPrepaymentRequest) validate() error {
	var errs []string
	if s.Prepayment == nil {
		errs = append(errs, "required field `prepayment` is missing for type `Subscription Group Prepayment Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
