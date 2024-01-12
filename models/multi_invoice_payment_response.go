package models

import (
    "encoding/json"
)

// MultiInvoicePaymentResponse represents a MultiInvoicePaymentResponse struct.
type MultiInvoicePaymentResponse struct {
    Payment MultiInvoicePayment `json:"payment"`
}

// MarshalJSON implements the json.Marshaler interface for MultiInvoicePaymentResponse.
// It customizes the JSON marshaling process for MultiInvoicePaymentResponse objects.
func (m *MultiInvoicePaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MultiInvoicePaymentResponse object to a map representation for JSON marshaling.
func (m *MultiInvoicePaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["payment"] = m.Payment
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MultiInvoicePaymentResponse.
// It customizes the JSON unmarshaling process for MultiInvoicePaymentResponse objects.
func (m *MultiInvoicePaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Payment MultiInvoicePayment `json:"payment"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    m.Payment = temp.Payment
    return nil
}
