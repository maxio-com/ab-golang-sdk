/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
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
    EuVat                 *bool                         `json:"eu_vat,omitempty"`
    Type                  *string                       `json:"type,omitempty"`
    TaxExemptAmount       *string                       `json:"tax_exempt_amount,omitempty"`
    AdditionalProperties  map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceTax.
// It customizes the JSON marshaling process for InvoiceTax objects.
func (i InvoiceTax) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceTax object to a map representation for JSON marshaling.
func (i InvoiceTax) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Uid != nil {
        structMap["uid"] = i.Uid
    }
    if i.Title != nil {
        structMap["title"] = i.Title
    }
    if i.Description.IsValueSet() {
        if i.Description.Value() != nil {
            structMap["description"] = i.Description.Value()
        } else {
            structMap["description"] = nil
        }
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
    if i.EuVat != nil {
        structMap["eu_vat"] = i.EuVat
    }
    if i.Type != nil {
        structMap["type"] = i.Type
    }
    if i.TaxExemptAmount != nil {
        structMap["tax_exempt_amount"] = i.TaxExemptAmount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceTax.
// It customizes the JSON unmarshaling process for InvoiceTax objects.
func (i *InvoiceTax) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceTax
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "title", "description", "source_type", "source_id", "percentage", "taxable_amount", "tax_amount", "transaction_id", "line_item_breakouts", "tax_component_breakouts", "eu_vat", "type", "tax_exempt_amount")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
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
    i.EuVat = temp.EuVat
    i.Type = temp.Type
    i.TaxExemptAmount = temp.TaxExemptAmount
    return nil
}

// tempInvoiceTax is a temporary struct used for validating the fields of InvoiceTax.
type tempInvoiceTax  struct {
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
    EuVat                 *bool                         `json:"eu_vat,omitempty"`
    Type                  *string                       `json:"type,omitempty"`
    TaxExemptAmount       *string                       `json:"tax_exempt_amount,omitempty"`
}
