/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ListInvoicesResponse represents a ListInvoicesResponse struct.
type ListInvoicesResponse struct {
    Invoices             []Invoice              `json:"invoices"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListInvoicesResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListInvoicesResponse) String() string {
    return fmt.Sprintf(
    	"ListInvoicesResponse[Invoices=%v, AdditionalProperties=%v]",
    	l.Invoices, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListInvoicesResponse.
// It customizes the JSON marshaling process for ListInvoicesResponse objects.
func (l ListInvoicesResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "invoices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListInvoicesResponse object to a map representation for JSON marshaling.
func (l ListInvoicesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["invoices"] = l.Invoices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListInvoicesResponse.
// It customizes the JSON unmarshaling process for ListInvoicesResponse objects.
func (l *ListInvoicesResponse) UnmarshalJSON(input []byte) error {
    var temp tempListInvoicesResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "invoices")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Invoices = *temp.Invoices
    return nil
}

// tempListInvoicesResponse is a temporary struct used for validating the fields of ListInvoicesResponse.
type tempListInvoicesResponse  struct {
    Invoices *[]Invoice `json:"invoices"`
}

func (l *tempListInvoicesResponse) validate() error {
    var errs []string
    if l.Invoices == nil {
        errs = append(errs, "required field `invoices` is missing for type `List Invoices Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
