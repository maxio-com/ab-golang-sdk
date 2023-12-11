package models

import (
	"encoding/json"
)

// OverrideSubscription represents a OverrideSubscription struct.
type OverrideSubscription struct {
	// Can be used to record an external signup date. Chargify uses this field to record when a subscription first goes active (either at signup or at trial end)
	ActivatedAt *string `json:"activated_at,omitempty"`
	// Can be used to record an external cancellation date. Chargify sets this field automatically when a subscription is canceled, whether by request or via dunning.
	CanceledAt *string `json:"canceled_at,omitempty"`
	// Can be used to record a reason for the original cancellation.
	CancellationMessage *string `json:"cancellation_message,omitempty"`
	// Can be used to record an external expiration date. Chargify sets this field automatically when a subscription expires (ceases billing) after a prescribed amount of time.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// Can only be used when a subscription is unbilled, which happens when a future initial billing date is passed at subscription creation. The value passed must be before the current date and time. Allows you to set when the period started so mid period component allocations have the correct proration.
	CurrentPeriodStartsAt *string `json:"current_period_starts_at,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for OverrideSubscription.
// It customizes the JSON marshaling process for OverrideSubscription objects.
func (o *OverrideSubscription) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OverrideSubscription object to a map representation for JSON marshaling.
func (o *OverrideSubscription) toMap() map[string]any {
	structMap := make(map[string]any)
	if o.ActivatedAt != nil {
		structMap["activated_at"] = o.ActivatedAt
	}
	if o.CanceledAt != nil {
		structMap["canceled_at"] = o.CanceledAt
	}
	if o.CancellationMessage != nil {
		structMap["cancellation_message"] = o.CancellationMessage
	}
	if o.ExpiresAt != nil {
		structMap["expires_at"] = o.ExpiresAt
	}
	if o.CurrentPeriodStartsAt != nil {
		structMap["current_period_starts_at"] = o.CurrentPeriodStartsAt
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OverrideSubscription.
// It customizes the JSON unmarshaling process for OverrideSubscription objects.
func (o *OverrideSubscription) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ActivatedAt           *string `json:"activated_at,omitempty"`
		CanceledAt            *string `json:"canceled_at,omitempty"`
		CancellationMessage   *string `json:"cancellation_message,omitempty"`
		ExpiresAt             *string `json:"expires_at,omitempty"`
		CurrentPeriodStartsAt *string `json:"current_period_starts_at,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	o.ActivatedAt = temp.ActivatedAt
	o.CanceledAt = temp.CanceledAt
	o.CancellationMessage = temp.CancellationMessage
	o.ExpiresAt = temp.ExpiresAt
	o.CurrentPeriodStartsAt = temp.CurrentPeriodStartsAt
	return nil
}
