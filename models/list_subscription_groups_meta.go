// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ListSubscriptionGroupsMeta represents a ListSubscriptionGroupsMeta struct.
type ListSubscriptionGroupsMeta struct {
    CurrentPage          *int                   `json:"current_page,omitempty"`
    TotalCount           *int                   `json:"total_count,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListSubscriptionGroupsMeta,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListSubscriptionGroupsMeta) String() string {
    return fmt.Sprintf(
    	"ListSubscriptionGroupsMeta[CurrentPage=%v, TotalCount=%v, AdditionalProperties=%v]",
    	l.CurrentPage, l.TotalCount, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionGroupsMeta.
// It customizes the JSON marshaling process for ListSubscriptionGroupsMeta objects.
func (l ListSubscriptionGroupsMeta) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "current_page", "total_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionGroupsMeta object to a map representation for JSON marshaling.
func (l ListSubscriptionGroupsMeta) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.CurrentPage != nil {
        structMap["current_page"] = l.CurrentPage
    }
    if l.TotalCount != nil {
        structMap["total_count"] = l.TotalCount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionGroupsMeta.
// It customizes the JSON unmarshaling process for ListSubscriptionGroupsMeta objects.
func (l *ListSubscriptionGroupsMeta) UnmarshalJSON(input []byte) error {
    var temp tempListSubscriptionGroupsMeta
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "current_page", "total_count")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.CurrentPage = temp.CurrentPage
    l.TotalCount = temp.TotalCount
    return nil
}

// tempListSubscriptionGroupsMeta is a temporary struct used for validating the fields of ListSubscriptionGroupsMeta.
type tempListSubscriptionGroupsMeta  struct {
    CurrentPage *int `json:"current_page,omitempty"`
    TotalCount  *int `json:"total_count,omitempty"`
}
