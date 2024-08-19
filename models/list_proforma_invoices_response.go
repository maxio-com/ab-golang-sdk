/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ListProformaInvoicesResponse represents a ListProformaInvoicesResponse struct.
type ListProformaInvoicesResponse struct {
    ProformaInvoices     []ProformaInvoice         `json:"proforma_invoices,omitempty"`
    Meta                 *ListProformaInvoicesMeta `json:"meta,omitempty"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListProformaInvoicesResponse.
// It customizes the JSON marshaling process for ListProformaInvoicesResponse objects.
func (l ListProformaInvoicesResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListProformaInvoicesResponse object to a map representation for JSON marshaling.
func (l ListProformaInvoicesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.ProformaInvoices != nil {
        structMap["proforma_invoices"] = l.ProformaInvoices
    }
    if l.Meta != nil {
        structMap["meta"] = l.Meta.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListProformaInvoicesResponse.
// It customizes the JSON unmarshaling process for ListProformaInvoicesResponse objects.
func (l *ListProformaInvoicesResponse) UnmarshalJSON(input []byte) error {
    var temp tempListProformaInvoicesResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "proforma_invoices", "meta")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.ProformaInvoices = temp.ProformaInvoices
    l.Meta = temp.Meta
    return nil
}

// tempListProformaInvoicesResponse is a temporary struct used for validating the fields of ListProformaInvoicesResponse.
type tempListProformaInvoicesResponse  struct {
    ProformaInvoices []ProformaInvoice         `json:"proforma_invoices,omitempty"`
    Meta             *ListProformaInvoicesMeta `json:"meta,omitempty"`
}
