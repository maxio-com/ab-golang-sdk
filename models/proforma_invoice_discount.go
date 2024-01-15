package models

import (
    "encoding/json"
)

// ProformaInvoiceDiscount represents a ProformaInvoiceDiscount struct.
type ProformaInvoiceDiscount struct {
    Title             *string                           `json:"title,omitempty"`
    SourceType        *string                           `json:"source_type,omitempty"`
    DiscountType      *string                           `json:"discount_type,omitempty"`
    EligibleAmount    *string                           `json:"eligible_amount,omitempty"`
    DiscountAmount    *string                           `json:"discount_amount,omitempty"`
    LineItemBreakouts []ProformaInvoiceDiscountBreakout `json:"line_item_breakouts,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceDiscount.
// It customizes the JSON marshaling process for ProformaInvoiceDiscount objects.
func (p *ProformaInvoiceDiscount) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProformaInvoiceDiscount object to a map representation for JSON marshaling.
func (p *ProformaInvoiceDiscount) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.Title != nil {
        structMap["title"] = p.Title
    }
    if p.SourceType != nil {
        structMap["source_type"] = p.SourceType
    }
    if p.DiscountType != nil {
        structMap["discount_type"] = p.DiscountType
    }
    if p.EligibleAmount != nil {
        structMap["eligible_amount"] = p.EligibleAmount
    }
    if p.DiscountAmount != nil {
        structMap["discount_amount"] = p.DiscountAmount
    }
    if p.LineItemBreakouts != nil {
        structMap["line_item_breakouts"] = p.LineItemBreakouts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoiceDiscount.
// It customizes the JSON unmarshaling process for ProformaInvoiceDiscount objects.
func (p *ProformaInvoiceDiscount) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Title             *string                           `json:"title,omitempty"`
        SourceType        *string                           `json:"source_type,omitempty"`
        DiscountType      *string                           `json:"discount_type,omitempty"`
        EligibleAmount    *string                           `json:"eligible_amount,omitempty"`
        DiscountAmount    *string                           `json:"discount_amount,omitempty"`
        LineItemBreakouts []ProformaInvoiceDiscountBreakout `json:"line_item_breakouts,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Title = temp.Title
    p.SourceType = temp.SourceType
    p.DiscountType = temp.DiscountType
    p.EligibleAmount = temp.EligibleAmount
    p.DiscountAmount = temp.DiscountAmount
    p.LineItemBreakouts = temp.LineItemBreakouts
    return nil
}
