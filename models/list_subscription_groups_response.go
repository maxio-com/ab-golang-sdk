// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ListSubscriptionGroupsResponse represents a ListSubscriptionGroupsResponse struct.
type ListSubscriptionGroupsResponse struct {
    SubscriptionGroups   []ListSubscriptionGroupsItem `json:"subscription_groups,omitempty"`
    Meta                 *ListSubscriptionGroupsMeta  `json:"meta,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for ListSubscriptionGroupsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListSubscriptionGroupsResponse) String() string {
    return fmt.Sprintf(
    	"ListSubscriptionGroupsResponse[SubscriptionGroups=%v, Meta=%v, AdditionalProperties=%v]",
    	l.SubscriptionGroups, l.Meta, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionGroupsResponse.
// It customizes the JSON marshaling process for ListSubscriptionGroupsResponse objects.
func (l ListSubscriptionGroupsResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "subscription_groups", "meta"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionGroupsResponse object to a map representation for JSON marshaling.
func (l ListSubscriptionGroupsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.SubscriptionGroups != nil {
        structMap["subscription_groups"] = l.SubscriptionGroups
    }
    if l.Meta != nil {
        structMap["meta"] = l.Meta.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionGroupsResponse.
// It customizes the JSON unmarshaling process for ListSubscriptionGroupsResponse objects.
func (l *ListSubscriptionGroupsResponse) UnmarshalJSON(input []byte) error {
    var temp tempListSubscriptionGroupsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription_groups", "meta")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.SubscriptionGroups = temp.SubscriptionGroups
    l.Meta = temp.Meta
    return nil
}

// tempListSubscriptionGroupsResponse is a temporary struct used for validating the fields of ListSubscriptionGroupsResponse.
type tempListSubscriptionGroupsResponse  struct {
    SubscriptionGroups []ListSubscriptionGroupsItem `json:"subscription_groups,omitempty"`
    Meta               *ListSubscriptionGroupsMeta  `json:"meta,omitempty"`
}
