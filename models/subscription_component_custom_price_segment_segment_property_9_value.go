// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionComponentCustomPriceSegmentSegmentProperty9Value represents a SubscriptionComponentCustomPriceSegmentSegmentProperty9Value struct.
// This is a container for one-of cases.
type SubscriptionComponentCustomPriceSegmentSegmentProperty9Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentCustomPriceSegmentSegmentProperty9Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty9Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty9Value.
// It customizes the JSON marshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty9Value objects.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty9Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentCustomPriceSegmentSegmentProperty9ValueContainer.From*` functions to initialize the SubscriptionComponentCustomPriceSegmentSegmentProperty9Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentCustomPriceSegmentSegmentProperty9Value object to a map representation for JSON marshaling.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty9Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty9Value.
// It customizes the JSON unmarshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty9Value objects.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty9Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty9Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty9Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty9Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionComponentCustomPriceSegmentSegmentProperty9Value represents a subscriptionComponentCustomPriceSegmentSegmentProperty9Value struct.
// This is a container for one-of cases.
type internalSubscriptionComponentCustomPriceSegmentSegmentProperty9Value struct {}

var SubscriptionComponentCustomPriceSegmentSegmentProperty9ValueContainer internalSubscriptionComponentCustomPriceSegmentSegmentProperty9Value

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty9Value instance, wrapping the provided string value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty9Value) FromString(val string) SubscriptionComponentCustomPriceSegmentSegmentProperty9Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty9Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty9Value instance, wrapping the provided float64 value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty9Value) FromPrecision(val float64) SubscriptionComponentCustomPriceSegmentSegmentProperty9Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty9Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty9Value instance, wrapping the provided int value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty9Value) FromNumber(val int) SubscriptionComponentCustomPriceSegmentSegmentProperty9Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty9Value{value: &val}
}
