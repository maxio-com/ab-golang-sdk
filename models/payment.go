package models

import (
	"encoding/json"
)

// Payment represents a Payment struct.
type Payment struct {
	// The uid of the paid invoice
	InvoiceUid *string `json:"invoice_uid,omitempty"`
	// The current status of the invoice. See [Invoice Statuses](https://chargify.zendesk.com/hc/en-us/articles/4407737494171#line-item-breakdowns) for more.
	Status *StatusEnum `json:"status,omitempty"`
	// The remaining due amount on the invoice
	DueAmount *string `json:"due_amount,omitempty"`
	// The total amount paid on this invoice (including any prior payments)
	PaidAmount *string `json:"paid_amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Payment.
// It customizes the JSON marshaling process for Payment objects.
func (p *Payment) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the Payment object to a map representation for JSON marshaling.
func (p *Payment) toMap() map[string]any {
	structMap := make(map[string]any)
	if p.InvoiceUid != nil {
		structMap["invoice_uid"] = p.InvoiceUid
	}
	if p.Status != nil {
		structMap["status"] = p.Status
	}
	if p.DueAmount != nil {
		structMap["due_amount"] = p.DueAmount
	}
	if p.PaidAmount != nil {
		structMap["paid_amount"] = p.PaidAmount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Payment.
// It customizes the JSON unmarshaling process for Payment objects.
func (p *Payment) UnmarshalJSON(input []byte) error {
	temp := &struct {
		InvoiceUid *string     `json:"invoice_uid,omitempty"`
		Status     *StatusEnum `json:"status,omitempty"`
		DueAmount  *string     `json:"due_amount,omitempty"`
		PaidAmount *string     `json:"paid_amount,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.InvoiceUid = temp.InvoiceUid
	p.Status = temp.Status
	p.DueAmount = temp.DueAmount
	p.PaidAmount = temp.PaidAmount
	return nil
}
