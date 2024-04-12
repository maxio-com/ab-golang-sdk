package models

import (
    "encoding/json"
)

// CancelGroupedSubscriptionsRequest represents a CancelGroupedSubscriptionsRequest struct.
type CancelGroupedSubscriptionsRequest struct {
    ChargeUnbilledUsage  *bool          `json:"charge_unbilled_usage,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CancelGroupedSubscriptionsRequest.
// It customizes the JSON marshaling process for CancelGroupedSubscriptionsRequest objects.
func (c CancelGroupedSubscriptionsRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CancelGroupedSubscriptionsRequest object to a map representation for JSON marshaling.
func (c CancelGroupedSubscriptionsRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ChargeUnbilledUsage != nil {
        structMap["charge_unbilled_usage"] = c.ChargeUnbilledUsage
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CancelGroupedSubscriptionsRequest.
// It customizes the JSON unmarshaling process for CancelGroupedSubscriptionsRequest objects.
func (c *CancelGroupedSubscriptionsRequest) UnmarshalJSON(input []byte) error {
    var temp cancelGroupedSubscriptionsRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "charge_unbilled_usage")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.ChargeUnbilledUsage = temp.ChargeUnbilledUsage
    return nil
}

// TODO
type cancelGroupedSubscriptionsRequest  struct {
    ChargeUnbilledUsage *bool `json:"charge_unbilled_usage,omitempty"`
}
