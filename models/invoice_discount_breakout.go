package models

import (
	"encoding/json"
)

// InvoiceDiscountBreakout represents a InvoiceDiscountBreakout struct.
type InvoiceDiscountBreakout struct {
	Uid            *string `json:"uid,omitempty"`
	EligibleAmount *string `json:"eligible_amount,omitempty"`
	DiscountAmount *string `json:"discount_amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceDiscountBreakout.
// It customizes the JSON marshaling process for InvoiceDiscountBreakout objects.
func (i *InvoiceDiscountBreakout) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceDiscountBreakout object to a map representation for JSON marshaling.
func (i *InvoiceDiscountBreakout) toMap() map[string]any {
	structMap := make(map[string]any)
	if i.Uid != nil {
		structMap["uid"] = i.Uid
	}
	if i.EligibleAmount != nil {
		structMap["eligible_amount"] = i.EligibleAmount
	}
	if i.DiscountAmount != nil {
		structMap["discount_amount"] = i.DiscountAmount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceDiscountBreakout.
// It customizes the JSON unmarshaling process for InvoiceDiscountBreakout objects.
func (i *InvoiceDiscountBreakout) UnmarshalJSON(input []byte) error {
	var temp invoiceDiscountBreakout
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	i.Uid = temp.Uid
	i.EligibleAmount = temp.EligibleAmount
	i.DiscountAmount = temp.DiscountAmount
	return nil
}

// TODO
type invoiceDiscountBreakout struct {
	Uid            *string `json:"uid,omitempty"`
	EligibleAmount *string `json:"eligible_amount,omitempty"`
	DiscountAmount *string `json:"discount_amount,omitempty"`
}
