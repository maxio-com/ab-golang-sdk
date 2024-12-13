/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "log"
    "time"
)

// OverrideSubscription represents a OverrideSubscription struct.
type OverrideSubscription struct {
    // Can be used to record an external signup date. Chargify uses this field to record when a subscription first goes active (either at signup or at trial end). Only ISO8601 format is supported.
    ActivatedAt           *time.Time             `json:"activated_at,omitempty"`
    // Can be used to record an external cancellation date. Chargify sets this field automatically when a subscription is canceled, whether by request or via dunning. Only ISO8601 format is supported.
    CanceledAt            *time.Time             `json:"canceled_at,omitempty"`
    // Can be used to record a reason for the original cancellation.
    CancellationMessage   *string                `json:"cancellation_message,omitempty"`
    // Can be used to record an external expiration date. Chargify sets this field automatically when a subscription expires (ceases billing) after a prescribed amount of time. Only ISO8601 format is supported.
    ExpiresAt             *time.Time             `json:"expires_at,omitempty"`
    // Can only be used when a subscription is unbilled, which happens when a future initial billing date is passed at subscription creation. The value passed must be before the current date and time. Allows you to set when the period started so mid period component allocations have the correct proration. Only ISO8601 format is supported.
    CurrentPeriodStartsAt *time.Time             `json:"current_period_starts_at,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OverrideSubscription.
// It customizes the JSON marshaling process for OverrideSubscription objects.
func (o OverrideSubscription) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "activated_at", "canceled_at", "cancellation_message", "expires_at", "current_period_starts_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OverrideSubscription object to a map representation for JSON marshaling.
func (o OverrideSubscription) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.ActivatedAt != nil {
        structMap["activated_at"] = o.ActivatedAt.Format(time.RFC3339)
    }
    if o.CanceledAt != nil {
        structMap["canceled_at"] = o.CanceledAt.Format(time.RFC3339)
    }
    if o.CancellationMessage != nil {
        structMap["cancellation_message"] = o.CancellationMessage
    }
    if o.ExpiresAt != nil {
        structMap["expires_at"] = o.ExpiresAt.Format(time.RFC3339)
    }
    if o.CurrentPeriodStartsAt != nil {
        structMap["current_period_starts_at"] = o.CurrentPeriodStartsAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OverrideSubscription.
// It customizes the JSON unmarshaling process for OverrideSubscription objects.
func (o *OverrideSubscription) UnmarshalJSON(input []byte) error {
    var temp tempOverrideSubscription
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "activated_at", "canceled_at", "cancellation_message", "expires_at", "current_period_starts_at")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    if temp.ActivatedAt != nil {
        ActivatedAtVal, err := time.Parse(time.RFC3339, *temp.ActivatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse activated_at as % s format.", time.RFC3339)
        }
        o.ActivatedAt = &ActivatedAtVal
    }
    if temp.CanceledAt != nil {
        CanceledAtVal, err := time.Parse(time.RFC3339, *temp.CanceledAt)
        if err != nil {
            log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
        }
        o.CanceledAt = &CanceledAtVal
    }
    o.CancellationMessage = temp.CancellationMessage
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        o.ExpiresAt = &ExpiresAtVal
    }
    if temp.CurrentPeriodStartsAt != nil {
        CurrentPeriodStartsAtVal, err := time.Parse(time.RFC3339, *temp.CurrentPeriodStartsAt)
        if err != nil {
            log.Fatalf("Cannot Parse current_period_starts_at as % s format.", time.RFC3339)
        }
        o.CurrentPeriodStartsAt = &CurrentPeriodStartsAtVal
    }
    return nil
}

// tempOverrideSubscription is a temporary struct used for validating the fields of OverrideSubscription.
type tempOverrideSubscription  struct {
    ActivatedAt           *string `json:"activated_at,omitempty"`
    CanceledAt            *string `json:"canceled_at,omitempty"`
    CancellationMessage   *string `json:"cancellation_message,omitempty"`
    ExpiresAt             *string `json:"expires_at,omitempty"`
    CurrentPeriodStartsAt *string `json:"current_period_starts_at,omitempty"`
}
