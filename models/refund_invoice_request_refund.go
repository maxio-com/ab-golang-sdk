// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// RefundInvoiceRequestRefund represents a RefundInvoiceRequestRefund struct.
// This is a container for any-of cases.
type RefundInvoiceRequestRefund struct {
	value                       any
	isRefundInvoice             bool
	isRefundConsolidatedInvoice bool
}

// String implements the fmt.Stringer interface for RefundInvoiceRequestRefund,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RefundInvoiceRequestRefund) String() string {
	return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RefundInvoiceRequestRefund.
// It customizes the JSON marshaling process for RefundInvoiceRequestRefund objects.
func (r RefundInvoiceRequestRefund) MarshalJSON() (
	[]byte,
	error) {
	if r.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.RefundInvoiceRequestRefundContainer.From*` functions to initialize the RefundInvoiceRequestRefund object.")
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RefundInvoiceRequestRefund object to a map representation for JSON marshaling.
func (r *RefundInvoiceRequestRefund) toMap() any {
	switch obj := r.value.(type) {
	case *RefundInvoice:
		return obj.toMap()
	case *RefundConsolidatedInvoice:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundInvoiceRequestRefund.
// It customizes the JSON unmarshaling process for RefundInvoiceRequestRefund objects.
func (r *RefundInvoiceRequestRefund) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(&RefundInvoice{}, false, &r.isRefundInvoice),
		NewTypeHolder(&RefundConsolidatedInvoice{}, false, &r.isRefundConsolidatedInvoice),
	)

	r.value = result
	return err
}

func (r *RefundInvoiceRequestRefund) AsRefundInvoice() (
	*RefundInvoice,
	bool) {
	if !r.isRefundInvoice {
		return nil, false
	}
	return r.value.(*RefundInvoice), true
}

func (r *RefundInvoiceRequestRefund) AsRefundConsolidatedInvoice() (
	*RefundConsolidatedInvoice,
	bool) {
	if !r.isRefundConsolidatedInvoice {
		return nil, false
	}
	return r.value.(*RefundConsolidatedInvoice), true
}

// internalRefundInvoiceRequestRefund represents a refundInvoiceRequestRefund struct.
// This is a container for any-of cases.
type internalRefundInvoiceRequestRefund struct{}

var RefundInvoiceRequestRefundContainer internalRefundInvoiceRequestRefund

// The internalRefundInvoiceRequestRefund instance, wrapping the provided RefundInvoice value.
func (r *internalRefundInvoiceRequestRefund) FromRefundInvoice(val RefundInvoice) RefundInvoiceRequestRefund {
	return RefundInvoiceRequestRefund{value: &val}
}

// The internalRefundInvoiceRequestRefund instance, wrapping the provided RefundConsolidatedInvoice value.
func (r *internalRefundInvoiceRequestRefund) FromRefundConsolidatedInvoice(val RefundConsolidatedInvoice) RefundInvoiceRequestRefund {
	return RefundInvoiceRequestRefund{value: &val}
}
