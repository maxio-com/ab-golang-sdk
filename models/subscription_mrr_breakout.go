package models

import (
	"encoding/json"
)

// SubscriptionMRRBreakout represents a SubscriptionMRRBreakout struct.
type SubscriptionMRRBreakout struct {
	PlanAmountInCents  int64 `json:"plan_amount_in_cents"`
	UsageAmountInCents int64 `json:"usage_amount_in_cents"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMRRBreakout.
// It customizes the JSON marshaling process for SubscriptionMRRBreakout objects.
func (s *SubscriptionMRRBreakout) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMRRBreakout object to a map representation for JSON marshaling.
func (s *SubscriptionMRRBreakout) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["plan_amount_in_cents"] = s.PlanAmountInCents
	structMap["usage_amount_in_cents"] = s.UsageAmountInCents
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMRRBreakout.
// It customizes the JSON unmarshaling process for SubscriptionMRRBreakout objects.
func (s *SubscriptionMRRBreakout) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PlanAmountInCents  int64 `json:"plan_amount_in_cents"`
		UsageAmountInCents int64 `json:"usage_amount_in_cents"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.PlanAmountInCents = temp.PlanAmountInCents
	s.UsageAmountInCents = temp.UsageAmountInCents
	return nil
}
