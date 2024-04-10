package models

import (
    "encoding/json"
)

// BulkCreateSegments represents a BulkCreateSegments struct.
type BulkCreateSegments struct {
    Segments             []CreateSegment `json:"segments,omitempty"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BulkCreateSegments.
// It customizes the JSON marshaling process for BulkCreateSegments objects.
func (b BulkCreateSegments) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BulkCreateSegments object to a map representation for JSON marshaling.
func (b BulkCreateSegments) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    if b.Segments != nil {
        structMap["segments"] = b.Segments
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkCreateSegments.
// It customizes the JSON unmarshaling process for BulkCreateSegments objects.
func (b *BulkCreateSegments) UnmarshalJSON(input []byte) error {
    var temp bulkCreateSegments
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

// TODO
type bulkCreateSegments  struct {
    Segments []CreateSegment `json:"segments,omitempty"`
}
