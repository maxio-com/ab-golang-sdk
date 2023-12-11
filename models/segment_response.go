package models

import (
	"encoding/json"
)

// SegmentResponse represents a SegmentResponse struct.
type SegmentResponse struct {
	Segment *Segment `json:"segment,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SegmentResponse.
// It customizes the JSON marshaling process for SegmentResponse objects.
func (s *SegmentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SegmentResponse object to a map representation for JSON marshaling.
func (s *SegmentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.Segment != nil {
		structMap["segment"] = s.Segment
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SegmentResponse.
// It customizes the JSON unmarshaling process for SegmentResponse objects.
func (s *SegmentResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Segment *Segment `json:"segment,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Segment = temp.Segment
	return nil
}
