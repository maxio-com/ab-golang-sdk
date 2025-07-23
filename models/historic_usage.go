// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// HistoricUsage represents a HistoricUsage struct.
// Optional for Event Based Components. If the `include=historic_usages` query param is provided, the last ten billing periods will be returned.
type HistoricUsage struct {
	// Total usage of a component for billing period
	TotalUsageQuantity *float64 `json:"total_usage_quantity,omitempty"`
	// Start date of billing period
	BillingPeriodStartsAt *time.Time `json:"billing_period_starts_at,omitempty"`
	// End date of billing period
	BillingPeriodEndsAt  *time.Time             `json:"billing_period_ends_at,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for HistoricUsage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (h HistoricUsage) String() string {
	return fmt.Sprintf(
		"HistoricUsage[TotalUsageQuantity=%v, BillingPeriodStartsAt=%v, BillingPeriodEndsAt=%v, AdditionalProperties=%v]",
		h.TotalUsageQuantity, h.BillingPeriodStartsAt, h.BillingPeriodEndsAt, h.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for HistoricUsage.
// It customizes the JSON marshaling process for HistoricUsage objects.
func (h HistoricUsage) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(h.AdditionalProperties,
		"total_usage_quantity", "billing_period_starts_at", "billing_period_ends_at"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(h.toMap())
}

// toMap converts the HistoricUsage object to a map representation for JSON marshaling.
func (h HistoricUsage) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, h.AdditionalProperties)
	if h.TotalUsageQuantity != nil {
		structMap["total_usage_quantity"] = h.TotalUsageQuantity
	}
	if h.BillingPeriodStartsAt != nil {
		structMap["billing_period_starts_at"] = h.BillingPeriodStartsAt.Format(time.RFC3339)
	}
	if h.BillingPeriodEndsAt != nil {
		structMap["billing_period_ends_at"] = h.BillingPeriodEndsAt.Format(time.RFC3339)
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for HistoricUsage.
// It customizes the JSON unmarshaling process for HistoricUsage objects.
func (h *HistoricUsage) UnmarshalJSON(input []byte) error {
	var temp tempHistoricUsage
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "total_usage_quantity", "billing_period_starts_at", "billing_period_ends_at")
	if err != nil {
		return err
	}
	h.AdditionalProperties = additionalProperties

	h.TotalUsageQuantity = temp.TotalUsageQuantity
	if temp.BillingPeriodStartsAt != nil {
		BillingPeriodStartsAtVal, err := time.Parse(time.RFC3339, *temp.BillingPeriodStartsAt)
		if err != nil {
			log.Fatalf("Cannot Parse billing_period_starts_at as % s format.", time.RFC3339)
		}
		h.BillingPeriodStartsAt = &BillingPeriodStartsAtVal
	}
	if temp.BillingPeriodEndsAt != nil {
		BillingPeriodEndsAtVal, err := time.Parse(time.RFC3339, *temp.BillingPeriodEndsAt)
		if err != nil {
			log.Fatalf("Cannot Parse billing_period_ends_at as % s format.", time.RFC3339)
		}
		h.BillingPeriodEndsAt = &BillingPeriodEndsAtVal
	}
	return nil
}

// tempHistoricUsage is a temporary struct used for validating the fields of HistoricUsage.
type tempHistoricUsage struct {
	TotalUsageQuantity    *float64 `json:"total_usage_quantity,omitempty"`
	BillingPeriodStartsAt *string  `json:"billing_period_starts_at,omitempty"`
	BillingPeriodEndsAt   *string  `json:"billing_period_ends_at,omitempty"`
}
