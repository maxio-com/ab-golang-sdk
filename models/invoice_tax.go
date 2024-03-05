package models

import (
    "encoding/json"
)

// InvoiceTax represents a InvoiceTax struct.
type InvoiceTax struct {
    Uid                   *string                       `json:"uid,omitempty"`
    Title                 *string                       `json:"title,omitempty"`
    Description           Optional[string]              `json:"description"`
    SourceType            *ProformaInvoiceTaxSourceType `json:"source_type,omitempty"`
    SourceId              *int                          `json:"source_id,omitempty"`
    Percentage            *string                       `json:"percentage,omitempty"`
    TaxableAmount         *string                       `json:"taxable_amount,omitempty"`
    TaxAmount             *string                       `json:"tax_amount,omitempty"`
    TransactionId         *int                          `json:"transaction_id,omitempty"`
    LineItemBreakouts     []InvoiceTaxBreakout          `json:"line_item_breakouts,omitempty"`
    TaxComponentBreakouts []InvoiceTaxComponentBreakout `json:"tax_component_breakouts,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceTax.
// It customizes the JSON marshaling process for InvoiceTax objects.
func (i *InvoiceTax) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceTax object to a map representation for JSON marshaling.
func (i *InvoiceTax) toMap() map[string]any {
    structMap := make(map[string]any)
    if i.Uid != nil {
        structMap["uid"] = i.Uid
    }
    if i.Title != nil {
        structMap["title"] = i.Title
    }
    if i.Description.IsValueSet() {
        structMap["description"] = i.Description.Value()
    }
    if i.SourceType != nil {
        structMap["source_type"] = i.SourceType
    }
    if i.SourceId != nil {
        structMap["source_id"] = i.SourceId
    }
    if i.Percentage != nil {
        structMap["percentage"] = i.Percentage
    }
    if i.TaxableAmount != nil {
        structMap["taxable_amount"] = i.TaxableAmount
    }
    if i.TaxAmount != nil {
        structMap["tax_amount"] = i.TaxAmount
    }
    if i.TransactionId != nil {
        structMap["transaction_id"] = i.TransactionId
    }
    if i.LineItemBreakouts != nil {
        structMap["line_item_breakouts"] = i.LineItemBreakouts
    }
    if i.TaxComponentBreakouts != nil {
        structMap["tax_component_breakouts"] = i.TaxComponentBreakouts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceTax.
// It customizes the JSON unmarshaling process for InvoiceTax objects.
func (i *InvoiceTax) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Uid                   *string                       `json:"uid,omitempty"`
        Title                 *string                       `json:"title,omitempty"`
        Description           Optional[string]              `json:"description"`
        SourceType            *ProformaInvoiceTaxSourceType `json:"source_type,omitempty"`
        SourceId              *int                          `json:"source_id,omitempty"`
        Percentage            *string                       `json:"percentage,omitempty"`
        TaxableAmount         *string                       `json:"taxable_amount,omitempty"`
        TaxAmount             *string                       `json:"tax_amount,omitempty"`
        TransactionId         *int                          `json:"transaction_id,omitempty"`
        LineItemBreakouts     []InvoiceTaxBreakout          `json:"line_item_breakouts,omitempty"`
        TaxComponentBreakouts []InvoiceTaxComponentBreakout `json:"tax_component_breakouts,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    i.Uid = temp.Uid
    i.Title = temp.Title
    i.Description = temp.Description
    i.SourceType = temp.SourceType
    i.SourceId = temp.SourceId
    i.Percentage = temp.Percentage
    i.TaxableAmount = temp.TaxableAmount
    i.TaxAmount = temp.TaxAmount
    i.TransactionId = temp.TransactionId
    i.LineItemBreakouts = temp.LineItemBreakouts
    i.TaxComponentBreakouts = temp.TaxComponentBreakouts
    return nil
}
