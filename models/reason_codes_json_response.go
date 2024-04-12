package models

import (
    "encoding/json"
)

// ReasonCodesJsonResponse represents a ReasonCodesJsonResponse struct.
type ReasonCodesJsonResponse struct {
    Ok                   *string        `json:"ok,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReasonCodesJsonResponse.
// It customizes the JSON marshaling process for ReasonCodesJsonResponse objects.
func (r ReasonCodesJsonResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReasonCodesJsonResponse object to a map representation for JSON marshaling.
func (r ReasonCodesJsonResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Ok != nil {
        structMap["ok"] = r.Ok
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReasonCodesJsonResponse.
// It customizes the JSON unmarshaling process for ReasonCodesJsonResponse objects.
func (r *ReasonCodesJsonResponse) UnmarshalJSON(input []byte) error {
    var temp reasonCodesJsonResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ok")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Ok = temp.Ok
    return nil
}

// TODO
type reasonCodesJsonResponse  struct {
    Ok *string `json:"ok,omitempty"`
}
