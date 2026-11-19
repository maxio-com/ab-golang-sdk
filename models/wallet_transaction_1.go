// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletTransaction1 represents a WalletTransaction1 struct.
type WalletTransaction1 struct {
    // Determines the direction of the movement. The amount itself is always sent positive.
    Kind                 Kind1                    `json:"kind"`
    // Always positive; `kind` determines the sign written to the ledger. Encoded as a decimal string with at most 8 fractional digits. Send fractional values as strings: a JSON number with a decimal point is rejected with `422`, because a double-precision float may already have lost precision before the request arrives.
    Amount               WalletTransaction1Amount `json:"amount"`
    Memo                 Optional[string]         `json:"memo"`
    // An expiration for the grant this creates. Only accepted with `kind: grant`; sending it with `kind: debit` is rejected. Omit it for an evergreen grant. Accepts either a bare calendar date (`YYYY-MM-DD`), resolved to the end of that day in the site's time zone, or a full timestamp. A calendar day that doesn't exist, such as `2027-02-30`, is rejected with `422` rather than rolled into the next month.
    ExpiresAt            Optional[string]         `json:"expires_at"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for WalletTransaction1,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletTransaction1) String() string {
    return fmt.Sprintf(
    	"WalletTransaction1[Kind=%v, Amount=%v, Memo=%v, ExpiresAt=%v, AdditionalProperties=%v]",
    	w.Kind, w.Amount, w.Memo, w.ExpiresAt, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletTransaction1.
// It customizes the JSON marshaling process for WalletTransaction1 objects.
func (w WalletTransaction1) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "kind", "amount", "memo", "expires_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletTransaction1 object to a map representation for JSON marshaling.
func (w WalletTransaction1) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["kind"] = w.Kind
    structMap["amount"] = w.Amount.toMap()
    if w.Memo.IsValueSet() {
        if w.Memo.Value() != nil {
            structMap["memo"] = w.Memo.Value()
        } else {
            structMap["memo"] = nil
        }
    }
    if w.ExpiresAt.IsValueSet() {
        if w.ExpiresAt.Value() != nil {
            structMap["expires_at"] = w.ExpiresAt.Value()
        } else {
            structMap["expires_at"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletTransaction1.
// It customizes the JSON unmarshaling process for WalletTransaction1 objects.
func (w *WalletTransaction1) UnmarshalJSON(input []byte) error {
    var temp tempWalletTransaction1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "kind", "amount", "memo", "expires_at")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Kind = *temp.Kind
    w.Amount = *temp.Amount
    w.Memo = temp.Memo
    w.ExpiresAt = temp.ExpiresAt
    return nil
}

// tempWalletTransaction1 is a temporary struct used for validating the fields of WalletTransaction1.
type tempWalletTransaction1  struct {
    Kind      *Kind1                    `json:"kind"`
    Amount    *WalletTransaction1Amount `json:"amount"`
    Memo      Optional[string]          `json:"memo"`
    ExpiresAt Optional[string]          `json:"expires_at"`
}

func (w *tempWalletTransaction1) validate() error {
    var errs []string
    if w.Kind == nil {
        errs = append(errs, "required field `kind` is missing for type `WalletTransaction1`")
    }
    if w.Amount == nil {
        errs = append(errs, "required field `amount` is missing for type `WalletTransaction1`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
