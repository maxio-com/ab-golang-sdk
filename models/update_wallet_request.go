// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UpdateWalletRequest represents a UpdateWalletRequest struct.
type UpdateWalletRequest struct {
    Wallet               Wallet2                `json:"wallet"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateWalletRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateWalletRequest) String() string {
    return fmt.Sprintf(
    	"UpdateWalletRequest[Wallet=%v, AdditionalProperties=%v]",
    	u.Wallet, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateWalletRequest.
// It customizes the JSON marshaling process for UpdateWalletRequest objects.
func (u UpdateWalletRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "wallet"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateWalletRequest object to a map representation for JSON marshaling.
func (u UpdateWalletRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["wallet"] = u.Wallet.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateWalletRequest.
// It customizes the JSON unmarshaling process for UpdateWalletRequest objects.
func (u *UpdateWalletRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateWalletRequest
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
    u.AdditionalProperties = additionalProperties
    
    u.Wallet = *temp.Wallet
    return nil
}

// tempUpdateWalletRequest is a temporary struct used for validating the fields of UpdateWalletRequest.
type tempUpdateWalletRequest  struct {
    Wallet *Wallet2 `json:"wallet"`
}

func (u *tempUpdateWalletRequest) validate() error {
    var errs []string
    if u.Wallet == nil {
        errs = append(errs, "required field `wallet` is missing for type `Update Wallet Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
