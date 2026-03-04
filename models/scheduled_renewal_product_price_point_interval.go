// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ScheduledRenewalProductPricePointInterval represents a ScheduledRenewalProductPricePointInterval struct.
// This is a container for one-of cases.
type ScheduledRenewalProductPricePointInterval struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for ScheduledRenewalProductPricePointInterval,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalProductPricePointInterval) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalProductPricePointInterval.
// It customizes the JSON marshaling process for ScheduledRenewalProductPricePointInterval objects.
func (s ScheduledRenewalProductPricePointInterval) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ScheduledRenewalProductPricePointIntervalContainer.From*` functions to initialize the ScheduledRenewalProductPricePointInterval object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalProductPricePointInterval object to a map representation for JSON marshaling.
func (s *ScheduledRenewalProductPricePointInterval) toMap() any {
    switch obj := s.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalProductPricePointInterval.
// It customizes the JSON unmarshaling process for ScheduledRenewalProductPricePointInterval objects.
func (s *ScheduledRenewalProductPricePointInterval) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &s.isString),
        NewTypeHolder(new(int), false, &s.isNumber),
    )
    
    s.value = result
    return err
}

func (s *ScheduledRenewalProductPricePointInterval) AsString() (
    *string,
    bool) {
    if !s.isString {
        return nil, false
    }
    return s.value.(*string), true
}

func (s *ScheduledRenewalProductPricePointInterval) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

// internalScheduledRenewalProductPricePointInterval represents a scheduledRenewalProductPricePointInterval struct.
// This is a container for one-of cases.
type internalScheduledRenewalProductPricePointInterval struct {}

var ScheduledRenewalProductPricePointIntervalContainer internalScheduledRenewalProductPricePointInterval

// The internalScheduledRenewalProductPricePointInterval instance, wrapping the provided string value.
func (s *internalScheduledRenewalProductPricePointInterval) FromString(val string) ScheduledRenewalProductPricePointInterval {
    return ScheduledRenewalProductPricePointInterval{value: &val}
}

// The internalScheduledRenewalProductPricePointInterval instance, wrapping the provided int value.
func (s *internalScheduledRenewalProductPricePointInterval) FromNumber(val int) ScheduledRenewalProductPricePointInterval {
    return ScheduledRenewalProductPricePointInterval{value: &val}
}
