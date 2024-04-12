package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// InvoiceResponse represents a InvoiceResponse struct.
type InvoiceResponse struct {
    Invoice              Invoice        `json:"invoice"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceResponse.
// It customizes the JSON marshaling process for InvoiceResponse objects.
func (i InvoiceResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceResponse object to a map representation for JSON marshaling.
func (i InvoiceResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["invoice"] = i.Invoice.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceResponse.
// It customizes the JSON unmarshaling process for InvoiceResponse objects.
func (i *InvoiceResponse) UnmarshalJSON(input []byte) error {
    var temp invoiceResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "invoice")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Invoice = *temp.Invoice
    return nil
}

// TODO
type invoiceResponse  struct {
    Invoice *Invoice `json:"invoice"`
}

func (i *invoiceResponse) validate() error {
    var errs []string
    if i.Invoice == nil {
        errs = append(errs, "required field `invoice` is missing for type `Invoice Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
