/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// InvoiceEventPayment2 represents a InvoiceEventPayment2 struct.
// A nested data structure detailing the method of payment
type InvoiceEventPayment2 struct {
	value                      any
	isPaymentMethodApplePay    bool
	isPaymentMethodBankAccount bool
	isPaymentMethodCreditCard  bool
	isPaymentMethodExternal    bool
	isPaymentMethodPaypal      bool
}

// String converts the InvoiceEventPayment2 object to a string representation.
func (i InvoiceEventPayment2) String() string {
	if bytes, err := json.Marshal(i.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for InvoiceEventPayment2.
// It customizes the JSON marshaling process for InvoiceEventPayment2 objects.
func (i *InvoiceEventPayment2) MarshalJSON() (
	[]byte,
	error) {
	if i.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.InvoiceEventPaymentContainer.From*` functions to initialize the InvoiceEventPayment2 object.")
	}
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceEventPayment2 object to a map representation for JSON marshaling.
func (i *InvoiceEventPayment2) toMap() any {
	switch obj := i.value.(type) {
	case *PaymentMethodApplePay:
		return obj.toMap()
	case *PaymentMethodBankAccount:
		return obj.toMap()
	case *PaymentMethodCreditCard:
		return obj.toMap()
	case *PaymentMethodExternal:
		return obj.toMap()
	case *PaymentMethodPaypal:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEventPayment2.
// It customizes the JSON unmarshaling process for InvoiceEventPayment2 objects.
func (i *InvoiceEventPayment2) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOfWithDiscriminator(input, "type",
		NewTypeHolderDiscriminator(&PaymentMethodApplePay{}, false, &i.isPaymentMethodApplePay, "apple_pay"),
		NewTypeHolderDiscriminator(&PaymentMethodBankAccount{}, false, &i.isPaymentMethodBankAccount, "bank_account"),
		NewTypeHolderDiscriminator(&PaymentMethodCreditCard{}, false, &i.isPaymentMethodCreditCard, "credit_card"),
		NewTypeHolderDiscriminator(&PaymentMethodExternal{}, false, &i.isPaymentMethodExternal, "external"),
		NewTypeHolderDiscriminator(&PaymentMethodPaypal{}, false, &i.isPaymentMethodPaypal, "paypal_account"),
	)

	i.value = result
	return err
}

func (i *InvoiceEventPayment2) AsPaymentMethodApplePay() (
	*PaymentMethodApplePay,
	bool) {
	if !i.isPaymentMethodApplePay {
		return nil, false
	}
	return i.value.(*PaymentMethodApplePay), true
}

func (i *InvoiceEventPayment2) AsPaymentMethodBankAccount() (
	*PaymentMethodBankAccount,
	bool) {
	if !i.isPaymentMethodBankAccount {
		return nil, false
	}
	return i.value.(*PaymentMethodBankAccount), true
}

func (i *InvoiceEventPayment2) AsPaymentMethodCreditCard() (
	*PaymentMethodCreditCard,
	bool) {
	if !i.isPaymentMethodCreditCard {
		return nil, false
	}
	return i.value.(*PaymentMethodCreditCard), true
}

func (i *InvoiceEventPayment2) AsPaymentMethodExternal() (
	*PaymentMethodExternal,
	bool) {
	if !i.isPaymentMethodExternal {
		return nil, false
	}
	return i.value.(*PaymentMethodExternal), true
}

func (i *InvoiceEventPayment2) AsPaymentMethodPaypal() (
	*PaymentMethodPaypal,
	bool) {
	if !i.isPaymentMethodPaypal {
		return nil, false
	}
	return i.value.(*PaymentMethodPaypal), true
}

// internalInvoiceEventPayment2 represents a invoiceEventPayment2 struct.
// A nested data structure detailing the method of payment
type internalInvoiceEventPayment2 struct{}

var InvoiceEventPaymentContainer internalInvoiceEventPayment2

func (i *internalInvoiceEventPayment2) FromPaymentMethodApplePay(val PaymentMethodApplePay) InvoiceEventPayment2 {
	return InvoiceEventPayment2{value: &val}
}

func (i *internalInvoiceEventPayment2) FromPaymentMethodBankAccount(val PaymentMethodBankAccount) InvoiceEventPayment2 {
	return InvoiceEventPayment2{value: &val}
}

func (i *internalInvoiceEventPayment2) FromPaymentMethodCreditCard(val PaymentMethodCreditCard) InvoiceEventPayment2 {
	return InvoiceEventPayment2{value: &val}
}

func (i *internalInvoiceEventPayment2) FromPaymentMethodExternal(val PaymentMethodExternal) InvoiceEventPayment2 {
	return InvoiceEventPayment2{value: &val}
}

func (i *internalInvoiceEventPayment2) FromPaymentMethodPaypal(val PaymentMethodPaypal) InvoiceEventPayment2 {
	return InvoiceEventPayment2{value: &val}
}
