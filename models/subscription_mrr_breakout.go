// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SubscriptionMRRBreakout represents a SubscriptionMRRBreakout struct.
type SubscriptionMRRBreakout struct {
    PlanAmountInCents    int64                  `json:"plan_amount_in_cents"`
    UsageAmountInCents   int64                  `json:"usage_amount_in_cents"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionMRRBreakout,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionMRRBreakout) String() string {
    return fmt.Sprintf(
    	"SubscriptionMRRBreakout[PlanAmountInCents=%v, UsageAmountInCents=%v, AdditionalProperties=%v]",
    	s.PlanAmountInCents, s.UsageAmountInCents, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMRRBreakout.
// It customizes the JSON marshaling process for SubscriptionMRRBreakout objects.
func (s SubscriptionMRRBreakout) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "plan_amount_in_cents", "usage_amount_in_cents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMRRBreakout object to a map representation for JSON marshaling.
func (s SubscriptionMRRBreakout) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["plan_amount_in_cents"] = s.PlanAmountInCents
    structMap["usage_amount_in_cents"] = s.UsageAmountInCents
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMRRBreakout.
// It customizes the JSON unmarshaling process for SubscriptionMRRBreakout objects.
func (s *SubscriptionMRRBreakout) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionMRRBreakout
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "plan_amount_in_cents", "usage_amount_in_cents")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.PlanAmountInCents = *temp.PlanAmountInCents
    s.UsageAmountInCents = *temp.UsageAmountInCents
    return nil
}

// tempSubscriptionMRRBreakout is a temporary struct used for validating the fields of SubscriptionMRRBreakout.
type tempSubscriptionMRRBreakout  struct {
    PlanAmountInCents  *int64 `json:"plan_amount_in_cents"`
    UsageAmountInCents *int64 `json:"usage_amount_in_cents"`
}

func (s *tempSubscriptionMRRBreakout) validate() error {
    var errs []string
    if s.PlanAmountInCents == nil {
        errs = append(errs, "required field `plan_amount_in_cents` is missing for type `Subscription MRR Breakout`")
    }
    if s.UsageAmountInCents == nil {
        errs = append(errs, "required field `usage_amount_in_cents` is missing for type `Subscription MRR Breakout`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
