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

// CreateReasonCodeRequest represents a CreateReasonCodeRequest struct.
type CreateReasonCodeRequest struct {
    ReasonCode           CreateReasonCode       `json:"reason_code"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateReasonCodeRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateReasonCodeRequest) String() string {
    return fmt.Sprintf(
    	"CreateReasonCodeRequest[ReasonCode=%v, AdditionalProperties=%v]",
    	c.ReasonCode, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateReasonCodeRequest.
// It customizes the JSON marshaling process for CreateReasonCodeRequest objects.
func (c CreateReasonCodeRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "reason_code"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateReasonCodeRequest object to a map representation for JSON marshaling.
func (c CreateReasonCodeRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["reason_code"] = c.ReasonCode.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateReasonCodeRequest.
// It customizes the JSON unmarshaling process for CreateReasonCodeRequest objects.
func (c *CreateReasonCodeRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateReasonCodeRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "reason_code")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.ReasonCode = *temp.ReasonCode
    return nil
}

// tempCreateReasonCodeRequest is a temporary struct used for validating the fields of CreateReasonCodeRequest.
type tempCreateReasonCodeRequest  struct {
    ReasonCode *CreateReasonCode `json:"reason_code"`
}

func (c *tempCreateReasonCodeRequest) validate() error {
    var errs []string
    if c.ReasonCode == nil {
        errs = append(errs, "required field `reason_code` is missing for type `Create Reason Code Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
