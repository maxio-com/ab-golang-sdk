package models

import (
    "encoding/json"
)

// SubscriptionResponse represents a SubscriptionResponse struct.
type SubscriptionResponse struct {
    Subscription *Subscription `json:"subscription,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionResponse.
// It customizes the JSON marshaling process for SubscriptionResponse objects.
func (s *SubscriptionResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionResponse object to a map representation for JSON marshaling.
func (s *SubscriptionResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if s.Subscription != nil {
        structMap["subscription"] = s.Subscription
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionResponse.
// It customizes the JSON unmarshaling process for SubscriptionResponse objects.
func (s *SubscriptionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Subscription *Subscription `json:"subscription,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Subscription = temp.Subscription
    return nil
}
