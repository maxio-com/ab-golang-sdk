/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// SegmentResponse represents a SegmentResponse struct.
type SegmentResponse struct {
    Segment              *Segment               `json:"segment,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SegmentResponse.
// It customizes the JSON marshaling process for SegmentResponse objects.
func (s SegmentResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "segment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SegmentResponse object to a map representation for JSON marshaling.
func (s SegmentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Segment != nil {
        structMap["segment"] = s.Segment.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SegmentResponse.
// It customizes the JSON unmarshaling process for SegmentResponse objects.
func (s *SegmentResponse) UnmarshalJSON(input []byte) error {
    var temp tempSegmentResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "segment")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Segment = temp.Segment
    return nil
}

// tempSegmentResponse is a temporary struct used for validating the fields of SegmentResponse.
type tempSegmentResponse  struct {
    Segment *Segment `json:"segment,omitempty"`
}
