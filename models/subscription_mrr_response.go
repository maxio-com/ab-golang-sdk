package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionMRRResponse represents a SubscriptionMRRResponse struct.
type SubscriptionMRRResponse struct {
    SubscriptionsMrr     []SubscriptionMRR `json:"subscriptions_mrr"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMRRResponse.
// It customizes the JSON marshaling process for SubscriptionMRRResponse objects.
func (s SubscriptionMRRResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMRRResponse object to a map representation for JSON marshaling.
func (s SubscriptionMRRResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["subscriptions_mrr"] = s.SubscriptionsMrr
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMRRResponse.
// It customizes the JSON unmarshaling process for SubscriptionMRRResponse objects.
func (s *SubscriptionMRRResponse) UnmarshalJSON(input []byte) error {
    var temp subscriptionMRRResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "subscriptions_mrr")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.SubscriptionsMrr = *temp.SubscriptionsMrr
    return nil
}

// TODO
type subscriptionMRRResponse  struct {
    SubscriptionsMrr *[]SubscriptionMRR `json:"subscriptions_mrr"`
}

func (s *subscriptionMRRResponse) validate() error {
    var errs []string
    if s.SubscriptionsMrr == nil {
        errs = append(errs, "required field `subscriptions_mrr` is missing for type `Subscription MRR Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
