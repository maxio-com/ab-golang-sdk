package models

import (
    "encoding/json"
)

// Breakouts represents a Breakouts struct.
type Breakouts struct {
    PlanAmountInCents    *int64  `json:"plan_amount_in_cents,omitempty"`
    PlanAmountFormatted  *string `json:"plan_amount_formatted,omitempty"`
    UsageAmountInCents   *int64  `json:"usage_amount_in_cents,omitempty"`
    UsageAmountFormatted *string `json:"usage_amount_formatted,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Breakouts.
// It customizes the JSON marshaling process for Breakouts objects.
func (b *Breakouts) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the Breakouts object to a map representation for JSON marshaling.
func (b *Breakouts) toMap() map[string]any {
    structMap := make(map[string]any)
    if b.PlanAmountInCents != nil {
        structMap["plan_amount_in_cents"] = b.PlanAmountInCents
    }
    if b.PlanAmountFormatted != nil {
        structMap["plan_amount_formatted"] = b.PlanAmountFormatted
    }
    if b.UsageAmountInCents != nil {
        structMap["usage_amount_in_cents"] = b.UsageAmountInCents
    }
    if b.UsageAmountFormatted != nil {
        structMap["usage_amount_formatted"] = b.UsageAmountFormatted
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Breakouts.
// It customizes the JSON unmarshaling process for Breakouts objects.
func (b *Breakouts) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PlanAmountInCents    *int64  `json:"plan_amount_in_cents,omitempty"`
        PlanAmountFormatted  *string `json:"plan_amount_formatted,omitempty"`
        UsageAmountInCents   *int64  `json:"usage_amount_in_cents,omitempty"`
        UsageAmountFormatted *string `json:"usage_amount_formatted,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    b.PlanAmountInCents = temp.PlanAmountInCents
    b.PlanAmountFormatted = temp.PlanAmountFormatted
    b.UsageAmountInCents = temp.UsageAmountInCents
    b.UsageAmountFormatted = temp.UsageAmountFormatted
    return nil
}
