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

// MetafieldEnum represents a MetafieldEnum struct.
// This is a container for one-of cases.
type MetafieldEnum struct {
    value           any
    isString        bool
    isArrayOfString bool
}

// String converts the MetafieldEnum object to a string representation.
func (m MetafieldEnum) String() string {
    if bytes, err := json.Marshal(m.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for MetafieldEnum.
// It customizes the JSON marshaling process for MetafieldEnum objects.
func (m MetafieldEnum) MarshalJSON() (
    []byte,
    error) {
    if m.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.MetafieldEnumContainer.From*` functions to initialize the MetafieldEnum object.")
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MetafieldEnum object to a map representation for JSON marshaling.
func (m *MetafieldEnum) toMap() any {
    switch obj := m.value.(type) {
    case *string:
        return *obj
    case *[]string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for MetafieldEnum.
// It customizes the JSON unmarshaling process for MetafieldEnum objects.
func (m *MetafieldEnum) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &m.isString),
        NewTypeHolder(new([]string), false, &m.isArrayOfString),
    )
    
    m.value = result
    return err
}

func (m *MetafieldEnum) AsString() (
    *string,
    bool) {
    if !m.isString {
        return nil, false
    }
    return m.value.(*string), true
}

func (m *MetafieldEnum) AsArrayOfString() (
    *[]string,
    bool) {
    if !m.isArrayOfString {
        return nil, false
    }
    return m.value.(*[]string), true
}

// internalMetafieldEnum represents a metafieldEnum struct.
// This is a container for one-of cases.
type internalMetafieldEnum struct {}

var MetafieldEnumContainer internalMetafieldEnum

// The internalMetafieldEnum instance, wrapping the provided string value.
func (m *internalMetafieldEnum) FromString(val string) MetafieldEnum {
    return MetafieldEnum{value: &val}
}

// The internalMetafieldEnum instance, wrapping the provided []string value.
func (m *internalMetafieldEnum) FromArrayOfString(val []string) MetafieldEnum {
    return MetafieldEnum{value: &val}
}
