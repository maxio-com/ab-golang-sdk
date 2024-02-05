package models

import (
    "encoding/json"
)

// SubscriptionGroupMemberError represents a SubscriptionGroupMemberError struct.
type SubscriptionGroupMemberError struct {
    Id      *int    `json:"id,omitempty"`
    Type    *string `json:"type,omitempty"`
    Message *string `json:"message,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupMemberError.
// It customizes the JSON marshaling process for SubscriptionGroupMemberError objects.
func (s *SubscriptionGroupMemberError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupMemberError object to a map representation for JSON marshaling.
func (s *SubscriptionGroupMemberError) toMap() map[string]any {
    structMap := make(map[string]any)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Message != nil {
        structMap["message"] = s.Message
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupMemberError.
// It customizes the JSON unmarshaling process for SubscriptionGroupMemberError objects.
func (s *SubscriptionGroupMemberError) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id      *int    `json:"id,omitempty"`
        Type    *string `json:"type,omitempty"`
        Message *string `json:"message,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Id = temp.Id
    s.Type = temp.Type
    s.Message = temp.Message
    return nil
}
