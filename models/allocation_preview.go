package models

import (
    "encoding/json"
    "log"
    "time"
)

// AllocationPreview represents a AllocationPreview struct.
type AllocationPreview struct {
    StartDate              *time.Time                  `json:"start_date,omitempty"`
    EndDate                *time.Time                  `json:"end_date,omitempty"`
    SubtotalInCents        *int64                      `json:"subtotal_in_cents,omitempty"`
    TotalTaxInCents        *int64                      `json:"total_tax_in_cents,omitempty"`
    TotalDiscountInCents   *int64                      `json:"total_discount_in_cents,omitempty"`
    TotalInCents           *int64                      `json:"total_in_cents,omitempty"`
    Direction              *AllocationPreviewDirection `json:"direction,omitempty"`
    ProrationScheme        *string                     `json:"proration_scheme,omitempty"`
    LineItems              []AllocationPreviewLineItem `json:"line_items,omitempty"`
    AccrueCharge           *bool                       `json:"accrue_charge,omitempty"`
    Allocations            []AllocationPreviewItem     `json:"allocations,omitempty"`
    PeriodType             *string                     `json:"period_type,omitempty"`
    // An integer representing the amount of the subscription's current balance
    ExistingBalanceInCents *int64                      `json:"existing_balance_in_cents,omitempty"`
    AdditionalProperties   map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationPreview.
// It customizes the JSON marshaling process for AllocationPreview objects.
func (a AllocationPreview) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationPreview object to a map representation for JSON marshaling.
func (a AllocationPreview) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.StartDate != nil {
        structMap["start_date"] = a.StartDate.Format(time.RFC3339)
    }
    if a.EndDate != nil {
        structMap["end_date"] = a.EndDate.Format(time.RFC3339)
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
    var temp allocationPreview
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "start_date", "end_date", "subtotal_in_cents", "total_tax_in_cents", "total_discount_in_cents", "total_in_cents", "direction", "proration_scheme", "line_items", "accrue_charge", "allocations", "period_type", "existing_balance_in_cents")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    if temp.StartDate != nil {
        StartDateVal, err := time.Parse(time.RFC3339, *temp.StartDate)
        if err != nil {
            log.Fatalf("Cannot Parse start_date as % s format.", time.RFC3339)
        }
        a.StartDate = &StartDateVal
    }
    if temp.EndDate != nil {
        EndDateVal, err := time.Parse(time.RFC3339, *temp.EndDate)
        if err != nil {
            log.Fatalf("Cannot Parse end_date as % s format.", time.RFC3339)
        }
        a.EndDate = &EndDateVal
    }
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

// TODO
type allocationPreview  struct {
    StartDate              *string                     `json:"start_date,omitempty"`
    EndDate                *string                     `json:"end_date,omitempty"`
    SubtotalInCents        *int64                      `json:"subtotal_in_cents,omitempty"`
    TotalTaxInCents        *int64                      `json:"total_tax_in_cents,omitempty"`
    TotalDiscountInCents   *int64                      `json:"total_discount_in_cents,omitempty"`
    TotalInCents           *int64                      `json:"total_in_cents,omitempty"`
    Direction              *AllocationPreviewDirection `json:"direction,omitempty"`
    ProrationScheme        *string                     `json:"proration_scheme,omitempty"`
    LineItems              []AllocationPreviewLineItem `json:"line_items,omitempty"`
    AccrueCharge           *bool                       `json:"accrue_charge,omitempty"`
    Allocations            []AllocationPreviewItem     `json:"allocations,omitempty"`
    PeriodType             *string                     `json:"period_type,omitempty"`
    ExistingBalanceInCents *int64                      `json:"existing_balance_in_cents,omitempty"`
}
