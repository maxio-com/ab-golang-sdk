package models

import (
    "encoding/json"
)

// BulkCreateSegments represents a BulkCreateSegments struct.
type BulkCreateSegments struct {
    Segments []CreateSegment `json:"segments,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BulkCreateSegments.
// It customizes the JSON marshaling process for BulkCreateSegments objects.
func (b *BulkCreateSegments) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BulkCreateSegments object to a map representation for JSON marshaling.
func (b *BulkCreateSegments) toMap() map[string]any {
    structMap := make(map[string]any)
    if b.Segments != nil {
        structMap["segments"] = b.Segments
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkCreateSegments.
// It customizes the JSON unmarshaling process for BulkCreateSegments objects.
func (b *BulkCreateSegments) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Segments []CreateSegment `json:"segments,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    b.Segments = temp.Segments
    return nil
}
