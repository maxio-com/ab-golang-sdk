package models

import (
    "encoding/json"
)

// CreateSegmentRequest represents a CreateSegmentRequest struct.
type CreateSegmentRequest struct {
    Segment CreateSegment `json:"segment"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSegmentRequest.
// It customizes the JSON marshaling process for CreateSegmentRequest objects.
func (c *CreateSegmentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSegmentRequest object to a map representation for JSON marshaling.
func (c *CreateSegmentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["segment"] = c.Segment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSegmentRequest.
// It customizes the JSON unmarshaling process for CreateSegmentRequest objects.
func (c *CreateSegmentRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Segment CreateSegment `json:"segment"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Segment = temp.Segment
    return nil
}
