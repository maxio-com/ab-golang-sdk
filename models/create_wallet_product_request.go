// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateWalletProductRequest represents a CreateWalletProductRequest struct.
// Does not accept `product_family_id`: a Wallet Product never belongs to a product family. The underlying endpoint accepts the param, but only to reject it; sending one returns `422` ("Product family cannot be set on a wallet product") rather than silently ignoring it.
type CreateWalletProductRequest struct {
    WalletProduct        WalletProduct          `json:"wallet_product"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateWalletProductRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateWalletProductRequest) String() string {
    return fmt.Sprintf(
    	"CreateWalletProductRequest[WalletProduct=%v, AdditionalProperties=%v]",
    	c.WalletProduct, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateWalletProductRequest.
// It customizes the JSON marshaling process for CreateWalletProductRequest objects.
func (c CreateWalletProductRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "wallet_product"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateWalletProductRequest object to a map representation for JSON marshaling.
func (c CreateWalletProductRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["wallet_product"] = c.WalletProduct.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateWalletProductRequest.
// It customizes the JSON unmarshaling process for CreateWalletProductRequest objects.
func (c *CreateWalletProductRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateWalletProductRequest
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
    c.AdditionalProperties = additionalProperties
    
    c.WalletProduct = *temp.WalletProduct
    return nil
}

// tempCreateWalletProductRequest is a temporary struct used for validating the fields of CreateWalletProductRequest.
type tempCreateWalletProductRequest  struct {
    WalletProduct *WalletProduct `json:"wallet_product"`
}

func (c *tempCreateWalletProductRequest) validate() error {
    var errs []string
    if c.WalletProduct == nil {
        errs = append(errs, "required field `wallet_product` is missing for type `Create Wallet Product Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
