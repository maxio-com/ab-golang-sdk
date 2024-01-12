package models

import (
    "encoding/json"
)

// SubscriptionIncludedCoupon represents a SubscriptionIncludedCoupon struct.
type SubscriptionIncludedCoupon struct {
    Code          *string          `json:"code,omitempty"`
    UseCount      *int             `json:"use_count,omitempty"`
    UsesAllowed   *int             `json:"uses_allowed,omitempty"`
    ExpiresAt     Optional[string] `json:"expires_at"`
    Recurring     *bool            `json:"recurring,omitempty"`
    AmountInCents Optional[int64]  `json:"amount_in_cents"`
    Percentage    Optional[string] `json:"percentage"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionIncludedCoupon.
// It customizes the JSON marshaling process for SubscriptionIncludedCoupon objects.
func (s *SubscriptionIncludedCoupon) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionIncludedCoupon object to a map representation for JSON marshaling.
func (s *SubscriptionIncludedCoupon) toMap() map[string]any {
    structMap := make(map[string]any)
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
        structMap["expires_at"] = s.ExpiresAt.Value()
    }
    if s.Recurring != nil {
        structMap["recurring"] = s.Recurring
    }
    if s.AmountInCents.IsValueSet() {
        structMap["amount_in_cents"] = s.AmountInCents.Value()
    }
    if s.Percentage.IsValueSet() {
        structMap["percentage"] = s.Percentage.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionIncludedCoupon.
// It customizes the JSON unmarshaling process for SubscriptionIncludedCoupon objects.
func (s *SubscriptionIncludedCoupon) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Code          *string          `json:"code,omitempty"`
        UseCount      *int             `json:"use_count,omitempty"`
        UsesAllowed   *int             `json:"uses_allowed,omitempty"`
        ExpiresAt     Optional[string] `json:"expires_at"`
        Recurring     *bool            `json:"recurring,omitempty"`
        AmountInCents Optional[int64]  `json:"amount_in_cents"`
        Percentage    Optional[string] `json:"percentage"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Code = temp.Code
    s.UseCount = temp.UseCount
    s.UsesAllowed = temp.UsesAllowed
    s.ExpiresAt = temp.ExpiresAt
    s.Recurring = temp.Recurring
    s.AmountInCents = temp.AmountInCents
    s.Percentage = temp.Percentage
    return nil
}
