// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionComponentCustomPriceSegmentSegmentProperty1Value represents a SubscriptionComponentCustomPriceSegmentSegmentProperty1Value struct.
// This is a container for one-of cases.
type SubscriptionComponentCustomPriceSegmentSegmentProperty1Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentCustomPriceSegmentSegmentProperty1Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty1Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty1Value.
// It customizes the JSON marshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty1Value objects.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty1Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentCustomPriceSegmentSegmentProperty1ValueContainer.From*` functions to initialize the SubscriptionComponentCustomPriceSegmentSegmentProperty1Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentCustomPriceSegmentSegmentProperty1Value object to a map representation for JSON marshaling.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty1Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty1Value.
// It customizes the JSON unmarshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty1Value objects.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty1Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty1Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty1Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty1Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionComponentCustomPriceSegmentSegmentProperty1Value represents a subscriptionComponentCustomPriceSegmentSegmentProperty1Value struct.
// This is a container for one-of cases.
type internalSubscriptionComponentCustomPriceSegmentSegmentProperty1Value struct {}

var SubscriptionComponentCustomPriceSegmentSegmentProperty1ValueContainer internalSubscriptionComponentCustomPriceSegmentSegmentProperty1Value

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty1Value instance, wrapping the provided string value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty1Value) FromString(val string) SubscriptionComponentCustomPriceSegmentSegmentProperty1Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty1Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty1Value instance, wrapping the provided float64 value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty1Value) FromPrecision(val float64) SubscriptionComponentCustomPriceSegmentSegmentProperty1Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty1Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty1Value instance, wrapping the provided int value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty1Value) FromNumber(val int) SubscriptionComponentCustomPriceSegmentSegmentProperty1Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty1Value{value: &val}
}
