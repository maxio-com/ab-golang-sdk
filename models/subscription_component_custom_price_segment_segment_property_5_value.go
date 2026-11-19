// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionComponentCustomPriceSegmentSegmentProperty5Value represents a SubscriptionComponentCustomPriceSegmentSegmentProperty5Value struct.
// This is a container for one-of cases.
type SubscriptionComponentCustomPriceSegmentSegmentProperty5Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentCustomPriceSegmentSegmentProperty5Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty5Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty5Value.
// It customizes the JSON marshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty5Value objects.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty5Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentCustomPriceSegmentSegmentProperty5ValueContainer.From*` functions to initialize the SubscriptionComponentCustomPriceSegmentSegmentProperty5Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentCustomPriceSegmentSegmentProperty5Value object to a map representation for JSON marshaling.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty5Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty5Value.
// It customizes the JSON unmarshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty5Value objects.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty5Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty5Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty5Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty5Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionComponentCustomPriceSegmentSegmentProperty5Value represents a subscriptionComponentCustomPriceSegmentSegmentProperty5Value struct.
// This is a container for one-of cases.
type internalSubscriptionComponentCustomPriceSegmentSegmentProperty5Value struct {}

var SubscriptionComponentCustomPriceSegmentSegmentProperty5ValueContainer internalSubscriptionComponentCustomPriceSegmentSegmentProperty5Value

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty5Value instance, wrapping the provided string value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty5Value) FromString(val string) SubscriptionComponentCustomPriceSegmentSegmentProperty5Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty5Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty5Value instance, wrapping the provided float64 value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty5Value) FromPrecision(val float64) SubscriptionComponentCustomPriceSegmentSegmentProperty5Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty5Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty5Value instance, wrapping the provided int value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty5Value) FromNumber(val int) SubscriptionComponentCustomPriceSegmentSegmentProperty5Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty5Value{value: &val}
}
