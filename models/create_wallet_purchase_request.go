// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateWalletPurchaseRequest represents a CreateWalletPurchaseRequest struct.
// Purchases prepaid units for the wallet at a price point, charging the subscription's payment method.
type CreateWalletPurchaseRequest struct {
    Purchase             Purchase               `json:"purchase"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateWalletPurchaseRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateWalletPurchaseRequest) String() string {
    return fmt.Sprintf(
    	"CreateWalletPurchaseRequest[Purchase=%v, AdditionalProperties=%v]",
    	c.Purchase, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateWalletPurchaseRequest.
// It customizes the JSON marshaling process for CreateWalletPurchaseRequest objects.
func (c CreateWalletPurchaseRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "purchase"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateWalletPurchaseRequest object to a map representation for JSON marshaling.
func (c CreateWalletPurchaseRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["purchase"] = c.Purchase.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateWalletPurchaseRequest.
// It customizes the JSON unmarshaling process for CreateWalletPurchaseRequest objects.
func (c *CreateWalletPurchaseRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateWalletPurchaseRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "purchase")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Purchase = *temp.Purchase
    return nil
}

// tempCreateWalletPurchaseRequest is a temporary struct used for validating the fields of CreateWalletPurchaseRequest.
type tempCreateWalletPurchaseRequest  struct {
    Purchase *Purchase `json:"purchase"`
}

func (c *tempCreateWalletPurchaseRequest) validate() error {
    var errs []string
    if c.Purchase == nil {
        errs = append(errs, "required field `purchase` is missing for type `Create Wallet Purchase Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
