package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionMRRBreakout represents a SubscriptionMRRBreakout struct.
type SubscriptionMRRBreakout struct {
    PlanAmountInCents    int64          `json:"plan_amount_in_cents"`
    UsageAmountInCents   int64          `json:"usage_amount_in_cents"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMRRBreakout.
// It customizes the JSON marshaling process for SubscriptionMRRBreakout objects.
func (s SubscriptionMRRBreakout) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMRRBreakout object to a map representation for JSON marshaling.
func (s SubscriptionMRRBreakout) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["plan_amount_in_cents"] = s.PlanAmountInCents
    structMap["usage_amount_in_cents"] = s.UsageAmountInCents
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMRRBreakout.
// It customizes the JSON unmarshaling process for SubscriptionMRRBreakout objects.
func (s *SubscriptionMRRBreakout) UnmarshalJSON(input []byte) error {
    var temp subscriptionMRRBreakout
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "plan_amount_in_cents", "usage_amount_in_cents")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.PlanAmountInCents = *temp.PlanAmountInCents
    s.UsageAmountInCents = *temp.UsageAmountInCents
    return nil
}

// subscriptionMRRBreakout is a temporary struct used for validating the fields of SubscriptionMRRBreakout.
type subscriptionMRRBreakout  struct {
    PlanAmountInCents  *int64 `json:"plan_amount_in_cents"`
    UsageAmountInCents *int64 `json:"usage_amount_in_cents"`
}

func (s *subscriptionMRRBreakout) validate() error {
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
    return errors.New(strings.Join(errs, "\n"))
}
