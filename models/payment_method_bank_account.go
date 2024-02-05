package models

import (
    "encoding/json"
)

// PaymentMethodBankAccount represents a PaymentMethodBankAccount struct.
type PaymentMethodBankAccount struct {
    MaskedAccountNumber string                    `json:"masked_account_number"`
    MaskedRoutingNumber string                    `json:"masked_routing_number"`
    Type                InvoiceEventPaymentMethod `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodBankAccount.
// It customizes the JSON marshaling process for PaymentMethodBankAccount objects.
func (p *PaymentMethodBankAccount) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodBankAccount object to a map representation for JSON marshaling.
func (p *PaymentMethodBankAccount) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["masked_account_number"] = p.MaskedAccountNumber
    structMap["masked_routing_number"] = p.MaskedRoutingNumber
    structMap["type"] = p.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodBankAccount.
// It customizes the JSON unmarshaling process for PaymentMethodBankAccount objects.
func (p *PaymentMethodBankAccount) UnmarshalJSON(input []byte) error {
    temp := &struct {
        MaskedAccountNumber string                    `json:"masked_account_number"`
        MaskedRoutingNumber string                    `json:"masked_routing_number"`
        Type                InvoiceEventPaymentMethod `json:"type"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.MaskedAccountNumber = temp.MaskedAccountNumber
    p.MaskedRoutingNumber = temp.MaskedRoutingNumber
    p.Type = temp.Type
    return nil
}
