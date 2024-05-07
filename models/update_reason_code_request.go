package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UpdateReasonCodeRequest represents a UpdateReasonCodeRequest struct.
type UpdateReasonCodeRequest struct {
    ReasonCode           UpdateReasonCode `json:"reason_code"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateReasonCodeRequest.
// It customizes the JSON marshaling process for UpdateReasonCodeRequest objects.
func (u UpdateReasonCodeRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateReasonCodeRequest object to a map representation for JSON marshaling.
func (u UpdateReasonCodeRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["reason_code"] = u.ReasonCode.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateReasonCodeRequest.
// It customizes the JSON unmarshaling process for UpdateReasonCodeRequest objects.
func (u *UpdateReasonCodeRequest) UnmarshalJSON(input []byte) error {
    var temp updateReasonCodeRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "reason_code")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.ReasonCode = *temp.ReasonCode
    return nil
}

// updateReasonCodeRequest is a temporary struct used for validating the fields of UpdateReasonCodeRequest.
type updateReasonCodeRequest  struct {
    ReasonCode *UpdateReasonCode `json:"reason_code"`
}

func (u *updateReasonCodeRequest) validate() error {
    var errs []string
    if u.ReasonCode == nil {
        errs = append(errs, "required field `reason_code` is missing for type `Update Reason Code Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
