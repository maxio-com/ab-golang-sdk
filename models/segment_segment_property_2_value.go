/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SegmentSegmentProperty2Value represents a SegmentSegmentProperty2Value struct.
// This is a container for one-of cases.
type SegmentSegmentProperty2Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
    isBoolean   bool
}

// String implements the fmt.Stringer interface for SegmentSegmentProperty2Value,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SegmentSegmentProperty2Value) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SegmentSegmentProperty2Value.
// It customizes the JSON marshaling process for SegmentSegmentProperty2Value objects.
func (s SegmentSegmentProperty2Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SegmentSegmentProperty2ValueContainer.From*` functions to initialize the SegmentSegmentProperty2Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SegmentSegmentProperty2Value object to a map representation for JSON marshaling.
func (s *SegmentSegmentProperty2Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SegmentSegmentProperty2Value.
// It customizes the JSON unmarshaling process for SegmentSegmentProperty2Value objects.
func (s *SegmentSegmentProperty2Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
        NewTypeHolder(new(bool), false, &s.isBoolean),
    )
    
    s.value = result
    return err
}

func (s *SegmentSegmentProperty2Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SegmentSegmentProperty2Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SegmentSegmentProperty2Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

func (s *SegmentSegmentProperty2Value) AsBoolean() (
    *bool,
    bool) {
    if !s.isBoolean {
        return nil, false
    }
    return s.value.(*bool), true
}

// internalSegmentSegmentProperty2Value represents a segmentSegmentProperty2Value struct.
// This is a container for one-of cases.
type internalSegmentSegmentProperty2Value struct {}

var SegmentSegmentProperty2ValueContainer internalSegmentSegmentProperty2Value

// The internalSegmentSegmentProperty2Value instance, wrapping the provided string value.
func (s *internalSegmentSegmentProperty2Value) FromString(val string) SegmentSegmentProperty2Value {
    return SegmentSegmentProperty2Value{value: &val}
}

// The internalSegmentSegmentProperty2Value instance, wrapping the provided float64 value.
func (s *internalSegmentSegmentProperty2Value) FromPrecision(val float64) SegmentSegmentProperty2Value {
    return SegmentSegmentProperty2Value{value: &val}
}

// The internalSegmentSegmentProperty2Value instance, wrapping the provided int value.
func (s *internalSegmentSegmentProperty2Value) FromNumber(val int) SegmentSegmentProperty2Value {
    return SegmentSegmentProperty2Value{value: &val}
}

// The internalSegmentSegmentProperty2Value instance, wrapping the provided bool value.
func (s *internalSegmentSegmentProperty2Value) FromBoolean(val bool) SegmentSegmentProperty2Value {
    return SegmentSegmentProperty2Value{value: &val}
}
