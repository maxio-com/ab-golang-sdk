package models

import (
    "encoding/json"
)

// ProformaInvoiceTaxBreakout represents a ProformaInvoiceTaxBreakout struct.
type ProformaInvoiceTaxBreakout struct {
    TaxableAmount *string `json:"taxable_amount,omitempty"`
    TaxAmount     *string `json:"tax_amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceTaxBreakout.
// It customizes the JSON marshaling process for ProformaInvoiceTaxBreakout objects.
func (p *ProformaInvoiceTaxBreakout) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProformaInvoiceTaxBreakout object to a map representation for JSON marshaling.
func (p *ProformaInvoiceTaxBreakout) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.TaxableAmount != nil {
        structMap["taxable_amount"] = p.TaxableAmount
    }
    if p.TaxAmount != nil {
        structMap["tax_amount"] = p.TaxAmount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoiceTaxBreakout.
// It customizes the JSON unmarshaling process for ProformaInvoiceTaxBreakout objects.
func (p *ProformaInvoiceTaxBreakout) UnmarshalJSON(input []byte) error {
    temp := &struct {
        TaxableAmount *string `json:"taxable_amount,omitempty"`
        TaxAmount     *string `json:"tax_amount,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.TaxableAmount = temp.TaxableAmount
    p.TaxAmount = temp.TaxAmount
    return nil
}
