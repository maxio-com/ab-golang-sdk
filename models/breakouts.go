package models

import (
    "encoding/json"
)

// Breakouts represents a Breakouts struct.
type Breakouts struct {
    PlanAmountInCents    *int64         `json:"plan_amount_in_cents,omitempty"`
    PlanAmountFormatted  *string        `json:"plan_amount_formatted,omitempty"`
    UsageAmountInCents   *int64         `json:"usage_amount_in_cents,omitempty"`
    UsageAmountFormatted *string        `json:"usage_amount_formatted,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Breakouts.
// It customizes the JSON marshaling process for Breakouts objects.
func (b Breakouts) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the Breakouts object to a map representation for JSON marshaling.
func (b Breakouts) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
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
    var temp breakouts
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "plan_amount_in_cents", "plan_amount_formatted", "usage_amount_in_cents", "usage_amount_formatted")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.PlanAmountInCents = temp.PlanAmountInCents
    b.PlanAmountFormatted = temp.PlanAmountFormatted
    b.UsageAmountInCents = temp.UsageAmountInCents
    b.UsageAmountFormatted = temp.UsageAmountFormatted
    return nil
}

// breakouts is a temporary struct used for validating the fields of Breakouts.
type breakouts  struct {
    PlanAmountInCents    *int64  `json:"plan_amount_in_cents,omitempty"`
    PlanAmountFormatted  *string `json:"plan_amount_formatted,omitempty"`
    UsageAmountInCents   *int64  `json:"usage_amount_in_cents,omitempty"`
    UsageAmountFormatted *string `json:"usage_amount_formatted,omitempty"`
}
