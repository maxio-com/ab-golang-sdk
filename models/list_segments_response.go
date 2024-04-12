package models

import (
    "encoding/json"
)

// ListSegmentsResponse represents a ListSegmentsResponse struct.
type ListSegmentsResponse struct {
    Segments             []Segment      `json:"segments,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListSegmentsResponse.
// It customizes the JSON marshaling process for ListSegmentsResponse objects.
func (l ListSegmentsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListSegmentsResponse object to a map representation for JSON marshaling.
func (l ListSegmentsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Segments != nil {
        structMap["segments"] = l.Segments
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSegmentsResponse.
// It customizes the JSON unmarshaling process for ListSegmentsResponse objects.
func (l *ListSegmentsResponse) UnmarshalJSON(input []byte) error {
    var temp listSegmentsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "segments")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.Segments = temp.Segments
    return nil
}

// TODO
type listSegmentsResponse  struct {
    Segments []Segment `json:"segments,omitempty"`
}
