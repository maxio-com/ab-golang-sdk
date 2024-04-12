package models

import (
    "encoding/json"
)

// NestedSubscriptionGroup represents a NestedSubscriptionGroup struct.
type NestedSubscriptionGroup struct {
    // The UID for the group
    Uid                   *string        `json:"uid,omitempty"`
    // Whether the group is configured to rely on a primary subscription for billing. At this time, it will always be 1.
    Scheme                *int           `json:"scheme,omitempty"`
    // The subscription ID of the primary within the group. Applicable to scheme 1.
    PrimarySubscriptionId *int           `json:"primary_subscription_id,omitempty"`
    // A boolean indicating whether the subscription is the primary in the group. Applicable to scheme 1.
    Primary               *bool          `json:"primary,omitempty"`
    AdditionalProperties  map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NestedSubscriptionGroup.
// It customizes the JSON marshaling process for NestedSubscriptionGroup objects.
func (n NestedSubscriptionGroup) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NestedSubscriptionGroup object to a map representation for JSON marshaling.
func (n NestedSubscriptionGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
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
    var temp nestedSubscriptionGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "scheme", "primary_subscription_id", "primary")
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

// TODO
type nestedSubscriptionGroup  struct {
    Uid                   *string `json:"uid,omitempty"`
    Scheme                *int    `json:"scheme,omitempty"`
    PrimarySubscriptionId *int    `json:"primary_subscription_id,omitempty"`
    Primary               *bool   `json:"primary,omitempty"`
}
