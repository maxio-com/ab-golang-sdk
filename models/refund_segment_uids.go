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

// RefundSegmentUids represents a RefundSegmentUids struct.
// This is a container for one-of cases.
type RefundSegmentUids struct {
    value           any
    isArrayOfString bool
    isString        bool
}

// String converts the RefundSegmentUids object to a string representation.
func (r RefundSegmentUids) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for RefundSegmentUids.
// It customizes the JSON marshaling process for RefundSegmentUids objects.
func (r RefundSegmentUids) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.RefundSegmentUidsContainer.From*` functions to initialize the RefundSegmentUids object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RefundSegmentUids object to a map representation for JSON marshaling.
func (r *RefundSegmentUids) toMap() any {
    switch obj := r.value.(type) {
    case *[]string:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundSegmentUids.
// It customizes the JSON unmarshaling process for RefundSegmentUids objects.
func (r *RefundSegmentUids) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new([]string), false, &r.isArrayOfString),
        NewTypeHolder(new(string), false, &r.isString),
    )
    
    r.value = result
    return err
}

func (r *RefundSegmentUids) AsArrayOfString() (
    *[]string,
    bool) {
    if !r.isArrayOfString {
        return nil, false
    }
    return r.value.(*[]string), true
}

func (r *RefundSegmentUids) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

// internalRefundSegmentUids represents a refundSegmentUids struct.
// This is a container for one-of cases.
type internalRefundSegmentUids struct {}

var RefundSegmentUidsContainer internalRefundSegmentUids

// The internalRefundSegmentUids instance, wrapping the provided []string value.
func (r *internalRefundSegmentUids) FromArrayOfString(val []string) RefundSegmentUids {
    return RefundSegmentUids{value: &val}
}

// The internalRefundSegmentUids instance, wrapping the provided string value.
func (r *internalRefundSegmentUids) FromString(val string) RefundSegmentUids {
    return RefundSegmentUids{value: &val}
}
