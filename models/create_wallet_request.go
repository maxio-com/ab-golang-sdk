// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateWalletRequest represents a CreateWalletRequest struct.
type CreateWalletRequest struct {
    Wallet               Wallet1                `json:"wallet"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateWalletRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateWalletRequest) String() string {
    return fmt.Sprintf(
    	"CreateWalletRequest[Wallet=%v, AdditionalProperties=%v]",
    	c.Wallet, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateWalletRequest.
// It customizes the JSON marshaling process for CreateWalletRequest objects.
func (c CreateWalletRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "wallet"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateWalletRequest object to a map representation for JSON marshaling.
func (c CreateWalletRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["wallet"] = c.Wallet.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateWalletRequest.
// It customizes the JSON unmarshaling process for CreateWalletRequest objects.
func (c *CreateWalletRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateWalletRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallet")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Wallet = *temp.Wallet
    return nil
}

// tempCreateWalletRequest is a temporary struct used for validating the fields of CreateWalletRequest.
type tempCreateWalletRequest  struct {
    Wallet *Wallet1 `json:"wallet"`
}

func (c *tempCreateWalletRequest) validate() error {
    var errs []string
    if c.Wallet == nil {
        errs = append(errs, "required field `wallet` is missing for type `Create Wallet Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
