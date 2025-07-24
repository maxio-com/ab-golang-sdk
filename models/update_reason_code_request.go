// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UpdateReasonCodeRequest represents a UpdateReasonCodeRequest struct.
type UpdateReasonCodeRequest struct {
    ReasonCode           UpdateReasonCode       `json:"reason_code"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateReasonCodeRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateReasonCodeRequest) String() string {
    return fmt.Sprintf(
    	"UpdateReasonCodeRequest[ReasonCode=%v, AdditionalProperties=%v]",
    	u.ReasonCode, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateReasonCodeRequest.
// It customizes the JSON marshaling process for UpdateReasonCodeRequest objects.
func (u UpdateReasonCodeRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "reason_code"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateReasonCodeRequest object to a map representation for JSON marshaling.
func (u UpdateReasonCodeRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["reason_code"] = u.ReasonCode.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateReasonCodeRequest.
// It customizes the JSON unmarshaling process for UpdateReasonCodeRequest objects.
func (u *UpdateReasonCodeRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateReasonCodeRequest
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
    u.AdditionalProperties = additionalProperties
    
    u.ReasonCode = *temp.ReasonCode
    return nil
}

// tempUpdateReasonCodeRequest is a temporary struct used for validating the fields of UpdateReasonCodeRequest.
type tempUpdateReasonCodeRequest  struct {
    ReasonCode *UpdateReasonCode `json:"reason_code"`
}

func (u *tempUpdateReasonCodeRequest) validate() error {
    var errs []string
    if u.ReasonCode == nil {
        errs = append(errs, "required field `reason_code` is missing for type `Update Reason Code Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
