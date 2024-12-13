/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionStateChange represents a SubscriptionStateChange struct.
type SubscriptionStateChange struct {
    PreviousSubscriptionState string                 `json:"previous_subscription_state"`
    NewSubscriptionState      string                 `json:"new_subscription_state"`
    AdditionalProperties      map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionStateChange.
// It customizes the JSON marshaling process for SubscriptionStateChange objects.
func (s SubscriptionStateChange) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "previous_subscription_state", "new_subscription_state"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionStateChange object to a map representation for JSON marshaling.
func (s SubscriptionStateChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["previous_subscription_state"] = s.PreviousSubscriptionState
    structMap["new_subscription_state"] = s.NewSubscriptionState
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionStateChange.
// It customizes the JSON unmarshaling process for SubscriptionStateChange objects.
func (s *SubscriptionStateChange) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionStateChange
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "previous_subscription_state", "new_subscription_state")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.PreviousSubscriptionState = *temp.PreviousSubscriptionState
    s.NewSubscriptionState = *temp.NewSubscriptionState
    return nil
}

// tempSubscriptionStateChange is a temporary struct used for validating the fields of SubscriptionStateChange.
type tempSubscriptionStateChange  struct {
    PreviousSubscriptionState *string `json:"previous_subscription_state"`
    NewSubscriptionState      *string `json:"new_subscription_state"`
}

func (s *tempSubscriptionStateChange) validate() error {
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
    return errors.New(strings.Join (errs, "\n"))
}
