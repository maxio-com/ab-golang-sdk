/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ConsolidatedInvoice represents a ConsolidatedInvoice struct.
type ConsolidatedInvoice struct {
    Invoices             []Invoice              `json:"invoices,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConsolidatedInvoice.
// It customizes the JSON marshaling process for ConsolidatedInvoice objects.
func (c ConsolidatedInvoice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "invoices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConsolidatedInvoice object to a map representation for JSON marshaling.
func (c ConsolidatedInvoice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Invoices != nil {
        structMap["invoices"] = c.Invoices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConsolidatedInvoice.
// It customizes the JSON unmarshaling process for ConsolidatedInvoice objects.
func (c *ConsolidatedInvoice) UnmarshalJSON(input []byte) error {
    var temp tempConsolidatedInvoice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "invoices")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Invoices = temp.Invoices
    return nil
}

// tempConsolidatedInvoice is a temporary struct used for validating the fields of ConsolidatedInvoice.
type tempConsolidatedInvoice  struct {
    Invoices []Invoice `json:"invoices,omitempty"`
}
