package models

import (
    "encoding/json"
)

// ListMrrFilter represents a ListMrrFilter struct.
type ListMrrFilter struct {
    // Submit ids in order to limit results. Use in query: `filter[subscription_ids]=1,2,3`.
    SubscriptionIds      []int          `json:"subscription_ids,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListMrrFilter.
// It customizes the JSON marshaling process for ListMrrFilter objects.
func (l ListMrrFilter) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListMrrFilter object to a map representation for JSON marshaling.
func (l ListMrrFilter) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.SubscriptionIds != nil {
        structMap["subscription_ids"] = l.SubscriptionIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListMrrFilter.
// It customizes the JSON unmarshaling process for ListMrrFilter objects.
func (l *ListMrrFilter) UnmarshalJSON(input []byte) error {
    var temp listMrrFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "subscription_ids")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.SubscriptionIds = temp.SubscriptionIds
    return nil
}

// listMrrFilter is a temporary struct used for validating the fields of ListMrrFilter.
type listMrrFilter  struct {
    SubscriptionIds []int `json:"subscription_ids,omitempty"`
}
