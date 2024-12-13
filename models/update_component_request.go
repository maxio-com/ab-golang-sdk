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

// UpdateComponentRequest represents a UpdateComponentRequest struct.
type UpdateComponentRequest struct {
    Component            UpdateComponent        `json:"component"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponentRequest.
// It customizes the JSON marshaling process for UpdateComponentRequest objects.
func (u UpdateComponentRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "component"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponentRequest object to a map representation for JSON marshaling.
func (u UpdateComponentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["component"] = u.Component.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponentRequest.
// It customizes the JSON unmarshaling process for UpdateComponentRequest objects.
func (u *UpdateComponentRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateComponentRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "component")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Component = *temp.Component
    return nil
}

// tempUpdateComponentRequest is a temporary struct used for validating the fields of UpdateComponentRequest.
type tempUpdateComponentRequest  struct {
    Component *UpdateComponent `json:"component"`
}

func (u *tempUpdateComponentRequest) validate() error {
    var errs []string
    if u.Component == nil {
        errs = append(errs, "required field `component` is missing for type `Update Component Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
