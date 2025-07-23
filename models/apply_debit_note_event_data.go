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

// ApplyDebitNoteEventData represents a ApplyDebitNoteEventData struct.
// Example schema for an `apply_debit_note` event
type ApplyDebitNoteEventData struct {
	// A unique, identifying string that appears on the debit note and in places it is referenced.
	DebitNoteNumber string `json:"debit_note_number"`
	// Unique identifier for the debit note. It is generated automatically by Chargify and has the prefix "db_" followed by alphanumeric characters.
	DebitNoteUid string `json:"debit_note_uid"`
	// The full, original amount of the debit note.
	OriginalAmount string `json:"original_amount"`
	// The amount of the debit note applied to invoice.
	AppliedAmount string `json:"applied_amount"`
	// The debit note memo.
	Memo Optional[string] `json:"memo"`
	// The time the debit note was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z"
	TransactionTime      Optional[time.Time]    `json:"transaction_time"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApplyDebitNoteEventData,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApplyDebitNoteEventData) String() string {
	return fmt.Sprintf(
		"ApplyDebitNoteEventData[DebitNoteNumber=%v, DebitNoteUid=%v, OriginalAmount=%v, AppliedAmount=%v, Memo=%v, TransactionTime=%v, AdditionalProperties=%v]",
		a.DebitNoteNumber, a.DebitNoteUid, a.OriginalAmount, a.AppliedAmount, a.Memo, a.TransactionTime, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApplyDebitNoteEventData.
// It customizes the JSON marshaling process for ApplyDebitNoteEventData objects.
func (a ApplyDebitNoteEventData) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"debit_note_number", "debit_note_uid", "original_amount", "applied_amount", "memo", "transaction_time"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the ApplyDebitNoteEventData object to a map representation for JSON marshaling.
func (a ApplyDebitNoteEventData) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	structMap["debit_note_number"] = a.DebitNoteNumber
	structMap["debit_note_uid"] = a.DebitNoteUid
	structMap["original_amount"] = a.OriginalAmount
	structMap["applied_amount"] = a.AppliedAmount
	if a.Memo.IsValueSet() {
		if a.Memo.Value() != nil {
			structMap["memo"] = a.Memo.Value()
		} else {
			structMap["memo"] = nil
		}
	}
	if a.TransactionTime.IsValueSet() {
		var TransactionTimeVal *string = nil
		if a.TransactionTime.Value() != nil {
			val := a.TransactionTime.Value().Format(time.RFC3339)
			TransactionTimeVal = &val
		}
		if a.TransactionTime.Value() != nil {
			structMap["transaction_time"] = TransactionTimeVal
		} else {
			structMap["transaction_time"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApplyDebitNoteEventData.
// It customizes the JSON unmarshaling process for ApplyDebitNoteEventData objects.
func (a *ApplyDebitNoteEventData) UnmarshalJSON(input []byte) error {
	var temp tempApplyDebitNoteEventData
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "debit_note_number", "debit_note_uid", "original_amount", "applied_amount", "memo", "transaction_time")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.DebitNoteNumber = *temp.DebitNoteNumber
	a.DebitNoteUid = *temp.DebitNoteUid
	a.OriginalAmount = *temp.OriginalAmount
	a.AppliedAmount = *temp.AppliedAmount
	a.Memo = temp.Memo
	a.TransactionTime.ShouldSetValue(temp.TransactionTime.IsValueSet())
	if temp.TransactionTime.Value() != nil {
		TransactionTimeVal, err := time.Parse(time.RFC3339, (*temp.TransactionTime.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
		}
		a.TransactionTime.SetValue(&TransactionTimeVal)
	}
	return nil
}

// tempApplyDebitNoteEventData is a temporary struct used for validating the fields of ApplyDebitNoteEventData.
type tempApplyDebitNoteEventData struct {
	DebitNoteNumber *string          `json:"debit_note_number"`
	DebitNoteUid    *string          `json:"debit_note_uid"`
	OriginalAmount  *string          `json:"original_amount"`
	AppliedAmount   *string          `json:"applied_amount"`
	Memo            Optional[string] `json:"memo"`
	TransactionTime Optional[string] `json:"transaction_time"`
}

func (a *tempApplyDebitNoteEventData) validate() error {
	var errs []string
	if a.DebitNoteNumber == nil {
		errs = append(errs, "required field `debit_note_number` is missing for type `Apply Debit Note Event Data`")
	}
	if a.DebitNoteUid == nil {
		errs = append(errs, "required field `debit_note_uid` is missing for type `Apply Debit Note Event Data`")
	}
	if a.OriginalAmount == nil {
		errs = append(errs, "required field `original_amount` is missing for type `Apply Debit Note Event Data`")
	}
	if a.AppliedAmount == nil {
		errs = append(errs, "required field `applied_amount` is missing for type `Apply Debit Note Event Data`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
