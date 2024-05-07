package models

import (
    "encoding/json"
)

// ProformaInvoiceTax represents a ProformaInvoiceTax struct.
type ProformaInvoiceTax struct {
    Uid                  *string                       `json:"uid,omitempty"`
    Title                *string                       `json:"title,omitempty"`
    SourceType           *ProformaInvoiceTaxSourceType `json:"source_type,omitempty"`
    Percentage           *string                       `json:"percentage,omitempty"`
    TaxableAmount        *string                       `json:"taxable_amount,omitempty"`
    TaxAmount            *string                       `json:"tax_amount,omitempty"`
    LineItemBreakouts    []InvoiceTaxBreakout          `json:"line_item_breakouts,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceTax.
// It customizes the JSON marshaling process for ProformaInvoiceTax objects.
func (p ProformaInvoiceTax) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProformaInvoiceTax object to a map representation for JSON marshaling.
func (p ProformaInvoiceTax) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
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
    var temp proformaInvoiceTax
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "title", "source_type", "percentage", "taxable_amount", "tax_amount", "line_item_breakouts")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Uid = temp.Uid
    p.Title = temp.Title
    p.SourceType = temp.SourceType
    p.Percentage = temp.Percentage
    p.TaxableAmount = temp.TaxableAmount
    p.TaxAmount = temp.TaxAmount
    p.LineItemBreakouts = temp.LineItemBreakouts
    return nil
}

// proformaInvoiceTax is a temporary struct used for validating the fields of ProformaInvoiceTax.
type proformaInvoiceTax  struct {
    Uid               *string                       `json:"uid,omitempty"`
    Title             *string                       `json:"title,omitempty"`
    SourceType        *ProformaInvoiceTaxSourceType `json:"source_type,omitempty"`
    Percentage        *string                       `json:"percentage,omitempty"`
    TaxableAmount     *string                       `json:"taxable_amount,omitempty"`
    TaxAmount         *string                       `json:"tax_amount,omitempty"`
    LineItemBreakouts []InvoiceTaxBreakout          `json:"line_item_breakouts,omitempty"`
}
