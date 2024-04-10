package models

import (
    "encoding/json"
)

// AccountBalances represents a AccountBalances struct.
type AccountBalances struct {
    // The balance, in cents, of the sum of the subscription's  open, payable invoices.
    OpenInvoices         *AccountBalance `json:"open_invoices,omitempty"`
    // The balance, in cents, of the sum of the subscription's  pending, payable invoices.
    PendingInvoices      *AccountBalance `json:"pending_invoices,omitempty"`
    // The balance, in cents, of the subscription's Pending Discount account.
    PendingDiscounts     *AccountBalance `json:"pending_discounts,omitempty"`
    // The balance, in cents, of the subscription's Service Credit account.
    ServiceCredits       *AccountBalance `json:"service_credits,omitempty"`
    // The balance, in cents, of the subscription's Prepayment account.
    Prepayments          *AccountBalance `json:"prepayments,omitempty"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountBalances.
// It customizes the JSON marshaling process for AccountBalances objects.
func (a AccountBalances) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AccountBalances object to a map representation for JSON marshaling.
func (a AccountBalances) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.OpenInvoices != nil {
        structMap["open_invoices"] = a.OpenInvoices.toMap()
    }
    if a.PendingInvoices != nil {
        structMap["pending_invoices"] = a.PendingInvoices.toMap()
    }
    if a.PendingDiscounts != nil {
        structMap["pending_discounts"] = a.PendingDiscounts.toMap()
    }
    if a.ServiceCredits != nil {
        structMap["service_credits"] = a.ServiceCredits.toMap()
    }
    if a.Prepayments != nil {
        structMap["prepayments"] = a.Prepayments.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountBalances.
// It customizes the JSON unmarshaling process for AccountBalances objects.
func (a *AccountBalances) UnmarshalJSON(input []byte) error {
    var temp accountBalances
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "open_invoices", "pending_invoices", "pending_discounts", "service_credits", "prepayments")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.OpenInvoices = temp.OpenInvoices
    a.PendingInvoices = temp.PendingInvoices
    a.PendingDiscounts = temp.PendingDiscounts
    a.ServiceCredits = temp.ServiceCredits
    a.Prepayments = temp.Prepayments
    return nil
}

// TODO
type accountBalances  struct {
    OpenInvoices     *AccountBalance `json:"open_invoices,omitempty"`
    PendingInvoices  *AccountBalance `json:"pending_invoices,omitempty"`
    PendingDiscounts *AccountBalance `json:"pending_discounts,omitempty"`
    ServiceCredits   *AccountBalance `json:"service_credits,omitempty"`
    Prepayments      *AccountBalance `json:"prepayments,omitempty"`
}
