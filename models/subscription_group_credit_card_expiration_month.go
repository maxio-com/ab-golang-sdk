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

// SubscriptionGroupCreditCardExpirationMonth represents a SubscriptionGroupCreditCardExpirationMonth struct.
// This is a container for one-of cases.
type SubscriptionGroupCreditCardExpirationMonth struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the SubscriptionGroupCreditCardExpirationMonth object to a string representation.
func (s SubscriptionGroupCreditCardExpirationMonth) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupCreditCardExpirationMonth.
// It customizes the JSON marshaling process for SubscriptionGroupCreditCardExpirationMonth objects.
func (s SubscriptionGroupCreditCardExpirationMonth) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionGroupCreditCardExpirationMonthContainer.From*` functions to initialize the SubscriptionGroupCreditCardExpirationMonth object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupCreditCardExpirationMonth object to a map representation for JSON marshaling.
func (s *SubscriptionGroupCreditCardExpirationMonth) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupCreditCardExpirationMonth.
// It customizes the JSON unmarshaling process for SubscriptionGroupCreditCardExpirationMonth objects.
func (s *SubscriptionGroupCreditCardExpirationMonth) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionGroupCreditCardExpirationMonth) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionGroupCreditCardExpirationMonth) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionGroupCreditCardExpirationMonth represents a subscriptionGroupCreditCardExpirationMonth struct.
// This is a container for one-of cases.
type internalSubscriptionGroupCreditCardExpirationMonth struct {}

var SubscriptionGroupCreditCardExpirationMonthContainer internalSubscriptionGroupCreditCardExpirationMonth

// The internalSubscriptionGroupCreditCardExpirationMonth instance, wrapping the provided string value.
func (s *internalSubscriptionGroupCreditCardExpirationMonth) FromString(val string) SubscriptionGroupCreditCardExpirationMonth {
    return SubscriptionGroupCreditCardExpirationMonth{value: &val}
}

// The internalSubscriptionGroupCreditCardExpirationMonth instance, wrapping the provided int value.
func (s *internalSubscriptionGroupCreditCardExpirationMonth) FromNumber(val int) SubscriptionGroupCreditCardExpirationMonth {
    return SubscriptionGroupCreditCardExpirationMonth{value: &val}
}
