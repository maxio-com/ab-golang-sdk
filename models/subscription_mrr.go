// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SubscriptionMRR represents a SubscriptionMRR struct.
type SubscriptionMRR struct {
    SubscriptionId       int                      `json:"subscription_id"`
    MrrAmountInCents     int64                    `json:"mrr_amount_in_cents"`
    Breakouts            *SubscriptionMRRBreakout `json:"breakouts,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionMRR,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionMRR) String() string {
    return fmt.Sprintf(
    	"SubscriptionMRR[SubscriptionId=%v, MrrAmountInCents=%v, Breakouts=%v, AdditionalProperties=%v]",
    	s.SubscriptionId, s.MrrAmountInCents, s.Breakouts, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMRR.
// It customizes the JSON marshaling process for SubscriptionMRR objects.
func (s SubscriptionMRR) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "subscription_id", "mrr_amount_in_cents", "breakouts"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMRR object to a map representation for JSON marshaling.
func (s SubscriptionMRR) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSubscriptionMRR
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription_id", "mrr_amount_in_cents", "breakouts")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.SubscriptionId = *temp.SubscriptionId
    s.MrrAmountInCents = *temp.MrrAmountInCents
    s.Breakouts = temp.Breakouts
    return nil
}

// tempSubscriptionMRR is a temporary struct used for validating the fields of SubscriptionMRR.
type tempSubscriptionMRR  struct {
    SubscriptionId   *int                     `json:"subscription_id"`
    MrrAmountInCents *int64                   `json:"mrr_amount_in_cents"`
    Breakouts        *SubscriptionMRRBreakout `json:"breakouts,omitempty"`
}

func (s *tempSubscriptionMRR) validate() error {
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
    return errors.New(strings.Join (errs, "\n"))
}
