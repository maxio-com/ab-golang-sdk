package models

import (
    "encoding/json"
)

// SiteStatistics represents a SiteStatistics struct.
type SiteStatistics struct {
    TotalSubscriptions         *int           `json:"total_subscriptions,omitempty"`
    SubscriptionsToday         *int           `json:"subscriptions_today,omitempty"`
    TotalRevenue               *string        `json:"total_revenue,omitempty"`
    RevenueToday               *string        `json:"revenue_today,omitempty"`
    RevenueThisMonth           *string        `json:"revenue_this_month,omitempty"`
    RevenueThisYear            *string        `json:"revenue_this_year,omitempty"`
    TotalCanceledSubscriptions *int           `json:"total_canceled_subscriptions,omitempty"`
    TotalActiveSubscriptions   *int           `json:"total_active_subscriptions,omitempty"`
    TotalPastDueSubscriptions  *int           `json:"total_past_due_subscriptions,omitempty"`
    TotalUnpaidSubscriptions   *int           `json:"total_unpaid_subscriptions,omitempty"`
    TotalDunningSubscriptions  *int           `json:"total_dunning_subscriptions,omitempty"`
    AdditionalProperties       map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteStatistics.
// It customizes the JSON marshaling process for SiteStatistics objects.
func (s SiteStatistics) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteStatistics object to a map representation for JSON marshaling.
func (s SiteStatistics) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.TotalSubscriptions != nil {
        structMap["total_subscriptions"] = s.TotalSubscriptions
    }
    if s.SubscriptionsToday != nil {
        structMap["subscriptions_today"] = s.SubscriptionsToday
    }
    if s.TotalRevenue != nil {
        structMap["total_revenue"] = s.TotalRevenue
    }
    if s.RevenueToday != nil {
        structMap["revenue_today"] = s.RevenueToday
    }
    if s.RevenueThisMonth != nil {
        structMap["revenue_this_month"] = s.RevenueThisMonth
    }
    if s.RevenueThisYear != nil {
        structMap["revenue_this_year"] = s.RevenueThisYear
    }
    if s.TotalCanceledSubscriptions != nil {
        structMap["total_canceled_subscriptions"] = s.TotalCanceledSubscriptions
    }
    if s.TotalActiveSubscriptions != nil {
        structMap["total_active_subscriptions"] = s.TotalActiveSubscriptions
    }
    if s.TotalPastDueSubscriptions != nil {
        structMap["total_past_due_subscriptions"] = s.TotalPastDueSubscriptions
    }
    if s.TotalUnpaidSubscriptions != nil {
        structMap["total_unpaid_subscriptions"] = s.TotalUnpaidSubscriptions
    }
    if s.TotalDunningSubscriptions != nil {
        structMap["total_dunning_subscriptions"] = s.TotalDunningSubscriptions
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteStatistics.
// It customizes the JSON unmarshaling process for SiteStatistics objects.
func (s *SiteStatistics) UnmarshalJSON(input []byte) error {
    var temp siteStatistics
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "total_subscriptions", "subscriptions_today", "total_revenue", "revenue_today", "revenue_this_month", "revenue_this_year", "total_canceled_subscriptions", "total_active_subscriptions", "total_past_due_subscriptions", "total_unpaid_subscriptions", "total_dunning_subscriptions")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.TotalSubscriptions = temp.TotalSubscriptions
    s.SubscriptionsToday = temp.SubscriptionsToday
    s.TotalRevenue = temp.TotalRevenue
    s.RevenueToday = temp.RevenueToday
    s.RevenueThisMonth = temp.RevenueThisMonth
    s.RevenueThisYear = temp.RevenueThisYear
    s.TotalCanceledSubscriptions = temp.TotalCanceledSubscriptions
    s.TotalActiveSubscriptions = temp.TotalActiveSubscriptions
    s.TotalPastDueSubscriptions = temp.TotalPastDueSubscriptions
    s.TotalUnpaidSubscriptions = temp.TotalUnpaidSubscriptions
    s.TotalDunningSubscriptions = temp.TotalDunningSubscriptions
    return nil
}

// siteStatistics is a temporary struct used for validating the fields of SiteStatistics.
type siteStatistics  struct {
    TotalSubscriptions         *int    `json:"total_subscriptions,omitempty"`
    SubscriptionsToday         *int    `json:"subscriptions_today,omitempty"`
    TotalRevenue               *string `json:"total_revenue,omitempty"`
    RevenueToday               *string `json:"revenue_today,omitempty"`
    RevenueThisMonth           *string `json:"revenue_this_month,omitempty"`
    RevenueThisYear            *string `json:"revenue_this_year,omitempty"`
    TotalCanceledSubscriptions *int    `json:"total_canceled_subscriptions,omitempty"`
    TotalActiveSubscriptions   *int    `json:"total_active_subscriptions,omitempty"`
    TotalPastDueSubscriptions  *int    `json:"total_past_due_subscriptions,omitempty"`
    TotalUnpaidSubscriptions   *int    `json:"total_unpaid_subscriptions,omitempty"`
    TotalDunningSubscriptions  *int    `json:"total_dunning_subscriptions,omitempty"`
}
