/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// AccountBalance represents a AccountBalance struct.
type AccountBalance struct {
    // The balance in cents.
    BalanceInCents           *int64                 `json:"balance_in_cents,omitempty"`
    // The automatic balance in cents.
    AutomaticBalanceInCents  Optional[int64]        `json:"automatic_balance_in_cents"`
    // The remittance balance in cents.
    RemittanceBalanceInCents Optional[int64]        `json:"remittance_balance_in_cents"`
    AdditionalProperties     map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountBalance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountBalance) String() string {
    return fmt.Sprintf(
    	"AccountBalance[BalanceInCents=%v, AutomaticBalanceInCents=%v, RemittanceBalanceInCents=%v, AdditionalProperties=%v]",
    	a.BalanceInCents, a.AutomaticBalanceInCents, a.RemittanceBalanceInCents, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountBalance.
// It customizes the JSON marshaling process for AccountBalance objects.
func (a AccountBalance) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "balance_in_cents", "automatic_balance_in_cents", "remittance_balance_in_cents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountBalance object to a map representation for JSON marshaling.
func (a AccountBalance) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
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
    var temp tempAccountBalance
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "balance_in_cents", "automatic_balance_in_cents", "remittance_balance_in_cents")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.BalanceInCents = temp.BalanceInCents
    a.AutomaticBalanceInCents = temp.AutomaticBalanceInCents
    a.RemittanceBalanceInCents = temp.RemittanceBalanceInCents
    return nil
}

// tempAccountBalance is a temporary struct used for validating the fields of AccountBalance.
type tempAccountBalance  struct {
    BalanceInCents           *int64          `json:"balance_in_cents,omitempty"`
    AutomaticBalanceInCents  Optional[int64] `json:"automatic_balance_in_cents"`
    RemittanceBalanceInCents Optional[int64] `json:"remittance_balance_in_cents"`
}
