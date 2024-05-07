package models

import (
    "encoding/json"
)

// ProformaInvoiceCredit represents a ProformaInvoiceCredit struct.
type ProformaInvoiceCredit struct {
    Uid                  *string        `json:"uid,omitempty"`
    Memo                 *string        `json:"memo,omitempty"`
    OriginalAmount       *string        `json:"original_amount,omitempty"`
    AppliedAmount        *string        `json:"applied_amount,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceCredit.
// It customizes the JSON marshaling process for ProformaInvoiceCredit objects.
func (p ProformaInvoiceCredit) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProformaInvoiceCredit object to a map representation for JSON marshaling.
func (p ProformaInvoiceCredit) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Uid != nil {
        structMap["uid"] = p.Uid
    }
    if p.Memo != nil {
        structMap["memo"] = p.Memo
    }
    if p.OriginalAmount != nil {
        structMap["original_amount"] = p.OriginalAmount
    }
    if p.AppliedAmount != nil {
        structMap["applied_amount"] = p.AppliedAmount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoiceCredit.
// It customizes the JSON unmarshaling process for ProformaInvoiceCredit objects.
func (p *ProformaInvoiceCredit) UnmarshalJSON(input []byte) error {
    var temp proformaInvoiceCredit
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "memo", "original_amount", "applied_amount")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Uid = temp.Uid
    p.Memo = temp.Memo
    p.OriginalAmount = temp.OriginalAmount
    p.AppliedAmount = temp.AppliedAmount
    return nil
}

// proformaInvoiceCredit is a temporary struct used for validating the fields of ProformaInvoiceCredit.
type proformaInvoiceCredit  struct {
    Uid            *string `json:"uid,omitempty"`
    Memo           *string `json:"memo,omitempty"`
    OriginalAmount *string `json:"original_amount,omitempty"`
    AppliedAmount  *string `json:"applied_amount,omitempty"`
}
