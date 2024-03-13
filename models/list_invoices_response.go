package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ListInvoicesResponse represents a ListInvoicesResponse struct.
type ListInvoicesResponse struct {
	Invoices []Invoice `json:"invoices"`
}

// MarshalJSON implements the json.Marshaler interface for ListInvoicesResponse.
// It customizes the JSON marshaling process for ListInvoicesResponse objects.
func (l *ListInvoicesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListInvoicesResponse object to a map representation for JSON marshaling.
func (l *ListInvoicesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["invoices"] = l.Invoices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListInvoicesResponse.
// It customizes the JSON unmarshaling process for ListInvoicesResponse objects.
func (l *ListInvoicesResponse) UnmarshalJSON(input []byte) error {
	var temp listInvoicesResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	l.Invoices = *temp.Invoices
	return nil
}

// TODO
type listInvoicesResponse struct {
	Invoices *[]Invoice `json:"invoices"`
}

func (l *listInvoicesResponse) validate() error {
	var errs []string
	if l.Invoices == nil {
		errs = append(errs, "required field `invoices` is missing for type `List Invoices Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
