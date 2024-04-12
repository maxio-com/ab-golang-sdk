package models

import (
    "encoding/json"
)

// ListSubscriptionComponentsForSiteFilter represents a ListSubscriptionComponentsForSiteFilter struct.
type ListSubscriptionComponentsForSiteFilter struct {
    // Allows fetching components allocation with matching currency based on provided values. Use in query `filter[currencies]=USD,EUR`.
    Currencies           []string            `json:"currencies,omitempty"`
    // Allows fetching components allocation with matching use_site_exchange_rate based on provided value. Use in query `filter[use_site_exchange_rate]=true`.
    UseSiteExchangeRate  *bool               `json:"use_site_exchange_rate,omitempty"`
    // Nested filter used for List Subscription Components For Site Filter
    Subscription         *SubscriptionFilter `json:"subscription,omitempty"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionComponentsForSiteFilter.
// It customizes the JSON marshaling process for ListSubscriptionComponentsForSiteFilter objects.
func (l ListSubscriptionComponentsForSiteFilter) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionComponentsForSiteFilter object to a map representation for JSON marshaling.
func (l ListSubscriptionComponentsForSiteFilter) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Currencies != nil {
        structMap["currencies"] = l.Currencies
    }
    if l.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = l.UseSiteExchangeRate
    }
    if l.Subscription != nil {
        structMap["subscription"] = l.Subscription.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionComponentsForSiteFilter.
// It customizes the JSON unmarshaling process for ListSubscriptionComponentsForSiteFilter objects.
func (l *ListSubscriptionComponentsForSiteFilter) UnmarshalJSON(input []byte) error {
    var temp listSubscriptionComponentsForSiteFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "currencies", "use_site_exchange_rate", "subscription")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.Currencies = temp.Currencies
    l.UseSiteExchangeRate = temp.UseSiteExchangeRate
    l.Subscription = temp.Subscription
    return nil
}

// TODO
type listSubscriptionComponentsForSiteFilter  struct {
    Currencies          []string            `json:"currencies,omitempty"`
    UseSiteExchangeRate *bool               `json:"use_site_exchange_rate,omitempty"`
    Subscription        *SubscriptionFilter `json:"subscription,omitempty"`
}
