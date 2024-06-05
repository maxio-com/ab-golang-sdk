package models

import (
    "encoding/json"
    "errors"
    "log"
    "strings"
    "time"
)

// ApplyCreditNoteEventData represents a ApplyCreditNoteEventData struct.
// Example schema for an `apply_credit_note` event
type ApplyCreditNoteEventData struct {
    // Unique identifier for the credit note application. It is generated automatically by Chargify and has the prefix "cdt_" followed by alphanumeric characters.
    Uid                  string                  `json:"uid"`
    // A unique, identifying string that appears on the credit note and in places it is referenced.
    CreditNoteNumber     string                  `json:"credit_note_number"`
    // Unique identifier for the credit note. It is generated automatically by Chargify and has the prefix "cn_" followed by alphanumeric characters.
    CreditNoteUid        string                  `json:"credit_note_uid"`
    // The full, original amount of the credit note.
    OriginalAmount       string                  `json:"original_amount"`
    // The amount of the credit note applied to invoice.
    AppliedAmount        string                  `json:"applied_amount"`
    // The time the credit note was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z"
    TransactionTime      *time.Time              `json:"transaction_time,omitempty"`
    // The credit note memo.
    Memo                 Optional[string]        `json:"memo"`
    // The role of the credit note (e.g. 'general')
    Role                 *string                 `json:"role,omitempty"`
    // Shows whether it was applied to consolidated invoice or not
    ConsolidatedInvoice  *bool                   `json:"consolidated_invoice,omitempty"`
    // List of credit notes applied to children invoices (if consolidated invoice)
    AppliedCreditNotes   []AppliedCreditNoteData `json:"applied_credit_notes,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApplyCreditNoteEventData.
// It customizes the JSON marshaling process for ApplyCreditNoteEventData objects.
func (a ApplyCreditNoteEventData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApplyCreditNoteEventData object to a map representation for JSON marshaling.
func (a ApplyCreditNoteEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["uid"] = a.Uid
    structMap["credit_note_number"] = a.CreditNoteNumber
    structMap["credit_note_uid"] = a.CreditNoteUid
    structMap["original_amount"] = a.OriginalAmount
    structMap["applied_amount"] = a.AppliedAmount
    if a.TransactionTime != nil {
        structMap["transaction_time"] = a.TransactionTime.Format(time.RFC3339)
    }
    if a.Memo.IsValueSet() {
        if a.Memo.Value() != nil {
            structMap["memo"] = a.Memo.Value()
        } else {
            structMap["memo"] = nil
        }
    }
    if a.Role != nil {
        structMap["role"] = a.Role
    }
    if a.ConsolidatedInvoice != nil {
        structMap["consolidated_invoice"] = a.ConsolidatedInvoice
    }
    if a.AppliedCreditNotes != nil {
        structMap["applied_credit_notes"] = a.AppliedCreditNotes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApplyCreditNoteEventData.
// It customizes the JSON unmarshaling process for ApplyCreditNoteEventData objects.
func (a *ApplyCreditNoteEventData) UnmarshalJSON(input []byte) error {
    var temp applyCreditNoteEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "credit_note_number", "credit_note_uid", "original_amount", "applied_amount", "transaction_time", "memo", "role", "consolidated_invoice", "applied_credit_notes")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Uid = *temp.Uid
    a.CreditNoteNumber = *temp.CreditNoteNumber
    a.CreditNoteUid = *temp.CreditNoteUid
    a.OriginalAmount = *temp.OriginalAmount
    a.AppliedAmount = *temp.AppliedAmount
    if temp.TransactionTime != nil {
        TransactionTimeVal, err := time.Parse(time.RFC3339, *temp.TransactionTime)
        if err != nil {
            log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
        }
        a.TransactionTime = &TransactionTimeVal
    }
    a.Memo = temp.Memo
    a.Role = temp.Role
    a.ConsolidatedInvoice = temp.ConsolidatedInvoice
    a.AppliedCreditNotes = temp.AppliedCreditNotes
    return nil
}

// applyCreditNoteEventData is a temporary struct used for validating the fields of ApplyCreditNoteEventData.
type applyCreditNoteEventData  struct {
    Uid                 *string                 `json:"uid"`
    CreditNoteNumber    *string                 `json:"credit_note_number"`
    CreditNoteUid       *string                 `json:"credit_note_uid"`
    OriginalAmount      *string                 `json:"original_amount"`
    AppliedAmount       *string                 `json:"applied_amount"`
    TransactionTime     *string                 `json:"transaction_time,omitempty"`
    Memo                Optional[string]        `json:"memo"`
    Role                *string                 `json:"role,omitempty"`
    ConsolidatedInvoice *bool                   `json:"consolidated_invoice,omitempty"`
    AppliedCreditNotes  []AppliedCreditNoteData `json:"applied_credit_notes,omitempty"`
}

func (a *applyCreditNoteEventData) validate() error {
    var errs []string
    if a.Uid == nil {
        errs = append(errs, "required field `uid` is missing for type `Apply Credit Note Event Data`")
    }
    if a.CreditNoteNumber == nil {
        errs = append(errs, "required field `credit_note_number` is missing for type `Apply Credit Note Event Data`")
    }
    if a.CreditNoteUid == nil {
        errs = append(errs, "required field `credit_note_uid` is missing for type `Apply Credit Note Event Data`")
    }
    if a.OriginalAmount == nil {
        errs = append(errs, "required field `original_amount` is missing for type `Apply Credit Note Event Data`")
    }
    if a.AppliedAmount == nil {
        errs = append(errs, "required field `applied_amount` is missing for type `Apply Credit Note Event Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
