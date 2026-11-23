// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletPurchasePreviewResponse represents a WalletPurchasePreviewResponse struct.
type WalletPurchasePreviewResponse struct {
    WalletPurchasePreviews []WalletPurchasePreview `json:"wallet_purchase_previews"`
    AdditionalProperties   map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for WalletPurchasePreviewResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletPurchasePreviewResponse) String() string {
    return fmt.Sprintf(
    	"WalletPurchasePreviewResponse[WalletPurchasePreviews=%v, AdditionalProperties=%v]",
    	w.WalletPurchasePreviews, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletPurchasePreviewResponse.
// It customizes the JSON marshaling process for WalletPurchasePreviewResponse objects.
func (w WalletPurchasePreviewResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "wallet_purchase_previews"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletPurchasePreviewResponse object to a map representation for JSON marshaling.
func (w WalletPurchasePreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["wallet_purchase_previews"] = w.WalletPurchasePreviews
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletPurchasePreviewResponse.
// It customizes the JSON unmarshaling process for WalletPurchasePreviewResponse objects.
func (w *WalletPurchasePreviewResponse) UnmarshalJSON(input []byte) error {
    var temp tempWalletPurchasePreviewResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallet_purchase_previews")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.WalletPurchasePreviews = *temp.WalletPurchasePreviews
    return nil
}

// tempWalletPurchasePreviewResponse is a temporary struct used for validating the fields of WalletPurchasePreviewResponse.
type tempWalletPurchasePreviewResponse  struct {
    WalletPurchasePreviews *[]WalletPurchasePreview `json:"wallet_purchase_previews"`
}

func (w *tempWalletPurchasePreviewResponse) validate() error {
    var errs []string
    if w.WalletPurchasePreviews == nil {
        errs = append(errs, "required field `wallet_purchase_previews` is missing for type `Wallet Purchase Preview Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
