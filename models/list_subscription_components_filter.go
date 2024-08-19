/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ListSubscriptionComponentsFilter represents a ListSubscriptionComponentsFilter struct.
type ListSubscriptionComponentsFilter struct {
    // Allows fetching components allocation with matching currency based on provided values. Use in query `filter[currencies]=EUR,USD`.
    Currencies           []string       `json:"currencies,omitempty"`
    // Allows fetching components allocation with matching use_site_exchange_rate based on provided value. Use in query `filter[use_site_exchange_rate]=true`.
    UseSiteExchangeRate  *bool          `json:"use_site_exchange_rate,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionComponentsFilter.
// It customizes the JSON marshaling process for ListSubscriptionComponentsFilter objects.
func (l ListSubscriptionComponentsFilter) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionComponentsFilter object to a map representation for JSON marshaling.
func (l ListSubscriptionComponentsFilter) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Currencies != nil {
        structMap["currencies"] = l.Currencies
    }
    if l.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = l.UseSiteExchangeRate
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionComponentsFilter.
// It customizes the JSON unmarshaling process for ListSubscriptionComponentsFilter objects.
func (l *ListSubscriptionComponentsFilter) UnmarshalJSON(input []byte) error {
    var temp tempListSubscriptionComponentsFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "currencies", "use_site_exchange_rate")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.Currencies = temp.Currencies
    l.UseSiteExchangeRate = temp.UseSiteExchangeRate
    return nil
}

// tempListSubscriptionComponentsFilter is a temporary struct used for validating the fields of ListSubscriptionComponentsFilter.
type tempListSubscriptionComponentsFilter  struct {
    Currencies          []string `json:"currencies,omitempty"`
    UseSiteExchangeRate *bool    `json:"use_site_exchange_rate,omitempty"`
}
