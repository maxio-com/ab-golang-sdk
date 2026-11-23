// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateWalletTransactionRequest represents a CreateWalletTransactionRequest struct.
type CreateWalletTransactionRequest struct {
    WalletTransaction    WalletTransaction1     `json:"wallet_transaction"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateWalletTransactionRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateWalletTransactionRequest) String() string {
    return fmt.Sprintf(
    	"CreateWalletTransactionRequest[WalletTransaction=%v, AdditionalProperties=%v]",
    	c.WalletTransaction, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateWalletTransactionRequest.
// It customizes the JSON marshaling process for CreateWalletTransactionRequest objects.
func (c CreateWalletTransactionRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "wallet_transaction"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateWalletTransactionRequest object to a map representation for JSON marshaling.
func (c CreateWalletTransactionRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["wallet_transaction"] = c.WalletTransaction.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateWalletTransactionRequest.
// It customizes the JSON unmarshaling process for CreateWalletTransactionRequest objects.
func (c *CreateWalletTransactionRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateWalletTransactionRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallet_transaction")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.WalletTransaction = *temp.WalletTransaction
    return nil
}

// tempCreateWalletTransactionRequest is a temporary struct used for validating the fields of CreateWalletTransactionRequest.
type tempCreateWalletTransactionRequest  struct {
    WalletTransaction *WalletTransaction1 `json:"wallet_transaction"`
}

func (c *tempCreateWalletTransactionRequest) validate() error {
    var errs []string
    if c.WalletTransaction == nil {
        errs = append(errs, "required field `wallet_transaction` is missing for type `Create Wallet Transaction Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
