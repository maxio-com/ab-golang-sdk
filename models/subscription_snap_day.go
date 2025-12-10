// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// SubscriptionSnapDay represents a SubscriptionSnapDay struct.
// This is a container for one-of cases.
type SubscriptionSnapDay struct {
    value     any
    isNumber  bool
    isSnapDay bool
}

// String implements the fmt.Stringer interface for SubscriptionSnapDay,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionSnapDay) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionSnapDay.
// It customizes the JSON marshaling process for SubscriptionSnapDay objects.
func (s SubscriptionSnapDay) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.SubscriptionSnapDayContainer.From*` functions to initialize the SubscriptionSnapDay object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionSnapDay object to a map representation for JSON marshaling.
func (s *SubscriptionSnapDay) toMap() any {
    switch obj := s.value.(type) {
    case *int:
        return *obj
    case *SnapDay:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionSnapDay.
// It customizes the JSON unmarshaling process for SubscriptionSnapDay objects.
func (s *SubscriptionSnapDay) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &s.isNumber),
        NewTypeHolder(new(SnapDay), false, &s.isSnapDay),
    )
    
    s.value = result
    return err
}

func (s *SubscriptionSnapDay) AsNumber() (
    *int,
    bool) {
    if !s.isNumber {
        return nil, false
    }
    return s.value.(*int), true
}

func (s *SubscriptionSnapDay) AsSnapDay() (
    *SnapDay,
    bool) {
    if !s.isSnapDay {
        return nil, false
    }
    return s.value.(*SnapDay), true
}

// internalSubscriptionSnapDay represents a subscriptionSnapDay struct.
// This is a container for one-of cases.
type internalSubscriptionSnapDay struct {}

var SubscriptionSnapDayContainer internalSubscriptionSnapDay

// The internalSubscriptionSnapDay instance, wrapping the provided int value.
func (s *internalSubscriptionSnapDay) FromNumber(val int) SubscriptionSnapDay {
    return SubscriptionSnapDay{value: &val}
}

// The internalSubscriptionSnapDay instance, wrapping the provided SnapDay value.
func (s *internalSubscriptionSnapDay) FromSnapDay(val SnapDay) SubscriptionSnapDay {
    return SubscriptionSnapDay{value: &val}
}
