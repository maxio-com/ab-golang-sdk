package models

import (
	"encoding/json"
)

// ListCreditNotesResponse represents a ListCreditNotesResponse struct.
type ListCreditNotesResponse struct {
	CreditNotes []CreditNote `json:"credit_notes"`
}

// MarshalJSON implements the json.Marshaler interface for ListCreditNotesResponse.
// It customizes the JSON marshaling process for ListCreditNotesResponse objects.
func (l *ListCreditNotesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListCreditNotesResponse object to a map representation for JSON marshaling.
func (l *ListCreditNotesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["credit_notes"] = l.CreditNotes
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListCreditNotesResponse.
// It customizes the JSON unmarshaling process for ListCreditNotesResponse objects.
func (l *ListCreditNotesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CreditNotes []CreditNote `json:"credit_notes"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.CreditNotes = temp.CreditNotes
	return nil
}
