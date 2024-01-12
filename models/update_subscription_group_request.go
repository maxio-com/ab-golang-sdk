package models

import (
    "encoding/json"
)

// UpdateSubscriptionGroupRequest represents a UpdateSubscriptionGroupRequest struct.
type UpdateSubscriptionGroupRequest struct {
    SubscriptionGroup UpdateSubscriptionGroup `json:"subscription_group"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionGroupRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionGroupRequest objects.
func (u *UpdateSubscriptionGroupRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionGroupRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionGroupRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["subscription_group"] = u.SubscriptionGroup
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionGroupRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionGroupRequest objects.
func (u *UpdateSubscriptionGroupRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        SubscriptionGroup UpdateSubscriptionGroup `json:"subscription_group"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.SubscriptionGroup = temp.SubscriptionGroup
    return nil
}
