package models

import (
    "encoding/json"
)

// VoidInvoiceRequest represents a VoidInvoiceRequest struct.
type VoidInvoiceRequest struct {
    Void VoidInvoice `json:"void"`
}

// MarshalJSON implements the json.Marshaler interface for VoidInvoiceRequest.
// It customizes the JSON marshaling process for VoidInvoiceRequest objects.
func (v *VoidInvoiceRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VoidInvoiceRequest object to a map representation for JSON marshaling.
func (v *VoidInvoiceRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["void"] = v.Void.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoidInvoiceRequest.
// It customizes the JSON unmarshaling process for VoidInvoiceRequest objects.
func (v *VoidInvoiceRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Void VoidInvoice `json:"void"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    v.Void = temp.Void
    return nil
}
