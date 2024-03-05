package models

import (
    "encoding/json"
)

// RecordPaymentResponse represents a RecordPaymentResponse struct.
type RecordPaymentResponse struct {
    PaidInvoices []PaidInvoice               `json:"paid_invoices,omitempty"`
    Prepayment   Optional[InvoicePrePayment] `json:"prepayment"`
}

// MarshalJSON implements the json.Marshaler interface for RecordPaymentResponse.
// It customizes the JSON marshaling process for RecordPaymentResponse objects.
func (r *RecordPaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RecordPaymentResponse object to a map representation for JSON marshaling.
func (r *RecordPaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if r.PaidInvoices != nil {
        structMap["paid_invoices"] = r.PaidInvoices
    }
    if r.Prepayment.IsValueSet() {
        structMap["prepayment"] = r.Prepayment.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RecordPaymentResponse.
// It customizes the JSON unmarshaling process for RecordPaymentResponse objects.
func (r *RecordPaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PaidInvoices []PaidInvoice               `json:"paid_invoices,omitempty"`
        Prepayment   Optional[InvoicePrePayment] `json:"prepayment"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    r.PaidInvoices = temp.PaidInvoices
    r.Prepayment = temp.Prepayment
    return nil
}
