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

// SubscriptionCustomPriceTrialPriceInCents represents a SubscriptionCustomPriceTrialPriceInCents struct.
// This is a container for one-of cases.
type SubscriptionCustomPriceTrialPriceInCents struct {
    value    any
    isString bool
    isLong   bool
}

// String converts the SubscriptionCustomPriceTrialPriceInCents object to a string representation.
func (s SubscriptionCustomPriceTrialPriceInCents) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionCustomPriceTrialPriceInCents.
// It customizes the JSON marshaling process for SubscriptionCustomPriceTrialPriceInCents objects.
func (s SubscriptionCustomPriceTrialPriceInCents) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionCustomPriceTrialPriceInCentsContainer.From*` functions to initialize the SubscriptionCustomPriceTrialPriceInCents object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionCustomPriceTrialPriceInCents object to a map representation for JSON marshaling.
func (s *SubscriptionCustomPriceTrialPriceInCents) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int64:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionCustomPriceTrialPriceInCents.
// It customizes the JSON unmarshaling process for SubscriptionCustomPriceTrialPriceInCents objects.
func (s *SubscriptionCustomPriceTrialPriceInCents) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int64), false, &s.isLong),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionCustomPriceTrialPriceInCents) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionCustomPriceTrialPriceInCents) AsLong() (
    *int64,
    bool) {
    if !s.isLong {
        return nil, false
    }
    return s.value.(*int64), true
}

// internalSubscriptionCustomPriceTrialPriceInCents represents a subscriptionCustomPriceTrialPriceInCents struct.
// This is a container for one-of cases.
type internalSubscriptionCustomPriceTrialPriceInCents struct {}

var SubscriptionCustomPriceTrialPriceInCentsContainer internalSubscriptionCustomPriceTrialPriceInCents

// The internalSubscriptionCustomPriceTrialPriceInCents instance, wrapping the provided string value.
func (s *internalSubscriptionCustomPriceTrialPriceInCents) FromString(val string) SubscriptionCustomPriceTrialPriceInCents {
    return SubscriptionCustomPriceTrialPriceInCents{value: &val}
}

// The internalSubscriptionCustomPriceTrialPriceInCents instance, wrapping the provided int64 value.
func (s *internalSubscriptionCustomPriceTrialPriceInCents) FromLong(val int64) SubscriptionCustomPriceTrialPriceInCents {
    return SubscriptionCustomPriceTrialPriceInCents{value: &val}
}
