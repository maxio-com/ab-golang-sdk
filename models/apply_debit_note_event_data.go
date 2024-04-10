package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ApplyDebitNoteEventData represents a ApplyDebitNoteEventData struct.
// Example schema for an `apply_debit_note` event
type ApplyDebitNoteEventData struct {
    // A unique, identifying string that appears on the debit note and in places it is referenced.
    DebitNoteNumber      string         `json:"debit_note_number"`
    // Unique identifier for the debit note. It is generated automatically by Chargify and has the prefix "db_" followed by alphanumeric characters.
    DebitNoteUid         string         `json:"debit_note_uid"`
    // The full, original amount of the debit note.
    OriginalAmount       string         `json:"original_amount"`
    // The amount of the debit note applied to invoice.
    AppliedAmount        string         `json:"applied_amount"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApplyDebitNoteEventData.
// It customizes the JSON marshaling process for ApplyDebitNoteEventData objects.
func (a ApplyDebitNoteEventData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApplyDebitNoteEventData object to a map representation for JSON marshaling.
func (a ApplyDebitNoteEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["debit_note_number"] = a.DebitNoteNumber
    structMap["debit_note_uid"] = a.DebitNoteUid
    structMap["original_amount"] = a.OriginalAmount
    structMap["applied_amount"] = a.AppliedAmount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApplyDebitNoteEventData.
// It customizes the JSON unmarshaling process for ApplyDebitNoteEventData objects.
func (a *ApplyDebitNoteEventData) UnmarshalJSON(input []byte) error {
    var temp applyDebitNoteEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "debit_note_number", "debit_note_uid", "original_amount", "applied_amount")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.DebitNoteNumber = *temp.DebitNoteNumber
    a.DebitNoteUid = *temp.DebitNoteUid
    a.OriginalAmount = *temp.OriginalAmount
    a.AppliedAmount = *temp.AppliedAmount
    return nil
}

// TODO
type applyDebitNoteEventData  struct {
    DebitNoteNumber *string `json:"debit_note_number"`
    DebitNoteUid    *string `json:"debit_note_uid"`
    OriginalAmount  *string `json:"original_amount"`
    AppliedAmount   *string `json:"applied_amount"`
}

func (a *applyDebitNoteEventData) validate() error {
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
