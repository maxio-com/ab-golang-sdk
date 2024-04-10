package models

import (
    "encoding/json"
)

// SignupProformaPreview represents a SignupProformaPreview struct.
type SignupProformaPreview struct {
    CurrentProformaInvoice *ProformaInvoice `json:"current_proforma_invoice,omitempty"`
    NextProformaInvoice    *ProformaInvoice `json:"next_proforma_invoice,omitempty"`
    AdditionalProperties   map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SignupProformaPreview.
// It customizes the JSON marshaling process for SignupProformaPreview objects.
func (s SignupProformaPreview) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SignupProformaPreview object to a map representation for JSON marshaling.
func (s SignupProformaPreview) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CurrentProformaInvoice != nil {
        structMap["current_proforma_invoice"] = s.CurrentProformaInvoice.toMap()
    }
    if s.NextProformaInvoice != nil {
        structMap["next_proforma_invoice"] = s.NextProformaInvoice.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SignupProformaPreview.
// It customizes the JSON unmarshaling process for SignupProformaPreview objects.
func (s *SignupProformaPreview) UnmarshalJSON(input []byte) error {
    var temp signupProformaPreview
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "current_proforma_invoice", "next_proforma_invoice")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.CurrentProformaInvoice = temp.CurrentProformaInvoice
    s.NextProformaInvoice = temp.NextProformaInvoice
    return nil
}

// TODO
type signupProformaPreview  struct {
    CurrentProformaInvoice *ProformaInvoice `json:"current_proforma_invoice,omitempty"`
    NextProformaInvoice    *ProformaInvoice `json:"next_proforma_invoice,omitempty"`
}
