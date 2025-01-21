/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// ServiceCredit represents a ServiceCredit struct.
type ServiceCredit struct {
    Id                   *int                   `json:"id,omitempty"`
    // The amount in cents of the entry
    AmountInCents        *int64                 `json:"amount_in_cents,omitempty"`
    // The new balance for the credit account
    EndingBalanceInCents *int64                 `json:"ending_balance_in_cents,omitempty"`
    // The type of entry
    EntryType            *ServiceCreditType     `json:"entry_type,omitempty"`
    // The memo attached to the entry
    Memo                 *string                `json:"memo,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ServiceCredit,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServiceCredit) String() string {
    return fmt.Sprintf(
    	"ServiceCredit[Id=%v, AmountInCents=%v, EndingBalanceInCents=%v, EntryType=%v, Memo=%v, AdditionalProperties=%v]",
    	s.Id, s.AmountInCents, s.EndingBalanceInCents, s.EntryType, s.Memo, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServiceCredit.
// It customizes the JSON marshaling process for ServiceCredit objects.
func (s ServiceCredit) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "amount_in_cents", "ending_balance_in_cents", "entry_type", "memo"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ServiceCredit object to a map representation for JSON marshaling.
func (s ServiceCredit) toMap() map[string]any {
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceCredit.
// It customizes the JSON unmarshaling process for ServiceCredit objects.
func (s *ServiceCredit) UnmarshalJSON(input []byte) error {
    var temp tempServiceCredit
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "amount_in_cents", "ending_balance_in_cents", "entry_type", "memo")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Id = temp.Id
    s.AmountInCents = temp.AmountInCents
    s.EndingBalanceInCents = temp.EndingBalanceInCents
    s.EntryType = temp.EntryType
    s.Memo = temp.Memo
    return nil
}

// tempServiceCredit is a temporary struct used for validating the fields of ServiceCredit.
type tempServiceCredit  struct {
    Id                   *int               `json:"id,omitempty"`
    AmountInCents        *int64             `json:"amount_in_cents,omitempty"`
    EndingBalanceInCents *int64             `json:"ending_balance_in_cents,omitempty"`
    EntryType            *ServiceCreditType `json:"entry_type,omitempty"`
    Memo                 *string            `json:"memo,omitempty"`
}
