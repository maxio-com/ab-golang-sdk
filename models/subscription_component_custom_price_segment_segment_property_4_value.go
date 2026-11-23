// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionComponentCustomPriceSegmentSegmentProperty4Value represents a SubscriptionComponentCustomPriceSegmentSegmentProperty4Value struct.
// This is a container for one-of cases.
type SubscriptionComponentCustomPriceSegmentSegmentProperty4Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentCustomPriceSegmentSegmentProperty4Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty4Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty4Value.
// It customizes the JSON marshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty4Value objects.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty4Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentCustomPriceSegmentSegmentProperty4ValueContainer.From*` functions to initialize the SubscriptionComponentCustomPriceSegmentSegmentProperty4Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentCustomPriceSegmentSegmentProperty4Value object to a map representation for JSON marshaling.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty4Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty4Value.
// It customizes the JSON unmarshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty4Value objects.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty4Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty4Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty4Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty4Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionComponentCustomPriceSegmentSegmentProperty4Value represents a subscriptionComponentCustomPriceSegmentSegmentProperty4Value struct.
// This is a container for one-of cases.
type internalSubscriptionComponentCustomPriceSegmentSegmentProperty4Value struct {}

var SubscriptionComponentCustomPriceSegmentSegmentProperty4ValueContainer internalSubscriptionComponentCustomPriceSegmentSegmentProperty4Value

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty4Value instance, wrapping the provided string value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty4Value) FromString(val string) SubscriptionComponentCustomPriceSegmentSegmentProperty4Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty4Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty4Value instance, wrapping the provided float64 value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty4Value) FromPrecision(val float64) SubscriptionComponentCustomPriceSegmentSegmentProperty4Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty4Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty4Value instance, wrapping the provided int value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty4Value) FromNumber(val int) SubscriptionComponentCustomPriceSegmentSegmentProperty4Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty4Value{value: &val}
}
