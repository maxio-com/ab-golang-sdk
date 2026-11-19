// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateWalletAdjustmentRequest represents a CreateWalletAdjustmentRequest struct.
// A signed correction to a specific grant's balance, or to an overdraft entry's usage, intended for support and ops use.
type CreateWalletAdjustmentRequest struct {
    Adjustment           Adjustment             `json:"adjustment"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateWalletAdjustmentRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateWalletAdjustmentRequest) String() string {
    return fmt.Sprintf(
    	"CreateWalletAdjustmentRequest[Adjustment=%v, AdditionalProperties=%v]",
    	c.Adjustment, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateWalletAdjustmentRequest.
// It customizes the JSON marshaling process for CreateWalletAdjustmentRequest objects.
func (c CreateWalletAdjustmentRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "adjustment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateWalletAdjustmentRequest object to a map representation for JSON marshaling.
func (c CreateWalletAdjustmentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["adjustment"] = c.Adjustment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateWalletAdjustmentRequest.
// It customizes the JSON unmarshaling process for CreateWalletAdjustmentRequest objects.
func (c *CreateWalletAdjustmentRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateWalletAdjustmentRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "adjustment")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Adjustment = *temp.Adjustment
    return nil
}

// tempCreateWalletAdjustmentRequest is a temporary struct used for validating the fields of CreateWalletAdjustmentRequest.
type tempCreateWalletAdjustmentRequest  struct {
    Adjustment *Adjustment `json:"adjustment"`
}

func (c *tempCreateWalletAdjustmentRequest) validate() error {
    var errs []string
    if c.Adjustment == nil {
        errs = append(errs, "required field `adjustment` is missing for type `Create Wallet Adjustment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
