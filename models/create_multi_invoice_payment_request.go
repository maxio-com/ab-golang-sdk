package models

import (
    "encoding/json"
)

// CreateMultiInvoicePaymentRequest represents a CreateMultiInvoicePaymentRequest struct.
type CreateMultiInvoicePaymentRequest struct {
    Payment CreateMultiInvoicePayment `json:"payment"`
}

// MarshalJSON implements the json.Marshaler interface for CreateMultiInvoicePaymentRequest.
// It customizes the JSON marshaling process for CreateMultiInvoicePaymentRequest objects.
func (c *CreateMultiInvoicePaymentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMultiInvoicePaymentRequest object to a map representation for JSON marshaling.
func (c *CreateMultiInvoicePaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["payment"] = c.Payment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMultiInvoicePaymentRequest.
// It customizes the JSON unmarshaling process for CreateMultiInvoicePaymentRequest objects.
func (c *CreateMultiInvoicePaymentRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Payment CreateMultiInvoicePayment `json:"payment"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Payment = temp.Payment
    return nil
}
