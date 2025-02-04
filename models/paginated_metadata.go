/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// PaginatedMetadata represents a PaginatedMetadata struct.
type PaginatedMetadata struct {
    TotalCount           *int                   `json:"total_count,omitempty"`
    CurrentPage          *int                   `json:"current_page,omitempty"`
    TotalPages           *int                   `json:"total_pages,omitempty"`
    PerPage              *int                   `json:"per_page,omitempty"`
    Metadata             []Metadata             `json:"metadata,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PaginatedMetadata,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PaginatedMetadata) String() string {
    return fmt.Sprintf(
    	"PaginatedMetadata[TotalCount=%v, CurrentPage=%v, TotalPages=%v, PerPage=%v, Metadata=%v, AdditionalProperties=%v]",
    	p.TotalCount, p.CurrentPage, p.TotalPages, p.PerPage, p.Metadata, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PaginatedMetadata.
// It customizes the JSON marshaling process for PaginatedMetadata objects.
func (p PaginatedMetadata) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "total_count", "current_page", "total_pages", "per_page", "metadata"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PaginatedMetadata object to a map representation for JSON marshaling.
func (p PaginatedMetadata) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.TotalCount != nil {
        structMap["total_count"] = p.TotalCount
    }
    if p.CurrentPage != nil {
        structMap["current_page"] = p.CurrentPage
    }
    if p.TotalPages != nil {
        structMap["total_pages"] = p.TotalPages
    }
    if p.PerPage != nil {
        structMap["per_page"] = p.PerPage
    }
    if p.Metadata != nil {
        structMap["metadata"] = p.Metadata
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaginatedMetadata.
// It customizes the JSON unmarshaling process for PaginatedMetadata objects.
func (p *PaginatedMetadata) UnmarshalJSON(input []byte) error {
    var temp tempPaginatedMetadata
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "total_count", "current_page", "total_pages", "per_page", "metadata")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.TotalCount = temp.TotalCount
    p.CurrentPage = temp.CurrentPage
    p.TotalPages = temp.TotalPages
    p.PerPage = temp.PerPage
    p.Metadata = temp.Metadata
    return nil
}

// tempPaginatedMetadata is a temporary struct used for validating the fields of PaginatedMetadata.
type tempPaginatedMetadata  struct {
    TotalCount  *int       `json:"total_count,omitempty"`
    CurrentPage *int       `json:"current_page,omitempty"`
    TotalPages  *int       `json:"total_pages,omitempty"`
    PerPage     *int       `json:"per_page,omitempty"`
    Metadata    []Metadata `json:"metadata,omitempty"`
}
