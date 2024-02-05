package models

import (
    "encoding/json"
)

// ProformaInvoiceTax represents a ProformaInvoiceTax struct.
type ProformaInvoiceTax struct {
    Uid               *string                      `json:"uid,omitempty"`
    Title             *string                      `json:"title,omitempty"`
    SourceType        *string                      `json:"source_type,omitempty"`
    Percentage        *string                      `json:"percentage,omitempty"`
    TaxableAmount     *string                      `json:"taxable_amount,omitempty"`
    TaxAmount         *string                      `json:"tax_amount,omitempty"`
    LineItemBreakouts []ProformaInvoiceTaxBreakout `json:"line_item_breakouts,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceTax.
// It customizes the JSON marshaling process for ProformaInvoiceTax objects.
func (p *ProformaInvoiceTax) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProformaInvoiceTax object to a map representation for JSON marshaling.
func (p *ProformaInvoiceTax) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.Uid != nil {
        structMap["uid"] = p.Uid
    }
    if p.Title != nil {
        structMap["title"] = p.Title
    }
    if p.SourceType != nil {
        structMap["source_type"] = p.SourceType
    }
    if p.Percentage != nil {
        structMap["percentage"] = p.Percentage
    }
    if p.TaxableAmount != nil {
        structMap["taxable_amount"] = p.TaxableAmount
    }
    if p.TaxAmount != nil {
        structMap["tax_amount"] = p.TaxAmount
    }
    if p.LineItemBreakouts != nil {
        structMap["line_item_breakouts"] = p.LineItemBreakouts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoiceTax.
// It customizes the JSON unmarshaling process for ProformaInvoiceTax objects.
func (p *ProformaInvoiceTax) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Uid               *string                      `json:"uid,omitempty"`
        Title             *string                      `json:"title,omitempty"`
        SourceType        *string                      `json:"source_type,omitempty"`
        Percentage        *string                      `json:"percentage,omitempty"`
        TaxableAmount     *string                      `json:"taxable_amount,omitempty"`
        TaxAmount         *string                      `json:"tax_amount,omitempty"`
        LineItemBreakouts []ProformaInvoiceTaxBreakout `json:"line_item_breakouts,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Uid = temp.Uid
    p.Title = temp.Title
    p.SourceType = temp.SourceType
    p.Percentage = temp.Percentage
    p.TaxableAmount = temp.TaxableAmount
    p.TaxAmount = temp.TaxAmount
    p.LineItemBreakouts = temp.LineItemBreakouts
    return nil
}
