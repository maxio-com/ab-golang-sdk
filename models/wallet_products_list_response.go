// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletProductsListResponse represents a WalletProductsListResponse struct.
type WalletProductsListResponse struct {
    WalletProducts       []WalletProductListItem `json:"wallet_products"`
    // Total number of Wallet Products matching the filters, across all pages.
    TotalCount           int                     `json:"total_count"`
    // Not a simple count of archived products matching the filters. With `status=archived`, equals `total_count` (there's nothing else to report). With any other `status`, it's `0` whenever `total_count` is non-zero, and only when `total_count` is `0` does it become the number of archived Wallet Products matching the same `q` search term (ignoring `status`): an empty-state hint, not a running total.
    ArchivedCount        int                     `json:"archived_count"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for WalletProductsListResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletProductsListResponse) String() string {
    return fmt.Sprintf(
    	"WalletProductsListResponse[WalletProducts=%v, TotalCount=%v, ArchivedCount=%v, AdditionalProperties=%v]",
    	w.WalletProducts, w.TotalCount, w.ArchivedCount, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletProductsListResponse.
// It customizes the JSON marshaling process for WalletProductsListResponse objects.
func (w WalletProductsListResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "wallet_products", "total_count", "archived_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletProductsListResponse object to a map representation for JSON marshaling.
func (w WalletProductsListResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["wallet_products"] = w.WalletProducts
    structMap["total_count"] = w.TotalCount
    structMap["archived_count"] = w.ArchivedCount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletProductsListResponse.
// It customizes the JSON unmarshaling process for WalletProductsListResponse objects.
func (w *WalletProductsListResponse) UnmarshalJSON(input []byte) error {
    var temp tempWalletProductsListResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallet_products", "total_count", "archived_count")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.WalletProducts = *temp.WalletProducts
    w.TotalCount = *temp.TotalCount
    w.ArchivedCount = *temp.ArchivedCount
    return nil
}

// tempWalletProductsListResponse is a temporary struct used for validating the fields of WalletProductsListResponse.
type tempWalletProductsListResponse  struct {
    WalletProducts *[]WalletProductListItem `json:"wallet_products"`
    TotalCount     *int                     `json:"total_count"`
    ArchivedCount  *int                     `json:"archived_count"`
}

func (w *tempWalletProductsListResponse) validate() error {
    var errs []string
    if w.WalletProducts == nil {
        errs = append(errs, "required field `wallet_products` is missing for type `Wallet Products List Response`")
    }
    if w.TotalCount == nil {
        errs = append(errs, "required field `total_count` is missing for type `Wallet Products List Response`")
    }
    if w.ArchivedCount == nil {
        errs = append(errs, "required field `archived_count` is missing for type `Wallet Products List Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
