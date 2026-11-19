// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletProductResponse represents a WalletProductResponse struct.
// A Wallet Product is a Component created with no product family, so it returns the full Component field set under a `wallet_product` root instead of `component`.
type WalletProductResponse struct {
    WalletProduct        Component              `json:"wallet_product"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletProductResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletProductResponse) String() string {
    return fmt.Sprintf(
    	"WalletProductResponse[WalletProduct=%v, AdditionalProperties=%v]",
    	w.WalletProduct, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletProductResponse.
// It customizes the JSON marshaling process for WalletProductResponse objects.
func (w WalletProductResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "wallet_product"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletProductResponse object to a map representation for JSON marshaling.
func (w WalletProductResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["wallet_product"] = w.WalletProduct.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletProductResponse.
// It customizes the JSON unmarshaling process for WalletProductResponse objects.
func (w *WalletProductResponse) UnmarshalJSON(input []byte) error {
    var temp tempWalletProductResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallet_product")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.WalletProduct = *temp.WalletProduct
    return nil
}

// tempWalletProductResponse is a temporary struct used for validating the fields of WalletProductResponse.
type tempWalletProductResponse  struct {
    WalletProduct *Component `json:"wallet_product"`
}

func (w *tempWalletProductResponse) validate() error {
    var errs []string
    if w.WalletProduct == nil {
        errs = append(errs, "required field `wallet_product` is missing for type `Wallet Product Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
