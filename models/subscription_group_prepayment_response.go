package models

import (
    "encoding/json"
)

// SubscriptionGroupPrepaymentResponse represents a SubscriptionGroupPrepaymentResponse struct.
type SubscriptionGroupPrepaymentResponse struct {
    Id                   *int               `json:"id,omitempty"`
    // The amount in cents of the entry.
    AmountInCents        *int64             `json:"amount_in_cents,omitempty"`
    // The ending balance in cents of the account.
    EndingBalanceInCents *int64             `json:"ending_balance_in_cents,omitempty"`
    // The type of entry
    EntryType            *ServiceCreditType `json:"entry_type,omitempty"`
    // A memo attached to the entry.
    Memo                 *string            `json:"memo,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupPrepaymentResponse.
// It customizes the JSON marshaling process for SubscriptionGroupPrepaymentResponse objects.
func (s *SubscriptionGroupPrepaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupPrepaymentResponse object to a map representation for JSON marshaling.
func (s *SubscriptionGroupPrepaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
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

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupPrepaymentResponse.
// It customizes the JSON unmarshaling process for SubscriptionGroupPrepaymentResponse objects.
func (s *SubscriptionGroupPrepaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                   *int               `json:"id,omitempty"`
        AmountInCents        *int64             `json:"amount_in_cents,omitempty"`
        EndingBalanceInCents *int64             `json:"ending_balance_in_cents,omitempty"`
        EntryType            *ServiceCreditType `json:"entry_type,omitempty"`
        Memo                 *string            `json:"memo,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Id = temp.Id
    s.AmountInCents = temp.AmountInCents
    s.EndingBalanceInCents = temp.EndingBalanceInCents
    s.EntryType = temp.EntryType
    s.Memo = temp.Memo
    return nil
}
