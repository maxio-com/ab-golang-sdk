// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletResponse represents a WalletResponse struct.
type WalletResponse struct {
    // A balance, denominated in the unit of a Wallet Product, owned by a subscription. Grants and recurring funding schedules are embedded so a single read shows the full balance breakdown.
    Wallet               Wallet                 `json:"wallet"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletResponse) String() string {
    return fmt.Sprintf(
    	"WalletResponse[Wallet=%v, AdditionalProperties=%v]",
    	w.Wallet, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletResponse.
// It customizes the JSON marshaling process for WalletResponse objects.
func (w WalletResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "wallet"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletResponse object to a map representation for JSON marshaling.
func (w WalletResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["wallet"] = w.Wallet.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletResponse.
// It customizes the JSON unmarshaling process for WalletResponse objects.
func (w *WalletResponse) UnmarshalJSON(input []byte) error {
    var temp tempWalletResponse
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
    w.AdditionalProperties = additionalProperties
    
    w.Wallet = *temp.Wallet
    return nil
}

// tempWalletResponse is a temporary struct used for validating the fields of WalletResponse.
type tempWalletResponse  struct {
    Wallet *Wallet `json:"wallet"`
}

func (w *tempWalletResponse) validate() error {
    var errs []string
    if w.Wallet == nil {
        errs = append(errs, "required field `wallet` is missing for type `Wallet Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
