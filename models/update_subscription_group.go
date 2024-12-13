/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// UpdateSubscriptionGroup represents a UpdateSubscriptionGroup struct.
type UpdateSubscriptionGroup struct {
    MemberIds            []int                  `json:"member_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionGroup.
// It customizes the JSON marshaling process for UpdateSubscriptionGroup objects.
func (u UpdateSubscriptionGroup) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "member_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionGroup object to a map representation for JSON marshaling.
func (u UpdateSubscriptionGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.MemberIds != nil {
        structMap["member_ids"] = u.MemberIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionGroup.
// It customizes the JSON unmarshaling process for UpdateSubscriptionGroup objects.
func (u *UpdateSubscriptionGroup) UnmarshalJSON(input []byte) error {
    var temp tempUpdateSubscriptionGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "member_ids")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.MemberIds = temp.MemberIds
    return nil
}

// tempUpdateSubscriptionGroup is a temporary struct used for validating the fields of UpdateSubscriptionGroup.
type tempUpdateSubscriptionGroup  struct {
    MemberIds []int `json:"member_ids,omitempty"`
}
