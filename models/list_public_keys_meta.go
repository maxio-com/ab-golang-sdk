package models

import (
    "encoding/json"
)

// ListPublicKeysMeta represents a ListPublicKeysMeta struct.
type ListPublicKeysMeta struct {
    TotalCount           *int           `json:"total_count,omitempty"`
    CurrentPage          *int           `json:"current_page,omitempty"`
    TotalPages           *int           `json:"total_pages,omitempty"`
    PerPage              *int           `json:"per_page,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListPublicKeysMeta.
// It customizes the JSON marshaling process for ListPublicKeysMeta objects.
func (l ListPublicKeysMeta) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListPublicKeysMeta object to a map representation for JSON marshaling.
func (l ListPublicKeysMeta) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
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
    var temp listPublicKeysMeta
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "total_count", "current_page", "total_pages", "per_page")
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

// TODO
type listPublicKeysMeta  struct {
    TotalCount  *int `json:"total_count,omitempty"`
    CurrentPage *int `json:"current_page,omitempty"`
    TotalPages  *int `json:"total_pages,omitempty"`
    PerPage     *int `json:"per_page,omitempty"`
}
