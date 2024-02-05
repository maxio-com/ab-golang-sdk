package models

import (
    "encoding/json"
)

// BulkUpdateSegments represents a BulkUpdateSegments struct.
type BulkUpdateSegments struct {
    Segments []BulkUpdateSegmentsItem `json:"segments,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BulkUpdateSegments.
// It customizes the JSON marshaling process for BulkUpdateSegments objects.
func (b *BulkUpdateSegments) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BulkUpdateSegments object to a map representation for JSON marshaling.
func (b *BulkUpdateSegments) toMap() map[string]any {
    structMap := make(map[string]any)
    if b.Segments != nil {
        structMap["segments"] = b.Segments
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkUpdateSegments.
// It customizes the JSON unmarshaling process for BulkUpdateSegments objects.
func (b *BulkUpdateSegments) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Segments []BulkUpdateSegmentsItem `json:"segments,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    b.Segments = temp.Segments
    return nil
}
