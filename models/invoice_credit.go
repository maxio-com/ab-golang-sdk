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

// InvoiceCredit represents a InvoiceCredit struct.
type InvoiceCredit struct {
    Uid                  *string                `json:"uid,omitempty"`
    CreditNoteNumber     *string                `json:"credit_note_number,omitempty"`
    CreditNoteUid        *string                `json:"credit_note_uid,omitempty"`
    TransactionTime      *time.Time             `json:"transaction_time,omitempty"`
    Memo                 *string                `json:"memo,omitempty"`
    OriginalAmount       *string                `json:"original_amount,omitempty"`
    AppliedAmount        *string                `json:"applied_amount,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoiceCredit,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceCredit) String() string {
    return fmt.Sprintf(
    	"InvoiceCredit[Uid=%v, CreditNoteNumber=%v, CreditNoteUid=%v, TransactionTime=%v, Memo=%v, OriginalAmount=%v, AppliedAmount=%v, AdditionalProperties=%v]",
    	i.Uid, i.CreditNoteNumber, i.CreditNoteUid, i.TransactionTime, i.Memo, i.OriginalAmount, i.AppliedAmount, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceCredit.
// It customizes the JSON marshaling process for InvoiceCredit objects.
func (i InvoiceCredit) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "uid", "credit_note_number", "credit_note_uid", "transaction_time", "memo", "original_amount", "applied_amount"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceCredit object to a map representation for JSON marshaling.
func (i InvoiceCredit) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Uid != nil {
        structMap["uid"] = i.Uid
    }
    if i.CreditNoteNumber != nil {
        structMap["credit_note_number"] = i.CreditNoteNumber
    }
    if i.CreditNoteUid != nil {
        structMap["credit_note_uid"] = i.CreditNoteUid
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

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceCredit.
// It customizes the JSON unmarshaling process for InvoiceCredit objects.
func (i *InvoiceCredit) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceCredit
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "credit_note_number", "credit_note_uid", "transaction_time", "memo", "original_amount", "applied_amount")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Uid = temp.Uid
    i.CreditNoteNumber = temp.CreditNoteNumber
    i.CreditNoteUid = temp.CreditNoteUid
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

// tempInvoiceCredit is a temporary struct used for validating the fields of InvoiceCredit.
type tempInvoiceCredit  struct {
    Uid              *string `json:"uid,omitempty"`
    CreditNoteNumber *string `json:"credit_note_number,omitempty"`
    CreditNoteUid    *string `json:"credit_note_uid,omitempty"`
    TransactionTime  *string `json:"transaction_time,omitempty"`
    Memo             *string `json:"memo,omitempty"`
    OriginalAmount   *string `json:"original_amount,omitempty"`
    AppliedAmount    *string `json:"applied_amount,omitempty"`
}
