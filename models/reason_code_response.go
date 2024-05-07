package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ReasonCodeResponse represents a ReasonCodeResponse struct.
type ReasonCodeResponse struct {
    ReasonCode           ReasonCode     `json:"reason_code"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReasonCodeResponse.
// It customizes the JSON marshaling process for ReasonCodeResponse objects.
func (r ReasonCodeResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReasonCodeResponse object to a map representation for JSON marshaling.
func (r ReasonCodeResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["reason_code"] = r.ReasonCode.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReasonCodeResponse.
// It customizes the JSON unmarshaling process for ReasonCodeResponse objects.
func (r *ReasonCodeResponse) UnmarshalJSON(input []byte) error {
    var temp reasonCodeResponse
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
    
    r.AdditionalProperties = additionalProperties
    r.ReasonCode = *temp.ReasonCode
    return nil
}

// reasonCodeResponse is a temporary struct used for validating the fields of ReasonCodeResponse.
type reasonCodeResponse  struct {
    ReasonCode *ReasonCode `json:"reason_code"`
}

func (r *reasonCodeResponse) validate() error {
    var errs []string
    if r.ReasonCode == nil {
        errs = append(errs, "required field `reason_code` is missing for type `Reason Code Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
