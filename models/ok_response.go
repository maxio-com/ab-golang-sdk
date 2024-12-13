/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// OkResponse represents a OkResponse struct.
type OkResponse struct {
    Ok                   *string                `json:"ok,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OkResponse.
// It customizes the JSON marshaling process for OkResponse objects.
func (o OkResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "ok"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OkResponse object to a map representation for JSON marshaling.
func (o OkResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Ok != nil {
        structMap["ok"] = o.Ok
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OkResponse.
// It customizes the JSON unmarshaling process for OkResponse objects.
func (o *OkResponse) UnmarshalJSON(input []byte) error {
    var temp tempOkResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ok")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Ok = temp.Ok
    return nil
}

// tempOkResponse is a temporary struct used for validating the fields of OkResponse.
type tempOkResponse  struct {
    Ok *string `json:"ok,omitempty"`
}
