package models

import (
    "encoding/json"
)

// ConsolidatedInvoice represents a ConsolidatedInvoice struct.
type ConsolidatedInvoice struct {
    Invoices []Invoice `json:"invoices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ConsolidatedInvoice.
// It customizes the JSON marshaling process for ConsolidatedInvoice objects.
func (c *ConsolidatedInvoice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConsolidatedInvoice object to a map representation for JSON marshaling.
func (c *ConsolidatedInvoice) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Invoices != nil {
        structMap["invoices"] = c.Invoices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConsolidatedInvoice.
// It customizes the JSON unmarshaling process for ConsolidatedInvoice objects.
func (c *ConsolidatedInvoice) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Invoices []Invoice `json:"invoices,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Invoices = temp.Invoices
    return nil
}
