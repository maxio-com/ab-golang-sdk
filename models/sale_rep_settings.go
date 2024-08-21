/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// SaleRepSettings represents a SaleRepSettings struct.
type SaleRepSettings struct {
    CustomerName         *string        `json:"customer_name,omitempty"`
    SubscriptionId       *int           `json:"subscription_id,omitempty"`
    SiteLink             *string        `json:"site_link,omitempty"`
    SiteName             *string        `json:"site_name,omitempty"`
    SubscriptionMrr      *string        `json:"subscription_mrr,omitempty"`
    SalesRepId           *int           `json:"sales_rep_id,omitempty"`
    SalesRepName         *string        `json:"sales_rep_name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SaleRepSettings.
// It customizes the JSON marshaling process for SaleRepSettings objects.
func (s SaleRepSettings) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SaleRepSettings object to a map representation for JSON marshaling.
func (s SaleRepSettings) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CustomerName != nil {
        structMap["customer_name"] = s.CustomerName
    }
    if s.SubscriptionId != nil {
        structMap["subscription_id"] = s.SubscriptionId
    }
    if s.SiteLink != nil {
        structMap["site_link"] = s.SiteLink
    }
    if s.SiteName != nil {
        structMap["site_name"] = s.SiteName
    }
    if s.SubscriptionMrr != nil {
        structMap["subscription_mrr"] = s.SubscriptionMrr
    }
    if s.SalesRepId != nil {
        structMap["sales_rep_id"] = s.SalesRepId
    }
    if s.SalesRepName != nil {
        structMap["sales_rep_name"] = s.SalesRepName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SaleRepSettings.
// It customizes the JSON unmarshaling process for SaleRepSettings objects.
func (s *SaleRepSettings) UnmarshalJSON(input []byte) error {
    var temp tempSaleRepSettings
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "customer_name", "subscription_id", "site_link", "site_name", "subscription_mrr", "sales_rep_id", "sales_rep_name")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.CustomerName = temp.CustomerName
    s.SubscriptionId = temp.SubscriptionId
    s.SiteLink = temp.SiteLink
    s.SiteName = temp.SiteName
    s.SubscriptionMrr = temp.SubscriptionMrr
    s.SalesRepId = temp.SalesRepId
    s.SalesRepName = temp.SalesRepName
    return nil
}

// tempSaleRepSettings is a temporary struct used for validating the fields of SaleRepSettings.
type tempSaleRepSettings  struct {
    CustomerName    *string `json:"customer_name,omitempty"`
    SubscriptionId  *int    `json:"subscription_id,omitempty"`
    SiteLink        *string `json:"site_link,omitempty"`
    SiteName        *string `json:"site_name,omitempty"`
    SubscriptionMrr *string `json:"subscription_mrr,omitempty"`
    SalesRepId      *int    `json:"sales_rep_id,omitempty"`
    SalesRepName    *string `json:"sales_rep_name,omitempty"`
}
