package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionGroupSignupFailure represents a SubscriptionGroupSignupFailure struct.
type SubscriptionGroupSignupFailure struct {
    SubscriptionGroup    SubscriptionGroupSignupFailureData `json:"subscription_group"`
    Customer             *string                            `json:"customer"`
    AdditionalProperties map[string]any                     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupFailure.
// It customizes the JSON marshaling process for SubscriptionGroupSignupFailure objects.
func (s SubscriptionGroupSignupFailure) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupFailure object to a map representation for JSON marshaling.
func (s SubscriptionGroupSignupFailure) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["subscription_group"] = s.SubscriptionGroup.toMap()
    if s.Customer != nil {
        structMap["customer"] = s.Customer
    } else {
        structMap["customer"] = nil
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupFailure.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupFailure objects.
func (s *SubscriptionGroupSignupFailure) UnmarshalJSON(input []byte) error {
    var temp subscriptionGroupSignupFailure
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "subscription_group", "customer")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.SubscriptionGroup = *temp.SubscriptionGroup
    s.Customer = temp.Customer
    return nil
}

// TODO
type subscriptionGroupSignupFailure  struct {
    SubscriptionGroup *SubscriptionGroupSignupFailureData `json:"subscription_group"`
    Customer          *string                             `json:"customer"`
}

func (s *subscriptionGroupSignupFailure) validate() error {
    var errs []string
    if s.SubscriptionGroup == nil {
        errs = append(errs, "required field `subscription_group` is missing for type `Subscription Group Signup Failure`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
