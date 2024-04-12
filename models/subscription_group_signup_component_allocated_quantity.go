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

// SubscriptionGroupSignupComponentAllocatedQuantity represents a SubscriptionGroupSignupComponentAllocatedQuantity struct.
// This is a container for one-of cases.
type SubscriptionGroupSignupComponentAllocatedQuantity struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the SubscriptionGroupSignupComponentAllocatedQuantity object to a string representation.
func (s SubscriptionGroupSignupComponentAllocatedQuantity) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupComponentAllocatedQuantity.
// It customizes the JSON marshaling process for SubscriptionGroupSignupComponentAllocatedQuantity objects.
func (s SubscriptionGroupSignupComponentAllocatedQuantity) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionGroupSignupComponentAllocatedQuantityContainer.From*` functions to initialize the SubscriptionGroupSignupComponentAllocatedQuantity object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupComponentAllocatedQuantity object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSignupComponentAllocatedQuantity) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupComponentAllocatedQuantity.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupComponentAllocatedQuantity objects.
func (s *SubscriptionGroupSignupComponentAllocatedQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionGroupSignupComponentAllocatedQuantity) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionGroupSignupComponentAllocatedQuantity) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionGroupSignupComponentAllocatedQuantity represents a subscriptionGroupSignupComponentAllocatedQuantity struct.
// This is a container for one-of cases.
type internalSubscriptionGroupSignupComponentAllocatedQuantity struct {}

var SubscriptionGroupSignupComponentAllocatedQuantityContainer internalSubscriptionGroupSignupComponentAllocatedQuantity

// The internalSubscriptionGroupSignupComponentAllocatedQuantity instance, wrapping the provided string value.
func (s *internalSubscriptionGroupSignupComponentAllocatedQuantity) FromString(val string) SubscriptionGroupSignupComponentAllocatedQuantity {
    return SubscriptionGroupSignupComponentAllocatedQuantity{value: &val}
}

// The internalSubscriptionGroupSignupComponentAllocatedQuantity instance, wrapping the provided int value.
func (s *internalSubscriptionGroupSignupComponentAllocatedQuantity) FromNumber(val int) SubscriptionGroupSignupComponentAllocatedQuantity {
    return SubscriptionGroupSignupComponentAllocatedQuantity{value: &val}
}
