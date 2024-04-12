package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// VoidInvoiceRequest represents a VoidInvoiceRequest struct.
type VoidInvoiceRequest struct {
    Void                 VoidInvoice    `json:"void"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VoidInvoiceRequest.
// It customizes the JSON marshaling process for VoidInvoiceRequest objects.
func (v VoidInvoiceRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VoidInvoiceRequest object to a map representation for JSON marshaling.
func (v VoidInvoiceRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    structMap["void"] = v.Void.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoidInvoiceRequest.
// It customizes the JSON unmarshaling process for VoidInvoiceRequest objects.
func (v *VoidInvoiceRequest) UnmarshalJSON(input []byte) error {
    var temp voidInvoiceRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "void")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.Void = *temp.Void
    return nil
}

// TODO
type voidInvoiceRequest  struct {
    Void *VoidInvoice `json:"void"`
}

func (v *voidInvoiceRequest) validate() error {
    var errs []string
    if v.Void == nil {
        errs = append(errs, "required field `void` is missing for type `Void Invoice Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
