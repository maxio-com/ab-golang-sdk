// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ListCreditNotesResponse represents a ListCreditNotesResponse struct.
type ListCreditNotesResponse struct {
    CreditNotes          []CreditNote           `json:"credit_notes"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListCreditNotesResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListCreditNotesResponse) String() string {
    return fmt.Sprintf(
    	"ListCreditNotesResponse[CreditNotes=%v, AdditionalProperties=%v]",
    	l.CreditNotes, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListCreditNotesResponse.
// It customizes the JSON marshaling process for ListCreditNotesResponse objects.
func (l ListCreditNotesResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "credit_notes"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListCreditNotesResponse object to a map representation for JSON marshaling.
func (l ListCreditNotesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["credit_notes"] = l.CreditNotes
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListCreditNotesResponse.
// It customizes the JSON unmarshaling process for ListCreditNotesResponse objects.
func (l *ListCreditNotesResponse) UnmarshalJSON(input []byte) error {
    var temp tempListCreditNotesResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "credit_notes")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.CreditNotes = *temp.CreditNotes
    return nil
}

// tempListCreditNotesResponse is a temporary struct used for validating the fields of ListCreditNotesResponse.
type tempListCreditNotesResponse  struct {
    CreditNotes *[]CreditNote `json:"credit_notes"`
}

func (l *tempListCreditNotesResponse) validate() error {
    var errs []string
    if l.CreditNotes == nil {
        errs = append(errs, "required field `credit_notes` is missing for type `List Credit Notes Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
