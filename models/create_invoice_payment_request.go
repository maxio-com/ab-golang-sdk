package models

import (
    "encoding/json"
)

// CreateInvoicePaymentRequest represents a CreateInvoicePaymentRequest struct.
type CreateInvoicePaymentRequest struct {
    Payment CreateInvoicePayment `json:"payment"`
    // The type of payment to be applied to an Invoice.
    Type    *InvoicePaymentType  `json:"type,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoicePaymentRequest.
// It customizes the JSON marshaling process for CreateInvoicePaymentRequest objects.
func (c *CreateInvoicePaymentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoicePaymentRequest object to a map representation for JSON marshaling.
func (c *CreateInvoicePaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["payment"] = c.Payment
    if c.Type != nil {
        structMap["type"] = c.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoicePaymentRequest.
// It customizes the JSON unmarshaling process for CreateInvoicePaymentRequest objects.
func (c *CreateInvoicePaymentRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Payment CreateInvoicePayment `json:"payment"`
        Type    *InvoicePaymentType  `json:"type,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Payment = temp.Payment
    c.Type = temp.Type
    return nil
}
