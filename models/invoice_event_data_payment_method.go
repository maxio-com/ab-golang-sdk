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

// InvoiceEventDataPaymentMethod represents a InvoiceEventDataPaymentMethod struct.
// This is a container for any-of cases.
type InvoiceEventDataPaymentMethod struct {
    value                      any
    isPaymentMethodApplePay    bool
    isPaymentMethodBankAccount bool
    isPaymentMethodCreditCard  bool
    isPaymentMethodExternal    bool
    isPaymentMethodPaypal      bool
}

// String converts the InvoiceEventDataPaymentMethod object to a string representation.
func (i InvoiceEventDataPaymentMethod) String() string {
    if bytes, err := json.Marshal(i.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for InvoiceEventDataPaymentMethod.
// It customizes the JSON marshaling process for InvoiceEventDataPaymentMethod objects.
func (i InvoiceEventDataPaymentMethod) MarshalJSON() (
    []byte,
    error) {
    if i.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.InvoiceEventDataPaymentMethodContainer.From*` functions to initialize the InvoiceEventDataPaymentMethod object.")
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceEventDataPaymentMethod object to a map representation for JSON marshaling.
func (i *InvoiceEventDataPaymentMethod) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEventDataPaymentMethod.
// It customizes the JSON unmarshaling process for InvoiceEventDataPaymentMethod objects.
func (i *InvoiceEventDataPaymentMethod) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&PaymentMethodApplePay{}, false, &i.isPaymentMethodApplePay),
        NewTypeHolder(&PaymentMethodBankAccount{}, false, &i.isPaymentMethodBankAccount),
        NewTypeHolder(&PaymentMethodCreditCard{}, false, &i.isPaymentMethodCreditCard),
        NewTypeHolder(&PaymentMethodExternal{}, false, &i.isPaymentMethodExternal),
        NewTypeHolder(&PaymentMethodPaypal{}, false, &i.isPaymentMethodPaypal),
    )
    
    i.value = result
    return err
}

func (i *InvoiceEventDataPaymentMethod) AsPaymentMethodApplePay() (
    *PaymentMethodApplePay,
    bool) {
    if !i.isPaymentMethodApplePay {
        return nil, false
    }
    return i.value.(*PaymentMethodApplePay), true
}

func (i *InvoiceEventDataPaymentMethod) AsPaymentMethodBankAccount() (
    *PaymentMethodBankAccount,
    bool) {
    if !i.isPaymentMethodBankAccount {
        return nil, false
    }
    return i.value.(*PaymentMethodBankAccount), true
}

func (i *InvoiceEventDataPaymentMethod) AsPaymentMethodCreditCard() (
    *PaymentMethodCreditCard,
    bool) {
    if !i.isPaymentMethodCreditCard {
        return nil, false
    }
    return i.value.(*PaymentMethodCreditCard), true
}

func (i *InvoiceEventDataPaymentMethod) AsPaymentMethodExternal() (
    *PaymentMethodExternal,
    bool) {
    if !i.isPaymentMethodExternal {
        return nil, false
    }
    return i.value.(*PaymentMethodExternal), true
}

func (i *InvoiceEventDataPaymentMethod) AsPaymentMethodPaypal() (
    *PaymentMethodPaypal,
    bool) {
    if !i.isPaymentMethodPaypal {
        return nil, false
    }
    return i.value.(*PaymentMethodPaypal), true
}

// internalInvoiceEventDataPaymentMethod represents a invoiceEventDataPaymentMethod struct.
// This is a container for any-of cases.
type internalInvoiceEventDataPaymentMethod struct {}

var InvoiceEventDataPaymentMethodContainer internalInvoiceEventDataPaymentMethod

// The internalInvoiceEventDataPaymentMethod instance, wrapping the provided PaymentMethodApplePay value.
func (i *internalInvoiceEventDataPaymentMethod) FromPaymentMethodApplePay(val PaymentMethodApplePay) InvoiceEventDataPaymentMethod {
    return InvoiceEventDataPaymentMethod{value: &val}
}

// The internalInvoiceEventDataPaymentMethod instance, wrapping the provided PaymentMethodBankAccount value.
func (i *internalInvoiceEventDataPaymentMethod) FromPaymentMethodBankAccount(val PaymentMethodBankAccount) InvoiceEventDataPaymentMethod {
    return InvoiceEventDataPaymentMethod{value: &val}
}

// The internalInvoiceEventDataPaymentMethod instance, wrapping the provided PaymentMethodCreditCard value.
func (i *internalInvoiceEventDataPaymentMethod) FromPaymentMethodCreditCard(val PaymentMethodCreditCard) InvoiceEventDataPaymentMethod {
    return InvoiceEventDataPaymentMethod{value: &val}
}

// The internalInvoiceEventDataPaymentMethod instance, wrapping the provided PaymentMethodExternal value.
func (i *internalInvoiceEventDataPaymentMethod) FromPaymentMethodExternal(val PaymentMethodExternal) InvoiceEventDataPaymentMethod {
    return InvoiceEventDataPaymentMethod{value: &val}
}

// The internalInvoiceEventDataPaymentMethod instance, wrapping the provided PaymentMethodPaypal value.
func (i *internalInvoiceEventDataPaymentMethod) FromPaymentMethodPaypal(val PaymentMethodPaypal) InvoiceEventDataPaymentMethod {
    return InvoiceEventDataPaymentMethod{value: &val}
}
