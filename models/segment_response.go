package models

import (
    "encoding/json"
)

// SegmentResponse represents a SegmentResponse struct.
type SegmentResponse struct {
    Segment              *Segment       `json:"segment,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SegmentResponse.
// It customizes the JSON marshaling process for SegmentResponse objects.
func (s SegmentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SegmentResponse object to a map representation for JSON marshaling.
func (s SegmentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Segment != nil {
        structMap["segment"] = s.Segment.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SegmentResponse.
// It customizes the JSON unmarshaling process for SegmentResponse objects.
func (s *SegmentResponse) UnmarshalJSON(input []byte) error {
    var temp segmentResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "segment")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Segment = temp.Segment
    return nil
}

// TODO
type segmentResponse  struct {
    Segment *Segment `json:"segment,omitempty"`
}
