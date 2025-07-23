// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ListSubscriptionComponentsForSiteFilter represents a ListSubscriptionComponentsForSiteFilter struct.
type ListSubscriptionComponentsForSiteFilter struct {
    // Allows fetching components allocation with matching currency based on provided values. Use in query `filter[currencies]=USD,EUR`.
    Currencies           []string               `json:"currencies,omitempty"`
    // Allows fetching components allocation with matching use_site_exchange_rate based on provided value. Use in query `filter[use_site_exchange_rate]=true`.
    UseSiteExchangeRate  *bool                  `json:"use_site_exchange_rate,omitempty"`
    // Nested filter used for List Subscription Components For Site Filter
    Subscription         *SubscriptionFilter    `json:"subscription,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListSubscriptionComponentsForSiteFilter,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListSubscriptionComponentsForSiteFilter) String() string {
    return fmt.Sprintf(
    	"ListSubscriptionComponentsForSiteFilter[Currencies=%v, UseSiteExchangeRate=%v, Subscription=%v, AdditionalProperties=%v]",
    	l.Currencies, l.UseSiteExchangeRate, l.Subscription, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionComponentsForSiteFilter.
// It customizes the JSON marshaling process for ListSubscriptionComponentsForSiteFilter objects.
func (l ListSubscriptionComponentsForSiteFilter) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "currencies", "use_site_exchange_rate", "subscription"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionComponentsForSiteFilter object to a map representation for JSON marshaling.
func (l ListSubscriptionComponentsForSiteFilter) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
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
    var temp tempListSubscriptionComponentsForSiteFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "currencies", "use_site_exchange_rate", "subscription")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Currencies = temp.Currencies
    l.UseSiteExchangeRate = temp.UseSiteExchangeRate
    l.Subscription = temp.Subscription
    return nil
}

// tempListSubscriptionComponentsForSiteFilter is a temporary struct used for validating the fields of ListSubscriptionComponentsForSiteFilter.
type tempListSubscriptionComponentsForSiteFilter  struct {
    Currencies          []string            `json:"currencies,omitempty"`
    UseSiteExchangeRate *bool               `json:"use_site_exchange_rate,omitempty"`
    Subscription        *SubscriptionFilter `json:"subscription,omitempty"`
}
