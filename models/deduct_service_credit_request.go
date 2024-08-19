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

// DeductServiceCreditRequest represents a DeductServiceCreditRequest struct.
type DeductServiceCreditRequest struct {
    Deduction            DeductServiceCredit `json:"deduction"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeductServiceCreditRequest.
// It customizes the JSON marshaling process for DeductServiceCreditRequest objects.
func (d DeductServiceCreditRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeductServiceCreditRequest object to a map representation for JSON marshaling.
func (d DeductServiceCreditRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["deduction"] = d.Deduction.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeductServiceCreditRequest.
// It customizes the JSON unmarshaling process for DeductServiceCreditRequest objects.
func (d *DeductServiceCreditRequest) UnmarshalJSON(input []byte) error {
    var temp tempDeductServiceCreditRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "deduction")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Deduction = *temp.Deduction
    return nil
}

// tempDeductServiceCreditRequest is a temporary struct used for validating the fields of DeductServiceCreditRequest.
type tempDeductServiceCreditRequest  struct {
    Deduction *DeductServiceCredit `json:"deduction"`
}

func (d *tempDeductServiceCreditRequest) validate() error {
    var errs []string
    if d.Deduction == nil {
        errs = append(errs, "required field `deduction` is missing for type `Deduct Service Credit Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
