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

// CreateSegmentRequest represents a CreateSegmentRequest struct.
type CreateSegmentRequest struct {
    Segment              CreateSegment          `json:"segment"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSegmentRequest.
// It customizes the JSON marshaling process for CreateSegmentRequest objects.
func (c CreateSegmentRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "segment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSegmentRequest object to a map representation for JSON marshaling.
func (c CreateSegmentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["segment"] = c.Segment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSegmentRequest.
// It customizes the JSON unmarshaling process for CreateSegmentRequest objects.
func (c *CreateSegmentRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateSegmentRequest
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
    c.AdditionalProperties = additionalProperties
    
    c.Segment = *temp.Segment
    return nil
}

// tempCreateSegmentRequest is a temporary struct used for validating the fields of CreateSegmentRequest.
type tempCreateSegmentRequest  struct {
    Segment *CreateSegment `json:"segment"`
}

func (c *tempCreateSegmentRequest) validate() error {
    var errs []string
    if c.Segment == nil {
        errs = append(errs, "required field `segment` is missing for type `Create Segment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
