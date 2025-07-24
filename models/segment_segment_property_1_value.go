// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SegmentSegmentProperty1Value represents a SegmentSegmentProperty1Value struct.
// This is a container for one-of cases.
type SegmentSegmentProperty1Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
    isBoolean   bool
}

// String implements the fmt.Stringer interface for SegmentSegmentProperty1Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SegmentSegmentProperty1Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SegmentSegmentProperty1Value.
// It customizes the JSON marshaling process for SegmentSegmentProperty1Value objects.
func (s SegmentSegmentProperty1Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SegmentSegmentProperty1ValueContainer.From*` functions to initialize the SegmentSegmentProperty1Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SegmentSegmentProperty1Value object to a map representation for JSON marshaling.
func (s *SegmentSegmentProperty1Value) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *float64:
        return *obj
    case *int:
        return *obj
    case *bool:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SegmentSegmentProperty1Value.
// It customizes the JSON unmarshaling process for SegmentSegmentProperty1Value objects.
func (s *SegmentSegmentProperty1Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
        NewTypeHolder(new(bool), false, &s.isBoolean),
    )
    
    s.value = result
    return err
}

func (s *SegmentSegmentProperty1Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SegmentSegmentProperty1Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SegmentSegmentProperty1Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

func (s *SegmentSegmentProperty1Value) AsBoolean() (
    *bool,
    bool) {
    if !s.isBoolean {
        return nil, false
    }
    return s.value.(*bool), true
}

// internalSegmentSegmentProperty1Value represents a segmentSegmentProperty1Value struct.
// This is a container for one-of cases.
type internalSegmentSegmentProperty1Value struct {}

var SegmentSegmentProperty1ValueContainer internalSegmentSegmentProperty1Value

// The internalSegmentSegmentProperty1Value instance, wrapping the provided string value.
func (s *internalSegmentSegmentProperty1Value) FromString(val string) SegmentSegmentProperty1Value {
    return SegmentSegmentProperty1Value{value: &val}
}

// The internalSegmentSegmentProperty1Value instance, wrapping the provided float64 value.
func (s *internalSegmentSegmentProperty1Value) FromPrecision(val float64) SegmentSegmentProperty1Value {
    return SegmentSegmentProperty1Value{value: &val}
}

// The internalSegmentSegmentProperty1Value instance, wrapping the provided int value.
func (s *internalSegmentSegmentProperty1Value) FromNumber(val int) SegmentSegmentProperty1Value {
    return SegmentSegmentProperty1Value{value: &val}
}

// The internalSegmentSegmentProperty1Value instance, wrapping the provided bool value.
func (s *internalSegmentSegmentProperty1Value) FromBoolean(val bool) SegmentSegmentProperty1Value {
    return SegmentSegmentProperty1Value{value: &val}
}
