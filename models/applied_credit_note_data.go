package models

import (
	"encoding/json"
)

// AppliedCreditNoteData represents a AppliedCreditNoteData struct.
type AppliedCreditNoteData struct {
	// The UID of the credit note
	Uid *string `json:"uid,omitempty"`
	// The number of the credit note
	Number *string `json:"number,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AppliedCreditNoteData.
// It customizes the JSON marshaling process for AppliedCreditNoteData objects.
func (a *AppliedCreditNoteData) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AppliedCreditNoteData object to a map representation for JSON marshaling.
func (a *AppliedCreditNoteData) toMap() map[string]any {
	structMap := make(map[string]any)
	if a.Uid != nil {
		structMap["uid"] = a.Uid
	}
	if a.Number != nil {
		structMap["number"] = a.Number
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AppliedCreditNoteData.
// It customizes the JSON unmarshaling process for AppliedCreditNoteData objects.
func (a *AppliedCreditNoteData) UnmarshalJSON(input []byte) error {
	var temp appliedCreditNoteData
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	a.Uid = temp.Uid
	a.Number = temp.Number
	return nil
}

// TODO
type appliedCreditNoteData struct {
	Uid    *string `json:"uid,omitempty"`
	Number *string `json:"number,omitempty"`
}
