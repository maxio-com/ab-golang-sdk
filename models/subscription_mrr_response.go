// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SubscriptionMRRResponse represents a SubscriptionMRRResponse struct.
type SubscriptionMRRResponse struct {
    SubscriptionsMrr     []SubscriptionMRR      `json:"subscriptions_mrr"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionMRRResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionMRRResponse) String() string {
    return fmt.Sprintf(
    	"SubscriptionMRRResponse[SubscriptionsMrr=%v, AdditionalProperties=%v]",
    	s.SubscriptionsMrr, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMRRResponse.
// It customizes the JSON marshaling process for SubscriptionMRRResponse objects.
func (s SubscriptionMRRResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "subscriptions_mrr"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMRRResponse object to a map representation for JSON marshaling.
func (s SubscriptionMRRResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["subscriptions_mrr"] = s.SubscriptionsMrr
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMRRResponse.
// It customizes the JSON unmarshaling process for SubscriptionMRRResponse objects.
func (s *SubscriptionMRRResponse) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionMRRResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscriptions_mrr")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.SubscriptionsMrr = *temp.SubscriptionsMrr
    return nil
}

// tempSubscriptionMRRResponse is a temporary struct used for validating the fields of SubscriptionMRRResponse.
type tempSubscriptionMRRResponse  struct {
    SubscriptionsMrr *[]SubscriptionMRR `json:"subscriptions_mrr"`
}

func (s *tempSubscriptionMRRResponse) validate() error {
    var errs []string
    if s.SubscriptionsMrr == nil {
        errs = append(errs, "required field `subscriptions_mrr` is missing for type `Subscription MRR Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
