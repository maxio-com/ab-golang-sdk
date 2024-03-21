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

// VoidInvoiceEventDataCreditNoteAttributes represents a VoidInvoiceEventDataCreditNoteAttributes struct.
// This is a container for one-of cases.
type VoidInvoiceEventDataCreditNoteAttributes struct {
	value        any
	isCreditNote bool
}

// String converts the VoidInvoiceEventDataCreditNoteAttributes object to a string representation.
func (v VoidInvoiceEventDataCreditNoteAttributes) String() string {
	if bytes, err := json.Marshal(v.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for VoidInvoiceEventDataCreditNoteAttributes.
// It customizes the JSON marshaling process for VoidInvoiceEventDataCreditNoteAttributes objects.
func (v *VoidInvoiceEventDataCreditNoteAttributes) MarshalJSON() (
	[]byte,
	error) {
	if v.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.VoidInvoiceEventDataCreditNoteAttributesContainer.From*` functions to initialize the VoidInvoiceEventDataCreditNoteAttributes object.")
	}
	return json.Marshal(v.toMap())
}

// toMap converts the VoidInvoiceEventDataCreditNoteAttributes object to a map representation for JSON marshaling.
func (v *VoidInvoiceEventDataCreditNoteAttributes) toMap() any {
	switch obj := v.value.(type) {
	case *CreditNote:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoidInvoiceEventDataCreditNoteAttributes.
// It customizes the JSON unmarshaling process for VoidInvoiceEventDataCreditNoteAttributes objects.
func (v *VoidInvoiceEventDataCreditNoteAttributes) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&CreditNote{}, false, &v.isCreditNote),
	)

	v.value = result
	return err
}

func (v *VoidInvoiceEventDataCreditNoteAttributes) AsCreditNote() (
	*CreditNote,
	bool) {
	if !v.isCreditNote {
		return nil, false
	}
	return v.value.(*CreditNote), true
}

// internalVoidInvoiceEventDataCreditNoteAttributes represents a voidInvoiceEventDataCreditNoteAttributes struct.
// This is a container for one-of cases.
type internalVoidInvoiceEventDataCreditNoteAttributes struct{}

var VoidInvoiceEventDataCreditNoteAttributesContainer internalVoidInvoiceEventDataCreditNoteAttributes

// The internalVoidInvoiceEventDataCreditNoteAttributes instance, wrapping the provided CreditNote value.
func (v *internalVoidInvoiceEventDataCreditNoteAttributes) FromCreditNote(val CreditNote) VoidInvoiceEventDataCreditNoteAttributes {
	return VoidInvoiceEventDataCreditNoteAttributes{value: &val}
}
