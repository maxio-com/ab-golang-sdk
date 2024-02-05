package models

import (
    "encoding/json"
)

// SubscriptionMRR represents a SubscriptionMRR struct.
type SubscriptionMRR struct {
    SubscriptionId   int                      `json:"subscription_id"`
    MrrAmountInCents int64                    `json:"mrr_amount_in_cents"`
    Breakouts        *SubscriptionMRRBreakout `json:"breakouts,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMRR.
// It customizes the JSON marshaling process for SubscriptionMRR objects.
func (s *SubscriptionMRR) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMRR object to a map representation for JSON marshaling.
func (s *SubscriptionMRR) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["subscription_id"] = s.SubscriptionId
    structMap["mrr_amount_in_cents"] = s.MrrAmountInCents
    if s.Breakouts != nil {
        structMap["breakouts"] = s.Breakouts.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMRR.
// It customizes the JSON unmarshaling process for SubscriptionMRR objects.
func (s *SubscriptionMRR) UnmarshalJSON(input []byte) error {
    temp := &struct {
        SubscriptionId   int                      `json:"subscription_id"`
        MrrAmountInCents int64                    `json:"mrr_amount_in_cents"`
        Breakouts        *SubscriptionMRRBreakout `json:"breakouts,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.SubscriptionId = temp.SubscriptionId
    s.MrrAmountInCents = temp.MrrAmountInCents
    s.Breakouts = temp.Breakouts
    return nil
}
