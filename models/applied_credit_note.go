package models

import (
	"encoding/json"
)

// AppliedCreditNote represents a AppliedCreditNote struct.
type AppliedCreditNote struct {
	// The UID of the credit note
	Uid *string `json:"uid,omitempty"`
	// The number of the credit note
	Number *string `json:"number,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AppliedCreditNote.
// It customizes the JSON marshaling process for AppliedCreditNote objects.
func (a *AppliedCreditNote) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AppliedCreditNote object to a map representation for JSON marshaling.
func (a *AppliedCreditNote) toMap() map[string]any {
	structMap := make(map[string]any)
	if a.Uid != nil {
		structMap["uid"] = a.Uid
	}
	if a.Number != nil {
		structMap["number"] = a.Number
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AppliedCreditNote.
// It customizes the JSON unmarshaling process for AppliedCreditNote objects.
func (a *AppliedCreditNote) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Uid    *string `json:"uid,omitempty"`
		Number *string `json:"number,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.Uid = temp.Uid
	a.Number = temp.Number
	return nil
}
