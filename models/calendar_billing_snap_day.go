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

// CalendarBillingSnapDay represents a CalendarBillingSnapDay struct.
// This is a container for one-of cases.
type CalendarBillingSnapDay struct {
    value    any
    isNumber bool
    isString bool
}

// String converts the CalendarBillingSnapDay object to a string representation.
func (c CalendarBillingSnapDay) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CalendarBillingSnapDay.
// It customizes the JSON marshaling process for CalendarBillingSnapDay objects.
func (c CalendarBillingSnapDay) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CalendarBillingSnapDayContainer.From*` functions to initialize the CalendarBillingSnapDay object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CalendarBillingSnapDay object to a map representation for JSON marshaling.
func (c *CalendarBillingSnapDay) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CalendarBillingSnapDay.
// It customizes the JSON unmarshaling process for CalendarBillingSnapDay objects.
func (c *CalendarBillingSnapDay) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CalendarBillingSnapDay) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CalendarBillingSnapDay) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCalendarBillingSnapDay represents a calendarBillingSnapDay struct.
// This is a container for one-of cases.
type internalCalendarBillingSnapDay struct {}

var CalendarBillingSnapDayContainer internalCalendarBillingSnapDay

// The internalCalendarBillingSnapDay instance, wrapping the provided int value.
func (c *internalCalendarBillingSnapDay) FromNumber(val int) CalendarBillingSnapDay {
    return CalendarBillingSnapDay{value: &val}
}

// The internalCalendarBillingSnapDay instance, wrapping the provided string value.
func (c *internalCalendarBillingSnapDay) FromString(val string) CalendarBillingSnapDay {
    return CalendarBillingSnapDay{value: &val}
}
