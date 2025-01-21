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

// InvoiceResponse represents a InvoiceResponse struct.
type InvoiceResponse struct {
    Invoice              Invoice                `json:"invoice"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoiceResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceResponse) String() string {
    return fmt.Sprintf(
    	"InvoiceResponse[Invoice=%v, AdditionalProperties=%v]",
    	i.Invoice, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceResponse.
// It customizes the JSON marshaling process for InvoiceResponse objects.
func (i InvoiceResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "invoice"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceResponse object to a map representation for JSON marshaling.
func (i InvoiceResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["invoice"] = i.Invoice.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceResponse.
// It customizes the JSON unmarshaling process for InvoiceResponse objects.
func (i *InvoiceResponse) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "invoice")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Invoice = *temp.Invoice
    return nil
}

// tempInvoiceResponse is a temporary struct used for validating the fields of InvoiceResponse.
type tempInvoiceResponse  struct {
    Invoice *Invoice `json:"invoice"`
}

func (i *tempInvoiceResponse) validate() error {
    var errs []string
    if i.Invoice == nil {
        errs = append(errs, "required field `invoice` is missing for type `Invoice Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
