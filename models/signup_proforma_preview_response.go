/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SignupProformaPreviewResponse represents a SignupProformaPreviewResponse struct.
type SignupProformaPreviewResponse struct {
    ProformaInvoicePreview SignupProformaPreview `json:"proforma_invoice_preview"`
    AdditionalProperties   map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SignupProformaPreviewResponse.
// It customizes the JSON marshaling process for SignupProformaPreviewResponse objects.
func (s SignupProformaPreviewResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SignupProformaPreviewResponse object to a map representation for JSON marshaling.
func (s SignupProformaPreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["proforma_invoice_preview"] = s.ProformaInvoicePreview.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SignupProformaPreviewResponse.
// It customizes the JSON unmarshaling process for SignupProformaPreviewResponse objects.
func (s *SignupProformaPreviewResponse) UnmarshalJSON(input []byte) error {
    var temp tempSignupProformaPreviewResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "proforma_invoice_preview")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ProformaInvoicePreview = *temp.ProformaInvoicePreview
    return nil
}

// tempSignupProformaPreviewResponse is a temporary struct used for validating the fields of SignupProformaPreviewResponse.
type tempSignupProformaPreviewResponse  struct {
    ProformaInvoicePreview *SignupProformaPreview `json:"proforma_invoice_preview"`
}

func (s *tempSignupProformaPreviewResponse) validate() error {
    var errs []string
    if s.ProformaInvoicePreview == nil {
        errs = append(errs, "required field `proforma_invoice_preview` is missing for type `Signup Proforma Preview Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
