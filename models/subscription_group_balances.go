// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// SubscriptionGroupBalances represents a SubscriptionGroupBalances struct.
type SubscriptionGroupBalances struct {
    Prepayments          *AccountBalance        `json:"prepayments,omitempty"`
    ServiceCredits       *AccountBalance        `json:"service_credits,omitempty"`
    OpenInvoices         *AccountBalance        `json:"open_invoices,omitempty"`
    PendingDiscounts     *AccountBalance        `json:"pending_discounts,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupBalances,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupBalances) String() string {
    return fmt.Sprintf(
    	"SubscriptionGroupBalances[Prepayments=%v, ServiceCredits=%v, OpenInvoices=%v, PendingDiscounts=%v, AdditionalProperties=%v]",
    	s.Prepayments, s.ServiceCredits, s.OpenInvoices, s.PendingDiscounts, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupBalances.
// It customizes the JSON marshaling process for SubscriptionGroupBalances objects.
func (s SubscriptionGroupBalances) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "prepayments", "service_credits", "open_invoices", "pending_discounts"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupBalances object to a map representation for JSON marshaling.
func (s SubscriptionGroupBalances) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "prepayments", "service_credits", "open_invoices", "pending_discounts")
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
