package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UpdateSegmentRequest represents a UpdateSegmentRequest struct.
type UpdateSegmentRequest struct {
    Segment              UpdateSegment  `json:"segment"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSegmentRequest.
// It customizes the JSON marshaling process for UpdateSegmentRequest objects.
func (u UpdateSegmentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSegmentRequest object to a map representation for JSON marshaling.
func (u UpdateSegmentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["segment"] = u.Segment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSegmentRequest.
// It customizes the JSON unmarshaling process for UpdateSegmentRequest objects.
func (u *UpdateSegmentRequest) UnmarshalJSON(input []byte) error {
    var temp updateSegmentRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "segment")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Segment = *temp.Segment
    return nil
}

// TODO
type updateSegmentRequest  struct {
    Segment *UpdateSegment `json:"segment"`
}

func (u *updateSegmentRequest) validate() error {
    var errs []string
    if u.Segment == nil {
        errs = append(errs, "required field `segment` is missing for type `Update Segment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
