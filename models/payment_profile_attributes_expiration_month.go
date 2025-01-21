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

// PaymentProfileAttributesExpirationMonth represents a PaymentProfileAttributesExpirationMonth struct.
// This is a container for one-of cases.
type PaymentProfileAttributesExpirationMonth struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for PaymentProfileAttributesExpirationMonth,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PaymentProfileAttributesExpirationMonth) String() string {
    return fmt.Sprintf("%v", p.value)
}

// MarshalJSON implements the json.Marshaler interface for PaymentProfileAttributesExpirationMonth.
// It customizes the JSON marshaling process for PaymentProfileAttributesExpirationMonth objects.
func (p PaymentProfileAttributesExpirationMonth) MarshalJSON() (
    []byte,
    error) {
    if p.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.PaymentProfileAttributesExpirationMonthContainer.From*` functions to initialize the PaymentProfileAttributesExpirationMonth object.")
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentProfileAttributesExpirationMonth object to a map representation for JSON marshaling.
func (p *PaymentProfileAttributesExpirationMonth) toMap() any {
    switch obj := p.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentProfileAttributesExpirationMonth.
// It customizes the JSON unmarshaling process for PaymentProfileAttributesExpirationMonth objects.
func (p *PaymentProfileAttributesExpirationMonth) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &p.isNumber),
        NewTypeHolder(new(string), false, &p.isString),
    )
    
    p.value = result
    return err
}

func (p *PaymentProfileAttributesExpirationMonth) AsNumber() (
    *int,
    bool) {
    if !p.isNumber {
        return nil, false
    }
    return p.value.(*int), true
}

func (p *PaymentProfileAttributesExpirationMonth) AsString() (
    *string,
    bool) {
    if !p.isString {
        return nil, false
    }
    return p.value.(*string), true
}

// internalPaymentProfileAttributesExpirationMonth represents a paymentProfileAttributesExpirationMonth struct.
// This is a container for one-of cases.
type internalPaymentProfileAttributesExpirationMonth struct {}

var PaymentProfileAttributesExpirationMonthContainer internalPaymentProfileAttributesExpirationMonth

// The internalPaymentProfileAttributesExpirationMonth instance, wrapping the provided int value.
func (p *internalPaymentProfileAttributesExpirationMonth) FromNumber(val int) PaymentProfileAttributesExpirationMonth {
    return PaymentProfileAttributesExpirationMonth{value: &val}
}

// The internalPaymentProfileAttributesExpirationMonth instance, wrapping the provided string value.
func (p *internalPaymentProfileAttributesExpirationMonth) FromString(val string) PaymentProfileAttributesExpirationMonth {
    return PaymentProfileAttributesExpirationMonth{value: &val}
}
