package models

import (
    "encoding/json"
)

// PaymentResponse represents a PaymentResponse struct.
type PaymentResponse struct {
    PaidInvoices []Payment          `json:"paid_invoices,omitempty"`
    Prepayment   *InvoicePrePayment `json:"prepayment,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentResponse.
// It customizes the JSON marshaling process for PaymentResponse objects.
func (p *PaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentResponse object to a map representation for JSON marshaling.
func (p *PaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.PaidInvoices != nil {
        structMap["paid_invoices"] = p.PaidInvoices
    }
    if p.Prepayment != nil {
        structMap["prepayment"] = p.Prepayment
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentResponse.
// It customizes the JSON unmarshaling process for PaymentResponse objects.
func (p *PaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PaidInvoices []Payment          `json:"paid_invoices,omitempty"`
        Prepayment   *InvoicePrePayment `json:"prepayment,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.PaidInvoices = temp.PaidInvoices
    p.Prepayment = temp.Prepayment
    return nil
}
