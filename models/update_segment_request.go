package models

import (
    "encoding/json"
)

// UpdateSegmentRequest represents a UpdateSegmentRequest struct.
type UpdateSegmentRequest struct {
    Segment UpdateSegment `json:"segment"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSegmentRequest.
// It customizes the JSON marshaling process for UpdateSegmentRequest objects.
func (u *UpdateSegmentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSegmentRequest object to a map representation for JSON marshaling.
func (u *UpdateSegmentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["segment"] = u.Segment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSegmentRequest.
// It customizes the JSON unmarshaling process for UpdateSegmentRequest objects.
func (u *UpdateSegmentRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Segment UpdateSegment `json:"segment"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.Segment = temp.Segment
    return nil
}
