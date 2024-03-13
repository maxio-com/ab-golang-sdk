package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// PaymentRelatedEvents represents a PaymentRelatedEvents struct.
type PaymentRelatedEvents struct {
	ProductId            int `json:"product_id"`
	AccountTransactionId int `json:"account_transaction_id"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentRelatedEvents.
// It customizes the JSON marshaling process for PaymentRelatedEvents objects.
func (p *PaymentRelatedEvents) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PaymentRelatedEvents object to a map representation for JSON marshaling.
func (p *PaymentRelatedEvents) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["product_id"] = p.ProductId
	structMap["account_transaction_id"] = p.AccountTransactionId
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentRelatedEvents.
// It customizes the JSON unmarshaling process for PaymentRelatedEvents objects.
func (p *PaymentRelatedEvents) UnmarshalJSON(input []byte) error {
	var temp paymentRelatedEvents
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	p.ProductId = *temp.ProductId
	p.AccountTransactionId = *temp.AccountTransactionId
	return nil
}

// TODO
type paymentRelatedEvents struct {
	ProductId            *int `json:"product_id"`
	AccountTransactionId *int `json:"account_transaction_id"`
}

func (p *paymentRelatedEvents) validate() error {
	var errs []string
	if p.ProductId == nil {
		errs = append(errs, "required field `product_id` is missing for type `Payment Related Events`")
	}
	if p.AccountTransactionId == nil {
		errs = append(errs, "required field `account_transaction_id` is missing for type `Payment Related Events`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
