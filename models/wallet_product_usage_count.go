// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletProductUsageCount represents a WalletProductUsageCount struct.
// A rollup of what's using this Wallet Product, for the catalog listing.
type WalletProductUsageCount struct {
    // Number of wallet templates billing through this product, whether or not any wallet has been created from them yet.
    WalletConfigurations int                    `json:"wallet_configurations"`
    // Number of active (non-archived) wallets currently denominated by this product.
    ActiveWallets        int                    `json:"active_wallets"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletProductUsageCount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletProductUsageCount) String() string {
    return fmt.Sprintf(
    	"WalletProductUsageCount[WalletConfigurations=%v, ActiveWallets=%v, AdditionalProperties=%v]",
    	w.WalletConfigurations, w.ActiveWallets, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletProductUsageCount.
// It customizes the JSON marshaling process for WalletProductUsageCount objects.
func (w WalletProductUsageCount) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "wallet_configurations", "active_wallets"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletProductUsageCount object to a map representation for JSON marshaling.
func (w WalletProductUsageCount) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["wallet_configurations"] = w.WalletConfigurations
    structMap["active_wallets"] = w.ActiveWallets
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletProductUsageCount.
// It customizes the JSON unmarshaling process for WalletProductUsageCount objects.
func (w *WalletProductUsageCount) UnmarshalJSON(input []byte) error {
    var temp tempWalletProductUsageCount
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallet_configurations", "active_wallets")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.WalletConfigurations = *temp.WalletConfigurations
    w.ActiveWallets = *temp.ActiveWallets
    return nil
}

// tempWalletProductUsageCount is a temporary struct used for validating the fields of WalletProductUsageCount.
type tempWalletProductUsageCount  struct {
    WalletConfigurations *int `json:"wallet_configurations"`
    ActiveWallets        *int `json:"active_wallets"`
}

func (w *tempWalletProductUsageCount) validate() error {
    var errs []string
    if w.WalletConfigurations == nil {
        errs = append(errs, "required field `wallet_configurations` is missing for type `Wallet Product Usage Count`")
    }
    if w.ActiveWallets == nil {
        errs = append(errs, "required field `active_wallets` is missing for type `Wallet Product Usage Count`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
