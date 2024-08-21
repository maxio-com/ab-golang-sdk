/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ActivateSubscriptionRequest represents a ActivateSubscriptionRequest struct.
type ActivateSubscriptionRequest struct {
    // You may choose how to handle the activation failure. `true` means do not change the subscription’s state and billing period. `false`  means to continue through with the activation and enter an end of life state. If this parameter is omitted or `null` is passed it will default to value set in the  site settings (default: `true`)
    RevertOnFailure      Optional[bool] `json:"revert_on_failure"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ActivateSubscriptionRequest.
// It customizes the JSON marshaling process for ActivateSubscriptionRequest objects.
func (a ActivateSubscriptionRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ActivateSubscriptionRequest object to a map representation for JSON marshaling.
func (a ActivateSubscriptionRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.RevertOnFailure.IsValueSet() {
        if a.RevertOnFailure.Value() != nil {
            structMap["revert_on_failure"] = a.RevertOnFailure.Value()
        } else {
            structMap["revert_on_failure"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ActivateSubscriptionRequest.
// It customizes the JSON unmarshaling process for ActivateSubscriptionRequest objects.
func (a *ActivateSubscriptionRequest) UnmarshalJSON(input []byte) error {
    var temp tempActivateSubscriptionRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "revert_on_failure")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.RevertOnFailure = temp.RevertOnFailure
    return nil
}

// tempActivateSubscriptionRequest is a temporary struct used for validating the fields of ActivateSubscriptionRequest.
type tempActivateSubscriptionRequest  struct {
    RevertOnFailure Optional[bool] `json:"revert_on_failure"`
}
