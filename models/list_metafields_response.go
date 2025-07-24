// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ListMetafieldsResponse represents a ListMetafieldsResponse struct.
type ListMetafieldsResponse struct {
    TotalCount           *int                   `json:"total_count,omitempty"`
    CurrentPage          *int                   `json:"current_page,omitempty"`
    TotalPages           *int                   `json:"total_pages,omitempty"`
    PerPage              *int                   `json:"per_page,omitempty"`
    Metafields           []Metafield            `json:"metafields,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListMetafieldsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListMetafieldsResponse) String() string {
    return fmt.Sprintf(
    	"ListMetafieldsResponse[TotalCount=%v, CurrentPage=%v, TotalPages=%v, PerPage=%v, Metafields=%v, AdditionalProperties=%v]",
    	l.TotalCount, l.CurrentPage, l.TotalPages, l.PerPage, l.Metafields, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListMetafieldsResponse.
// It customizes the JSON marshaling process for ListMetafieldsResponse objects.
func (l ListMetafieldsResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "total_count", "current_page", "total_pages", "per_page", "metafields"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListMetafieldsResponse object to a map representation for JSON marshaling.
func (l ListMetafieldsResponse) toMap() map[string]any {
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
    if l.Metafields != nil {
        structMap["metafields"] = l.Metafields
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListMetafieldsResponse.
// It customizes the JSON unmarshaling process for ListMetafieldsResponse objects.
func (l *ListMetafieldsResponse) UnmarshalJSON(input []byte) error {
    var temp tempListMetafieldsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "total_count", "current_page", "total_pages", "per_page", "metafields")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.TotalCount = temp.TotalCount
    l.CurrentPage = temp.CurrentPage
    l.TotalPages = temp.TotalPages
    l.PerPage = temp.PerPage
    l.Metafields = temp.Metafields
    return nil
}

// tempListMetafieldsResponse is a temporary struct used for validating the fields of ListMetafieldsResponse.
type tempListMetafieldsResponse  struct {
    TotalCount  *int        `json:"total_count,omitempty"`
    CurrentPage *int        `json:"current_page,omitempty"`
    TotalPages  *int        `json:"total_pages,omitempty"`
    PerPage     *int        `json:"per_page,omitempty"`
    Metafields  []Metafield `json:"metafields,omitempty"`
}
