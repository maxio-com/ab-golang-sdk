/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// PaymentProfile represents a PaymentProfile struct.
type PaymentProfile struct {
    value                       any
    isApplePayPaymentProfile    bool
    isBankAccountPaymentProfile bool
    isCreditCardPaymentProfile  bool
    isPaypalPaymentProfile      bool
}

// String implements the fmt.Stringer interface for PaymentProfile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PaymentProfile) String() string {
    return fmt.Sprintf("%v", p.value)
}

// MarshalJSON implements the json.Marshaler interface for PaymentProfile.
// It customizes the JSON marshaling process for PaymentProfile objects.
func (p PaymentProfile) MarshalJSON() (
    []byte,
    error) {
    if p.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.PaymentProfileContainer.From*` functions to initialize the PaymentProfile object.")
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentProfile object to a map representation for JSON marshaling.
func (p *PaymentProfile) toMap() any {
    switch obj := p.value.(type) {
    case *ApplePayPaymentProfile:
        return obj.toMap()
    case *BankAccountPaymentProfile:
        return obj.toMap()
    case *CreditCardPaymentProfile:
        return obj.toMap()
    case *PaypalPaymentProfile:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentProfile.
// It customizes the JSON unmarshaling process for PaymentProfile objects.
func (p *PaymentProfile) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOfWithDiscriminator(input, "payment_type",
        NewTypeHolderDiscriminator(&ApplePayPaymentProfile{}, false, &p.isApplePayPaymentProfile, "apple_pay"),
        NewTypeHolderDiscriminator(&BankAccountPaymentProfile{}, false, &p.isBankAccountPaymentProfile, "bank_account"),
        NewTypeHolderDiscriminator(&CreditCardPaymentProfile{}, false, &p.isCreditCardPaymentProfile, "credit_card"),
        NewTypeHolderDiscriminator(&PaypalPaymentProfile{}, false, &p.isPaypalPaymentProfile, "paypal_account"),
    )
    
    p.value = result
    return err
}

func (p *PaymentProfile) AsApplePayPaymentProfile() (
    *ApplePayPaymentProfile,
    bool) {
    if !p.isApplePayPaymentProfile {
        return nil, false
    }
    return p.value.(*ApplePayPaymentProfile), true
}

func (p *PaymentProfile) AsBankAccountPaymentProfile() (
    *BankAccountPaymentProfile,
    bool) {
    if !p.isBankAccountPaymentProfile {
        return nil, false
    }
    return p.value.(*BankAccountPaymentProfile), true
}

func (p *PaymentProfile) AsCreditCardPaymentProfile() (
    *CreditCardPaymentProfile,
    bool) {
    if !p.isCreditCardPaymentProfile {
        return nil, false
    }
    return p.value.(*CreditCardPaymentProfile), true
}

func (p *PaymentProfile) AsPaypalPaymentProfile() (
    *PaypalPaymentProfile,
    bool) {
    if !p.isPaypalPaymentProfile {
        return nil, false
    }
    return p.value.(*PaypalPaymentProfile), true
}

// internalPaymentProfile represents a paymentProfile struct.
type internalPaymentProfile struct {}

var PaymentProfileContainer internalPaymentProfile

// The internalPaymentProfile instance, wrapping the provided ApplePayPaymentProfile value.
func (p *internalPaymentProfile) FromApplePayPaymentProfile(val ApplePayPaymentProfile) PaymentProfile {
    return PaymentProfile{value: &val}
}

// The internalPaymentProfile instance, wrapping the provided BankAccountPaymentProfile value.
func (p *internalPaymentProfile) FromBankAccountPaymentProfile(val BankAccountPaymentProfile) PaymentProfile {
    return PaymentProfile{value: &val}
}

// The internalPaymentProfile instance, wrapping the provided CreditCardPaymentProfile value.
func (p *internalPaymentProfile) FromCreditCardPaymentProfile(val CreditCardPaymentProfile) PaymentProfile {
    return PaymentProfile{value: &val}
}

// The internalPaymentProfile instance, wrapping the provided PaypalPaymentProfile value.
func (p *internalPaymentProfile) FromPaypalPaymentProfile(val PaypalPaymentProfile) PaymentProfile {
    return PaymentProfile{value: &val}
}
