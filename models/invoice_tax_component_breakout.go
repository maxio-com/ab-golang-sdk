package models

import (
    "encoding/json"
)

// InvoiceTaxComponentBreakout represents a InvoiceTaxComponentBreakout struct.
type InvoiceTaxComponentBreakout struct {
    TaxRuleId       *int    `json:"tax_rule_id,omitempty"`
    Percentage      *string `json:"percentage,omitempty"`
    CountryCode     *string `json:"country_code,omitempty"`
    SubdivisionCode *string `json:"subdivision_code,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceTaxComponentBreakout.
// It customizes the JSON marshaling process for InvoiceTaxComponentBreakout objects.
func (i *InvoiceTaxComponentBreakout) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceTaxComponentBreakout object to a map representation for JSON marshaling.
func (i *InvoiceTaxComponentBreakout) toMap() map[string]any {
    structMap := make(map[string]any)
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceTaxComponentBreakout.
// It customizes the JSON unmarshaling process for InvoiceTaxComponentBreakout objects.
func (i *InvoiceTaxComponentBreakout) UnmarshalJSON(input []byte) error {
    temp := &struct {
        TaxRuleId       *int    `json:"tax_rule_id,omitempty"`
        Percentage      *string `json:"percentage,omitempty"`
        CountryCode     *string `json:"country_code,omitempty"`
        SubdivisionCode *string `json:"subdivision_code,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    i.TaxRuleId = temp.TaxRuleId
    i.Percentage = temp.Percentage
    i.CountryCode = temp.CountryCode
    i.SubdivisionCode = temp.SubdivisionCode
    return nil
}
