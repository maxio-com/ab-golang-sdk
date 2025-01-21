/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreatePrepaymentRequest represents a CreatePrepaymentRequest struct.
type CreatePrepaymentRequest struct {
    Prepayment           CreatePrepayment       `json:"prepayment"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreatePrepaymentRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreatePrepaymentRequest) String() string {
    return fmt.Sprintf(
    	"CreatePrepaymentRequest[Prepayment=%v, AdditionalProperties=%v]",
    	c.Prepayment, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreatePrepaymentRequest.
// It customizes the JSON marshaling process for CreatePrepaymentRequest objects.
func (c CreatePrepaymentRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "prepayment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePrepaymentRequest object to a map representation for JSON marshaling.
func (c CreatePrepaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["prepayment"] = c.Prepayment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrepaymentRequest.
// It customizes the JSON unmarshaling process for CreatePrepaymentRequest objects.
func (c *CreatePrepaymentRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreatePrepaymentRequest
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

// tempCreatePrepaymentRequest is a temporary struct used for validating the fields of CreatePrepaymentRequest.
type tempCreatePrepaymentRequest  struct {
    Prepayment *CreatePrepayment `json:"prepayment"`
}

func (c *tempCreatePrepaymentRequest) validate() error {
    var errs []string
    if c.Prepayment == nil {
        errs = append(errs, "required field `prepayment` is missing for type `Create Prepayment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
