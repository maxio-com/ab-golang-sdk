// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ScheduledRenewalProductPricePointPriceInCents represents a ScheduledRenewalProductPricePointPriceInCents struct.
// This is a container for one-of cases.
type ScheduledRenewalProductPricePointPriceInCents struct {
    value    any
    isString bool
    isLong   bool
}

// String implements the fmt.Stringer interface for ScheduledRenewalProductPricePointPriceInCents,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalProductPricePointPriceInCents) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalProductPricePointPriceInCents.
// It customizes the JSON marshaling process for ScheduledRenewalProductPricePointPriceInCents objects.
func (s ScheduledRenewalProductPricePointPriceInCents) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ScheduledRenewalProductPricePointPriceInCentsContainer.From*` functions to initialize the ScheduledRenewalProductPricePointPriceInCents object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalProductPricePointPriceInCents object to a map representation for JSON marshaling.
func (s *ScheduledRenewalProductPricePointPriceInCents) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int64:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalProductPricePointPriceInCents.
// It customizes the JSON unmarshaling process for ScheduledRenewalProductPricePointPriceInCents objects.
func (s *ScheduledRenewalProductPricePointPriceInCents) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int64), false, &s.isLong),
    )
    
    s.value = result
    return err
}

func (s *ScheduledRenewalProductPricePointPriceInCents) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *ScheduledRenewalProductPricePointPriceInCents) AsLong() (
    *int64,
    bool) {
    if !s.isLong {
        return nil, false
    }
    return s.value.(*int64), true
}

// internalScheduledRenewalProductPricePointPriceInCents represents a scheduledRenewalProductPricePointPriceInCents struct.
// This is a container for one-of cases.
type internalScheduledRenewalProductPricePointPriceInCents struct {}

var ScheduledRenewalProductPricePointPriceInCentsContainer internalScheduledRenewalProductPricePointPriceInCents

// The internalScheduledRenewalProductPricePointPriceInCents instance, wrapping the provided string value.
func (s *internalScheduledRenewalProductPricePointPriceInCents) FromString(val string) ScheduledRenewalProductPricePointPriceInCents {
    return ScheduledRenewalProductPricePointPriceInCents{value: &val}
}

// The internalScheduledRenewalProductPricePointPriceInCents instance, wrapping the provided int64 value.
func (s *internalScheduledRenewalProductPricePointPriceInCents) FromLong(val int64) ScheduledRenewalProductPricePointPriceInCents {
    return ScheduledRenewalProductPricePointPriceInCents{value: &val}
}
