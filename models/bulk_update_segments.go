/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// BulkUpdateSegments represents a BulkUpdateSegments struct.
type BulkUpdateSegments struct {
    Segments             []BulkUpdateSegmentsItem `json:"segments,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BulkUpdateSegments.
// It customizes the JSON marshaling process for BulkUpdateSegments objects.
func (b BulkUpdateSegments) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BulkUpdateSegments object to a map representation for JSON marshaling.
func (b BulkUpdateSegments) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    if b.Segments != nil {
        structMap["segments"] = b.Segments
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkUpdateSegments.
// It customizes the JSON unmarshaling process for BulkUpdateSegments objects.
func (b *BulkUpdateSegments) UnmarshalJSON(input []byte) error {
    var temp tempBulkUpdateSegments
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "segments")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.Segments = temp.Segments
    return nil
}

// tempBulkUpdateSegments is a temporary struct used for validating the fields of BulkUpdateSegments.
type tempBulkUpdateSegments  struct {
    Segments []BulkUpdateSegmentsItem `json:"segments,omitempty"`
}
