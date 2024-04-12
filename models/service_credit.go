package models

import (
    "encoding/json"
)

// ServiceCredit represents a ServiceCredit struct.
type ServiceCredit struct {
    Id                   *int               `json:"id,omitempty"`
    // The amount in cents of the entry
    AmountInCents        *int64             `json:"amount_in_cents,omitempty"`
    // The new balance for the credit account
    EndingBalanceInCents *int64             `json:"ending_balance_in_cents,omitempty"`
    // The type of entry
    EntryType            *ServiceCreditType `json:"entry_type,omitempty"`
    // The memo attached to the entry
    Memo                 *string            `json:"memo,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ServiceCredit.
// It customizes the JSON marshaling process for ServiceCredit objects.
func (s ServiceCredit) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the ServiceCredit object to a map representation for JSON marshaling.
func (s ServiceCredit) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp serviceCredit
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "amount_in_cents", "ending_balance_in_cents", "entry_type", "memo")
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

// TODO
type serviceCredit  struct {
    Id                   *int               `json:"id,omitempty"`
    AmountInCents        *int64             `json:"amount_in_cents,omitempty"`
    EndingBalanceInCents *int64             `json:"ending_balance_in_cents,omitempty"`
    EntryType            *ServiceCreditType `json:"entry_type,omitempty"`
    Memo                 *string            `json:"memo,omitempty"`
}
