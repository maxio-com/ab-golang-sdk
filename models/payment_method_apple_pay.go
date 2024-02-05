package models

import (
    "encoding/json"
)

// PaymentMethodApplePay represents a PaymentMethodApplePay struct.
type PaymentMethodApplePay struct {
    Type InvoiceEventPaymentMethod `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodApplePay.
// It customizes the JSON marshaling process for PaymentMethodApplePay objects.
func (p *PaymentMethodApplePay) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodApplePay object to a map representation for JSON marshaling.
func (p *PaymentMethodApplePay) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["type"] = p.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodApplePay.
// It customizes the JSON unmarshaling process for PaymentMethodApplePay objects.
func (p *PaymentMethodApplePay) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Type InvoiceEventPaymentMethod `json:"type"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Type = temp.Type
    return nil
}
