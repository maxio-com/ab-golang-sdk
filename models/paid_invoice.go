package models

import (
    "encoding/json"
)

// PaidInvoice represents a PaidInvoice struct.
type PaidInvoice struct {
    // The uid of the paid invoice
    InvoiceId            *string        `json:"invoice_id,omitempty"`
    // The current status of the invoice. See [Invoice Statuses](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405078794253-Introduction-to-Invoices#invoice-statuses) for more.
    Status               *InvoiceStatus `json:"status,omitempty"`
    // The remaining due amount on the invoice
    DueAmount            *string        `json:"due_amount,omitempty"`
    // The total amount paid on this invoice (including any prior payments)
    PaidAmount           *string        `json:"paid_amount,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PaidInvoice.
// It customizes the JSON marshaling process for PaidInvoice objects.
func (p PaidInvoice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaidInvoice object to a map representation for JSON marshaling.
func (p PaidInvoice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
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
    var temp paidInvoice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "invoice_id", "status", "due_amount", "paid_amount")
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

// paidInvoice is a temporary struct used for validating the fields of PaidInvoice.
type paidInvoice  struct {
    InvoiceId  *string        `json:"invoice_id,omitempty"`
    Status     *InvoiceStatus `json:"status,omitempty"`
    DueAmount  *string        `json:"due_amount,omitempty"`
    PaidAmount *string        `json:"paid_amount,omitempty"`
}
