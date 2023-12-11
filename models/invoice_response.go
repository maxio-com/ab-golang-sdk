package models

import (
	"encoding/json"
)

// InvoiceResponse represents a InvoiceResponse struct.
type InvoiceResponse struct {
	Invoice Invoice `json:"invoice"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceResponse.
// It customizes the JSON marshaling process for InvoiceResponse objects.
func (i *InvoiceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceResponse object to a map representation for JSON marshaling.
func (i *InvoiceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["invoice"] = i.Invoice
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceResponse.
// It customizes the JSON unmarshaling process for InvoiceResponse objects.
func (i *InvoiceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Invoice Invoice `json:"invoice"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	i.Invoice = temp.Invoice
	return nil
}
