/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// ListSaleRepItem represents a ListSaleRepItem struct.
type ListSaleRepItem struct {
    Id                   *int                      `json:"id,omitempty"`
    FullName             *string                   `json:"full_name,omitempty"`
    SubscriptionsCount   *int                      `json:"subscriptions_count,omitempty"`
    MrrData              map[string]SaleRepItemMrr `json:"mrr_data,omitempty"`
    TestMode             *bool                     `json:"test_mode,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for ListSaleRepItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListSaleRepItem) String() string {
    return fmt.Sprintf(
    	"ListSaleRepItem[Id=%v, FullName=%v, SubscriptionsCount=%v, MrrData=%v, TestMode=%v, AdditionalProperties=%v]",
    	l.Id, l.FullName, l.SubscriptionsCount, l.MrrData, l.TestMode, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListSaleRepItem.
// It customizes the JSON marshaling process for ListSaleRepItem objects.
func (l ListSaleRepItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "id", "full_name", "subscriptions_count", "mrr_data", "test_mode"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSaleRepItem object to a map representation for JSON marshaling.
func (l ListSaleRepItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Id != nil {
        structMap["id"] = l.Id
    }
    if l.FullName != nil {
        structMap["full_name"] = l.FullName
    }
    if l.SubscriptionsCount != nil {
        structMap["subscriptions_count"] = l.SubscriptionsCount
    }
    if l.MrrData != nil {
        structMap["mrr_data"] = l.MrrData
    }
    if l.TestMode != nil {
        structMap["test_mode"] = l.TestMode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSaleRepItem.
// It customizes the JSON unmarshaling process for ListSaleRepItem objects.
func (l *ListSaleRepItem) UnmarshalJSON(input []byte) error {
    var temp tempListSaleRepItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "full_name", "subscriptions_count", "mrr_data", "test_mode")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Id = temp.Id
    l.FullName = temp.FullName
    l.SubscriptionsCount = temp.SubscriptionsCount
    l.MrrData = temp.MrrData
    l.TestMode = temp.TestMode
    return nil
}

// tempListSaleRepItem is a temporary struct used for validating the fields of ListSaleRepItem.
type tempListSaleRepItem  struct {
    Id                 *int                      `json:"id,omitempty"`
    FullName           *string                   `json:"full_name,omitempty"`
    SubscriptionsCount *int                      `json:"subscriptions_count,omitempty"`
    MrrData            map[string]SaleRepItemMrr `json:"mrr_data,omitempty"`
    TestMode           *bool                     `json:"test_mode,omitempty"`
}
