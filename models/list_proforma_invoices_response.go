package models

import (
	"encoding/json"
)

// ListProformaInvoicesResponse represents a ListProformaInvoicesResponse struct.
type ListProformaInvoicesResponse struct {
	ProformaInvoices []ProformaInvoice         `json:"proforma_invoices,omitempty"`
	Meta             *ListProformaInvoicesMeta `json:"meta,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ListProformaInvoicesResponse.
// It customizes the JSON marshaling process for ListProformaInvoicesResponse objects.
func (l *ListProformaInvoicesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListProformaInvoicesResponse object to a map representation for JSON marshaling.
func (l *ListProformaInvoicesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
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
	var temp listProformaInvoicesResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	l.ProformaInvoices = temp.ProformaInvoices
	l.Meta = temp.Meta
	return nil
}

// TODO
type listProformaInvoicesResponse struct {
	ProformaInvoices []ProformaInvoice         `json:"proforma_invoices,omitempty"`
	Meta             *ListProformaInvoicesMeta `json:"meta,omitempty"`
}
