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

// ReasonCodeResponse represents a ReasonCodeResponse struct.
type ReasonCodeResponse struct {
    ReasonCode           ReasonCode             `json:"reason_code"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReasonCodeResponse.
// It customizes the JSON marshaling process for ReasonCodeResponse objects.
func (r ReasonCodeResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "reason_code"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ReasonCodeResponse object to a map representation for JSON marshaling.
func (r ReasonCodeResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["reason_code"] = r.ReasonCode.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReasonCodeResponse.
// It customizes the JSON unmarshaling process for ReasonCodeResponse objects.
func (r *ReasonCodeResponse) UnmarshalJSON(input []byte) error {
    var temp tempReasonCodeResponse
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
    r.AdditionalProperties = additionalProperties
    
    r.ReasonCode = *temp.ReasonCode
    return nil
}

// tempReasonCodeResponse is a temporary struct used for validating the fields of ReasonCodeResponse.
type tempReasonCodeResponse  struct {
    ReasonCode *ReasonCode `json:"reason_code"`
}

func (r *tempReasonCodeResponse) validate() error {
    var errs []string
    if r.ReasonCode == nil {
        errs = append(errs, "required field `reason_code` is missing for type `Reason Code Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
