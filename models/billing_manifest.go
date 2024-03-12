package models

import (
	"encoding/json"
	"log"
	"time"
)

// BillingManifest represents a BillingManifest struct.
type BillingManifest struct {
	LineItems              []BillingManifestItem `json:"line_items,omitempty"`
	TotalInCents           *int64                `json:"total_in_cents,omitempty"`
	TotalDiscountInCents   *int64                `json:"total_discount_in_cents,omitempty"`
	TotalTaxInCents        *int64                `json:"total_tax_in_cents,omitempty"`
	SubtotalInCents        *int64                `json:"subtotal_in_cents,omitempty"`
	StartDate              *time.Time            `json:"start_date,omitempty"`
	EndDate                *time.Time            `json:"end_date,omitempty"`
	PeriodType             *string               `json:"period_type,omitempty"`
	ExistingBalanceInCents *int64                `json:"existing_balance_in_cents,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BillingManifest.
// It customizes the JSON marshaling process for BillingManifest objects.
func (b *BillingManifest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BillingManifest object to a map representation for JSON marshaling.
func (b *BillingManifest) toMap() map[string]any {
	structMap := make(map[string]any)
	if b.LineItems != nil {
		structMap["line_items"] = b.LineItems
	}
	if b.TotalInCents != nil {
		structMap["total_in_cents"] = b.TotalInCents
	}
	if b.TotalDiscountInCents != nil {
		structMap["total_discount_in_cents"] = b.TotalDiscountInCents
	}
	if b.TotalTaxInCents != nil {
		structMap["total_tax_in_cents"] = b.TotalTaxInCents
	}
	if b.SubtotalInCents != nil {
		structMap["subtotal_in_cents"] = b.SubtotalInCents
	}
	if b.StartDate != nil {
		structMap["start_date"] = b.StartDate.Format(time.RFC3339)
	}
	if b.EndDate != nil {
		structMap["end_date"] = b.EndDate.Format(time.RFC3339)
	}
	if b.PeriodType != nil {
		structMap["period_type"] = b.PeriodType
	}
	if b.ExistingBalanceInCents != nil {
		structMap["existing_balance_in_cents"] = b.ExistingBalanceInCents
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BillingManifest.
// It customizes the JSON unmarshaling process for BillingManifest objects.
func (b *BillingManifest) UnmarshalJSON(input []byte) error {
	var temp billingManifest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	b.LineItems = temp.LineItems
	b.TotalInCents = temp.TotalInCents
	b.TotalDiscountInCents = temp.TotalDiscountInCents
	b.TotalTaxInCents = temp.TotalTaxInCents
	b.SubtotalInCents = temp.SubtotalInCents
	if temp.StartDate != nil {
		StartDateVal, err := time.Parse(time.RFC3339, *temp.StartDate)
		if err != nil {
			log.Fatalf("Cannot Parse start_date as % s format.", time.RFC3339)
		}
		b.StartDate = &StartDateVal
	}
	if temp.EndDate != nil {
		EndDateVal, err := time.Parse(time.RFC3339, *temp.EndDate)
		if err != nil {
			log.Fatalf("Cannot Parse end_date as % s format.", time.RFC3339)
		}
		b.EndDate = &EndDateVal
	}
	b.PeriodType = temp.PeriodType
	b.ExistingBalanceInCents = temp.ExistingBalanceInCents
	return nil
}

// TODO
type billingManifest struct {
	LineItems              []BillingManifestItem `json:"line_items,omitempty"`
	TotalInCents           *int64                `json:"total_in_cents,omitempty"`
	TotalDiscountInCents   *int64                `json:"total_discount_in_cents,omitempty"`
	TotalTaxInCents        *int64                `json:"total_tax_in_cents,omitempty"`
	SubtotalInCents        *int64                `json:"subtotal_in_cents,omitempty"`
	StartDate              *string               `json:"start_date,omitempty"`
	EndDate                *string               `json:"end_date,omitempty"`
	PeriodType             *string               `json:"period_type,omitempty"`
	ExistingBalanceInCents *int64                `json:"existing_balance_in_cents,omitempty"`
}
