package models

import (
    "encoding/json"
)

// ProformaInvoiceDiscount represents a ProformaInvoiceDiscount struct.
type ProformaInvoiceDiscount struct {
    Uid                  *string                            `json:"uid,omitempty"`
    Title                *string                            `json:"title,omitempty"`
    Code                 *string                            `json:"code,omitempty"`
    SourceType           *ProformaInvoiceDiscountSourceType `json:"source_type,omitempty"`
    DiscountType         *InvoiceDiscountType               `json:"discount_type,omitempty"`
    EligibleAmount       *string                            `json:"eligible_amount,omitempty"`
    DiscountAmount       *string                            `json:"discount_amount,omitempty"`
    LineItemBreakouts    []InvoiceDiscountBreakout          `json:"line_item_breakouts,omitempty"`
    AdditionalProperties map[string]any                     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceDiscount.
// It customizes the JSON marshaling process for ProformaInvoiceDiscount objects.
func (p ProformaInvoiceDiscount) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProformaInvoiceDiscount object to a map representation for JSON marshaling.
func (p ProformaInvoiceDiscount) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Uid != nil {
        structMap["uid"] = p.Uid
    }
    if p.Title != nil {
        structMap["title"] = p.Title
    }
    if p.Code != nil {
        structMap["code"] = p.Code
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
    var temp proformaInvoiceDiscount
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "title", "code", "source_type", "discount_type", "eligible_amount", "discount_amount", "line_item_breakouts")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Uid = temp.Uid
    p.Title = temp.Title
    p.Code = temp.Code
    p.SourceType = temp.SourceType
    p.DiscountType = temp.DiscountType
    p.EligibleAmount = temp.EligibleAmount
    p.DiscountAmount = temp.DiscountAmount
    p.LineItemBreakouts = temp.LineItemBreakouts
    return nil
}

// TODO
type proformaInvoiceDiscount  struct {
    Uid               *string                            `json:"uid,omitempty"`
    Title             *string                            `json:"title,omitempty"`
    Code              *string                            `json:"code,omitempty"`
    SourceType        *ProformaInvoiceDiscountSourceType `json:"source_type,omitempty"`
    DiscountType      *InvoiceDiscountType               `json:"discount_type,omitempty"`
    EligibleAmount    *string                            `json:"eligible_amount,omitempty"`
    DiscountAmount    *string                            `json:"discount_amount,omitempty"`
    LineItemBreakouts []InvoiceDiscountBreakout          `json:"line_item_breakouts,omitempty"`
}
