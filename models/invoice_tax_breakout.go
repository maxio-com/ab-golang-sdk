package models

import (
	"encoding/json"
)

// InvoiceTaxBreakout represents a InvoiceTaxBreakout struct.
type InvoiceTaxBreakout struct {
	Uid             *string `json:"uid,omitempty"`
	TaxableAmount   *string `json:"taxable_amount,omitempty"`
	TaxAmount       *string `json:"tax_amount,omitempty"`
	TaxExemptAmount *string `json:"tax_exempt_amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceTaxBreakout.
// It customizes the JSON marshaling process for InvoiceTaxBreakout objects.
func (i *InvoiceTaxBreakout) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceTaxBreakout object to a map representation for JSON marshaling.
func (i *InvoiceTaxBreakout) toMap() map[string]any {
	structMap := make(map[string]any)
	if i.Uid != nil {
		structMap["uid"] = i.Uid
	}
	if i.TaxableAmount != nil {
		structMap["taxable_amount"] = i.TaxableAmount
	}
	if i.TaxAmount != nil {
		structMap["tax_amount"] = i.TaxAmount
	}
	if i.TaxExemptAmount != nil {
		structMap["tax_exempt_amount"] = i.TaxExemptAmount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceTaxBreakout.
// It customizes the JSON unmarshaling process for InvoiceTaxBreakout objects.
func (i *InvoiceTaxBreakout) UnmarshalJSON(input []byte) error {
	var temp invoiceTaxBreakout
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	i.Uid = temp.Uid
	i.TaxableAmount = temp.TaxableAmount
	i.TaxAmount = temp.TaxAmount
	i.TaxExemptAmount = temp.TaxExemptAmount
	return nil
}

// TODO
type invoiceTaxBreakout struct {
	Uid             *string `json:"uid,omitempty"`
	TaxableAmount   *string `json:"taxable_amount,omitempty"`
	TaxAmount       *string `json:"tax_amount,omitempty"`
	TaxExemptAmount *string `json:"tax_exempt_amount,omitempty"`
}
