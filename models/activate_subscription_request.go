package models

import (
    "encoding/json"
)

// ActivateSubscriptionRequest represents a ActivateSubscriptionRequest struct.
type ActivateSubscriptionRequest struct {
    // You may choose how to handle the activation failure. `true` means do not change the subscription’s state and billing period. `false`  means to continue through with the activation and enter an end of life state. If this parameter is omitted or `null` is passed it will default to value set in the  site settings (default: `true`)
    RevertOnFailure Optional[bool] `json:"revert_on_failure"`
}

// MarshalJSON implements the json.Marshaler interface for ActivateSubscriptionRequest.
// It customizes the JSON marshaling process for ActivateSubscriptionRequest objects.
func (a *ActivateSubscriptionRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ActivateSubscriptionRequest object to a map representation for JSON marshaling.
func (a *ActivateSubscriptionRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    if a.RevertOnFailure.IsValueSet() {
        structMap["revert_on_failure"] = a.RevertOnFailure.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ActivateSubscriptionRequest.
// It customizes the JSON unmarshaling process for ActivateSubscriptionRequest objects.
func (a *ActivateSubscriptionRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        RevertOnFailure Optional[bool] `json:"revert_on_failure"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    a.RevertOnFailure = temp.RevertOnFailure
    return nil
}
