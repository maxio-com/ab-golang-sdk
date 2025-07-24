// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ListPublicKeysMeta represents a ListPublicKeysMeta struct.
type ListPublicKeysMeta struct {
    TotalCount           *int                   `json:"total_count,omitempty"`
    CurrentPage          *int                   `json:"current_page,omitempty"`
    TotalPages           *int                   `json:"total_pages,omitempty"`
    PerPage              *int                   `json:"per_page,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListPublicKeysMeta,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListPublicKeysMeta) String() string {
    return fmt.Sprintf(
    	"ListPublicKeysMeta[TotalCount=%v, CurrentPage=%v, TotalPages=%v, PerPage=%v, AdditionalProperties=%v]",
    	l.TotalCount, l.CurrentPage, l.TotalPages, l.PerPage, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListPublicKeysMeta.
// It customizes the JSON marshaling process for ListPublicKeysMeta objects.
func (l ListPublicKeysMeta) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "total_count", "current_page", "total_pages", "per_page"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListPublicKeysMeta object to a map representation for JSON marshaling.
func (l ListPublicKeysMeta) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.TotalCount != nil {
        structMap["total_count"] = l.TotalCount
    }
    if l.CurrentPage != nil {
        structMap["current_page"] = l.CurrentPage
    }
    if l.TotalPages != nil {
        structMap["total_pages"] = l.TotalPages
    }
    if l.PerPage != nil {
        structMap["per_page"] = l.PerPage
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListPublicKeysMeta.
// It customizes the JSON unmarshaling process for ListPublicKeysMeta objects.
func (l *ListPublicKeysMeta) UnmarshalJSON(input []byte) error {
    var temp tempListPublicKeysMeta
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "total_count", "current_page", "total_pages", "per_page")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.TotalCount = temp.TotalCount
    l.CurrentPage = temp.CurrentPage
    l.TotalPages = temp.TotalPages
    l.PerPage = temp.PerPage
    return nil
}

// tempListPublicKeysMeta is a temporary struct used for validating the fields of ListPublicKeysMeta.
type tempListPublicKeysMeta  struct {
    TotalCount  *int `json:"total_count,omitempty"`
    CurrentPage *int `json:"current_page,omitempty"`
    TotalPages  *int `json:"total_pages,omitempty"`
    PerPage     *int `json:"per_page,omitempty"`
}
