package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PaymentMethodBankAccount represents a PaymentMethodBankAccount struct.
type PaymentMethodBankAccount struct {
    MaskedAccountNumber  string                    `json:"masked_account_number"`
    MaskedRoutingNumber  string                    `json:"masked_routing_number"`
    Type                 InvoiceEventPaymentMethod `json:"type"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodBankAccount.
// It customizes the JSON marshaling process for PaymentMethodBankAccount objects.
func (p PaymentMethodBankAccount) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodBankAccount object to a map representation for JSON marshaling.
func (p PaymentMethodBankAccount) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["masked_account_number"] = p.MaskedAccountNumber
    structMap["masked_routing_number"] = p.MaskedRoutingNumber
    structMap["type"] = p.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodBankAccount.
// It customizes the JSON unmarshaling process for PaymentMethodBankAccount objects.
func (p *PaymentMethodBankAccount) UnmarshalJSON(input []byte) error {
    var temp paymentMethodBankAccount
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "masked_account_number", "masked_routing_number", "type")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.MaskedAccountNumber = *temp.MaskedAccountNumber
    p.MaskedRoutingNumber = *temp.MaskedRoutingNumber
    p.Type = *temp.Type
    return nil
}

// TODO
type paymentMethodBankAccount  struct {
    MaskedAccountNumber *string                    `json:"masked_account_number"`
    MaskedRoutingNumber *string                    `json:"masked_routing_number"`
    Type                *InvoiceEventPaymentMethod `json:"type"`
}

func (p *paymentMethodBankAccount) validate() error {
    var errs []string
    if p.MaskedAccountNumber == nil {
        errs = append(errs, "required field `masked_account_number` is missing for type `Payment Method Bank Account`")
    }
    if p.MaskedRoutingNumber == nil {
        errs = append(errs, "required field `masked_routing_number` is missing for type `Payment Method Bank Account`")
    }
    if p.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Payment Method Bank Account`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
