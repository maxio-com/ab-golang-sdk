package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionGroupMembersArrayError represents a SubscriptionGroupMembersArrayError struct.
type SubscriptionGroupMembersArrayError struct {
    Members              []string       `json:"members"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupMembersArrayError.
// It customizes the JSON marshaling process for SubscriptionGroupMembersArrayError objects.
func (s SubscriptionGroupMembersArrayError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupMembersArrayError object to a map representation for JSON marshaling.
func (s SubscriptionGroupMembersArrayError) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["members"] = s.Members
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupMembersArrayError.
// It customizes the JSON unmarshaling process for SubscriptionGroupMembersArrayError objects.
func (s *SubscriptionGroupMembersArrayError) UnmarshalJSON(input []byte) error {
    var temp subscriptionGroupMembersArrayError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "members")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Members = *temp.Members
    return nil
}

// TODO
type subscriptionGroupMembersArrayError  struct {
    Members *[]string `json:"members"`
}

func (s *subscriptionGroupMembersArrayError) validate() error {
    var errs []string
    if s.Members == nil {
        errs = append(errs, "required field `members` is missing for type `Subscription Group Members Array Error`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
