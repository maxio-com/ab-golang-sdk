package models

import (
    "encoding/json"
)

// CreatePayment represents a CreatePayment struct.
type CreatePayment struct {
    Amount         string                   `json:"amount"`
    Memo           string                   `json:"memo"`
    PaymentDetails string                   `json:"payment_details"`
    // The type of payment method used. Defaults to other.
    PaymentMethod  InvoicePaymentMethodType `json:"payment_method"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePayment.
// It customizes the JSON marshaling process for CreatePayment objects.
func (c *CreatePayment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePayment object to a map representation for JSON marshaling.
func (c *CreatePayment) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["amount"] = c.Amount
    structMap["memo"] = c.Memo
    structMap["payment_details"] = c.PaymentDetails
    structMap["payment_method"] = c.PaymentMethod
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePayment.
// It customizes the JSON unmarshaling process for CreatePayment objects.
func (c *CreatePayment) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Amount         string                   `json:"amount"`
        Memo           string                   `json:"memo"`
        PaymentDetails string                   `json:"payment_details"`
        PaymentMethod  InvoicePaymentMethodType `json:"payment_method"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Amount = temp.Amount
    c.Memo = temp.Memo
    c.PaymentDetails = temp.PaymentDetails
    c.PaymentMethod = temp.PaymentMethod
    return nil
}
