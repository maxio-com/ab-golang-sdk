// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletPurchaseResultResponse represents a WalletPurchaseResultResponse struct.
type WalletPurchaseResultResponse struct {
    WalletPurchases      []WalletPurchase       `json:"wallet_purchases"`
    // The wallet's new balance, as a decimal string.
    Balance              string                 `json:"balance"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletPurchaseResultResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletPurchaseResultResponse) String() string {
    return fmt.Sprintf(
    	"WalletPurchaseResultResponse[WalletPurchases=%v, Balance=%v, AdditionalProperties=%v]",
    	w.WalletPurchases, w.Balance, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletPurchaseResultResponse.
// It customizes the JSON marshaling process for WalletPurchaseResultResponse objects.
func (w WalletPurchaseResultResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "wallet_purchases", "balance"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletPurchaseResultResponse object to a map representation for JSON marshaling.
func (w WalletPurchaseResultResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["wallet_purchases"] = w.WalletPurchases
    structMap["balance"] = w.Balance
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletPurchaseResultResponse.
// It customizes the JSON unmarshaling process for WalletPurchaseResultResponse objects.
func (w *WalletPurchaseResultResponse) UnmarshalJSON(input []byte) error {
    var temp tempWalletPurchaseResultResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallet_purchases", "balance")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.WalletPurchases = *temp.WalletPurchases
    w.Balance = *temp.Balance
    return nil
}

// tempWalletPurchaseResultResponse is a temporary struct used for validating the fields of WalletPurchaseResultResponse.
type tempWalletPurchaseResultResponse  struct {
    WalletPurchases *[]WalletPurchase `json:"wallet_purchases"`
    Balance         *string           `json:"balance"`
}

func (w *tempWalletPurchaseResultResponse) validate() error {
    var errs []string
    if w.WalletPurchases == nil {
        errs = append(errs, "required field `wallet_purchases` is missing for type `Wallet Purchase Result Response`")
    }
    if w.Balance == nil {
        errs = append(errs, "required field `balance` is missing for type `Wallet Purchase Result Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
