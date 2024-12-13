/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ProformaInvoicePayment represents a ProformaInvoicePayment struct.
type ProformaInvoicePayment struct {
    Memo                 *string                `json:"memo,omitempty"`
    OriginalAmount       *string                `json:"original_amount,omitempty"`
    AppliedAmount        *string                `json:"applied_amount,omitempty"`
    Prepayment           *bool                  `json:"prepayment,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaInvoicePayment.
// It customizes the JSON marshaling process for ProformaInvoicePayment objects.
func (p ProformaInvoicePayment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "memo", "original_amount", "applied_amount", "prepayment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the ProformaInvoicePayment object to a map representation for JSON marshaling.
func (p ProformaInvoicePayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Memo != nil {
        structMap["memo"] = p.Memo
    }
    if p.OriginalAmount != nil {
        structMap["original_amount"] = p.OriginalAmount
    }
    if p.AppliedAmount != nil {
        structMap["applied_amount"] = p.AppliedAmount
    }
    if p.Prepayment != nil {
        structMap["prepayment"] = p.Prepayment
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoicePayment.
// It customizes the JSON unmarshaling process for ProformaInvoicePayment objects.
func (p *ProformaInvoicePayment) UnmarshalJSON(input []byte) error {
    var temp tempProformaInvoicePayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "memo", "original_amount", "applied_amount", "prepayment")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Memo = temp.Memo
    p.OriginalAmount = temp.OriginalAmount
    p.AppliedAmount = temp.AppliedAmount
    p.Prepayment = temp.Prepayment
    return nil
}

// tempProformaInvoicePayment is a temporary struct used for validating the fields of ProformaInvoicePayment.
type tempProformaInvoicePayment  struct {
    Memo           *string `json:"memo,omitempty"`
    OriginalAmount *string `json:"original_amount,omitempty"`
    AppliedAmount  *string `json:"applied_amount,omitempty"`
    Prepayment     *bool   `json:"prepayment,omitempty"`
}
