package models

import (
    "encoding/json"
)

// UpdateSubscriptionRequest represents a UpdateSubscriptionRequest struct.
type UpdateSubscriptionRequest struct {
    Subscription UpdateSubscription `json:"subscription"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionRequest objects.
func (u *UpdateSubscriptionRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["subscription"] = u.Subscription
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionRequest objects.
func (u *UpdateSubscriptionRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Subscription UpdateSubscription `json:"subscription"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.Subscription = temp.Subscription
    return nil
}
