// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionComponentCustomPriceSegmentSegmentProperty7Value represents a SubscriptionComponentCustomPriceSegmentSegmentProperty7Value struct.
// This is a container for one-of cases.
type SubscriptionComponentCustomPriceSegmentSegmentProperty7Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentCustomPriceSegmentSegmentProperty7Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty7Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty7Value.
// It customizes the JSON marshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty7Value objects.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty7Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentCustomPriceSegmentSegmentProperty7ValueContainer.From*` functions to initialize the SubscriptionComponentCustomPriceSegmentSegmentProperty7Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentCustomPriceSegmentSegmentProperty7Value object to a map representation for JSON marshaling.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty7Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty7Value.
// It customizes the JSON unmarshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty7Value objects.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty7Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty7Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty7Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty7Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionComponentCustomPriceSegmentSegmentProperty7Value represents a subscriptionComponentCustomPriceSegmentSegmentProperty7Value struct.
// This is a container for one-of cases.
type internalSubscriptionComponentCustomPriceSegmentSegmentProperty7Value struct {}

var SubscriptionComponentCustomPriceSegmentSegmentProperty7ValueContainer internalSubscriptionComponentCustomPriceSegmentSegmentProperty7Value

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty7Value instance, wrapping the provided string value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty7Value) FromString(val string) SubscriptionComponentCustomPriceSegmentSegmentProperty7Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty7Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty7Value instance, wrapping the provided float64 value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty7Value) FromPrecision(val float64) SubscriptionComponentCustomPriceSegmentSegmentProperty7Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty7Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty7Value instance, wrapping the provided int value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty7Value) FromNumber(val int) SubscriptionComponentCustomPriceSegmentSegmentProperty7Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty7Value{value: &val}
}
