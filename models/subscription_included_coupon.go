// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// SubscriptionIncludedCoupon represents a SubscriptionIncludedCoupon struct.
type SubscriptionIncludedCoupon struct {
	Code                 *string                `json:"code,omitempty"`
	UseCount             *int                   `json:"use_count,omitempty"`
	UsesAllowed          *int                   `json:"uses_allowed,omitempty"`
	ExpiresAt            Optional[string]       `json:"expires_at"`
	Recurring            *bool                  `json:"recurring,omitempty"`
	AmountInCents        Optional[int64]        `json:"amount_in_cents"`
	Percentage           Optional[string]       `json:"percentage"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionIncludedCoupon,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionIncludedCoupon) String() string {
	return fmt.Sprintf(
		"SubscriptionIncludedCoupon[Code=%v, UseCount=%v, UsesAllowed=%v, ExpiresAt=%v, Recurring=%v, AmountInCents=%v, Percentage=%v, AdditionalProperties=%v]",
		s.Code, s.UseCount, s.UsesAllowed, s.ExpiresAt, s.Recurring, s.AmountInCents, s.Percentage, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionIncludedCoupon.
// It customizes the JSON marshaling process for SubscriptionIncludedCoupon objects.
func (s SubscriptionIncludedCoupon) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"code", "use_count", "uses_allowed", "expires_at", "recurring", "amount_in_cents", "percentage"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionIncludedCoupon object to a map representation for JSON marshaling.
func (s SubscriptionIncludedCoupon) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Code != nil {
		structMap["code"] = s.Code
	}
	if s.UseCount != nil {
		structMap["use_count"] = s.UseCount
	}
	if s.UsesAllowed != nil {
		structMap["uses_allowed"] = s.UsesAllowed
	}
	if s.ExpiresAt.IsValueSet() {
		if s.ExpiresAt.Value() != nil {
			structMap["expires_at"] = s.ExpiresAt.Value()
		} else {
			structMap["expires_at"] = nil
		}
	}
	if s.Recurring != nil {
		structMap["recurring"] = s.Recurring
	}
	if s.AmountInCents.IsValueSet() {
		if s.AmountInCents.Value() != nil {
			structMap["amount_in_cents"] = s.AmountInCents.Value()
		} else {
			structMap["amount_in_cents"] = nil
		}
	}
	if s.Percentage.IsValueSet() {
		if s.Percentage.Value() != nil {
			structMap["percentage"] = s.Percentage.Value()
		} else {
			structMap["percentage"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionIncludedCoupon.
// It customizes the JSON unmarshaling process for SubscriptionIncludedCoupon objects.
func (s *SubscriptionIncludedCoupon) UnmarshalJSON(input []byte) error {
	var temp tempSubscriptionIncludedCoupon
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "code", "use_count", "uses_allowed", "expires_at", "recurring", "amount_in_cents", "percentage")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Code = temp.Code
	s.UseCount = temp.UseCount
	s.UsesAllowed = temp.UsesAllowed
	s.ExpiresAt = temp.ExpiresAt
	s.Recurring = temp.Recurring
	s.AmountInCents = temp.AmountInCents
	s.Percentage = temp.Percentage
	return nil
}

// tempSubscriptionIncludedCoupon is a temporary struct used for validating the fields of SubscriptionIncludedCoupon.
type tempSubscriptionIncludedCoupon struct {
	Code          *string          `json:"code,omitempty"`
	UseCount      *int             `json:"use_count,omitempty"`
	UsesAllowed   *int             `json:"uses_allowed,omitempty"`
	ExpiresAt     Optional[string] `json:"expires_at"`
	Recurring     *bool            `json:"recurring,omitempty"`
	AmountInCents Optional[int64]  `json:"amount_in_cents"`
	Percentage    Optional[string] `json:"percentage"`
}
