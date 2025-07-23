// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// ServiceCredit1 represents a ServiceCredit1 struct.
type ServiceCredit1 struct {
    Id                      *int                   `json:"id,omitempty"`
    // The amount in cents of the entry
    AmountInCents           *int64                 `json:"amount_in_cents,omitempty"`
    // The new balance for the credit account
    EndingBalanceInCents    *int64                 `json:"ending_balance_in_cents,omitempty"`
    // The type of entry
    EntryType               *ServiceCreditType     `json:"entry_type,omitempty"`
    // The memo attached to the entry
    Memo                    *string                `json:"memo,omitempty"`
    // The invoice uid associated with the entry. Only present for debit entries
    InvoiceUid              Optional[string]       `json:"invoice_uid"`
    // The remaining balance for the entry
    RemainingBalanceInCents *int64                 `json:"remaining_balance_in_cents,omitempty"`
    // The date and time the entry was created
    CreatedAt               *time.Time             `json:"created_at,omitempty"`
    AdditionalProperties    map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ServiceCredit1,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServiceCredit1) String() string {
    return fmt.Sprintf(
    	"ServiceCredit1[Id=%v, AmountInCents=%v, EndingBalanceInCents=%v, EntryType=%v, Memo=%v, InvoiceUid=%v, RemainingBalanceInCents=%v, CreatedAt=%v, AdditionalProperties=%v]",
    	s.Id, s.AmountInCents, s.EndingBalanceInCents, s.EntryType, s.Memo, s.InvoiceUid, s.RemainingBalanceInCents, s.CreatedAt, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServiceCredit1.
// It customizes the JSON marshaling process for ServiceCredit1 objects.
func (s ServiceCredit1) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "amount_in_cents", "ending_balance_in_cents", "entry_type", "memo", "invoice_uid", "remaining_balance_in_cents", "created_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ServiceCredit1 object to a map representation for JSON marshaling.
func (s ServiceCredit1) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.AmountInCents != nil {
        structMap["amount_in_cents"] = s.AmountInCents
    }
    if s.EndingBalanceInCents != nil {
        structMap["ending_balance_in_cents"] = s.EndingBalanceInCents
    }
    if s.EntryType != nil {
        structMap["entry_type"] = s.EntryType
    }
    if s.Memo != nil {
        structMap["memo"] = s.Memo
    }
    if s.InvoiceUid.IsValueSet() {
        if s.InvoiceUid.Value() != nil {
            structMap["invoice_uid"] = s.InvoiceUid.Value()
        } else {
            structMap["invoice_uid"] = nil
        }
    }
    if s.RemainingBalanceInCents != nil {
        structMap["remaining_balance_in_cents"] = s.RemainingBalanceInCents
    }
    if s.CreatedAt != nil {
        structMap["created_at"] = s.CreatedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceCredit1.
// It customizes the JSON unmarshaling process for ServiceCredit1 objects.
func (s *ServiceCredit1) UnmarshalJSON(input []byte) error {
    var temp tempServiceCredit1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "amount_in_cents", "ending_balance_in_cents", "entry_type", "memo", "invoice_uid", "remaining_balance_in_cents", "created_at")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Id = temp.Id
    s.AmountInCents = temp.AmountInCents
    s.EndingBalanceInCents = temp.EndingBalanceInCents
    s.EntryType = temp.EntryType
    s.Memo = temp.Memo
    s.InvoiceUid = temp.InvoiceUid
    s.RemainingBalanceInCents = temp.RemainingBalanceInCents
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        s.CreatedAt = &CreatedAtVal
    }
    return nil
}

// tempServiceCredit1 is a temporary struct used for validating the fields of ServiceCredit1.
type tempServiceCredit1  struct {
    Id                      *int               `json:"id,omitempty"`
    AmountInCents           *int64             `json:"amount_in_cents,omitempty"`
    EndingBalanceInCents    *int64             `json:"ending_balance_in_cents,omitempty"`
    EntryType               *ServiceCreditType `json:"entry_type,omitempty"`
    Memo                    *string            `json:"memo,omitempty"`
    InvoiceUid              Optional[string]   `json:"invoice_uid"`
    RemainingBalanceInCents *int64             `json:"remaining_balance_in_cents,omitempty"`
    CreatedAt               *string            `json:"created_at,omitempty"`
}
