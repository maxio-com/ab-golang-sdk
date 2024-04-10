package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionMRR represents a SubscriptionMRR struct.
type SubscriptionMRR struct {
    SubscriptionId       int                      `json:"subscription_id"`
    MrrAmountInCents     int64                    `json:"mrr_amount_in_cents"`
    Breakouts            *SubscriptionMRRBreakout `json:"breakouts,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMRR.
// It customizes the JSON marshaling process for SubscriptionMRR objects.
func (s SubscriptionMRR) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMRR object to a map representation for JSON marshaling.
func (s SubscriptionMRR) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp subscriptionMRR
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "subscription_id", "mrr_amount_in_cents", "breakouts")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.SubscriptionId = *temp.SubscriptionId
    s.MrrAmountInCents = *temp.MrrAmountInCents
    s.Breakouts = temp.Breakouts
    return nil
}

// TODO
type subscriptionMRR  struct {
    SubscriptionId   *int                     `json:"subscription_id"`
    MrrAmountInCents *int64                   `json:"mrr_amount_in_cents"`
    Breakouts        *SubscriptionMRRBreakout `json:"breakouts,omitempty"`
}

func (s *subscriptionMRR) validate() error {
    var errs []string
    if s.SubscriptionId == nil {
        errs = append(errs, "required field `subscription_id` is missing for type `Subscription MRR`")
    }
    if s.MrrAmountInCents == nil {
        errs = append(errs, "required field `mrr_amount_in_cents` is missing for type `Subscription MRR`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
