// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionComponentCustomPriceSegmentSegmentProperty6Value represents a SubscriptionComponentCustomPriceSegmentSegmentProperty6Value struct.
// This is a container for one-of cases.
type SubscriptionComponentCustomPriceSegmentSegmentProperty6Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentCustomPriceSegmentSegmentProperty6Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty6Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty6Value.
// It customizes the JSON marshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty6Value objects.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty6Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentCustomPriceSegmentSegmentProperty6ValueContainer.From*` functions to initialize the SubscriptionComponentCustomPriceSegmentSegmentProperty6Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentCustomPriceSegmentSegmentProperty6Value object to a map representation for JSON marshaling.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty6Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty6Value.
// It customizes the JSON unmarshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty6Value objects.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty6Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty6Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty6Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty6Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionComponentCustomPriceSegmentSegmentProperty6Value represents a subscriptionComponentCustomPriceSegmentSegmentProperty6Value struct.
// This is a container for one-of cases.
type internalSubscriptionComponentCustomPriceSegmentSegmentProperty6Value struct {}

var SubscriptionComponentCustomPriceSegmentSegmentProperty6ValueContainer internalSubscriptionComponentCustomPriceSegmentSegmentProperty6Value

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty6Value instance, wrapping the provided string value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty6Value) FromString(val string) SubscriptionComponentCustomPriceSegmentSegmentProperty6Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty6Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty6Value instance, wrapping the provided float64 value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty6Value) FromPrecision(val float64) SubscriptionComponentCustomPriceSegmentSegmentProperty6Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty6Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty6Value instance, wrapping the provided int value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty6Value) FromNumber(val int) SubscriptionComponentCustomPriceSegmentSegmentProperty6Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty6Value{value: &val}
}
