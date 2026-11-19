// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletTransactionsListResponse represents a WalletTransactionsListResponse struct.
type WalletTransactionsListResponse struct {
    WalletTransactions   []WalletTransaction    `json:"wallet_transactions"`
    Meta                 ListPublicKeysMeta     `json:"meta"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletTransactionsListResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletTransactionsListResponse) String() string {
    return fmt.Sprintf(
    	"WalletTransactionsListResponse[WalletTransactions=%v, Meta=%v, AdditionalProperties=%v]",
    	w.WalletTransactions, w.Meta, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletTransactionsListResponse.
// It customizes the JSON marshaling process for WalletTransactionsListResponse objects.
func (w WalletTransactionsListResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "wallet_transactions", "meta"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletTransactionsListResponse object to a map representation for JSON marshaling.
func (w WalletTransactionsListResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["wallet_transactions"] = w.WalletTransactions
    structMap["meta"] = w.Meta.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletTransactionsListResponse.
// It customizes the JSON unmarshaling process for WalletTransactionsListResponse objects.
func (w *WalletTransactionsListResponse) UnmarshalJSON(input []byte) error {
    var temp tempWalletTransactionsListResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallet_transactions", "meta")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.WalletTransactions = *temp.WalletTransactions
    w.Meta = *temp.Meta
    return nil
}

// tempWalletTransactionsListResponse is a temporary struct used for validating the fields of WalletTransactionsListResponse.
type tempWalletTransactionsListResponse  struct {
    WalletTransactions *[]WalletTransaction `json:"wallet_transactions"`
    Meta               *ListPublicKeysMeta  `json:"meta"`
}

func (w *tempWalletTransactionsListResponse) validate() error {
    var errs []string
    if w.WalletTransactions == nil {
        errs = append(errs, "required field `wallet_transactions` is missing for type `Wallet Transactions List Response`")
    }
    if w.Meta == nil {
        errs = append(errs, "required field `meta` is missing for type `Wallet Transactions List Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
