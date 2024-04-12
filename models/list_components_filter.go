package models

import (
    "encoding/json"
)

// ListComponentsFilter represents a ListComponentsFilter struct.
type ListComponentsFilter struct {
    // Allows fetching components with matching id based on provided value. Use in query `filter[ids]=1,2,3`.
    Ids                  []int          `json:"ids,omitempty"`
    // Allows fetching components with matching use_site_exchange_rate based on provided value (refers to default price point). Use in query `filter[use_site_exchange_rate]=true`.
    UseSiteExchangeRate  *bool          `json:"use_site_exchange_rate,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListComponentsFilter.
// It customizes the JSON marshaling process for ListComponentsFilter objects.
func (l ListComponentsFilter) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListComponentsFilter object to a map representation for JSON marshaling.
func (l ListComponentsFilter) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Ids != nil {
        structMap["ids"] = l.Ids
    }
    if l.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = l.UseSiteExchangeRate
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListComponentsFilter.
// It customizes the JSON unmarshaling process for ListComponentsFilter objects.
func (l *ListComponentsFilter) UnmarshalJSON(input []byte) error {
    var temp listComponentsFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ids", "use_site_exchange_rate")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.Ids = temp.Ids
    l.UseSiteExchangeRate = temp.UseSiteExchangeRate
    return nil
}

// TODO
type listComponentsFilter  struct {
    Ids                 []int `json:"ids,omitempty"`
    UseSiteExchangeRate *bool `json:"use_site_exchange_rate,omitempty"`
}
