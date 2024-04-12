package models

import (
    "encoding/json"
)

// SubscriptionGroupUpdateError represents a SubscriptionGroupUpdateError struct.
type SubscriptionGroupUpdateError struct {
    Members              []SubscriptionGroupMemberError `json:"members,omitempty"`
    AdditionalProperties map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupUpdateError.
// It customizes the JSON marshaling process for SubscriptionGroupUpdateError objects.
func (s SubscriptionGroupUpdateError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupUpdateError object to a map representation for JSON marshaling.
func (s SubscriptionGroupUpdateError) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Members != nil {
        structMap["members"] = s.Members
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupUpdateError.
// It customizes the JSON unmarshaling process for SubscriptionGroupUpdateError objects.
func (s *SubscriptionGroupUpdateError) UnmarshalJSON(input []byte) error {
    var temp subscriptionGroupUpdateError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "members")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Members = temp.Members
    return nil
}

// TODO
type subscriptionGroupUpdateError  struct {
    Members []SubscriptionGroupMemberError `json:"members,omitempty"`
}
