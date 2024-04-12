package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionStateChange represents a SubscriptionStateChange struct.
type SubscriptionStateChange struct {
    PreviousSubscriptionState string         `json:"previous_subscription_state"`
    NewSubscriptionState      string         `json:"new_subscription_state"`
    AdditionalProperties      map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionStateChange.
// It customizes the JSON marshaling process for SubscriptionStateChange objects.
func (s SubscriptionStateChange) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionStateChange object to a map representation for JSON marshaling.
func (s SubscriptionStateChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["previous_subscription_state"] = s.PreviousSubscriptionState
    structMap["new_subscription_state"] = s.NewSubscriptionState
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionStateChange.
// It customizes the JSON unmarshaling process for SubscriptionStateChange objects.
func (s *SubscriptionStateChange) UnmarshalJSON(input []byte) error {
    var temp subscriptionStateChange
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "previous_subscription_state", "new_subscription_state")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.PreviousSubscriptionState = *temp.PreviousSubscriptionState
    s.NewSubscriptionState = *temp.NewSubscriptionState
    return nil
}

// TODO
type subscriptionStateChange  struct {
    PreviousSubscriptionState *string `json:"previous_subscription_state"`
    NewSubscriptionState      *string `json:"new_subscription_state"`
}

func (s *subscriptionStateChange) validate() error {
    var errs []string
    if s.PreviousSubscriptionState == nil {
        errs = append(errs, "required field `previous_subscription_state` is missing for type `Subscription State Change`")
    }
    if s.NewSubscriptionState == nil {
        errs = append(errs, "required field `new_subscription_state` is missing for type `Subscription State Change`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
