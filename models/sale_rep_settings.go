/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// SaleRepSettings represents a SaleRepSettings struct.
type SaleRepSettings struct {
    CustomerName         *string                `json:"customer_name,omitempty"`
    SubscriptionId       *int                   `json:"subscription_id,omitempty"`
    SiteLink             *string                `json:"site_link,omitempty"`
    SiteName             *string                `json:"site_name,omitempty"`
    SubscriptionMrr      *string                `json:"subscription_mrr,omitempty"`
    SalesRepId           *int                   `json:"sales_rep_id,omitempty"`
    SalesRepName         *string                `json:"sales_rep_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SaleRepSettings,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SaleRepSettings) String() string {
    return fmt.Sprintf(
    	"SaleRepSettings[CustomerName=%v, SubscriptionId=%v, SiteLink=%v, SiteName=%v, SubscriptionMrr=%v, SalesRepId=%v, SalesRepName=%v, AdditionalProperties=%v]",
    	s.CustomerName, s.SubscriptionId, s.SiteLink, s.SiteName, s.SubscriptionMrr, s.SalesRepId, s.SalesRepName, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SaleRepSettings.
// It customizes the JSON marshaling process for SaleRepSettings objects.
func (s SaleRepSettings) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "customer_name", "subscription_id", "site_link", "site_name", "subscription_mrr", "sales_rep_id", "sales_rep_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SaleRepSettings object to a map representation for JSON marshaling.
func (s SaleRepSettings) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "customer_name", "subscription_id", "site_link", "site_name", "subscription_mrr", "sales_rep_id", "sales_rep_name")
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
