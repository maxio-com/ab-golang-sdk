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

// RefundConsolidatedInvoiceSegmentUids represents a RefundConsolidatedInvoiceSegmentUids struct.
// This is a container for one-of cases.
type RefundConsolidatedInvoiceSegmentUids struct {
    value           any
    isArrayOfString bool
    isString        bool
}

// String implements the fmt.Stringer interface for RefundConsolidatedInvoiceSegmentUids,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RefundConsolidatedInvoiceSegmentUids) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RefundConsolidatedInvoiceSegmentUids.
// It customizes the JSON marshaling process for RefundConsolidatedInvoiceSegmentUids objects.
func (r RefundConsolidatedInvoiceSegmentUids) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.RefundConsolidatedInvoiceSegmentUidsContainer.From*` functions to initialize the RefundConsolidatedInvoiceSegmentUids object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RefundConsolidatedInvoiceSegmentUids object to a map representation for JSON marshaling.
func (r *RefundConsolidatedInvoiceSegmentUids) toMap() any {
    switch obj := r.value.(type) {
    case *[]string:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundConsolidatedInvoiceSegmentUids.
// It customizes the JSON unmarshaling process for RefundConsolidatedInvoiceSegmentUids objects.
func (r *RefundConsolidatedInvoiceSegmentUids) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new([]string), false, &r.isArrayOfString),
        NewTypeHolder(new(string), false, &r.isString),
    )
    
    r.value = result
    return err
}

func (r *RefundConsolidatedInvoiceSegmentUids) AsArrayOfString() (
    *[]string,
    bool) {
    if !r.isArrayOfString {
        return nil, false
    }
    return r.value.(*[]string), true
}

func (r *RefundConsolidatedInvoiceSegmentUids) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

// internalRefundConsolidatedInvoiceSegmentUids represents a refundConsolidatedInvoiceSegmentUids struct.
// This is a container for one-of cases.
type internalRefundConsolidatedInvoiceSegmentUids struct {}

var RefundConsolidatedInvoiceSegmentUidsContainer internalRefundConsolidatedInvoiceSegmentUids

// The internalRefundConsolidatedInvoiceSegmentUids instance, wrapping the provided []string value.
func (r *internalRefundConsolidatedInvoiceSegmentUids) FromArrayOfString(val []string) RefundConsolidatedInvoiceSegmentUids {
    return RefundConsolidatedInvoiceSegmentUids{value: &val}
}

// The internalRefundConsolidatedInvoiceSegmentUids instance, wrapping the provided string value.
func (r *internalRefundConsolidatedInvoiceSegmentUids) FromString(val string) RefundConsolidatedInvoiceSegmentUids {
    return RefundConsolidatedInvoiceSegmentUids{value: &val}
}
