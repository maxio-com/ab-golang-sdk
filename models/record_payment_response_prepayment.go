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

// RecordPaymentResponsePrepayment represents a RecordPaymentResponsePrepayment struct.
// This is a container for one-of cases.
type RecordPaymentResponsePrepayment struct {
	value               any
	isInvoicePrePayment bool
}

// String converts the RecordPaymentResponsePrepayment object to a string representation.
func (r RecordPaymentResponsePrepayment) String() string {
	if bytes, err := json.Marshal(r.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for RecordPaymentResponsePrepayment.
// It customizes the JSON marshaling process for RecordPaymentResponsePrepayment objects.
func (r *RecordPaymentResponsePrepayment) MarshalJSON() (
	[]byte,
	error) {
	if r.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.RecordPaymentResponsePrepaymentContainer.From*` functions to initialize the RecordPaymentResponsePrepayment object.")
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RecordPaymentResponsePrepayment object to a map representation for JSON marshaling.
func (r *RecordPaymentResponsePrepayment) toMap() any {
	switch obj := r.value.(type) {
	case *InvoicePrePayment:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RecordPaymentResponsePrepayment.
// It customizes the JSON unmarshaling process for RecordPaymentResponsePrepayment objects.
func (r *RecordPaymentResponsePrepayment) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&InvoicePrePayment{}, false, &r.isInvoicePrePayment),
	)

	r.value = result
	return err
}

func (r *RecordPaymentResponsePrepayment) AsInvoicePrePayment() (
	*InvoicePrePayment,
	bool) {
	if !r.isInvoicePrePayment {
		return nil, false
	}
	return r.value.(*InvoicePrePayment), true
}

// internalRecordPaymentResponsePrepayment represents a recordPaymentResponsePrepayment struct.
// This is a container for one-of cases.
type internalRecordPaymentResponsePrepayment struct{}

var RecordPaymentResponsePrepaymentContainer internalRecordPaymentResponsePrepayment

func (r *internalRecordPaymentResponsePrepayment) FromInvoicePrePayment(val InvoicePrePayment) RecordPaymentResponsePrepayment {
	return RecordPaymentResponsePrepayment{value: &val}
}
