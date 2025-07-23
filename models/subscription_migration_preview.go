// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// SubscriptionMigrationPreview represents a SubscriptionMigrationPreview struct.
type SubscriptionMigrationPreview struct {
	// The amount of the prorated adjustment that would be issued for the current subscription.
	ProratedAdjustmentInCents *int64 `json:"prorated_adjustment_in_cents,omitempty"`
	// The amount of the charge that would be created for the new product.
	ChargeInCents *int64 `json:"charge_in_cents,omitempty"`
	// The amount of the payment due in the case of an upgrade.
	PaymentDueInCents *int64 `json:"payment_due_in_cents,omitempty"`
	// Represents a credit in cents that is applied to your subscription as part of a migration process for a specific product, which reduces the amount owed for the subscription.
	CreditAppliedInCents *int64                 `json:"credit_applied_in_cents,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionMigrationPreview,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionMigrationPreview) String() string {
	return fmt.Sprintf(
		"SubscriptionMigrationPreview[ProratedAdjustmentInCents=%v, ChargeInCents=%v, PaymentDueInCents=%v, CreditAppliedInCents=%v, AdditionalProperties=%v]",
		s.ProratedAdjustmentInCents, s.ChargeInCents, s.PaymentDueInCents, s.CreditAppliedInCents, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMigrationPreview.
// It customizes the JSON marshaling process for SubscriptionMigrationPreview objects.
func (s SubscriptionMigrationPreview) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"prorated_adjustment_in_cents", "charge_in_cents", "payment_due_in_cents", "credit_applied_in_cents"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMigrationPreview object to a map representation for JSON marshaling.
func (s SubscriptionMigrationPreview) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.ProratedAdjustmentInCents != nil {
		structMap["prorated_adjustment_in_cents"] = s.ProratedAdjustmentInCents
	}
	if s.ChargeInCents != nil {
		structMap["charge_in_cents"] = s.ChargeInCents
	}
	if s.PaymentDueInCents != nil {
		structMap["payment_due_in_cents"] = s.PaymentDueInCents
	}
	if s.CreditAppliedInCents != nil {
		structMap["credit_applied_in_cents"] = s.CreditAppliedInCents
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMigrationPreview.
// It customizes the JSON unmarshaling process for SubscriptionMigrationPreview objects.
func (s *SubscriptionMigrationPreview) UnmarshalJSON(input []byte) error {
	var temp tempSubscriptionMigrationPreview
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "prorated_adjustment_in_cents", "charge_in_cents", "payment_due_in_cents", "credit_applied_in_cents")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.ProratedAdjustmentInCents = temp.ProratedAdjustmentInCents
	s.ChargeInCents = temp.ChargeInCents
	s.PaymentDueInCents = temp.PaymentDueInCents
	s.CreditAppliedInCents = temp.CreditAppliedInCents
	return nil
}

// tempSubscriptionMigrationPreview is a temporary struct used for validating the fields of SubscriptionMigrationPreview.
type tempSubscriptionMigrationPreview struct {
	ProratedAdjustmentInCents *int64 `json:"prorated_adjustment_in_cents,omitempty"`
	ChargeInCents             *int64 `json:"charge_in_cents,omitempty"`
	PaymentDueInCents         *int64 `json:"payment_due_in_cents,omitempty"`
	CreditAppliedInCents      *int64 `json:"credit_applied_in_cents,omitempty"`
}
