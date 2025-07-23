// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// VoidInvoiceEventData represents a VoidInvoiceEventData struct.
// Example schema for an `void_invoice` event
type VoidInvoiceEventData struct {
	CreditNoteAttributes *CreditNote `json:"credit_note_attributes"`
	// The memo provided during invoice voiding.
	Memo *string `json:"memo"`
	// The amount of the void.
	AppliedAmount *string `json:"applied_amount"`
	// The time the refund was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z"
	TransactionTime *time.Time `json:"transaction_time"`
	// If true, the invoice is an advance invoice.
	IsAdvanceInvoice bool `json:"is_advance_invoice"`
	// The reason for the void.
	Reason               string                 `json:"reason"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VoidInvoiceEventData,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VoidInvoiceEventData) String() string {
	return fmt.Sprintf(
		"VoidInvoiceEventData[CreditNoteAttributes=%v, Memo=%v, AppliedAmount=%v, TransactionTime=%v, IsAdvanceInvoice=%v, Reason=%v, AdditionalProperties=%v]",
		v.CreditNoteAttributes, v.Memo, v.AppliedAmount, v.TransactionTime, v.IsAdvanceInvoice, v.Reason, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VoidInvoiceEventData.
// It customizes the JSON marshaling process for VoidInvoiceEventData objects.
func (v VoidInvoiceEventData) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(v.AdditionalProperties,
		"credit_note_attributes", "memo", "applied_amount", "transaction_time", "is_advance_invoice", "reason"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(v.toMap())
}

// toMap converts the VoidInvoiceEventData object to a map representation for JSON marshaling.
func (v VoidInvoiceEventData) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, v.AdditionalProperties)
	if v.CreditNoteAttributes != nil {
		structMap["credit_note_attributes"] = v.CreditNoteAttributes.toMap()
	} else {
		structMap["credit_note_attributes"] = nil
	}
	if v.Memo != nil {
		structMap["memo"] = v.Memo
	} else {
		structMap["memo"] = nil
	}
	if v.AppliedAmount != nil {
		structMap["applied_amount"] = v.AppliedAmount
	} else {
		structMap["applied_amount"] = nil
	}
	if v.TransactionTime != nil {
		structMap["transaction_time"] = v.TransactionTime.Format(time.RFC3339)
	} else {
		structMap["transaction_time"] = nil
	}
	structMap["is_advance_invoice"] = v.IsAdvanceInvoice
	structMap["reason"] = v.Reason
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoidInvoiceEventData.
// It customizes the JSON unmarshaling process for VoidInvoiceEventData objects.
func (v *VoidInvoiceEventData) UnmarshalJSON(input []byte) error {
	var temp tempVoidInvoiceEventData
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "credit_note_attributes", "memo", "applied_amount", "transaction_time", "is_advance_invoice", "reason")
	if err != nil {
		return err
	}
	v.AdditionalProperties = additionalProperties

	v.CreditNoteAttributes = temp.CreditNoteAttributes
	v.Memo = temp.Memo
	v.AppliedAmount = temp.AppliedAmount
	if temp.TransactionTime != nil {
		TransactionTimeVal, err := time.Parse(time.RFC3339, *temp.TransactionTime)
		if err != nil {
			log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
		}
		v.TransactionTime = &TransactionTimeVal
	}
	v.IsAdvanceInvoice = *temp.IsAdvanceInvoice
	v.Reason = *temp.Reason
	return nil
}

// tempVoidInvoiceEventData is a temporary struct used for validating the fields of VoidInvoiceEventData.
type tempVoidInvoiceEventData struct {
	CreditNoteAttributes *CreditNote `json:"credit_note_attributes"`
	Memo                 *string     `json:"memo"`
	AppliedAmount        *string     `json:"applied_amount"`
	TransactionTime      *string     `json:"transaction_time"`
	IsAdvanceInvoice     *bool       `json:"is_advance_invoice"`
	Reason               *string     `json:"reason"`
}

func (v *tempVoidInvoiceEventData) validate() error {
	var errs []string
	if v.IsAdvanceInvoice == nil {
		errs = append(errs, "required field `is_advance_invoice` is missing for type `Void Invoice Event Data`")
	}
	if v.Reason == nil {
		errs = append(errs, "required field `reason` is missing for type `Void Invoice Event Data`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
