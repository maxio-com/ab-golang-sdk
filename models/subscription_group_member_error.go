package models

import (
    "encoding/json"
)

// SubscriptionGroupMemberError represents a SubscriptionGroupMemberError struct.
type SubscriptionGroupMemberError struct {
    Id                   *int           `json:"id,omitempty"`
    Type                 *string        `json:"type,omitempty"`
    Message              *string        `json:"message,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupMemberError.
// It customizes the JSON marshaling process for SubscriptionGroupMemberError objects.
func (s SubscriptionGroupMemberError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupMemberError object to a map representation for JSON marshaling.
func (s SubscriptionGroupMemberError) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp subscriptionGroupMemberError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "type", "message")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Id = temp.Id
    s.Type = temp.Type
    s.Message = temp.Message
    return nil
}

// TODO
type subscriptionGroupMemberError  struct {
    Id      *int    `json:"id,omitempty"`
    Type    *string `json:"type,omitempty"`
    Message *string `json:"message,omitempty"`
}
