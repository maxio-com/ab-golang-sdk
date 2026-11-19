// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletTransactionResultResponse represents a WalletTransactionResultResponse struct.
// Returned by [Create Wallet Transaction]($e/Wallet%20Transactions/createWalletTransaction). `wallet_transactions` and `balance` are always present. `requested_amount`, `drained_amount`, and `rejected_amount` are present only when the request's `kind` was `debit`.
type WalletTransactionResultResponse struct {
    WalletTransactions   []WalletTransaction    `json:"wallet_transactions"`
    // `kind: debit` only. The amount that was requested, as a decimal string.
    RequestedAmount      *string                `json:"requested_amount,omitempty"`
    // `kind: debit` only. The amount actually drained from the wallet's grants, as a decimal string.
    DrainedAmount        *string                `json:"drained_amount,omitempty"`
    // `kind: debit` only, as a decimal string. Under `hard_stop_overflow: discard`, this is the requested amount minus what the wallet could cover (`"0.0"` on a full drain). Always `"0.0"` under `allow_negative`; nothing is rejected, and the balance goes negative instead.
    RejectedAmount       *string                `json:"rejected_amount,omitempty"`
    // The wallet's new balance, as a decimal string.
    Balance              string                 `json:"balance"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletTransactionResultResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletTransactionResultResponse) String() string {
    return fmt.Sprintf(
    	"WalletTransactionResultResponse[WalletTransactions=%v, RequestedAmount=%v, DrainedAmount=%v, RejectedAmount=%v, Balance=%v, AdditionalProperties=%v]",
    	w.WalletTransactions, w.RequestedAmount, w.DrainedAmount, w.RejectedAmount, w.Balance, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletTransactionResultResponse.
// It customizes the JSON marshaling process for WalletTransactionResultResponse objects.
func (w WalletTransactionResultResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "wallet_transactions", "requested_amount", "drained_amount", "rejected_amount", "balance"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletTransactionResultResponse object to a map representation for JSON marshaling.
func (w WalletTransactionResultResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["wallet_transactions"] = w.WalletTransactions
    if w.RequestedAmount != nil {
        structMap["requested_amount"] = w.RequestedAmount
    }
    if w.DrainedAmount != nil {
        structMap["drained_amount"] = w.DrainedAmount
    }
    if w.RejectedAmount != nil {
        structMap["rejected_amount"] = w.RejectedAmount
    }
    structMap["balance"] = w.Balance
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletTransactionResultResponse.
// It customizes the JSON unmarshaling process for WalletTransactionResultResponse objects.
func (w *WalletTransactionResultResponse) UnmarshalJSON(input []byte) error {
    var temp tempWalletTransactionResultResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallet_transactions", "requested_amount", "drained_amount", "rejected_amount", "balance")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.WalletTransactions = *temp.WalletTransactions
    w.RequestedAmount = temp.RequestedAmount
    w.DrainedAmount = temp.DrainedAmount
    w.RejectedAmount = temp.RejectedAmount
    w.Balance = *temp.Balance
    return nil
}

// tempWalletTransactionResultResponse is a temporary struct used for validating the fields of WalletTransactionResultResponse.
type tempWalletTransactionResultResponse  struct {
    WalletTransactions *[]WalletTransaction `json:"wallet_transactions"`
    RequestedAmount    *string              `json:"requested_amount,omitempty"`
    DrainedAmount      *string              `json:"drained_amount,omitempty"`
    RejectedAmount     *string              `json:"rejected_amount,omitempty"`
    Balance            *string              `json:"balance"`
}

func (w *tempWalletTransactionResultResponse) validate() error {
    var errs []string
    if w.WalletTransactions == nil {
        errs = append(errs, "required field `wallet_transactions` is missing for type `Wallet Transaction Result Response`")
    }
    if w.Balance == nil {
        errs = append(errs, "required field `balance` is missing for type `Wallet Transaction Result Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
