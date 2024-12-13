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

// UpdateSubscriptionGroupRequest represents a UpdateSubscriptionGroupRequest struct.
type UpdateSubscriptionGroupRequest struct {
    SubscriptionGroup    UpdateSubscriptionGroup `json:"subscription_group"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionGroupRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionGroupRequest objects.
func (u UpdateSubscriptionGroupRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "subscription_group"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionGroupRequest object to a map representation for JSON marshaling.
func (u UpdateSubscriptionGroupRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["subscription_group"] = u.SubscriptionGroup.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionGroupRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionGroupRequest objects.
func (u *UpdateSubscriptionGroupRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateSubscriptionGroupRequest
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
    u.AdditionalProperties = additionalProperties
    
    u.SubscriptionGroup = *temp.SubscriptionGroup
    return nil
}

// tempUpdateSubscriptionGroupRequest is a temporary struct used for validating the fields of UpdateSubscriptionGroupRequest.
type tempUpdateSubscriptionGroupRequest  struct {
    SubscriptionGroup *UpdateSubscriptionGroup `json:"subscription_group"`
}

func (u *tempUpdateSubscriptionGroupRequest) validate() error {
    var errs []string
    if u.SubscriptionGroup == nil {
        errs = append(errs, "required field `subscription_group` is missing for type `Update Subscription Group Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
