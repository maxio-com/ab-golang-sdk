/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SubscriptionGroupSingleError represents a SubscriptionGroupSingleError struct.
type SubscriptionGroupSingleError struct {
    SubscriptionGroup    string                 `json:"subscription_group"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupSingleError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupSingleError) String() string {
    return fmt.Sprintf(
    	"SubscriptionGroupSingleError[SubscriptionGroup=%v, AdditionalProperties=%v]",
    	s.SubscriptionGroup, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSingleError.
// It customizes the JSON marshaling process for SubscriptionGroupSingleError objects.
func (s SubscriptionGroupSingleError) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "subscription_group"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSingleError object to a map representation for JSON marshaling.
func (s SubscriptionGroupSingleError) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["subscription_group"] = s.SubscriptionGroup
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSingleError.
// It customizes the JSON unmarshaling process for SubscriptionGroupSingleError objects.
func (s *SubscriptionGroupSingleError) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroupSingleError
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

// tempSubscriptionGroupSingleError is a temporary struct used for validating the fields of SubscriptionGroupSingleError.
type tempSubscriptionGroupSingleError  struct {
    SubscriptionGroup *string `json:"subscription_group"`
}

func (s *tempSubscriptionGroupSingleError) validate() error {
    var errs []string
    if s.SubscriptionGroup == nil {
        errs = append(errs, "required field `subscription_group` is missing for type `Subscription Group Single Error`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
