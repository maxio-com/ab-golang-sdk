// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UpdateWalletProductRequest represents a UpdateWalletProductRequest struct.
// Only `name`, `description`, `taxable`, `tax_code`, `accounting_code`, and `item_category` are mutable. `unit_name`, `handle`, `price_point`, and `price_points` are not accepted here: the request still returns `200`, but those fields are silently ignored and nothing about them changes. `handle` stays fixed because it's the Wallet Product's only stable identifier on a surface with no product family in the path; `unit_name` stays fixed because a wallet's behavior snapshots it into the wallet's own `unit` field and never re-derives it.
type UpdateWalletProductRequest struct {
    WalletProduct        WalletProduct1         `json:"wallet_product"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateWalletProductRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateWalletProductRequest) String() string {
    return fmt.Sprintf(
    	"UpdateWalletProductRequest[WalletProduct=%v, AdditionalProperties=%v]",
    	u.WalletProduct, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateWalletProductRequest.
// It customizes the JSON marshaling process for UpdateWalletProductRequest objects.
func (u UpdateWalletProductRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "wallet_product"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateWalletProductRequest object to a map representation for JSON marshaling.
func (u UpdateWalletProductRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["wallet_product"] = u.WalletProduct.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateWalletProductRequest.
// It customizes the JSON unmarshaling process for UpdateWalletProductRequest objects.
func (u *UpdateWalletProductRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateWalletProductRequest
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
    u.AdditionalProperties = additionalProperties
    
    u.WalletProduct = *temp.WalletProduct
    return nil
}

// tempUpdateWalletProductRequest is a temporary struct used for validating the fields of UpdateWalletProductRequest.
type tempUpdateWalletProductRequest  struct {
    WalletProduct *WalletProduct1 `json:"wallet_product"`
}

func (u *tempUpdateWalletProductRequest) validate() error {
    var errs []string
    if u.WalletProduct == nil {
        errs = append(errs, "required field `wallet_product` is missing for type `Update Wallet Product Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
