package models

import (
	"encoding/json"
)

// SubscriptionGroupSignupRequest represents a SubscriptionGroupSignupRequest struct.
type SubscriptionGroupSignupRequest struct {
	SubscriptionGroup SubscriptionGroupSignup `json:"subscription_group"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupRequest.
// It customizes the JSON marshaling process for SubscriptionGroupSignupRequest objects.
func (s *SubscriptionGroupSignupRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupRequest object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSignupRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscription_group"] = s.SubscriptionGroup
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupRequest.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupRequest objects.
func (s *SubscriptionGroupSignupRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SubscriptionGroup SubscriptionGroupSignup `json:"subscription_group"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.SubscriptionGroup = temp.SubscriptionGroup
	return nil
}
