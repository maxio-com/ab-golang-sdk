// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionComponentCustomPriceSegmentSegmentProperty2Value represents a SubscriptionComponentCustomPriceSegmentSegmentProperty2Value struct.
// This is a container for one-of cases.
type SubscriptionComponentCustomPriceSegmentSegmentProperty2Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentCustomPriceSegmentSegmentProperty2Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty2Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty2Value.
// It customizes the JSON marshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty2Value objects.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty2Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentCustomPriceSegmentSegmentProperty2ValueContainer.From*` functions to initialize the SubscriptionComponentCustomPriceSegmentSegmentProperty2Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentCustomPriceSegmentSegmentProperty2Value object to a map representation for JSON marshaling.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty2Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty2Value.
// It customizes the JSON unmarshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty2Value objects.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty2Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty2Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty2Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty2Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionComponentCustomPriceSegmentSegmentProperty2Value represents a subscriptionComponentCustomPriceSegmentSegmentProperty2Value struct.
// This is a container for one-of cases.
type internalSubscriptionComponentCustomPriceSegmentSegmentProperty2Value struct {}

var SubscriptionComponentCustomPriceSegmentSegmentProperty2ValueContainer internalSubscriptionComponentCustomPriceSegmentSegmentProperty2Value

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty2Value instance, wrapping the provided string value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty2Value) FromString(val string) SubscriptionComponentCustomPriceSegmentSegmentProperty2Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty2Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty2Value instance, wrapping the provided float64 value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty2Value) FromPrecision(val float64) SubscriptionComponentCustomPriceSegmentSegmentProperty2Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty2Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty2Value instance, wrapping the provided int value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty2Value) FromNumber(val int) SubscriptionComponentCustomPriceSegmentSegmentProperty2Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty2Value{value: &val}
}
