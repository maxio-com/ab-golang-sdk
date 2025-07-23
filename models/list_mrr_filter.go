// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// ListMrrFilter represents a ListMrrFilter struct.
type ListMrrFilter struct {
	// Submit ids in order to limit results. Use in query: `filter[subscription_ids]=1,2,3`.
	SubscriptionIds      []int                  `json:"subscription_ids,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListMrrFilter,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListMrrFilter) String() string {
	return fmt.Sprintf(
		"ListMrrFilter[SubscriptionIds=%v, AdditionalProperties=%v]",
		l.SubscriptionIds, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListMrrFilter.
// It customizes the JSON marshaling process for ListMrrFilter objects.
func (l ListMrrFilter) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(l.AdditionalProperties,
		"subscription_ids"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(l.toMap())
}

// toMap converts the ListMrrFilter object to a map representation for JSON marshaling.
func (l ListMrrFilter) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, l.AdditionalProperties)
	if l.SubscriptionIds != nil {
		structMap["subscription_ids"] = l.SubscriptionIds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListMrrFilter.
// It customizes the JSON unmarshaling process for ListMrrFilter objects.
func (l *ListMrrFilter) UnmarshalJSON(input []byte) error {
	var temp tempListMrrFilter
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription_ids")
	if err != nil {
		return err
	}
	l.AdditionalProperties = additionalProperties

	l.SubscriptionIds = temp.SubscriptionIds
	return nil
}

// tempListMrrFilter is a temporary struct used for validating the fields of ListMrrFilter.
type tempListMrrFilter struct {
	SubscriptionIds []int `json:"subscription_ids,omitempty"`
}
