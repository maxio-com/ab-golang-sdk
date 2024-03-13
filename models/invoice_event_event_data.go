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

// InvoiceEventEventData represents a InvoiceEventEventData struct.
// This is a container for any-of cases.
type InvoiceEventEventData struct {
	value                                    any
	isApplyCreditNoteEventData               bool
	isApplyDebitNoteEventData                bool
	isApplyPaymentEventData                  bool
	isChangeInvoiceCollectionMethodEventData bool
	isIssueInvoiceEventData                  bool
	isRefundInvoiceEventData                 bool
	isRemovePaymentEventData                 bool
	isVoidInvoiceEventData                   bool
	isVoidRemainderEventData                 bool
}

// String converts the InvoiceEventEventData object to a string representation.
func (i InvoiceEventEventData) String() string {
	if bytes, err := json.Marshal(i.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for InvoiceEventEventData.
// It customizes the JSON marshaling process for InvoiceEventEventData objects.
func (i *InvoiceEventEventData) MarshalJSON() (
	[]byte,
	error) {
	if i.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.InvoiceEventEventDataContainer.From*` functions to initialize the InvoiceEventEventData object.")
	}
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceEventEventData object to a map representation for JSON marshaling.
func (i *InvoiceEventEventData) toMap() any {
	switch obj := i.value.(type) {
	case *ApplyCreditNoteEventData:
		return obj.toMap()
	case *ApplyDebitNoteEventData:
		return obj.toMap()
	case *ApplyPaymentEventData:
		return obj.toMap()
	case *ChangeInvoiceCollectionMethodEventData:
		return obj.toMap()
	case *IssueInvoiceEventData:
		return obj.toMap()
	case *RefundInvoiceEventData:
		return obj.toMap()
	case *RemovePaymentEventData:
		return obj.toMap()
	case *VoidInvoiceEventData:
		return obj.toMap()
	case *VoidRemainderEventData:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEventEventData.
// It customizes the JSON unmarshaling process for InvoiceEventEventData objects.
func (i *InvoiceEventEventData) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(&ApplyCreditNoteEventData{}, false, &i.isApplyCreditNoteEventData),
		NewTypeHolder(&ApplyDebitNoteEventData{}, false, &i.isApplyDebitNoteEventData),
		NewTypeHolder(&ApplyPaymentEventData{}, false, &i.isApplyPaymentEventData),
		NewTypeHolder(&ChangeInvoiceCollectionMethodEventData{}, false, &i.isChangeInvoiceCollectionMethodEventData),
		NewTypeHolder(&IssueInvoiceEventData{}, false, &i.isIssueInvoiceEventData),
		NewTypeHolder(&RefundInvoiceEventData{}, false, &i.isRefundInvoiceEventData),
		NewTypeHolder(&RemovePaymentEventData{}, false, &i.isRemovePaymentEventData),
		NewTypeHolder(&VoidInvoiceEventData{}, false, &i.isVoidInvoiceEventData),
		NewTypeHolder(&VoidRemainderEventData{}, false, &i.isVoidRemainderEventData),
	)

	i.value = result
	return err
}

func (i *InvoiceEventEventData) AsApplyCreditNoteEventData() (
	*ApplyCreditNoteEventData,
	bool) {
	if !i.isApplyCreditNoteEventData {
		return nil, false
	}
	return i.value.(*ApplyCreditNoteEventData), true
}

func (i *InvoiceEventEventData) AsApplyDebitNoteEventData() (
	*ApplyDebitNoteEventData,
	bool) {
	if !i.isApplyDebitNoteEventData {
		return nil, false
	}
	return i.value.(*ApplyDebitNoteEventData), true
}

func (i *InvoiceEventEventData) AsApplyPaymentEventData() (
	*ApplyPaymentEventData,
	bool) {
	if !i.isApplyPaymentEventData {
		return nil, false
	}
	return i.value.(*ApplyPaymentEventData), true
}

func (i *InvoiceEventEventData) AsChangeInvoiceCollectionMethodEventData() (
	*ChangeInvoiceCollectionMethodEventData,
	bool) {
	if !i.isChangeInvoiceCollectionMethodEventData {
		return nil, false
	}
	return i.value.(*ChangeInvoiceCollectionMethodEventData), true
}

func (i *InvoiceEventEventData) AsIssueInvoiceEventData() (
	*IssueInvoiceEventData,
	bool) {
	if !i.isIssueInvoiceEventData {
		return nil, false
	}
	return i.value.(*IssueInvoiceEventData), true
}

func (i *InvoiceEventEventData) AsRefundInvoiceEventData() (
	*RefundInvoiceEventData,
	bool) {
	if !i.isRefundInvoiceEventData {
		return nil, false
	}
	return i.value.(*RefundInvoiceEventData), true
}

func (i *InvoiceEventEventData) AsRemovePaymentEventData() (
	*RemovePaymentEventData,
	bool) {
	if !i.isRemovePaymentEventData {
		return nil, false
	}
	return i.value.(*RemovePaymentEventData), true
}

func (i *InvoiceEventEventData) AsVoidInvoiceEventData() (
	*VoidInvoiceEventData,
	bool) {
	if !i.isVoidInvoiceEventData {
		return nil, false
	}
	return i.value.(*VoidInvoiceEventData), true
}

func (i *InvoiceEventEventData) AsVoidRemainderEventData() (
	*VoidRemainderEventData,
	bool) {
	if !i.isVoidRemainderEventData {
		return nil, false
	}
	return i.value.(*VoidRemainderEventData), true
}

// internalInvoiceEventEventData represents a invoiceEventEventData struct.
// This is a container for any-of cases.
type internalInvoiceEventEventData struct{}

var InvoiceEventEventDataContainer internalInvoiceEventEventData

func (i *internalInvoiceEventEventData) FromApplyCreditNoteEventData(val ApplyCreditNoteEventData) InvoiceEventEventData {
	return InvoiceEventEventData{value: &val}
}

func (i *internalInvoiceEventEventData) FromApplyDebitNoteEventData(val ApplyDebitNoteEventData) InvoiceEventEventData {
	return InvoiceEventEventData{value: &val}
}

func (i *internalInvoiceEventEventData) FromApplyPaymentEventData(val ApplyPaymentEventData) InvoiceEventEventData {
	return InvoiceEventEventData{value: &val}
}

func (i *internalInvoiceEventEventData) FromChangeInvoiceCollectionMethodEventData(val ChangeInvoiceCollectionMethodEventData) InvoiceEventEventData {
	return InvoiceEventEventData{value: &val}
}

func (i *internalInvoiceEventEventData) FromIssueInvoiceEventData(val IssueInvoiceEventData) InvoiceEventEventData {
	return InvoiceEventEventData{value: &val}
}

func (i *internalInvoiceEventEventData) FromRefundInvoiceEventData(val RefundInvoiceEventData) InvoiceEventEventData {
	return InvoiceEventEventData{value: &val}
}

func (i *internalInvoiceEventEventData) FromRemovePaymentEventData(val RemovePaymentEventData) InvoiceEventEventData {
	return InvoiceEventEventData{value: &val}
}

func (i *internalInvoiceEventEventData) FromVoidInvoiceEventData(val VoidInvoiceEventData) InvoiceEventEventData {
	return InvoiceEventEventData{value: &val}
}

func (i *internalInvoiceEventEventData) FromVoidRemainderEventData(val VoidRemainderEventData) InvoiceEventEventData {
	return InvoiceEventEventData{value: &val}
}
