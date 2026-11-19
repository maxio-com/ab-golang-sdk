// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionComponentCustomPriceSegmentSegmentProperty3Value represents a SubscriptionComponentCustomPriceSegmentSegmentProperty3Value struct.
// This is a container for one-of cases.
type SubscriptionComponentCustomPriceSegmentSegmentProperty3Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
}

// String implements the fmt.Stringer interface for SubscriptionComponentCustomPriceSegmentSegmentProperty3Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty3Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty3Value.
// It customizes the JSON marshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty3Value objects.
func (s SubscriptionComponentCustomPriceSegmentSegmentProperty3Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionComponentCustomPriceSegmentSegmentProperty3ValueContainer.From*` functions to initialize the SubscriptionComponentCustomPriceSegmentSegmentProperty3Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentCustomPriceSegmentSegmentProperty3Value object to a map representation for JSON marshaling.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty3Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentCustomPriceSegmentSegmentProperty3Value.
// It customizes the JSON unmarshaling process for SubscriptionComponentCustomPriceSegmentSegmentProperty3Value objects.
func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty3Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty3Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty3Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SubscriptionComponentCustomPriceSegmentSegmentProperty3Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalSubscriptionComponentCustomPriceSegmentSegmentProperty3Value represents a subscriptionComponentCustomPriceSegmentSegmentProperty3Value struct.
// This is a container for one-of cases.
type internalSubscriptionComponentCustomPriceSegmentSegmentProperty3Value struct {}

var SubscriptionComponentCustomPriceSegmentSegmentProperty3ValueContainer internalSubscriptionComponentCustomPriceSegmentSegmentProperty3Value

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty3Value instance, wrapping the provided string value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty3Value) FromString(val string) SubscriptionComponentCustomPriceSegmentSegmentProperty3Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty3Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty3Value instance, wrapping the provided float64 value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty3Value) FromPrecision(val float64) SubscriptionComponentCustomPriceSegmentSegmentProperty3Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty3Value{value: &val}
}

// The internalSubscriptionComponentCustomPriceSegmentSegmentProperty3Value instance, wrapping the provided int value.
func (s *internalSubscriptionComponentCustomPriceSegmentSegmentProperty3Value) FromNumber(val int) SubscriptionComponentCustomPriceSegmentSegmentProperty3Value {
    return SubscriptionComponentCustomPriceSegmentSegmentProperty3Value{value: &val}
}
