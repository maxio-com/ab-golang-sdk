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

// PaymentProfileResponsePaymentProfile represents a PaymentProfileResponsePaymentProfile struct.
// This is a container for one-of cases.
type PaymentProfileResponsePaymentProfile struct {
    value                       any
    isBankAccountPaymentProfile bool
    isCreditCardPaymentProfile  bool
}

// String converts the PaymentProfileResponsePaymentProfile object to a string representation.
func (p PaymentProfileResponsePaymentProfile) String() string {
    if bytes, err := json.Marshal(p.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for PaymentProfileResponsePaymentProfile.
// It customizes the JSON marshaling process for PaymentProfileResponsePaymentProfile objects.
func (p PaymentProfileResponsePaymentProfile) MarshalJSON() (
    []byte,
    error) {
    if p.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.PaymentProfileResponsePaymentProfileContainer.From*` functions to initialize the PaymentProfileResponsePaymentProfile object.")
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentProfileResponsePaymentProfile object to a map representation for JSON marshaling.
func (p *PaymentProfileResponsePaymentProfile) toMap() any {
    switch obj := p.value.(type) {
    case *BankAccountPaymentProfile:
        return obj.toMap()
    case *CreditCardPaymentProfile:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentProfileResponsePaymentProfile.
// It customizes the JSON unmarshaling process for PaymentProfileResponsePaymentProfile objects.
func (p *PaymentProfileResponsePaymentProfile) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&BankAccountPaymentProfile{}, false, &p.isBankAccountPaymentProfile),
        NewTypeHolder(&CreditCardPaymentProfile{}, false, &p.isCreditCardPaymentProfile),
    )
    
    p.value = result
    return err
}

func (p *PaymentProfileResponsePaymentProfile) AsBankAccountPaymentProfile() (
    *BankAccountPaymentProfile,
    bool) {
    if !p.isBankAccountPaymentProfile {
        return nil, false
    }
    return p.value.(*BankAccountPaymentProfile), true
}

func (p *PaymentProfileResponsePaymentProfile) AsCreditCardPaymentProfile() (
    *CreditCardPaymentProfile,
    bool) {
    if !p.isCreditCardPaymentProfile {
        return nil, false
    }
    return p.value.(*CreditCardPaymentProfile), true
}

// internalPaymentProfileResponsePaymentProfile represents a paymentProfileResponsePaymentProfile struct.
// This is a container for one-of cases.
type internalPaymentProfileResponsePaymentProfile struct {}

var PaymentProfileResponsePaymentProfileContainer internalPaymentProfileResponsePaymentProfile

// The internalPaymentProfileResponsePaymentProfile instance, wrapping the provided BankAccountPaymentProfile value.
func (p *internalPaymentProfileResponsePaymentProfile) FromBankAccountPaymentProfile(val BankAccountPaymentProfile) PaymentProfileResponsePaymentProfile {
    return PaymentProfileResponsePaymentProfile{value: &val}
}

// The internalPaymentProfileResponsePaymentProfile instance, wrapping the provided CreditCardPaymentProfile value.
func (p *internalPaymentProfileResponsePaymentProfile) FromCreditCardPaymentProfile(val CreditCardPaymentProfile) PaymentProfileResponsePaymentProfile {
    return PaymentProfileResponsePaymentProfile{value: &val}
}
