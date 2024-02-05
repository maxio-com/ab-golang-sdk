package models

import (
    "encoding/json"
)

// BankAccountVerification represents a BankAccountVerification struct.
type BankAccountVerification struct {
    Deposit1InCents *int64 `json:"deposit_1_in_cents,omitempty"`
    Deposit2InCents *int64 `json:"deposit_2_in_cents,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BankAccountVerification.
// It customizes the JSON marshaling process for BankAccountVerification objects.
func (b *BankAccountVerification) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BankAccountVerification object to a map representation for JSON marshaling.
func (b *BankAccountVerification) toMap() map[string]any {
    structMap := make(map[string]any)
    if b.Deposit1InCents != nil {
        structMap["deposit_1_in_cents"] = b.Deposit1InCents
    }
    if b.Deposit2InCents != nil {
        structMap["deposit_2_in_cents"] = b.Deposit2InCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountVerification.
// It customizes the JSON unmarshaling process for BankAccountVerification objects.
func (b *BankAccountVerification) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Deposit1InCents *int64 `json:"deposit_1_in_cents,omitempty"`
        Deposit2InCents *int64 `json:"deposit_2_in_cents,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    b.Deposit1InCents = temp.Deposit1InCents
    b.Deposit2InCents = temp.Deposit2InCents
    return nil
}
