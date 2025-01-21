/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// SubscriptionPreview represents a SubscriptionPreview struct.
type SubscriptionPreview struct {
    CurrentBillingManifest *BillingManifest       `json:"current_billing_manifest,omitempty"`
    NextBillingManifest    *BillingManifest       `json:"next_billing_manifest,omitempty"`
    AdditionalProperties   map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionPreview,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionPreview) String() string {
    return fmt.Sprintf(
    	"SubscriptionPreview[CurrentBillingManifest=%v, NextBillingManifest=%v, AdditionalProperties=%v]",
    	s.CurrentBillingManifest, s.NextBillingManifest, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionPreview.
// It customizes the JSON marshaling process for SubscriptionPreview objects.
func (s SubscriptionPreview) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "current_billing_manifest", "next_billing_manifest"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionPreview object to a map representation for JSON marshaling.
func (s SubscriptionPreview) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CurrentBillingManifest != nil {
        structMap["current_billing_manifest"] = s.CurrentBillingManifest.toMap()
    }
    if s.NextBillingManifest != nil {
        structMap["next_billing_manifest"] = s.NextBillingManifest.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionPreview.
// It customizes the JSON unmarshaling process for SubscriptionPreview objects.
func (s *SubscriptionPreview) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionPreview
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "current_billing_manifest", "next_billing_manifest")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CurrentBillingManifest = temp.CurrentBillingManifest
    s.NextBillingManifest = temp.NextBillingManifest
    return nil
}

// tempSubscriptionPreview is a temporary struct used for validating the fields of SubscriptionPreview.
type tempSubscriptionPreview  struct {
    CurrentBillingManifest *BillingManifest `json:"current_billing_manifest,omitempty"`
    NextBillingManifest    *BillingManifest `json:"next_billing_manifest,omitempty"`
}
