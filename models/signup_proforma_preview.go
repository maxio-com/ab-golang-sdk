package models

import (
	"encoding/json"
)

// SignupProformaPreview represents a SignupProformaPreview struct.
type SignupProformaPreview struct {
	CurrentProformaInvoice *ProformaInvoice `json:"current_proforma_invoice,omitempty"`
	NextProformaInvoice    *ProformaInvoice `json:"next_proforma_invoice,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SignupProformaPreview.
// It customizes the JSON marshaling process for SignupProformaPreview objects.
func (s *SignupProformaPreview) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SignupProformaPreview object to a map representation for JSON marshaling.
func (s *SignupProformaPreview) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.CurrentProformaInvoice != nil {
		structMap["current_proforma_invoice"] = s.CurrentProformaInvoice
	}
	if s.NextProformaInvoice != nil {
		structMap["next_proforma_invoice"] = s.NextProformaInvoice
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SignupProformaPreview.
// It customizes the JSON unmarshaling process for SignupProformaPreview objects.
func (s *SignupProformaPreview) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CurrentProformaInvoice *ProformaInvoice `json:"current_proforma_invoice,omitempty"`
		NextProformaInvoice    *ProformaInvoice `json:"next_proforma_invoice,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.CurrentProformaInvoice = temp.CurrentProformaInvoice
	s.NextProformaInvoice = temp.NextProformaInvoice
	return nil
}
