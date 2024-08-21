/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// SubscriptionGroupBalances represents a SubscriptionGroupBalances struct.
type SubscriptionGroupBalances struct {
    Prepayments          *AccountBalance `json:"prepayments,omitempty"`
    ServiceCredits       *AccountBalance `json:"service_credits,omitempty"`
    OpenInvoices         *AccountBalance `json:"open_invoices,omitempty"`
    PendingDiscounts     *AccountBalance `json:"pending_discounts,omitempty"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupBalances.
// It customizes the JSON marshaling process for SubscriptionGroupBalances objects.
func (s SubscriptionGroupBalances) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupBalances object to a map representation for JSON marshaling.
func (s SubscriptionGroupBalances) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Prepayments != nil {
        structMap["prepayments"] = s.Prepayments.toMap()
    }
    if s.ServiceCredits != nil {
        structMap["service_credits"] = s.ServiceCredits.toMap()
    }
    if s.OpenInvoices != nil {
        structMap["open_invoices"] = s.OpenInvoices.toMap()
    }
    if s.PendingDiscounts != nil {
        structMap["pending_discounts"] = s.PendingDiscounts.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupBalances.
// It customizes the JSON unmarshaling process for SubscriptionGroupBalances objects.
func (s *SubscriptionGroupBalances) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroupBalances
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "prepayments", "service_credits", "open_invoices", "pending_discounts")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Prepayments = temp.Prepayments
    s.ServiceCredits = temp.ServiceCredits
    s.OpenInvoices = temp.OpenInvoices
    s.PendingDiscounts = temp.PendingDiscounts
    return nil
}

// tempSubscriptionGroupBalances is a temporary struct used for validating the fields of SubscriptionGroupBalances.
type tempSubscriptionGroupBalances  struct {
    Prepayments      *AccountBalance `json:"prepayments,omitempty"`
    ServiceCredits   *AccountBalance `json:"service_credits,omitempty"`
    OpenInvoices     *AccountBalance `json:"open_invoices,omitempty"`
    PendingDiscounts *AccountBalance `json:"pending_discounts,omitempty"`
}
