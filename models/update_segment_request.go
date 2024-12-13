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

// UpdateSegmentRequest represents a UpdateSegmentRequest struct.
type UpdateSegmentRequest struct {
    Segment              UpdateSegment          `json:"segment"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSegmentRequest.
// It customizes the JSON marshaling process for UpdateSegmentRequest objects.
func (u UpdateSegmentRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "segment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSegmentRequest object to a map representation for JSON marshaling.
func (u UpdateSegmentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["segment"] = u.Segment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSegmentRequest.
// It customizes the JSON unmarshaling process for UpdateSegmentRequest objects.
func (u *UpdateSegmentRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateSegmentRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "segment")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Segment = *temp.Segment
    return nil
}

// tempUpdateSegmentRequest is a temporary struct used for validating the fields of UpdateSegmentRequest.
type tempUpdateSegmentRequest  struct {
    Segment *UpdateSegment `json:"segment"`
}

func (u *tempUpdateSegmentRequest) validate() error {
    var errs []string
    if u.Segment == nil {
        errs = append(errs, "required field `segment` is missing for type `Update Segment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
