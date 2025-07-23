// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// InvoiceEventPayment represents a InvoiceEventPayment struct.
// A nested data structure detailing the method of payment
type InvoiceEventPayment struct {
	value                      any
	isPaymentMethodApplePay    bool
	isPaymentMethodBankAccount bool
	isPaymentMethodCreditCard  bool
	isPaymentMethodExternal    bool
	isPaymentMethodPaypal      bool
}

// String implements the fmt.Stringer interface for InvoiceEventPayment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceEventPayment) String() string {
	return fmt.Sprintf("%v", i.value)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceEventPayment.
// It customizes the JSON marshaling process for InvoiceEventPayment objects.
func (i InvoiceEventPayment) MarshalJSON() (
	[]byte,
	error) {
	if i.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.InvoiceEventPaymentContainer.From*` functions to initialize the InvoiceEventPayment object.")
	}
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceEventPayment object to a map representation for JSON marshaling.
func (i *InvoiceEventPayment) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEventPayment.
// It customizes the JSON unmarshaling process for InvoiceEventPayment objects.
func (i *InvoiceEventPayment) UnmarshalJSON(input []byte) error {
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

func (i *InvoiceEventPayment) AsPaymentMethodApplePay() (
	*PaymentMethodApplePay,
	bool) {
	if !i.isPaymentMethodApplePay {
		return nil, false
	}
	return i.value.(*PaymentMethodApplePay), true
}

func (i *InvoiceEventPayment) AsPaymentMethodBankAccount() (
	*PaymentMethodBankAccount,
	bool) {
	if !i.isPaymentMethodBankAccount {
		return nil, false
	}
	return i.value.(*PaymentMethodBankAccount), true
}

func (i *InvoiceEventPayment) AsPaymentMethodCreditCard() (
	*PaymentMethodCreditCard,
	bool) {
	if !i.isPaymentMethodCreditCard {
		return nil, false
	}
	return i.value.(*PaymentMethodCreditCard), true
}

func (i *InvoiceEventPayment) AsPaymentMethodExternal() (
	*PaymentMethodExternal,
	bool) {
	if !i.isPaymentMethodExternal {
		return nil, false
	}
	return i.value.(*PaymentMethodExternal), true
}

func (i *InvoiceEventPayment) AsPaymentMethodPaypal() (
	*PaymentMethodPaypal,
	bool) {
	if !i.isPaymentMethodPaypal {
		return nil, false
	}
	return i.value.(*PaymentMethodPaypal), true
}

// internalInvoiceEventPayment represents a invoiceEventPayment struct.
// A nested data structure detailing the method of payment
type internalInvoiceEventPayment struct{}

var InvoiceEventPaymentContainer internalInvoiceEventPayment

// The internalInvoiceEventPayment instance, wrapping the provided PaymentMethodApplePay value.
func (i *internalInvoiceEventPayment) FromPaymentMethodApplePay(val PaymentMethodApplePay) InvoiceEventPayment {
	return InvoiceEventPayment{value: &val}
}

// The internalInvoiceEventPayment instance, wrapping the provided PaymentMethodBankAccount value.
func (i *internalInvoiceEventPayment) FromPaymentMethodBankAccount(val PaymentMethodBankAccount) InvoiceEventPayment {
	return InvoiceEventPayment{value: &val}
}

// The internalInvoiceEventPayment instance, wrapping the provided PaymentMethodCreditCard value.
func (i *internalInvoiceEventPayment) FromPaymentMethodCreditCard(val PaymentMethodCreditCard) InvoiceEventPayment {
	return InvoiceEventPayment{value: &val}
}

// The internalInvoiceEventPayment instance, wrapping the provided PaymentMethodExternal value.
func (i *internalInvoiceEventPayment) FromPaymentMethodExternal(val PaymentMethodExternal) InvoiceEventPayment {
	return InvoiceEventPayment{value: &val}
}

// The internalInvoiceEventPayment instance, wrapping the provided PaymentMethodPaypal value.
func (i *internalInvoiceEventPayment) FromPaymentMethodPaypal(val PaymentMethodPaypal) InvoiceEventPayment {
	return InvoiceEventPayment{value: &val}
}
