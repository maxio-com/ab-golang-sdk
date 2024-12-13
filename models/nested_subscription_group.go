/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// NestedSubscriptionGroup represents a NestedSubscriptionGroup struct.
type NestedSubscriptionGroup struct {
    // The UID for the group
    Uid                   *string                `json:"uid,omitempty"`
    // Whether the group is configured to rely on a primary subscription for billing. At this time, it will always be 1.
    Scheme                *int                   `json:"scheme,omitempty"`
    // The subscription ID of the primary within the group. Applicable to scheme 1.
    PrimarySubscriptionId *int                   `json:"primary_subscription_id,omitempty"`
    // A boolean indicating whether the subscription is the primary in the group. Applicable to scheme 1.
    Primary               *bool                  `json:"primary,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NestedSubscriptionGroup.
// It customizes the JSON marshaling process for NestedSubscriptionGroup objects.
func (n NestedSubscriptionGroup) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "uid", "scheme", "primary_subscription_id", "primary"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NestedSubscriptionGroup object to a map representation for JSON marshaling.
func (n NestedSubscriptionGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
    if n.Uid != nil {
        structMap["uid"] = n.Uid
    }
    if n.Scheme != nil {
        structMap["scheme"] = n.Scheme
    }
    if n.PrimarySubscriptionId != nil {
        structMap["primary_subscription_id"] = n.PrimarySubscriptionId
    }
    if n.Primary != nil {
        structMap["primary"] = n.Primary
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NestedSubscriptionGroup.
// It customizes the JSON unmarshaling process for NestedSubscriptionGroup objects.
func (n *NestedSubscriptionGroup) UnmarshalJSON(input []byte) error {
    var temp tempNestedSubscriptionGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "scheme", "primary_subscription_id", "primary")
    if err != nil {
    	return err
    }
    n.AdditionalProperties = additionalProperties
    
    n.Uid = temp.Uid
    n.Scheme = temp.Scheme
    n.PrimarySubscriptionId = temp.PrimarySubscriptionId
    n.Primary = temp.Primary
    return nil
}

// tempNestedSubscriptionGroup is a temporary struct used for validating the fields of NestedSubscriptionGroup.
type tempNestedSubscriptionGroup  struct {
    Uid                   *string `json:"uid,omitempty"`
    Scheme                *int    `json:"scheme,omitempty"`
    PrimarySubscriptionId *int    `json:"primary_subscription_id,omitempty"`
    Primary               *bool   `json:"primary,omitempty"`
}
