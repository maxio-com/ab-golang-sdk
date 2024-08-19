/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// SaleRepSubscription represents a SaleRepSubscription struct.
type SaleRepSubscription struct {
    Id                   *int             `json:"id,omitempty"`
    SiteName             *string          `json:"site_name,omitempty"`
    SubscriptionUrl      *string          `json:"subscription_url,omitempty"`
    CustomerName         *string          `json:"customer_name,omitempty"`
    CreatedAt            *string          `json:"created_at,omitempty"`
    Mrr                  *string          `json:"mrr,omitempty"`
    Usage                *string          `json:"usage,omitempty"`
    Recurring            *string          `json:"recurring,omitempty"`
    LastPayment          *string          `json:"last_payment,omitempty"`
    ChurnDate            Optional[string] `json:"churn_date"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SaleRepSubscription.
// It customizes the JSON marshaling process for SaleRepSubscription objects.
func (s SaleRepSubscription) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SaleRepSubscription object to a map representation for JSON marshaling.
func (s SaleRepSubscription) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.SiteName != nil {
        structMap["site_name"] = s.SiteName
    }
    if s.SubscriptionUrl != nil {
        structMap["subscription_url"] = s.SubscriptionUrl
    }
    if s.CustomerName != nil {
        structMap["customer_name"] = s.CustomerName
    }
    if s.CreatedAt != nil {
        structMap["created_at"] = s.CreatedAt
    }
    if s.Mrr != nil {
        structMap["mrr"] = s.Mrr
    }
    if s.Usage != nil {
        structMap["usage"] = s.Usage
    }
    if s.Recurring != nil {
        structMap["recurring"] = s.Recurring
    }
    if s.LastPayment != nil {
        structMap["last_payment"] = s.LastPayment
    }
    if s.ChurnDate.IsValueSet() {
        if s.ChurnDate.Value() != nil {
            structMap["churn_date"] = s.ChurnDate.Value()
        } else {
            structMap["churn_date"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SaleRepSubscription.
// It customizes the JSON unmarshaling process for SaleRepSubscription objects.
func (s *SaleRepSubscription) UnmarshalJSON(input []byte) error {
    var temp tempSaleRepSubscription
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "site_name", "subscription_url", "customer_name", "created_at", "mrr", "usage", "recurring", "last_payment", "churn_date")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Id = temp.Id
    s.SiteName = temp.SiteName
    s.SubscriptionUrl = temp.SubscriptionUrl
    s.CustomerName = temp.CustomerName
    s.CreatedAt = temp.CreatedAt
    s.Mrr = temp.Mrr
    s.Usage = temp.Usage
    s.Recurring = temp.Recurring
    s.LastPayment = temp.LastPayment
    s.ChurnDate = temp.ChurnDate
    return nil
}

// tempSaleRepSubscription is a temporary struct used for validating the fields of SaleRepSubscription.
type tempSaleRepSubscription  struct {
    Id              *int             `json:"id,omitempty"`
    SiteName        *string          `json:"site_name,omitempty"`
    SubscriptionUrl *string          `json:"subscription_url,omitempty"`
    CustomerName    *string          `json:"customer_name,omitempty"`
    CreatedAt       *string          `json:"created_at,omitempty"`
    Mrr             *string          `json:"mrr,omitempty"`
    Usage           *string          `json:"usage,omitempty"`
    Recurring       *string          `json:"recurring,omitempty"`
    LastPayment     *string          `json:"last_payment,omitempty"`
    ChurnDate       Optional[string] `json:"churn_date"`
}
