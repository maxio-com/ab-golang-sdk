package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionGroupSignupRequest represents a SubscriptionGroupSignupRequest struct.
type SubscriptionGroupSignupRequest struct {
    SubscriptionGroup    SubscriptionGroupSignup `json:"subscription_group"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupRequest.
// It customizes the JSON marshaling process for SubscriptionGroupSignupRequest objects.
func (s SubscriptionGroupSignupRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupRequest object to a map representation for JSON marshaling.
func (s SubscriptionGroupSignupRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["subscription_group"] = s.SubscriptionGroup.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupRequest.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupRequest objects.
func (s *SubscriptionGroupSignupRequest) UnmarshalJSON(input []byte) error {
    var temp subscriptionGroupSignupRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "subscription_group")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.SubscriptionGroup = *temp.SubscriptionGroup
    return nil
}

// TODO
type subscriptionGroupSignupRequest  struct {
    SubscriptionGroup *SubscriptionGroupSignup `json:"subscription_group"`
}

func (s *subscriptionGroupSignupRequest) validate() error {
    var errs []string
    if s.SubscriptionGroup == nil {
        errs = append(errs, "required field `subscription_group` is missing for type `Subscription Group Signup Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
