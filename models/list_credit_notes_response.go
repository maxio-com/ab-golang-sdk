package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListCreditNotesResponse represents a ListCreditNotesResponse struct.
type ListCreditNotesResponse struct {
    CreditNotes          []CreditNote   `json:"credit_notes"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListCreditNotesResponse.
// It customizes the JSON marshaling process for ListCreditNotesResponse objects.
func (l ListCreditNotesResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListCreditNotesResponse object to a map representation for JSON marshaling.
func (l ListCreditNotesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["credit_notes"] = l.CreditNotes
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListCreditNotesResponse.
// It customizes the JSON unmarshaling process for ListCreditNotesResponse objects.
func (l *ListCreditNotesResponse) UnmarshalJSON(input []byte) error {
    var temp listCreditNotesResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "credit_notes")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.CreditNotes = *temp.CreditNotes
    return nil
}

// TODO
type listCreditNotesResponse  struct {
    CreditNotes *[]CreditNote `json:"credit_notes"`
}

func (l *listCreditNotesResponse) validate() error {
    var errs []string
    if l.CreditNotes == nil {
        errs = append(errs, "required field `credit_notes` is missing for type `List Credit Notes Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
