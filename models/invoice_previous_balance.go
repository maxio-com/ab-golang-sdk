// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// InvoicePreviousBalance represents a InvoicePreviousBalance struct.
type InvoicePreviousBalance struct {
    CapturedAt           *time.Time             `json:"captured_at,omitempty"`
    Invoices             []InvoiceBalanceItem   `json:"invoices,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoicePreviousBalance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoicePreviousBalance) String() string {
    return fmt.Sprintf(
    	"InvoicePreviousBalance[CapturedAt=%v, Invoices=%v, AdditionalProperties=%v]",
    	i.CapturedAt, i.Invoices, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoicePreviousBalance.
// It customizes the JSON marshaling process for InvoicePreviousBalance objects.
func (i InvoicePreviousBalance) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "captured_at", "invoices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoicePreviousBalance object to a map representation for JSON marshaling.
func (i InvoicePreviousBalance) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
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
    var temp tempInvoicePreviousBalance
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "captured_at", "invoices")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
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

// tempInvoicePreviousBalance is a temporary struct used for validating the fields of InvoicePreviousBalance.
type tempInvoicePreviousBalance  struct {
    CapturedAt *string              `json:"captured_at,omitempty"`
    Invoices   []InvoiceBalanceItem `json:"invoices,omitempty"`
}
