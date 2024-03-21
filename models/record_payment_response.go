package models

import (
	"encoding/json"
)

// RecordPaymentResponse represents a RecordPaymentResponse struct.
type RecordPaymentResponse struct {
	PaidInvoices []PaidInvoice                             `json:"paid_invoices,omitempty"`
	Prepayment   Optional[RecordPaymentResponsePrepayment] `json:"prepayment"`
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
		if r.Prepayment.Value() != nil {
			structMap["prepayment"] = r.Prepayment.Value().toMap()
		} else {
			structMap["prepayment"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RecordPaymentResponse.
// It customizes the JSON unmarshaling process for RecordPaymentResponse objects.
func (r *RecordPaymentResponse) UnmarshalJSON(input []byte) error {
	var temp recordPaymentResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	r.PaidInvoices = temp.PaidInvoices
	r.Prepayment = temp.Prepayment
	return nil
}

// TODO
type recordPaymentResponse struct {
	PaidInvoices []PaidInvoice                             `json:"paid_invoices,omitempty"`
	Prepayment   Optional[RecordPaymentResponsePrepayment] `json:"prepayment"`
}
