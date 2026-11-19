// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionComponentCustomPriceSegmentSegmentProperty8Value represents a SubscriptionComponentCustomPriceSegmentSegmentProperty8Value struct.
// This is a container for one-of cases.
type SubscriptionComponentCustomPriceSegmentSegmentProperty8Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentCustomPriceSegmentSegmentProperty8Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty8Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty8Value.
// It customizes the JSON marshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty8Value objects.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty8Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentCustomPriceSegmentSegmentProperty8ValueContainer.From*` functions to initialize the SubscriptionComponentCustomPriceSegmentSegmentProperty8Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentCustomPriceSegmentSegmentProperty8Value object to a map representation for JSON marshaling.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty8Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty8Value.
// It customizes the JSON unmarshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty8Value objects.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty8Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty8Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty8Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty8Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionComponentCustomPriceSegmentSegmentProperty8Value represents a subscriptionComponentCustomPriceSegmentSegmentProperty8Value struct.
// This is a container for one-of cases.
type internalSubscriptionComponentCustomPriceSegmentSegmentProperty8Value struct {}

var SubscriptionComponentCustomPriceSegmentSegmentProperty8ValueContainer internalSubscriptionComponentCustomPriceSegmentSegmentProperty8Value

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty8Value instance, wrapping the provided string value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty8Value) FromString(val string) SubscriptionComponentCustomPriceSegmentSegmentProperty8Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty8Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty8Value instance, wrapping the provided float64 value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty8Value) FromPrecision(val float64) SubscriptionComponentCustomPriceSegmentSegmentProperty8Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty8Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty8Value instance, wrapping the provided int value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty8Value) FromNumber(val int) SubscriptionComponentCustomPriceSegmentSegmentProperty8Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty8Value{value: &val}
}
