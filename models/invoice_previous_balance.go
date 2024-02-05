package models

import (
    "encoding/json"
    "log"
    "time"
)

// InvoicePreviousBalance represents a InvoicePreviousBalance struct.
type InvoicePreviousBalance struct {
    CapturedAt *time.Time           `json:"captured_at,omitempty"`
    Invoices   []InvoiceBalanceItem `json:"invoices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoicePreviousBalance.
// It customizes the JSON marshaling process for InvoicePreviousBalance objects.
func (i *InvoicePreviousBalance) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoicePreviousBalance object to a map representation for JSON marshaling.
func (i *InvoicePreviousBalance) toMap() map[string]any {
    structMap := make(map[string]any)
    if i.CapturedAt != nil {
        structMap["captured_at"] = i.CapturedAt.Format(time.RFC3339)
    }
    if i.Invoices != nil {
        structMap["invoices"] = i.Invoices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePreviousBalance.
// It customizes the JSON unmarshaling process for InvoicePreviousBalance objects.
func (i *InvoicePreviousBalance) UnmarshalJSON(input []byte) error {
    temp := &struct {
        CapturedAt *string              `json:"captured_at,omitempty"`
        Invoices   []InvoiceBalanceItem `json:"invoices,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    if temp.CapturedAt != nil {
        CapturedAtVal, err := time.Parse(time.RFC3339, *temp.CapturedAt)
        if err != nil {
            log.Fatalf("Cannot Parse captured_at as % s format.", time.RFC3339)
        }
        i.CapturedAt = &CapturedAtVal
    }
    i.Invoices = temp.Invoices
    return nil
}
