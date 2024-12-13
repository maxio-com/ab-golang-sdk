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

// SubscriptionGroupMembersArrayError represents a SubscriptionGroupMembersArrayError struct.
type SubscriptionGroupMembersArrayError struct {
    Members              []string               `json:"members"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupMembersArrayError.
// It customizes the JSON marshaling process for SubscriptionGroupMembersArrayError objects.
func (s SubscriptionGroupMembersArrayError) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "members"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupMembersArrayError object to a map representation for JSON marshaling.
func (s SubscriptionGroupMembersArrayError) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["members"] = s.Members
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupMembersArrayError.
// It customizes the JSON unmarshaling process for SubscriptionGroupMembersArrayError objects.
func (s *SubscriptionGroupMembersArrayError) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroupMembersArrayError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "members")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Members = *temp.Members
    return nil
}

// tempSubscriptionGroupMembersArrayError is a temporary struct used for validating the fields of SubscriptionGroupMembersArrayError.
type tempSubscriptionGroupMembersArrayError  struct {
    Members *[]string `json:"members"`
}

func (s *tempSubscriptionGroupMembersArrayError) validate() error {
    var errs []string
    if s.Members == nil {
        errs = append(errs, "required field `members` is missing for type `Subscription Group Members Array Error`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
