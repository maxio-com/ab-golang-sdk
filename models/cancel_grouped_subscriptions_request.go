package models

import (
	"encoding/json"
)

// CancelGroupedSubscriptionsRequest represents a CancelGroupedSubscriptionsRequest struct.
type CancelGroupedSubscriptionsRequest struct {
	ChargeUnbilledUsage *bool `json:"charge_unbilled_usage,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CancelGroupedSubscriptionsRequest.
// It customizes the JSON marshaling process for CancelGroupedSubscriptionsRequest objects.
func (c *CancelGroupedSubscriptionsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CancelGroupedSubscriptionsRequest object to a map representation for JSON marshaling.
func (c *CancelGroupedSubscriptionsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.ChargeUnbilledUsage != nil {
		structMap["charge_unbilled_usage"] = c.ChargeUnbilledUsage
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CancelGroupedSubscriptionsRequest.
// It customizes the JSON unmarshaling process for CancelGroupedSubscriptionsRequest objects.
func (c *CancelGroupedSubscriptionsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ChargeUnbilledUsage *bool `json:"charge_unbilled_usage,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.ChargeUnbilledUsage = temp.ChargeUnbilledUsage
	return nil
}
