package models

import (
	"encoding/json"
)

// ProformaInvoiceDiscountBreakout represents a ProformaInvoiceDiscountBreakout struct.
type ProformaInvoiceDiscountBreakout struct {
	EligibleAmount *string `json:"eligible_amount,omitempty"`
	DiscountAmount *string `json:"discount_amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceDiscountBreakout.
// It customizes the JSON marshaling process for ProformaInvoiceDiscountBreakout objects.
func (p *ProformaInvoiceDiscountBreakout) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the ProformaInvoiceDiscountBreakout object to a map representation for JSON marshaling.
func (p *ProformaInvoiceDiscountBreakout) toMap() map[string]any {
	structMap := make(map[string]any)
	if p.EligibleAmount != nil {
		structMap["eligible_amount"] = p.EligibleAmount
	}
	if p.DiscountAmount != nil {
		structMap["discount_amount"] = p.DiscountAmount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoiceDiscountBreakout.
// It customizes the JSON unmarshaling process for ProformaInvoiceDiscountBreakout objects.
func (p *ProformaInvoiceDiscountBreakout) UnmarshalJSON(input []byte) error {
	temp := &struct {
		EligibleAmount *string `json:"eligible_amount,omitempty"`
		DiscountAmount *string `json:"discount_amount,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.EligibleAmount = temp.EligibleAmount
	p.DiscountAmount = temp.DiscountAmount
	return nil
}
