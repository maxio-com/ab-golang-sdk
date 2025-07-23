// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// PaidInvoice represents a PaidInvoice struct.
type PaidInvoice struct {
	// The uid of the paid invoice
	InvoiceId *string `json:"invoice_id,omitempty"`
	// The current status of the invoice. See [Invoice Statuses](https://maxio.zendesk.com/hc/en-us/articles/24252287829645-Advanced-Billing-Invoices-Overview#invoice-statuses) for more.
	Status *InvoiceStatus `json:"status,omitempty"`
	// The remaining due amount on the invoice
	DueAmount *string `json:"due_amount,omitempty"`
	// The total amount paid on this invoice (including any prior payments)
	PaidAmount           *string                `json:"paid_amount,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PaidInvoice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PaidInvoice) String() string {
	return fmt.Sprintf(
		"PaidInvoice[InvoiceId=%v, Status=%v, DueAmount=%v, PaidAmount=%v, AdditionalProperties=%v]",
		p.InvoiceId, p.Status, p.DueAmount, p.PaidAmount, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PaidInvoice.
// It customizes the JSON marshaling process for PaidInvoice objects.
func (p PaidInvoice) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"invoice_id", "status", "due_amount", "paid_amount"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PaidInvoice object to a map representation for JSON marshaling.
func (p PaidInvoice) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, p.AdditionalProperties)
	if p.InvoiceId != nil {
		structMap["invoice_id"] = p.InvoiceId
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

// UnmarshalJSON implements the json.Unmarshaler interface for PaidInvoice.
// It customizes the JSON unmarshaling process for PaidInvoice objects.
func (p *PaidInvoice) UnmarshalJSON(input []byte) error {
	var temp tempPaidInvoice
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "invoice_id", "status", "due_amount", "paid_amount")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.InvoiceId = temp.InvoiceId
	p.Status = temp.Status
	p.DueAmount = temp.DueAmount
	p.PaidAmount = temp.PaidAmount
	return nil
}

// tempPaidInvoice is a temporary struct used for validating the fields of PaidInvoice.
type tempPaidInvoice struct {
	InvoiceId  *string        `json:"invoice_id,omitempty"`
	Status     *InvoiceStatus `json:"status,omitempty"`
	DueAmount  *string        `json:"due_amount,omitempty"`
	PaidAmount *string        `json:"paid_amount,omitempty"`
}
