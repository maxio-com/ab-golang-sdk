package models

import (
	"encoding/json"
)

// SubscriptionMRRResponse represents a SubscriptionMRRResponse struct.
type SubscriptionMRRResponse struct {
	SubscriptionsMrr []SubscriptionMRR `json:"subscriptions_mrr"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMRRResponse.
// It customizes the JSON marshaling process for SubscriptionMRRResponse objects.
func (s *SubscriptionMRRResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMRRResponse object to a map representation for JSON marshaling.
func (s *SubscriptionMRRResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscriptions_mrr"] = s.SubscriptionsMrr
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMRRResponse.
// It customizes the JSON unmarshaling process for SubscriptionMRRResponse objects.
func (s *SubscriptionMRRResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SubscriptionsMrr []SubscriptionMRR `json:"subscriptions_mrr"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.SubscriptionsMrr = temp.SubscriptionsMrr
	return nil
}
