// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SubscriptionGroupSignupRequest represents a SubscriptionGroupSignupRequest struct.
type SubscriptionGroupSignupRequest struct {
    SubscriptionGroup    SubscriptionGroupSignup `json:"subscription_group"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupSignupRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupSignupRequest) String() string {
    return fmt.Sprintf(
    	"SubscriptionGroupSignupRequest[SubscriptionGroup=%v, AdditionalProperties=%v]",
    	s.SubscriptionGroup, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupRequest.
// It customizes the JSON marshaling process for SubscriptionGroupSignupRequest objects.
func (s SubscriptionGroupSignupRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "subscription_group"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupRequest object to a map representation for JSON marshaling.
func (s SubscriptionGroupSignupRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["subscription_group"] = s.SubscriptionGroup.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupRequest.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupRequest objects.
func (s *SubscriptionGroupSignupRequest) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroupSignupRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription_group")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.SubscriptionGroup = *temp.SubscriptionGroup
    return nil
}

// tempSubscriptionGroupSignupRequest is a temporary struct used for validating the fields of SubscriptionGroupSignupRequest.
type tempSubscriptionGroupSignupRequest  struct {
    SubscriptionGroup *SubscriptionGroupSignup `json:"subscription_group"`
}

func (s *tempSubscriptionGroupSignupRequest) validate() error {
    var errs []string
    if s.SubscriptionGroup == nil {
        errs = append(errs, "required field `subscription_group` is missing for type `Subscription Group Signup Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
