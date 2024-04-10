package models

import (
    "encoding/json"
)

// AccountBalance represents a AccountBalance struct.
type AccountBalance struct {
    // The balance in cents.
    BalanceInCents           *int64          `json:"balance_in_cents,omitempty"`
    // The automatic balance in cents.
    AutomaticBalanceInCents  Optional[int64] `json:"automatic_balance_in_cents"`
    // The remittance balance in cents.
    RemittanceBalanceInCents Optional[int64] `json:"remittance_balance_in_cents"`
    AdditionalProperties     map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountBalance.
// It customizes the JSON marshaling process for AccountBalance objects.
func (a AccountBalance) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AccountBalance object to a map representation for JSON marshaling.
func (a AccountBalance) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.BalanceInCents != nil {
        structMap["balance_in_cents"] = a.BalanceInCents
    }
    if a.AutomaticBalanceInCents.IsValueSet() {
        if a.AutomaticBalanceInCents.Value() != nil {
            structMap["automatic_balance_in_cents"] = a.AutomaticBalanceInCents.Value()
        } else {
            structMap["automatic_balance_in_cents"] = nil
        }
    }
    if a.RemittanceBalanceInCents.IsValueSet() {
        if a.RemittanceBalanceInCents.Value() != nil {
            structMap["remittance_balance_in_cents"] = a.RemittanceBalanceInCents.Value()
        } else {
            structMap["remittance_balance_in_cents"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountBalance.
// It customizes the JSON unmarshaling process for AccountBalance objects.
func (a *AccountBalance) UnmarshalJSON(input []byte) error {
    var temp accountBalance
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "balance_in_cents", "automatic_balance_in_cents", "remittance_balance_in_cents")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.BalanceInCents = temp.BalanceInCents
    a.AutomaticBalanceInCents = temp.AutomaticBalanceInCents
    a.RemittanceBalanceInCents = temp.RemittanceBalanceInCents
    return nil
}

// TODO
type accountBalance  struct {
    BalanceInCents           *int64          `json:"balance_in_cents,omitempty"`
    AutomaticBalanceInCents  Optional[int64] `json:"automatic_balance_in_cents"`
    RemittanceBalanceInCents Optional[int64] `json:"remittance_balance_in_cents"`
}
