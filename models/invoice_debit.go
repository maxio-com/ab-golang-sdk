/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// InvoiceDebit represents a InvoiceDebit struct.
type InvoiceDebit struct {
    Uid                  *string                `json:"uid,omitempty"`
    DebitNoteNumber      *string                `json:"debit_note_number,omitempty"`
    DebitNoteUid         *string                `json:"debit_note_uid,omitempty"`
    // The role of the debit note.
    Role                 *DebitNoteRole         `json:"role,omitempty"`
    TransactionTime      *time.Time             `json:"transaction_time,omitempty"`
    Memo                 *string                `json:"memo,omitempty"`
    OriginalAmount       *string                `json:"original_amount,omitempty"`
    AppliedAmount        *string                `json:"applied_amount,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoiceDebit,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceDebit) String() string {
    return fmt.Sprintf(
    	"InvoiceDebit[Uid=%v, DebitNoteNumber=%v, DebitNoteUid=%v, Role=%v, TransactionTime=%v, Memo=%v, OriginalAmount=%v, AppliedAmount=%v, AdditionalProperties=%v]",
    	i.Uid, i.DebitNoteNumber, i.DebitNoteUid, i.Role, i.TransactionTime, i.Memo, i.OriginalAmount, i.AppliedAmount, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceDebit.
// It customizes the JSON marshaling process for InvoiceDebit objects.
func (i InvoiceDebit) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "uid", "debit_note_number", "debit_note_uid", "role", "transaction_time", "memo", "original_amount", "applied_amount"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceDebit object to a map representation for JSON marshaling.
func (i InvoiceDebit) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Uid != nil {
        structMap["uid"] = i.Uid
    }
    if i.DebitNoteNumber != nil {
        structMap["debit_note_number"] = i.DebitNoteNumber
    }
    if i.DebitNoteUid != nil {
        structMap["debit_note_uid"] = i.DebitNoteUid
    }
    if i.Role != nil {
        structMap["role"] = i.Role
    }
    if i.TransactionTime != nil {
        structMap["transaction_time"] = i.TransactionTime.Format(time.RFC3339)
    }
    if i.Memo != nil {
        structMap["memo"] = i.Memo
    }
    if i.OriginalAmount != nil {
        structMap["original_amount"] = i.OriginalAmount
    }
    if i.AppliedAmount != nil {
        structMap["applied_amount"] = i.AppliedAmount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceDebit.
// It customizes the JSON unmarshaling process for InvoiceDebit objects.
func (i *InvoiceDebit) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceDebit
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "debit_note_number", "debit_note_uid", "role", "transaction_time", "memo", "original_amount", "applied_amount")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Uid = temp.Uid
    i.DebitNoteNumber = temp.DebitNoteNumber
    i.DebitNoteUid = temp.DebitNoteUid
    i.Role = temp.Role
    if temp.TransactionTime != nil {
        TransactionTimeVal, err := time.Parse(time.RFC3339, *temp.TransactionTime)
        if err != nil {
            log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
        }
        i.TransactionTime = &TransactionTimeVal
    }
    i.Memo = temp.Memo
    i.OriginalAmount = temp.OriginalAmount
    i.AppliedAmount = temp.AppliedAmount
    return nil
}

// tempInvoiceDebit is a temporary struct used for validating the fields of InvoiceDebit.
type tempInvoiceDebit  struct {
    Uid             *string        `json:"uid,omitempty"`
    DebitNoteNumber *string        `json:"debit_note_number,omitempty"`
    DebitNoteUid    *string        `json:"debit_note_uid,omitempty"`
    Role            *DebitNoteRole `json:"role,omitempty"`
    TransactionTime *string        `json:"transaction_time,omitempty"`
    Memo            *string        `json:"memo,omitempty"`
    OriginalAmount  *string        `json:"original_amount,omitempty"`
    AppliedAmount   *string        `json:"applied_amount,omitempty"`
}
