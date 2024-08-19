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

// VoidInvoice represents a VoidInvoice struct.
type VoidInvoice struct {
    Reason               string         `json:"reason"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VoidInvoice.
// It customizes the JSON marshaling process for VoidInvoice objects.
func (v VoidInvoice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VoidInvoice object to a map representation for JSON marshaling.
func (v VoidInvoice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    structMap["reason"] = v.Reason
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoidInvoice.
// It customizes the JSON unmarshaling process for VoidInvoice objects.
func (v *VoidInvoice) UnmarshalJSON(input []byte) error {
    var temp tempVoidInvoice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "reason")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.Reason = *temp.Reason
    return nil
}

// tempVoidInvoice is a temporary struct used for validating the fields of VoidInvoice.
type tempVoidInvoice  struct {
    Reason *string `json:"reason"`
}

func (v *tempVoidInvoice) validate() error {
    var errs []string
    if v.Reason == nil {
        errs = append(errs, "required field `reason` is missing for type `Void Invoice`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
