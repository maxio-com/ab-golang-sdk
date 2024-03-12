package models

import (
	"encoding/json"
	"errors"
	"strings"
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
	var temp paymentMethodPaypal
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	p.Email = *temp.Email
	p.Type = *temp.Type
	return nil
}

// TODO
type paymentMethodPaypal struct {
	Email *string                    `json:"email"`
	Type  *InvoiceEventPaymentMethod `json:"type"`
}

func (p *paymentMethodPaypal) validate() error {
	var errs []string
	if p.Email == nil {
		errs = append(errs, "required field `email` is missing for type `Payment Method Paypal`")
	}
	if p.Type == nil {
		errs = append(errs, "required field `type` is missing for type `Payment Method Paypal`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
