/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// RenewalPreview represents a RenewalPreview struct.
type RenewalPreview struct {
    // The timestamp for the subscription’s next renewal
    NextAssessmentAt       *time.Time               `json:"next_assessment_at,omitempty"`
    // An integer representing the amount of the total pre-tax, pre-discount charges that will be assessed at the next renewal
    SubtotalInCents        *int64                   `json:"subtotal_in_cents,omitempty"`
    // An integer representing the total tax charges that will be assessed at the next renewal
    TotalTaxInCents        *int64                   `json:"total_tax_in_cents,omitempty"`
    // An integer representing the amount of the coupon discounts that will be applied to the next renewal
    TotalDiscountInCents   *int64                   `json:"total_discount_in_cents,omitempty"`
    // An integer representing the total amount owed, less any discounts, that will be assessed at the next renewal
    TotalInCents           *int64                   `json:"total_in_cents,omitempty"`
    // An integer representing the amount of the subscription’s current balance
    ExistingBalanceInCents *int64                   `json:"existing_balance_in_cents,omitempty"`
    // An integer representing the existing_balance_in_cents plus the total_in_cents
    TotalAmountDueInCents  *int64                   `json:"total_amount_due_in_cents,omitempty"`
    // A boolean indicating whether or not additional taxes will be calculated at the time of renewal. This will be true if you are using Avalara and the address of the subscription is in one of your defined taxable regions.
    UncalculatedTaxes      *bool                    `json:"uncalculated_taxes,omitempty"`
    // An array of objects representing the individual transactions that will be created at the next renewal
    LineItems              []RenewalPreviewLineItem `json:"line_items,omitempty"`
    AdditionalProperties   map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for RenewalPreview,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RenewalPreview) String() string {
    return fmt.Sprintf(
    	"RenewalPreview[NextAssessmentAt=%v, SubtotalInCents=%v, TotalTaxInCents=%v, TotalDiscountInCents=%v, TotalInCents=%v, ExistingBalanceInCents=%v, TotalAmountDueInCents=%v, UncalculatedTaxes=%v, LineItems=%v, AdditionalProperties=%v]",
    	r.NextAssessmentAt, r.SubtotalInCents, r.TotalTaxInCents, r.TotalDiscountInCents, r.TotalInCents, r.ExistingBalanceInCents, r.TotalAmountDueInCents, r.UncalculatedTaxes, r.LineItems, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RenewalPreview.
// It customizes the JSON marshaling process for RenewalPreview objects.
func (r RenewalPreview) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "next_assessment_at", "subtotal_in_cents", "total_tax_in_cents", "total_discount_in_cents", "total_in_cents", "existing_balance_in_cents", "total_amount_due_in_cents", "uncalculated_taxes", "line_items"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RenewalPreview object to a map representation for JSON marshaling.
func (r RenewalPreview) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.NextAssessmentAt != nil {
        structMap["next_assessment_at"] = r.NextAssessmentAt.Format(time.RFC3339)
    }
    if r.SubtotalInCents != nil {
        structMap["subtotal_in_cents"] = r.SubtotalInCents
    }
    if r.TotalTaxInCents != nil {
        structMap["total_tax_in_cents"] = r.TotalTaxInCents
    }
    if r.TotalDiscountInCents != nil {
        structMap["total_discount_in_cents"] = r.TotalDiscountInCents
    }
    if r.TotalInCents != nil {
        structMap["total_in_cents"] = r.TotalInCents
    }
    if r.ExistingBalanceInCents != nil {
        structMap["existing_balance_in_cents"] = r.ExistingBalanceInCents
    }
    if r.TotalAmountDueInCents != nil {
        structMap["total_amount_due_in_cents"] = r.TotalAmountDueInCents
    }
    if r.UncalculatedTaxes != nil {
        structMap["uncalculated_taxes"] = r.UncalculatedTaxes
    }
    if r.LineItems != nil {
        structMap["line_items"] = r.LineItems
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RenewalPreview.
// It customizes the JSON unmarshaling process for RenewalPreview objects.
func (r *RenewalPreview) UnmarshalJSON(input []byte) error {
    var temp tempRenewalPreview
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "next_assessment_at", "subtotal_in_cents", "total_tax_in_cents", "total_discount_in_cents", "total_in_cents", "existing_balance_in_cents", "total_amount_due_in_cents", "uncalculated_taxes", "line_items")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    if temp.NextAssessmentAt != nil {
        NextAssessmentAtVal, err := time.Parse(time.RFC3339, *temp.NextAssessmentAt)
        if err != nil {
            log.Fatalf("Cannot Parse next_assessment_at as % s format.", time.RFC3339)
        }
        r.NextAssessmentAt = &NextAssessmentAtVal
    }
    r.SubtotalInCents = temp.SubtotalInCents
    r.TotalTaxInCents = temp.TotalTaxInCents
    r.TotalDiscountInCents = temp.TotalDiscountInCents
    r.TotalInCents = temp.TotalInCents
    r.ExistingBalanceInCents = temp.ExistingBalanceInCents
    r.TotalAmountDueInCents = temp.TotalAmountDueInCents
    r.UncalculatedTaxes = temp.UncalculatedTaxes
    r.LineItems = temp.LineItems
    return nil
}

// tempRenewalPreview is a temporary struct used for validating the fields of RenewalPreview.
type tempRenewalPreview  struct {
    NextAssessmentAt       *string                  `json:"next_assessment_at,omitempty"`
    SubtotalInCents        *int64                   `json:"subtotal_in_cents,omitempty"`
    TotalTaxInCents        *int64                   `json:"total_tax_in_cents,omitempty"`
    TotalDiscountInCents   *int64                   `json:"total_discount_in_cents,omitempty"`
    TotalInCents           *int64                   `json:"total_in_cents,omitempty"`
    ExistingBalanceInCents *int64                   `json:"existing_balance_in_cents,omitempty"`
    TotalAmountDueInCents  *int64                   `json:"total_amount_due_in_cents,omitempty"`
    UncalculatedTaxes      *bool                    `json:"uncalculated_taxes,omitempty"`
    LineItems              []RenewalPreviewLineItem `json:"line_items,omitempty"`
}
