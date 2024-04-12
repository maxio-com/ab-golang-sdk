package models

import (
    "encoding/json"
)

// ReactivationBilling represents a ReactivationBilling struct.
// These values are only applicable to subscriptions using calendar billing
type ReactivationBilling struct {
    // You may choose how to handle the reactivation charge for that subscription: 1) `prorated` A prorated charge for the product price will be attempted for to complete the period 2) `immediate` A full-price charge for the product price will be attempted immediately 3) `delayed` A full-price charge for the product price will be attempted at the next renewal
    ReactivationCharge   *ReactivationCharge `json:"reactivation_charge,omitempty"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReactivationBilling.
// It customizes the JSON marshaling process for ReactivationBilling objects.
func (r ReactivationBilling) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReactivationBilling object to a map representation for JSON marshaling.
func (r ReactivationBilling) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.ReactivationCharge != nil {
        structMap["reactivation_charge"] = r.ReactivationCharge
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReactivationBilling.
// It customizes the JSON unmarshaling process for ReactivationBilling objects.
func (r *ReactivationBilling) UnmarshalJSON(input []byte) error {
    var temp reactivationBilling
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "reactivation_charge")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.ReactivationCharge = temp.ReactivationCharge
    return nil
}

// TODO
type reactivationBilling  struct {
    ReactivationCharge *ReactivationCharge `json:"reactivation_charge,omitempty"`
}
