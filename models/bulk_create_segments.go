/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// BulkCreateSegments represents a BulkCreateSegments struct.
type BulkCreateSegments struct {
    Segments             []CreateSegment        `json:"segments,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BulkCreateSegments.
// It customizes the JSON marshaling process for BulkCreateSegments objects.
func (b BulkCreateSegments) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "segments"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BulkCreateSegments object to a map representation for JSON marshaling.
func (b BulkCreateSegments) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
    if b.Segments != nil {
        structMap["segments"] = b.Segments
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkCreateSegments.
// It customizes the JSON unmarshaling process for BulkCreateSegments objects.
func (b *BulkCreateSegments) UnmarshalJSON(input []byte) error {
    var temp tempBulkCreateSegments
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "segments")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.Segments = temp.Segments
    return nil
}

// tempBulkCreateSegments is a temporary struct used for validating the fields of BulkCreateSegments.
type tempBulkCreateSegments  struct {
    Segments []CreateSegment `json:"segments,omitempty"`
}
