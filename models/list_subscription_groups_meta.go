package models

import (
    "encoding/json"
)

// ListSubscriptionGroupsMeta represents a ListSubscriptionGroupsMeta struct.
type ListSubscriptionGroupsMeta struct {
    CurrentPage          *int           `json:"current_page,omitempty"`
    TotalCount           *int           `json:"total_count,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionGroupsMeta.
// It customizes the JSON marshaling process for ListSubscriptionGroupsMeta objects.
func (l ListSubscriptionGroupsMeta) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionGroupsMeta object to a map representation for JSON marshaling.
func (l ListSubscriptionGroupsMeta) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
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
    var temp listSubscriptionGroupsMeta
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "current_page", "total_count")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.CurrentPage = temp.CurrentPage
    l.TotalCount = temp.TotalCount
    return nil
}

// TODO
type listSubscriptionGroupsMeta  struct {
    CurrentPage *int `json:"current_page,omitempty"`
    TotalCount  *int `json:"total_count,omitempty"`
}
