/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateInvoiceRequest represents a CreateInvoiceRequest struct.
type CreateInvoiceRequest struct {
    Invoice              CreateInvoice  `json:"invoice"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceRequest.
// It customizes the JSON marshaling process for CreateInvoiceRequest objects.
func (c CreateInvoiceRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceRequest object to a map representation for JSON marshaling.
func (c CreateInvoiceRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["invoice"] = c.Invoice.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceRequest.
// It customizes the JSON unmarshaling process for CreateInvoiceRequest objects.
func (c *CreateInvoiceRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateInvoiceRequest
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
    
    c.AdditionalProperties = additionalProperties
    c.Invoice = *temp.Invoice
    return nil
}

// tempCreateInvoiceRequest is a temporary struct used for validating the fields of CreateInvoiceRequest.
type tempCreateInvoiceRequest  struct {
    Invoice *CreateInvoice `json:"invoice"`
}

func (c *tempCreateInvoiceRequest) validate() error {
    var errs []string
    if c.Invoice == nil {
        errs = append(errs, "required field `invoice` is missing for type `Create Invoice Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
