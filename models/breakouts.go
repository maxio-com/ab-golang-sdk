// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// Breakouts represents a Breakouts struct.
type Breakouts struct {
    PlanAmountInCents    *int64                 `json:"plan_amount_in_cents,omitempty"`
    PlanAmountFormatted  *string                `json:"plan_amount_formatted,omitempty"`
    UsageAmountInCents   *int64                 `json:"usage_amount_in_cents,omitempty"`
    UsageAmountFormatted *string                `json:"usage_amount_formatted,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Breakouts,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b Breakouts) String() string {
    return fmt.Sprintf(
    	"Breakouts[PlanAmountInCents=%v, PlanAmountFormatted=%v, UsageAmountInCents=%v, UsageAmountFormatted=%v, AdditionalProperties=%v]",
    	b.PlanAmountInCents, b.PlanAmountFormatted, b.UsageAmountInCents, b.UsageAmountFormatted, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Breakouts.
// It customizes the JSON marshaling process for Breakouts objects.
func (b Breakouts) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "plan_amount_in_cents", "plan_amount_formatted", "usage_amount_in_cents", "usage_amount_formatted"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the Breakouts object to a map representation for JSON marshaling.
func (b Breakouts) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
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
    var temp tempBreakouts
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "plan_amount_in_cents", "plan_amount_formatted", "usage_amount_in_cents", "usage_amount_formatted")
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

// tempBreakouts is a temporary struct used for validating the fields of Breakouts.
type tempBreakouts  struct {
    PlanAmountInCents    *int64  `json:"plan_amount_in_cents,omitempty"`
    PlanAmountFormatted  *string `json:"plan_amount_formatted,omitempty"`
    UsageAmountInCents   *int64  `json:"usage_amount_in_cents,omitempty"`
    UsageAmountFormatted *string `json:"usage_amount_formatted,omitempty"`
}
