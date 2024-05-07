package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListInvoicesResponse represents a ListInvoicesResponse struct.
type ListInvoicesResponse struct {
    Invoices             []Invoice      `json:"invoices"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListInvoicesResponse.
// It customizes the JSON marshaling process for ListInvoicesResponse objects.
func (l ListInvoicesResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListInvoicesResponse object to a map representation for JSON marshaling.
func (l ListInvoicesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "invoices")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.Invoices = *temp.Invoices
    return nil
}

// listInvoicesResponse is a temporary struct used for validating the fields of ListInvoicesResponse.
type listInvoicesResponse  struct {
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
