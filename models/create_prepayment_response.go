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

// CreatePrepaymentResponse represents a CreatePrepaymentResponse struct.
type CreatePrepaymentResponse struct {
    Prepayment           CreatedPrepayment      `json:"prepayment"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePrepaymentResponse.
// It customizes the JSON marshaling process for CreatePrepaymentResponse objects.
func (c CreatePrepaymentResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "prepayment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePrepaymentResponse object to a map representation for JSON marshaling.
func (c CreatePrepaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["prepayment"] = c.Prepayment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrepaymentResponse.
// It customizes the JSON unmarshaling process for CreatePrepaymentResponse objects.
func (c *CreatePrepaymentResponse) UnmarshalJSON(input []byte) error {
    var temp tempCreatePrepaymentResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "prepayment")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Prepayment = *temp.Prepayment
    return nil
}

// tempCreatePrepaymentResponse is a temporary struct used for validating the fields of CreatePrepaymentResponse.
type tempCreatePrepaymentResponse  struct {
    Prepayment *CreatedPrepayment `json:"prepayment"`
}

func (c *tempCreatePrepaymentResponse) validate() error {
    var errs []string
    if c.Prepayment == nil {
        errs = append(errs, "required field `prepayment` is missing for type `Create Prepayment Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
