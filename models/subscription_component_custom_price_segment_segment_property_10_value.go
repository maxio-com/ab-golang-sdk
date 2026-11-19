// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionComponentCustomPriceSegmentSegmentProperty10Value represents a SubscriptionComponentCustomPriceSegmentSegmentProperty10Value struct.
// This is a container for one-of cases.
type SubscriptionComponentCustomPriceSegmentSegmentProperty10Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentCustomPriceSegmentSegmentProperty10Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty10Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty10Value.
// It customizes the JSON marshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty10Value objects.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty10Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentCustomPriceSegmentSegmentProperty10ValueContainer.From*` functions to initialize the SubscriptionComponentCustomPriceSegmentSegmentProperty10Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentCustomPriceSegmentSegmentProperty10Value object to a map representation for JSON marshaling.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty10Value) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *float64:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty10Value.
// It customizes the JSON unmarshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty10Value objects.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty10Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty10Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty10Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty10Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionComponentCustomPriceSegmentSegmentProperty10Value represents a subscriptionComponentCustomPriceSegmentSegmentProperty10Value struct.
// This is a container for one-of cases.
type internalSubscriptionComponentCustomPriceSegmentSegmentProperty10Value struct {}

var SubscriptionComponentCustomPriceSegmentSegmentProperty10ValueContainer internalSubscriptionComponentCustomPriceSegmentSegmentProperty10Value

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty10Value instance, wrapping the provided string value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty10Value) FromString(val string) SubscriptionComponentCustomPriceSegmentSegmentProperty10Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty10Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty10Value instance, wrapping the provided float64 value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty10Value) FromPrecision(val float64) SubscriptionComponentCustomPriceSegmentSegmentProperty10Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty10Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty10Value instance, wrapping the provided int value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty10Value) FromNumber(val int) SubscriptionComponentCustomPriceSegmentSegmentProperty10Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty10Value{value: &val}
}
