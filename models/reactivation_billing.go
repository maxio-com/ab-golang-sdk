// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ReactivationBilling represents a ReactivationBilling struct.
// These values are only applicable to subscriptions using calendar billing
type ReactivationBilling struct {
    // You may choose how to handle the reactivation charge for that subscription: 1) `prorated` A prorated charge for the product price will be attempted for to complete the period 2) `immediate` A full-price charge for the product price will be attempted immediately 3) `delayed` A full-price charge for the product price will be attempted at the next renewal
    ReactivationCharge   *ReactivationCharge    `json:"reactivation_charge,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ReactivationBilling,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ReactivationBilling) String() string {
    return fmt.Sprintf(
    	"ReactivationBilling[ReactivationCharge=%v, AdditionalProperties=%v]",
    	r.ReactivationCharge, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ReactivationBilling.
// It customizes the JSON marshaling process for ReactivationBilling objects.
func (r ReactivationBilling) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "reactivation_charge"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ReactivationBilling object to a map representation for JSON marshaling.
func (r ReactivationBilling) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.ReactivationCharge != nil {
        structMap["reactivation_charge"] = r.ReactivationCharge
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReactivationBilling.
// It customizes the JSON unmarshaling process for ReactivationBilling objects.
func (r *ReactivationBilling) UnmarshalJSON(input []byte) error {
    var temp tempReactivationBilling
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "reactivation_charge")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.ReactivationCharge = temp.ReactivationCharge
    return nil
}

// tempReactivationBilling is a temporary struct used for validating the fields of ReactivationBilling.
type tempReactivationBilling  struct {
    ReactivationCharge *ReactivationCharge `json:"reactivation_charge,omitempty"`
}
