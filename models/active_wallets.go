// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ActiveWallets represents a ActiveWallets struct.
type ActiveWallets struct {
    // Total count of active wallets drawing on this product.
    Total                int                         `json:"total"`
    // Up to 50 of the most recently created active wallets, newest first.
    Sample               []WalletProductActiveWallet `json:"sample"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for ActiveWallets,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ActiveWallets) String() string {
    return fmt.Sprintf(
    	"ActiveWallets[Total=%v, Sample=%v, AdditionalProperties=%v]",
    	a.Total, a.Sample, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ActiveWallets.
// It customizes the JSON marshaling process for ActiveWallets objects.
func (a ActiveWallets) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "total", "sample"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ActiveWallets object to a map representation for JSON marshaling.
func (a ActiveWallets) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["total"] = a.Total
    structMap["sample"] = a.Sample
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ActiveWallets.
// It customizes the JSON unmarshaling process for ActiveWallets objects.
func (a *ActiveWallets) UnmarshalJSON(input []byte) error {
    var temp tempActiveWallets
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "total", "sample")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Total = *temp.Total
    a.Sample = *temp.Sample
    return nil
}

// tempActiveWallets is a temporary struct used for validating the fields of ActiveWallets.
type tempActiveWallets  struct {
    Total  *int                         `json:"total"`
    Sample *[]WalletProductActiveWallet `json:"sample"`
}

func (a *tempActiveWallets) validate() error {
    var errs []string
    if a.Total == nil {
        errs = append(errs, "required field `total` is missing for type `ActiveWallets`")
    }
    if a.Sample == nil {
        errs = append(errs, "required field `sample` is missing for type `ActiveWallets`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
