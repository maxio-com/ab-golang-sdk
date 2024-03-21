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

// RefundPrepaymentAmount represents a RefundPrepaymentAmount struct.
// This is a container for one-of cases.
type RefundPrepaymentAmount struct {
	value       any
	isString    bool
	isPrecision bool
}

// String converts the RefundPrepaymentAmount object to a string representation.
func (r RefundPrepaymentAmount) String() string {
	if bytes, err := json.Marshal(r.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for RefundPrepaymentAmount.
// It customizes the JSON marshaling process for RefundPrepaymentAmount objects.
func (r *RefundPrepaymentAmount) MarshalJSON() (
	[]byte,
	error) {
	if r.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.RefundPrepaymentAmountContainer.From*` functions to initialize the RefundPrepaymentAmount object.")
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RefundPrepaymentAmount object to a map representation for JSON marshaling.
func (r *RefundPrepaymentAmount) toMap() any {
	switch obj := r.value.(type) {
	case *string:
		return *obj
	case *float64:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundPrepaymentAmount.
// It customizes the JSON unmarshaling process for RefundPrepaymentAmount objects.
func (r *RefundPrepaymentAmount) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &r.isString),
		NewTypeHolder(new(float64), false, &r.isPrecision),
	)

	r.value = result
	return err
}

func (r *RefundPrepaymentAmount) AsString() (
	*string,
	bool) {
	if !r.isString {
		return nil, false
	}
	return r.value.(*string), true
}

func (r *RefundPrepaymentAmount) AsPrecision() (
	*float64,
	bool) {
	if !r.isPrecision {
		return nil, false
	}
	return r.value.(*float64), true
}

// internalRefundPrepaymentAmount represents a refundPrepaymentAmount struct.
// This is a container for one-of cases.
type internalRefundPrepaymentAmount struct{}

var RefundPrepaymentAmountContainer internalRefundPrepaymentAmount

// The internalRefundPrepaymentAmount instance, wrapping the provided string value.
func (r *internalRefundPrepaymentAmount) FromString(val string) RefundPrepaymentAmount {
	return RefundPrepaymentAmount{value: &val}
}

// The internalRefundPrepaymentAmount instance, wrapping the provided float64 value.
func (r *internalRefundPrepaymentAmount) FromPrecision(val float64) RefundPrepaymentAmount {
	return RefundPrepaymentAmount{value: &val}
}
