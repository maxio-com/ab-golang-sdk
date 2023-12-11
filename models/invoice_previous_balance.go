package models

import (
	"encoding/json"
)

// InvoicePreviousBalance represents a InvoicePreviousBalance struct.
type InvoicePreviousBalance struct {
	CaptureDate *string              `json:"capture_date,omitempty"`
	Invoices    []InvoiceBalanceItem `json:"invoices,omitempty"`
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
	if i.CaptureDate != nil {
		structMap["capture_date"] = i.CaptureDate
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
		CaptureDate *string              `json:"capture_date,omitempty"`
		Invoices    []InvoiceBalanceItem `json:"invoices,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	i.CaptureDate = temp.CaptureDate
	i.Invoices = temp.Invoices
	return nil
}
