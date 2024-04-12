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

// SegmentSegmentProperty3Value represents a SegmentSegmentProperty3Value struct.
// This is a container for one-of cases.
type SegmentSegmentProperty3Value struct {
    value       any
    isString    bool
    isPrecision bool
    isNumber    bool
    isBoolean   bool
}

// String converts the SegmentSegmentProperty3Value object to a string representation.
func (s SegmentSegmentProperty3Value) String() string {
    if bytes, err := json.Marshal(s.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for SegmentSegmentProperty3Value.
// It customizes the JSON marshaling process for SegmentSegmentProperty3Value objects.
func (s SegmentSegmentProperty3Value) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SegmentSegmentProperty3ValueContainer.From*` functions to initialize the SegmentSegmentProperty3Value object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SegmentSegmentProperty3Value object to a map representation for JSON marshaling.
func (s *SegmentSegmentProperty3Value) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SegmentSegmentProperty3Value.
// It customizes the JSON unmarshaling process for SegmentSegmentProperty3Value objects.
func (s *SegmentSegmentProperty3Value) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(float64), false, &s.isPrecision),
        NewTypeHolder(new(int), false, &s.isNumber),
        NewTypeHolder(new(bool), false, &s.isBoolean),
    )
    
    s.value = result
    return err
}

func (s *SegmentSegmentProperty3Value) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *SegmentSegmentProperty3Value) AsPrecision() (
    *float64,
    bool) {
    if !s.isPrecision {
        return nil, false
    }
    return s.value.(*float64), true
}

func (s *SegmentSegmentProperty3Value) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

func (s *SegmentSegmentProperty3Value) AsBoolean() (
    *bool,
    bool) {
    if !s.isBoolean {
        return nil, false
    }
    return s.value.(*bool), true
}

// internalSegmentSegmentProperty3Value represents a segmentSegmentProperty3Value struct.
// This is a container for one-of cases.
type internalSegmentSegmentProperty3Value struct {}

var SegmentSegmentProperty3ValueContainer internalSegmentSegmentProperty3Value

// The internalSegmentSegmentProperty3Value instance, wrapping the provided string value.
func (s *internalSegmentSegmentProperty3Value) FromString(val string) SegmentSegmentProperty3Value {
    return SegmentSegmentProperty3Value{value: &val}
}

// The internalSegmentSegmentProperty3Value instance, wrapping the provided float64 value.
func (s *internalSegmentSegmentProperty3Value) FromPrecision(val float64) SegmentSegmentProperty3Value {
    return SegmentSegmentProperty3Value{value: &val}
}

// The internalSegmentSegmentProperty3Value instance, wrapping the provided int value.
func (s *internalSegmentSegmentProperty3Value) FromNumber(val int) SegmentSegmentProperty3Value {
    return SegmentSegmentProperty3Value{value: &val}
}

// The internalSegmentSegmentProperty3Value instance, wrapping the provided bool value.
func (s *internalSegmentSegmentProperty3Value) FromBoolean(val bool) SegmentSegmentProperty3Value {
    return SegmentSegmentProperty3Value{value: &val}
}
