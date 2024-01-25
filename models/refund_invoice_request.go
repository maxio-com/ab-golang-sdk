package models

import (
	"encoding/json"
)

// RefundInvoiceRequest represents a RefundInvoiceRequest struct.
type RefundInvoiceRequest struct {
	Refund Refund `json:"refund"`
}

// MarshalJSON implements the json.Marshaler interface for RefundInvoiceRequest.
// It customizes the JSON marshaling process for RefundInvoiceRequest objects.
func (r *RefundInvoiceRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RefundInvoiceRequest object to a map representation for JSON marshaling.
func (r *RefundInvoiceRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["refund"] = r.Refund
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundInvoiceRequest.
// It customizes the JSON unmarshaling process for RefundInvoiceRequest objects.
func (r *RefundInvoiceRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Refund Refund `json:"refund"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Refund = temp.Refund
	return nil
}
