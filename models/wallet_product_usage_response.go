// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletProductUsageResponse represents a WalletProductUsageResponse struct.
type WalletProductUsageResponse struct {
    // Wallet templates billing through this product at its default price point.
    WalletTemplates      []WalletProductUsageTemplate `json:"wallet_templates"`
    // Wallet templates billing through this product at a price point other than the default.
    PricePointOverrides  []WalletProductUsageTemplate `json:"price_point_overrides"`
    ActiveWallets        ActiveWallets                `json:"active_wallets"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for WalletProductUsageResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletProductUsageResponse) String() string {
    return fmt.Sprintf(
    	"WalletProductUsageResponse[WalletTemplates=%v, PricePointOverrides=%v, ActiveWallets=%v, AdditionalProperties=%v]",
    	w.WalletTemplates, w.PricePointOverrides, w.ActiveWallets, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletProductUsageResponse.
// It customizes the JSON marshaling process for WalletProductUsageResponse objects.
func (w WalletProductUsageResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "wallet_templates", "price_point_overrides", "active_wallets"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletProductUsageResponse object to a map representation for JSON marshaling.
func (w WalletProductUsageResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["wallet_templates"] = w.WalletTemplates
    structMap["price_point_overrides"] = w.PricePointOverrides
    structMap["active_wallets"] = w.ActiveWallets.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletProductUsageResponse.
// It customizes the JSON unmarshaling process for WalletProductUsageResponse objects.
func (w *WalletProductUsageResponse) UnmarshalJSON(input []byte) error {
    var temp tempWalletProductUsageResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallet_templates", "price_point_overrides", "active_wallets")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.WalletTemplates = *temp.WalletTemplates
    w.PricePointOverrides = *temp.PricePointOverrides
    w.ActiveWallets = *temp.ActiveWallets
    return nil
}

// tempWalletProductUsageResponse is a temporary struct used for validating the fields of WalletProductUsageResponse.
type tempWalletProductUsageResponse  struct {
    WalletTemplates     *[]WalletProductUsageTemplate `json:"wallet_templates"`
    PricePointOverrides *[]WalletProductUsageTemplate `json:"price_point_overrides"`
    ActiveWallets       *ActiveWallets                `json:"active_wallets"`
}

func (w *tempWalletProductUsageResponse) validate() error {
    var errs []string
    if w.WalletTemplates == nil {
        errs = append(errs, "required field `wallet_templates` is missing for type `Wallet Product Usage Response`")
    }
    if w.PricePointOverrides == nil {
        errs = append(errs, "required field `price_point_overrides` is missing for type `Wallet Product Usage Response`")
    }
    if w.ActiveWallets == nil {
        errs = append(errs, "required field `active_wallets` is missing for type `Wallet Product Usage Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
