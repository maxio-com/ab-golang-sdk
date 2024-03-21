package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateMultiInvoicePayment represents a CreateMultiInvoicePayment struct.
type CreateMultiInvoicePayment struct {
	// A description to be attached to the payment.
	Memo *string `json:"memo,omitempty"`
	// Additional information related to the payment method (eg. Check #).
	Details *string `json:"details,omitempty"`
	// The type of payment method used. Defaults to other.
	Method *InvoicePaymentMethodType `json:"method,omitempty"`
	// Dollar amount of the sum of the invoices payment (eg. "10.50" => $10.50).
	Amount CreateMultiInvoicePaymentAmount `json:"amount"`
	// Date reflecting when the payment was received from a customer. Must be in the past.
	ReceivedOn   *string                           `json:"received_on,omitempty"`
	Applications []CreateInvoicePaymentApplication `json:"applications"`
}

// MarshalJSON implements the json.Marshaler interface for CreateMultiInvoicePayment.
// It customizes the JSON marshaling process for CreateMultiInvoicePayment objects.
func (c *CreateMultiInvoicePayment) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateMultiInvoicePayment object to a map representation for JSON marshaling.
func (c *CreateMultiInvoicePayment) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Memo != nil {
		structMap["memo"] = c.Memo
	}
	if c.Details != nil {
		structMap["details"] = c.Details
	}
	if c.Method != nil {
		structMap["method"] = c.Method
	}
	structMap["amount"] = c.Amount.toMap()
	if c.ReceivedOn != nil {
		structMap["received_on"] = c.ReceivedOn
	}
	structMap["applications"] = c.Applications
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMultiInvoicePayment.
// It customizes the JSON unmarshaling process for CreateMultiInvoicePayment objects.
func (c *CreateMultiInvoicePayment) UnmarshalJSON(input []byte) error {
	var temp createMultiInvoicePayment
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.Memo = temp.Memo
	c.Details = temp.Details
	c.Method = temp.Method
	c.Amount = *temp.Amount
	c.ReceivedOn = temp.ReceivedOn
	c.Applications = *temp.Applications
	return nil
}

// TODO
type createMultiInvoicePayment struct {
	Memo         *string                            `json:"memo,omitempty"`
	Details      *string                            `json:"details,omitempty"`
	Method       *InvoicePaymentMethodType          `json:"method,omitempty"`
	Amount       *CreateMultiInvoicePaymentAmount   `json:"amount"`
	ReceivedOn   *string                            `json:"received_on,omitempty"`
	Applications *[]CreateInvoicePaymentApplication `json:"applications"`
}

func (c *createMultiInvoicePayment) validate() error {
	var errs []string
	if c.Amount == nil {
		errs = append(errs, "required field `amount` is missing for type `Create Multi Invoice Payment`")
	}
	if c.Applications == nil {
		errs = append(errs, "required field `applications` is missing for type `Create Multi Invoice Payment`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
