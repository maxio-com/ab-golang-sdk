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

// BillingManifest represents a BillingManifest struct.
type BillingManifest struct {
    LineItems              []BillingManifestItem  `json:"line_items,omitempty"`
    TotalInCents           *int64                 `json:"total_in_cents,omitempty"`
    TotalDiscountInCents   *int64                 `json:"total_discount_in_cents,omitempty"`
    TotalTaxInCents        *int64                 `json:"total_tax_in_cents,omitempty"`
    SubtotalInCents        *int64                 `json:"subtotal_in_cents,omitempty"`
    StartDate              Optional[time.Time]    `json:"start_date"`
    EndDate                Optional[time.Time]    `json:"end_date"`
    PeriodType             Optional[string]       `json:"period_type"`
    ExistingBalanceInCents *int64                 `json:"existing_balance_in_cents,omitempty"`
    AdditionalProperties   map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BillingManifest.
// It customizes the JSON marshaling process for BillingManifest objects.
func (b BillingManifest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "line_items", "total_in_cents", "total_discount_in_cents", "total_tax_in_cents", "subtotal_in_cents", "start_date", "end_date", "period_type", "existing_balance_in_cents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BillingManifest object to a map representation for JSON marshaling.
func (b BillingManifest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
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
    if b.StartDate.IsValueSet() {
        var StartDateVal *string = nil
        if b.StartDate.Value() != nil {
            val := b.StartDate.Value().Format(time.RFC3339)
            StartDateVal = &val
        }
        if b.StartDate.Value() != nil {
            structMap["start_date"] = StartDateVal
        } else {
            structMap["start_date"] = nil
        }
    }
    if b.EndDate.IsValueSet() {
        var EndDateVal *string = nil
        if b.EndDate.Value() != nil {
            val := b.EndDate.Value().Format(time.RFC3339)
            EndDateVal = &val
        }
        if b.EndDate.Value() != nil {
            structMap["end_date"] = EndDateVal
        } else {
            structMap["end_date"] = nil
        }
    }
    if b.PeriodType.IsValueSet() {
        if b.PeriodType.Value() != nil {
            structMap["period_type"] = b.PeriodType.Value()
        } else {
            structMap["period_type"] = nil
        }
    }
    if b.ExistingBalanceInCents != nil {
        structMap["existing_balance_in_cents"] = b.ExistingBalanceInCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BillingManifest.
// It customizes the JSON unmarshaling process for BillingManifest objects.
func (b *BillingManifest) UnmarshalJSON(input []byte) error {
    var temp tempBillingManifest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "line_items", "total_in_cents", "total_discount_in_cents", "total_tax_in_cents", "subtotal_in_cents", "start_date", "end_date", "period_type", "existing_balance_in_cents")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.LineItems = temp.LineItems
    b.TotalInCents = temp.TotalInCents
    b.TotalDiscountInCents = temp.TotalDiscountInCents
    b.TotalTaxInCents = temp.TotalTaxInCents
    b.SubtotalInCents = temp.SubtotalInCents
    b.StartDate.ShouldSetValue(temp.StartDate.IsValueSet())
    if temp.StartDate.Value() != nil {
        StartDateVal, err := time.Parse(time.RFC3339, (*temp.StartDate.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse start_date as % s format.", time.RFC3339)
        }
        b.StartDate.SetValue(&StartDateVal)
    }
    b.EndDate.ShouldSetValue(temp.EndDate.IsValueSet())
    if temp.EndDate.Value() != nil {
        EndDateVal, err := time.Parse(time.RFC3339, (*temp.EndDate.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse end_date as % s format.", time.RFC3339)
        }
        b.EndDate.SetValue(&EndDateVal)
    }
    b.PeriodType = temp.PeriodType
    b.ExistingBalanceInCents = temp.ExistingBalanceInCents
    return nil
}

// tempBillingManifest is a temporary struct used for validating the fields of BillingManifest.
type tempBillingManifest  struct {
    LineItems              []BillingManifestItem `json:"line_items,omitempty"`
    TotalInCents           *int64                `json:"total_in_cents,omitempty"`
    TotalDiscountInCents   *int64                `json:"total_discount_in_cents,omitempty"`
    TotalTaxInCents        *int64                `json:"total_tax_in_cents,omitempty"`
    SubtotalInCents        *int64                `json:"subtotal_in_cents,omitempty"`
    StartDate              Optional[string]      `json:"start_date"`
    EndDate                Optional[string]      `json:"end_date"`
    PeriodType             Optional[string]      `json:"period_type"`
    ExistingBalanceInCents *int64                `json:"existing_balance_in_cents,omitempty"`
}
