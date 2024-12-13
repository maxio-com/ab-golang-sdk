/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ListSegmentsResponse represents a ListSegmentsResponse struct.
type ListSegmentsResponse struct {
    Segments             []Segment              `json:"segments,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListSegmentsResponse.
// It customizes the JSON marshaling process for ListSegmentsResponse objects.
func (l ListSegmentsResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "segments"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSegmentsResponse object to a map representation for JSON marshaling.
func (l ListSegmentsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Segments != nil {
        structMap["segments"] = l.Segments
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSegmentsResponse.
// It customizes the JSON unmarshaling process for ListSegmentsResponse objects.
func (l *ListSegmentsResponse) UnmarshalJSON(input []byte) error {
    var temp tempListSegmentsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "segments")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Segments = temp.Segments
    return nil
}

// tempListSegmentsResponse is a temporary struct used for validating the fields of ListSegmentsResponse.
type tempListSegmentsResponse  struct {
    Segments []Segment `json:"segments,omitempty"`
}
