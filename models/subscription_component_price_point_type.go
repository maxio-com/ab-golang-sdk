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

// SubscriptionComponentPricePointType represents a SubscriptionComponentPricePointType struct.
// This is a container for one-of cases.
type SubscriptionComponentPricePointType struct {
    value            any
    isPricePointType bool
}

// String converts the SubscriptionComponentPricePointType object to a string representation.
func (s SubscriptionComponentPricePointType) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentPricePointType.
// It customizes the JSON marshaling process for SubscriptionComponentPricePointType objects.
func (s SubscriptionComponentPricePointType) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentPricePointTypeContainer.From*` functions to initialize the SubscriptionComponentPricePointType object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentPricePointType object to a map representation for JSON marshaling.
func (s *SubscriptionComponentPricePointType) toMap() any {
    switch obj := s.value.(type) {
    case *PricePointType:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentPricePointType.
// It customizes the JSON unmarshaling process for SubscriptionComponentPricePointType objects.
func (s *SubscriptionComponentPricePointType) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(PricePointType), false, &s.isPricePointType),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentPricePointType) AsPricePointType() (
    *PricePointType,
    bool) {
    if !s.isPricePointType {
        return nil, false
    }
    return s.value.(*PricePointType), true
}

// internalSubscriptionComponentPricePointType represents a subscriptionComponentPricePointType struct.
// This is a container for one-of cases.
type internalSubscriptionComponentPricePointType struct {}

var SubscriptionComponentPricePointTypeContainer internalSubscriptionComponentPricePointType

// The internalSubscriptionComponentPricePointType instance, wrapping the provided PricePointType value.
func (s *internalSubscriptionComponentPricePointType) FromPricePointType(val PricePointType) SubscriptionComponentPricePointType {
    return SubscriptionComponentPricePointType{value: &val}
}
