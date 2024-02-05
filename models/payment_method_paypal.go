package models

import (
    "encoding/json"
)

// PaymentMethodPaypal represents a PaymentMethodPaypal struct.
type PaymentMethodPaypal struct {
    Email string                    `json:"email"`
    Type  InvoiceEventPaymentMethod `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodPaypal.
// It customizes the JSON marshaling process for PaymentMethodPaypal objects.
func (p *PaymentMethodPaypal) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodPaypal object to a map representation for JSON marshaling.
func (p *PaymentMethodPaypal) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["email"] = p.Email
    structMap["type"] = p.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodPaypal.
// It customizes the JSON unmarshaling process for PaymentMethodPaypal objects.
func (p *PaymentMethodPaypal) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Email string                    `json:"email"`
        Type  InvoiceEventPaymentMethod `json:"type"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Email = temp.Email
    p.Type = temp.Type
    return nil
}
