/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SegmentSegmentProperty4Value represents a SegmentSegmentProperty4Value struct.
// This is a container for one-of cases.
type SegmentSegmentProperty4Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
    isBoolean   bool
}

// String converts the SegmentSegmentProperty4Value object to a string representation.
func (s SegmentSegmentProperty4Value) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for SegmentSegmentProperty4Value.
// It customizes the JSON marshaling process for SegmentSegmentProperty4Value objects.
func (s SegmentSegmentProperty4Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SegmentSegmentProperty4ValueContainer.From*` functions to initialize the SegmentSegmentProperty4Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SegmentSegmentProperty4Value object to a map representation for JSON marshaling.
func (s *SegmentSegmentProperty4Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SegmentSegmentProperty4Value.
// It customizes the JSON unmarshaling process for SegmentSegmentProperty4Value objects.
func (s *SegmentSegmentProperty4Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
        NewTypeHolder(new(bool), false, &s.isBoolean),
    )
    
    s.value = result
    return err
}

func (s *SegmentSegmentProperty4Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SegmentSegmentProperty4Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SegmentSegmentProperty4Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

func (s *SegmentSegmentProperty4Value) AsBoolean() (
    *bool,
    bool) {
    if !s.isBoolean {
        return nil, false
    }
    return s.value.(*bool), true
}

// internalSegmentSegmentProperty4Value represents a segmentSegmentProperty4Value struct.
// This is a container for one-of cases.
type internalSegmentSegmentProperty4Value struct {}

var SegmentSegmentProperty4ValueContainer internalSegmentSegmentProperty4Value

// The internalSegmentSegmentProperty4Value instance, wrapping the provided string value.
func (s *internalSegmentSegmentProperty4Value) FromString(val string) SegmentSegmentProperty4Value {
    return SegmentSegmentProperty4Value{value: &val}
}

// The internalSegmentSegmentProperty4Value instance, wrapping the provided float64 value.
func (s *internalSegmentSegmentProperty4Value) FromPrecision(val float64) SegmentSegmentProperty4Value {
    return SegmentSegmentProperty4Value{value: &val}
}

// The internalSegmentSegmentProperty4Value instance, wrapping the provided int value.
func (s *internalSegmentSegmentProperty4Value) FromNumber(val int) SegmentSegmentProperty4Value {
    return SegmentSegmentProperty4Value{value: &val}
}

// The internalSegmentSegmentProperty4Value instance, wrapping the provided bool value.
func (s *internalSegmentSegmentProperty4Value) FromBoolean(val bool) SegmentSegmentProperty4Value {
    return SegmentSegmentProperty4Value{value: &val}
}
