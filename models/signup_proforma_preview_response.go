package models

import (
	"encoding/json"
)

// SignupProformaPreviewResponse represents a SignupProformaPreviewResponse struct.
type SignupProformaPreviewResponse struct {
	ProformaInvoicePreview SignupProformaPreview `json:"proforma_invoice_preview"`
}

// MarshalJSON implements the json.Marshaler interface for SignupProformaPreviewResponse.
// It customizes the JSON marshaling process for SignupProformaPreviewResponse objects.
func (s *SignupProformaPreviewResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SignupProformaPreviewResponse object to a map representation for JSON marshaling.
func (s *SignupProformaPreviewResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["proforma_invoice_preview"] = s.ProformaInvoicePreview
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SignupProformaPreviewResponse.
// It customizes the JSON unmarshaling process for SignupProformaPreviewResponse objects.
func (s *SignupProformaPreviewResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ProformaInvoicePreview SignupProformaPreview `json:"proforma_invoice_preview"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.ProformaInvoicePreview = temp.ProformaInvoicePreview
	return nil
}
