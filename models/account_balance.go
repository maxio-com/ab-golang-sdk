package models

import (
    "encoding/json"
)

// AccountBalance represents a AccountBalance struct.
type AccountBalance struct {
    // The balance in cents.
    BalanceInCents *int64 `json:"balance_in_cents,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AccountBalance.
// It customizes the JSON marshaling process for AccountBalance objects.
func (a *AccountBalance) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AccountBalance object to a map representation for JSON marshaling.
func (a *AccountBalance) toMap() map[string]any {
    structMap := make(map[string]any)
    if a.BalanceInCents != nil {
        structMap["balance_in_cents"] = a.BalanceInCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountBalance.
// It customizes the JSON unmarshaling process for AccountBalance objects.
func (a *AccountBalance) UnmarshalJSON(input []byte) error {
    temp := &struct {
        BalanceInCents *int64 `json:"balance_in_cents,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    a.BalanceInCents = temp.BalanceInCents
    return nil
}
