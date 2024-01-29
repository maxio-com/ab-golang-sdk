package models

import (
	"encoding/json"
)

// SubscriptionPreview represents a SubscriptionPreview struct.
type SubscriptionPreview struct {
	CurrentBillingManifest *BillingManifest `json:"current_billing_manifest,omitempty"`
	NextBillingManifest    *BillingManifest `json:"next_billing_manifest,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionPreview.
// It customizes the JSON marshaling process for SubscriptionPreview objects.
func (s *SubscriptionPreview) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionPreview object to a map representation for JSON marshaling.
func (s *SubscriptionPreview) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.CurrentBillingManifest != nil {
		structMap["current_billing_manifest"] = s.CurrentBillingManifest
	}
	if s.NextBillingManifest != nil {
		structMap["next_billing_manifest"] = s.NextBillingManifest
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionPreview.
// It customizes the JSON unmarshaling process for SubscriptionPreview objects.
func (s *SubscriptionPreview) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CurrentBillingManifest *BillingManifest `json:"current_billing_manifest,omitempty"`
		NextBillingManifest    *BillingManifest `json:"next_billing_manifest,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.CurrentBillingManifest = temp.CurrentBillingManifest
	s.NextBillingManifest = temp.NextBillingManifest
	return nil
}
