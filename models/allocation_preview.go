package models

import (
	"encoding/json"
)

// AllocationPreview represents a AllocationPreview struct.
type AllocationPreview struct {
	StartDate            *string                     `json:"start_date,omitempty"`
	EndDate              *string                     `json:"end_date,omitempty"`
	SubtotalInCents      *int64                      `json:"subtotal_in_cents,omitempty"`
	TotalTaxInCents      *int64                      `json:"total_tax_in_cents,omitempty"`
	TotalDiscountInCents *int64                      `json:"total_discount_in_cents,omitempty"`
	TotalInCents         *int64                      `json:"total_in_cents,omitempty"`
	Direction            *string                     `json:"direction,omitempty"`
	ProrationScheme      *string                     `json:"proration_scheme,omitempty"`
	LineItems            []AllocationPreviewLineItem `json:"line_items,omitempty"`
	AccrueCharge         *bool                       `json:"accrue_charge,omitempty"`
	Allocations          []AllocationPreviewItem     `json:"allocations,omitempty"`
	PeriodType           *string                     `json:"period_type,omitempty"`
	// An integer representing the amount of the subscription's current balance
	ExistingBalanceInCents *int64 `json:"existing_balance_in_cents,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationPreview.
// It customizes the JSON marshaling process for AllocationPreview objects.
func (a *AllocationPreview) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AllocationPreview object to a map representation for JSON marshaling.
func (a *AllocationPreview) toMap() map[string]any {
	structMap := make(map[string]any)
	if a.StartDate != nil {
		structMap["start_date"] = a.StartDate
	}
	if a.EndDate != nil {
		structMap["end_date"] = a.EndDate
	}
	if a.SubtotalInCents != nil {
		structMap["subtotal_in_cents"] = a.SubtotalInCents
	}
	if a.TotalTaxInCents != nil {
		structMap["total_tax_in_cents"] = a.TotalTaxInCents
	}
	if a.TotalDiscountInCents != nil {
		structMap["total_discount_in_cents"] = a.TotalDiscountInCents
	}
	if a.TotalInCents != nil {
		structMap["total_in_cents"] = a.TotalInCents
	}
	if a.Direction != nil {
		structMap["direction"] = a.Direction
	}
	if a.ProrationScheme != nil {
		structMap["proration_scheme"] = a.ProrationScheme
	}
	if a.LineItems != nil {
		structMap["line_items"] = a.LineItems
	}
	if a.AccrueCharge != nil {
		structMap["accrue_charge"] = a.AccrueCharge
	}
	if a.Allocations != nil {
		structMap["allocations"] = a.Allocations
	}
	if a.PeriodType != nil {
		structMap["period_type"] = a.PeriodType
	}
	if a.ExistingBalanceInCents != nil {
		structMap["existing_balance_in_cents"] = a.ExistingBalanceInCents
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreview.
// It customizes the JSON unmarshaling process for AllocationPreview objects.
func (a *AllocationPreview) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StartDate              *string                     `json:"start_date,omitempty"`
		EndDate                *string                     `json:"end_date,omitempty"`
		SubtotalInCents        *int64                      `json:"subtotal_in_cents,omitempty"`
		TotalTaxInCents        *int64                      `json:"total_tax_in_cents,omitempty"`
		TotalDiscountInCents   *int64                      `json:"total_discount_in_cents,omitempty"`
		TotalInCents           *int64                      `json:"total_in_cents,omitempty"`
		Direction              *string                     `json:"direction,omitempty"`
		ProrationScheme        *string                     `json:"proration_scheme,omitempty"`
		LineItems              []AllocationPreviewLineItem `json:"line_items,omitempty"`
		AccrueCharge           *bool                       `json:"accrue_charge,omitempty"`
		Allocations            []AllocationPreviewItem     `json:"allocations,omitempty"`
		PeriodType             *string                     `json:"period_type,omitempty"`
		ExistingBalanceInCents *int64                      `json:"existing_balance_in_cents,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.StartDate = temp.StartDate
	a.EndDate = temp.EndDate
	a.SubtotalInCents = temp.SubtotalInCents
	a.TotalTaxInCents = temp.TotalTaxInCents
	a.TotalDiscountInCents = temp.TotalDiscountInCents
	a.TotalInCents = temp.TotalInCents
	a.Direction = temp.Direction
	a.ProrationScheme = temp.ProrationScheme
	a.LineItems = temp.LineItems
	a.AccrueCharge = temp.AccrueCharge
	a.Allocations = temp.Allocations
	a.PeriodType = temp.PeriodType
	a.ExistingBalanceInCents = temp.ExistingBalanceInCents
	return nil
}
