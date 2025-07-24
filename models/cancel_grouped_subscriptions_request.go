// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// CancelGroupedSubscriptionsRequest represents a CancelGroupedSubscriptionsRequest struct.
type CancelGroupedSubscriptionsRequest struct {
    ChargeUnbilledUsage  *bool                  `json:"charge_unbilled_usage,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CancelGroupedSubscriptionsRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CancelGroupedSubscriptionsRequest) String() string {
    return fmt.Sprintf(
    	"CancelGroupedSubscriptionsRequest[ChargeUnbilledUsage=%v, AdditionalProperties=%v]",
    	c.ChargeUnbilledUsage, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CancelGroupedSubscriptionsRequest.
// It customizes the JSON marshaling process for CancelGroupedSubscriptionsRequest objects.
func (c CancelGroupedSubscriptionsRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "charge_unbilled_usage"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CancelGroupedSubscriptionsRequest object to a map representation for JSON marshaling.
func (c CancelGroupedSubscriptionsRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ChargeUnbilledUsage != nil {
        structMap["charge_unbilled_usage"] = c.ChargeUnbilledUsage
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CancelGroupedSubscriptionsRequest.
// It customizes the JSON unmarshaling process for CancelGroupedSubscriptionsRequest objects.
func (c *CancelGroupedSubscriptionsRequest) UnmarshalJSON(input []byte) error {
    var temp tempCancelGroupedSubscriptionsRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "charge_unbilled_usage")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.ChargeUnbilledUsage = temp.ChargeUnbilledUsage
    return nil
}

// tempCancelGroupedSubscriptionsRequest is a temporary struct used for validating the fields of CancelGroupedSubscriptionsRequest.
type tempCancelGroupedSubscriptionsRequest  struct {
    ChargeUnbilledUsage *bool `json:"charge_unbilled_usage,omitempty"`
}
