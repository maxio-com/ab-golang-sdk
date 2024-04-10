package models

import (
    "encoding/json"
)

// ReactivateSubscriptionGroupRequest represents a ReactivateSubscriptionGroupRequest struct.
type ReactivateSubscriptionGroupRequest struct {
    Resume               *bool          `json:"resume,omitempty"`
    ResumeMembers        *bool          `json:"resume_members,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReactivateSubscriptionGroupRequest.
// It customizes the JSON marshaling process for ReactivateSubscriptionGroupRequest objects.
func (r ReactivateSubscriptionGroupRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReactivateSubscriptionGroupRequest object to a map representation for JSON marshaling.
func (r ReactivateSubscriptionGroupRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Resume != nil {
        structMap["resume"] = r.Resume
    }
    if r.ResumeMembers != nil {
        structMap["resume_members"] = r.ResumeMembers
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReactivateSubscriptionGroupRequest.
// It customizes the JSON unmarshaling process for ReactivateSubscriptionGroupRequest objects.
func (r *ReactivateSubscriptionGroupRequest) UnmarshalJSON(input []byte) error {
    var temp reactivateSubscriptionGroupRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "resume", "resume_members")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Resume = temp.Resume
    r.ResumeMembers = temp.ResumeMembers
    return nil
}

// TODO
type reactivateSubscriptionGroupRequest  struct {
    Resume        *bool `json:"resume,omitempty"`
    ResumeMembers *bool `json:"resume_members,omitempty"`
}
