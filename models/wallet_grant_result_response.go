// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletGrantResultResponse represents a WalletGrantResultResponse struct.
// Returned by an adjustment: the ledger row written, plus the wallet's new balance.
type WalletGrantResultResponse struct {
    WalletTransactions   []WalletTransaction    `json:"wallet_transactions"`
    // The wallet's new balance, as a decimal string.
    Balance              string                 `json:"balance"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletGrantResultResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletGrantResultResponse) String() string {
    return fmt.Sprintf(
    	"WalletGrantResultResponse[WalletTransactions=%v, Balance=%v, AdditionalProperties=%v]",
    	w.WalletTransactions, w.Balance, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletGrantResultResponse.
// It customizes the JSON marshaling process for WalletGrantResultResponse objects.
func (w WalletGrantResultResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "wallet_transactions", "balance"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletGrantResultResponse object to a map representation for JSON marshaling.
func (w WalletGrantResultResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["wallet_transactions"] = w.WalletTransactions
    structMap["balance"] = w.Balance
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletGrantResultResponse.
// It customizes the JSON unmarshaling process for WalletGrantResultResponse objects.
func (w *WalletGrantResultResponse) UnmarshalJSON(input []byte) error {
    var temp tempWalletGrantResultResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallet_transactions", "balance")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.WalletTransactions = *temp.WalletTransactions
    w.Balance = *temp.Balance
    return nil
}

// tempWalletGrantResultResponse is a temporary struct used for validating the fields of WalletGrantResultResponse.
type tempWalletGrantResultResponse  struct {
    WalletTransactions *[]WalletTransaction `json:"wallet_transactions"`
    Balance            *string              `json:"balance"`
}

func (w *tempWalletGrantResultResponse) validate() error {
    var errs []string
    if w.WalletTransactions == nil {
        errs = append(errs, "required field `wallet_transactions` is missing for type `Wallet Grant Result Response`")
    }
    if w.Balance == nil {
        errs = append(errs, "required field `balance` is missing for type `Wallet Grant Result Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
