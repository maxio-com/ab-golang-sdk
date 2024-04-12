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

// PaymentProfileAttributesExpirationYear represents a PaymentProfileAttributesExpirationYear struct.
// This is a container for one-of cases.
type PaymentProfileAttributesExpirationYear struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the PaymentProfileAttributesExpirationYear object to a string representation.
func (p PaymentProfileAttributesExpirationYear) String() string {
    if bytes, err := json.Marshal(p.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for PaymentProfileAttributesExpirationYear.
// It customizes the JSON marshaling process for PaymentProfileAttributesExpirationYear objects.
func (p PaymentProfileAttributesExpirationYear) MarshalJSON() (
    []byte,
    error) {
    if p.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.PaymentProfileAttributesExpirationYearContainer.From*` functions to initialize the PaymentProfileAttributesExpirationYear object.")
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentProfileAttributesExpirationYear object to a map representation for JSON marshaling.
func (p *PaymentProfileAttributesExpirationYear) toMap() any {
    switch obj := p.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentProfileAttributesExpirationYear.
// It customizes the JSON unmarshaling process for PaymentProfileAttributesExpirationYear objects.
func (p *PaymentProfileAttributesExpirationYear) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &p.isNumber),
        NewTypeHolder(new(string), false, &p.isString),
    )
    
    p.value = result
    return err
}

func (p *PaymentProfileAttributesExpirationYear) AsNumber() (
    *int,
    bool) {
    if !p.isNumber {
        return nil, false
    }
    return p.value.(*int), true
}

func (p *PaymentProfileAttributesExpirationYear) AsString() (
    *string,
    bool) {
    if !p.isString {
        return nil, false
    }
    return p.value.(*string), true
}

// internalPaymentProfileAttributesExpirationYear represents a paymentProfileAttributesExpirationYear struct.
// This is a container for one-of cases.
type internalPaymentProfileAttributesExpirationYear struct {}

var PaymentProfileAttributesExpirationYearContainer internalPaymentProfileAttributesExpirationYear

// The internalPaymentProfileAttributesExpirationYear instance, wrapping the provided int value.
func (p *internalPaymentProfileAttributesExpirationYear) FromNumber(val int) PaymentProfileAttributesExpirationYear {
    return PaymentProfileAttributesExpirationYear{value: &val}
}

// The internalPaymentProfileAttributesExpirationYear instance, wrapping the provided string value.
func (p *internalPaymentProfileAttributesExpirationYear) FromString(val string) PaymentProfileAttributesExpirationYear {
    return PaymentProfileAttributesExpirationYear{value: &val}
}
