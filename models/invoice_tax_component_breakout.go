/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// InvoiceTaxComponentBreakout represents a InvoiceTaxComponentBreakout struct.
type InvoiceTaxComponentBreakout struct {
    TaxRuleId            *int           `json:"tax_rule_id,omitempty"`
    Percentage           *string        `json:"percentage,omitempty"`
    CountryCode          *string        `json:"country_code,omitempty"`
    SubdivisionCode      *string        `json:"subdivision_code,omitempty"`
    TaxAmount            *string        `json:"tax_amount,omitempty"`
    TaxableAmount        *string        `json:"taxable_amount,omitempty"`
    TaxExemptAmount      *string        `json:"tax_exempt_amount,omitempty"`
    NonTaxableAmount     *string        `json:"non_taxable_amount,omitempty"`
    TaxName              *string        `json:"tax_name,omitempty"`
    TaxType              *string        `json:"tax_type,omitempty"`
    RateType             *string        `json:"rate_type,omitempty"`
    TaxAuthorityType     *int           `json:"tax_authority_type,omitempty"`
    StateAssignedNo      *string        `json:"state_assigned_no,omitempty"`
    TaxSubType           *string        `json:"tax_sub_type,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceTaxComponentBreakout.
// It customizes the JSON marshaling process for InvoiceTaxComponentBreakout objects.
func (i InvoiceTaxComponentBreakout) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceTaxComponentBreakout object to a map representation for JSON marshaling.
func (i InvoiceTaxComponentBreakout) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.TaxRuleId != nil {
        structMap["tax_rule_id"] = i.TaxRuleId
    }
    if i.Percentage != nil {
        structMap["percentage"] = i.Percentage
    }
    if i.CountryCode != nil {
        structMap["country_code"] = i.CountryCode
    }
    if i.SubdivisionCode != nil {
        structMap["subdivision_code"] = i.SubdivisionCode
    }
    if i.TaxAmount != nil {
        structMap["tax_amount"] = i.TaxAmount
    }
    if i.TaxableAmount != nil {
        structMap["taxable_amount"] = i.TaxableAmount
    }
    if i.TaxExemptAmount != nil {
        structMap["tax_exempt_amount"] = i.TaxExemptAmount
    }
    if i.NonTaxableAmount != nil {
        structMap["non_taxable_amount"] = i.NonTaxableAmount
    }
    if i.TaxName != nil {
        structMap["tax_name"] = i.TaxName
    }
    if i.TaxType != nil {
        structMap["tax_type"] = i.TaxType
    }
    if i.RateType != nil {
        structMap["rate_type"] = i.RateType
    }
    if i.TaxAuthorityType != nil {
        structMap["tax_authority_type"] = i.TaxAuthorityType
    }
    if i.StateAssignedNo != nil {
        structMap["state_assigned_no"] = i.StateAssignedNo
    }
    if i.TaxSubType != nil {
        structMap["tax_sub_type"] = i.TaxSubType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceTaxComponentBreakout.
// It customizes the JSON unmarshaling process for InvoiceTaxComponentBreakout objects.
func (i *InvoiceTaxComponentBreakout) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceTaxComponentBreakout
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "tax_rule_id", "percentage", "country_code", "subdivision_code", "tax_amount", "taxable_amount", "tax_exempt_amount", "non_taxable_amount", "tax_name", "tax_type", "rate_type", "tax_authority_type", "state_assigned_no", "tax_sub_type")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.TaxRuleId = temp.TaxRuleId
    i.Percentage = temp.Percentage
    i.CountryCode = temp.CountryCode
    i.SubdivisionCode = temp.SubdivisionCode
    i.TaxAmount = temp.TaxAmount
    i.TaxableAmount = temp.TaxableAmount
    i.TaxExemptAmount = temp.TaxExemptAmount
    i.NonTaxableAmount = temp.NonTaxableAmount
    i.TaxName = temp.TaxName
    i.TaxType = temp.TaxType
    i.RateType = temp.RateType
    i.TaxAuthorityType = temp.TaxAuthorityType
    i.StateAssignedNo = temp.StateAssignedNo
    i.TaxSubType = temp.TaxSubType
    return nil
}

// tempInvoiceTaxComponentBreakout is a temporary struct used for validating the fields of InvoiceTaxComponentBreakout.
type tempInvoiceTaxComponentBreakout  struct {
    TaxRuleId        *int    `json:"tax_rule_id,omitempty"`
    Percentage       *string `json:"percentage,omitempty"`
    CountryCode      *string `json:"country_code,omitempty"`
    SubdivisionCode  *string `json:"subdivision_code,omitempty"`
    TaxAmount        *string `json:"tax_amount,omitempty"`
    TaxableAmount    *string `json:"taxable_amount,omitempty"`
    TaxExemptAmount  *string `json:"tax_exempt_amount,omitempty"`
    NonTaxableAmount *string `json:"non_taxable_amount,omitempty"`
    TaxName          *string `json:"tax_name,omitempty"`
    TaxType          *string `json:"tax_type,omitempty"`
    RateType         *string `json:"rate_type,omitempty"`
    TaxAuthorityType *int    `json:"tax_authority_type,omitempty"`
    StateAssignedNo  *string `json:"state_assigned_no,omitempty"`
    TaxSubType       *string `json:"tax_sub_type,omitempty"`
}
