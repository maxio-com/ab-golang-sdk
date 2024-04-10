package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PaymentMethodApplePay represents a PaymentMethodApplePay struct.
type PaymentMethodApplePay struct {
    Type                 InvoiceEventPaymentMethod `json:"type"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodApplePay.
// It customizes the JSON marshaling process for PaymentMethodApplePay objects.
func (p PaymentMethodApplePay) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodApplePay object to a map representation for JSON marshaling.
func (p PaymentMethodApplePay) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["type"] = p.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodApplePay.
// It customizes the JSON unmarshaling process for PaymentMethodApplePay objects.
func (p *PaymentMethodApplePay) UnmarshalJSON(input []byte) error {
    var temp paymentMethodApplePay
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "type")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Type = *temp.Type
    return nil
}

// TODO
type paymentMethodApplePay  struct {
    Type *InvoiceEventPaymentMethod `json:"type"`
}

func (p *paymentMethodApplePay) validate() error {
    var errs []string
    if p.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Payment Method Apple Pay`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
